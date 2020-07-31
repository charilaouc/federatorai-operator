package controllers

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	datahubv1alpha1_event "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events"
	federatoraiv1alpha1 "github.com/containers-ai/federatorai-operator/api/v1alpha1"
	controlleruitl "github.com/containers-ai/federatorai-operator/controllers/util"
	client_datahub "github.com/containers-ai/federatorai-operator/pkg/client/datahub"
	"github.com/containers-ai/federatorai-operator/pkg/component"
	"github.com/containers-ai/federatorai-operator/pkg/processcrdspec/alamedaserviceparamter"
	repository_keycode "github.com/containers-ai/federatorai-operator/pkg/repository/keycode"
	repository_keycode_datahub "github.com/containers-ai/federatorai-operator/pkg/repository/keycode/datahub"
	"github.com/containers-ai/federatorai-operator/pkg/util"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Namespace string

const (
	retryLimitDuration                  time.Duration = 30 * time.Minute
	addKeycodeSuccessMessageTemplate                  = "Add keycode %s success"
	deleteKeycodeSuccessMessageTemplate               = "Delete keycode %s success"
	addKeycodeFailedMessageTemplate                   = "Add keycode %s failed"
	deleteKeycodeFailedMessageTemplate                = "Delete keycode %s failed"
)

var (
	requeueDuration     = 30 * time.Second
	keycodeSpecialCases = []string{
		"D3JXNLIFTQKQEZ3WZBNIDA3WZA7HKQ",
	}
)

type KeycodeStatus struct {
	codeNumber string
	state      federatoraiv1alpha1.KeycodeState
}

// AlamedaServiceKeycodeReconciler reconciles a AlamedaService object
type AlamedaServiceKeycodeReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	DatahubClientMap     map[Namespace]client_datahub.Client
	DatahubClientMapLock sync.Mutex

	FirstRetryTimeCache map[types.NamespacedName]*time.Time
	FirstRetryTimeLock  sync.Mutex

	ClusterID        string
	EventChanMap     map[Namespace]chan datahubv1alpha1_event.Event
	EventChanMapLock sync.Mutex

	LastReconcileTaskMap     map[Namespace]KeycodeStatus
	LastReconcileTaskMapLock sync.Mutex
}

// Reconcile reconcile AlamedaService's keycode
func (r *AlamedaServiceKeycodeReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {

	r.Log.Info("Reconcile Keycode")

	var reconcileResult = ctrl.Result{}
	alamedaService := &federatoraiv1alpha1.AlamedaService{}
	isAlamedaServiceOwningLock := false
	defer r.setLastReconcileTask(&isAlamedaServiceOwningLock, Namespace(req.Namespace), alamedaService)
	defer r.flushEvents(&isAlamedaServiceOwningLock, Namespace(req.Namespace), alamedaService)
	defer r.handleFirstRetryTime(&isAlamedaServiceOwningLock, &reconcileResult, req.NamespacedName)
	defer func(isAlamedaServiceOwningLock *bool) {

		if isAlamedaServiceOwningLock == nil || !(*isAlamedaServiceOwningLock) {
			return
		}

		instance := &federatoraiv1alpha1.AlamedaService{}
		err := r.Client.Get(context.TODO(), client.ObjectKey{Namespace: req.Namespace, Name: req.Name}, instance)
		if err != nil && k8sErrors.IsNotFound(err) {
			addr, err := r.getDatahubAddressByNamespace(req.Namespace)
			if err != nil {
				r.Log.V(-1).Info("Get datahub address failed, skip deleting datahub client", "AlamedaService.Namespace", req.Namespace, "AlamedaService.Name", req.Name, "error", err.Error())
			}
			err = r.deleteDatahubClient(addr)
			if err != nil {
				r.Log.V(-1).Info("Deleting datahub client failed", "AlamedaService.Namespace", req.Namespace, "AlamedaService.Name", req.Name, "error", err.Error())
			}
			return
		} else if err != nil {
			r.Log.V(-1).Info("Get AlamedaService failed, skip writing status", "AlamedaService.Namespace", req.Namespace, "AlamedaService.Name", req.Name, "error", err.Error())
			return
		}

		instance.Spec.Keycode = alamedaService.Spec.Keycode
		instance.Status.KeycodeStatus = alamedaService.Status.KeycodeStatus

		// Get keycodeRepository
		keycodeRepository, err := r.getKeycodeRepository(req.Namespace)
		if err != nil {
			r.Log.V(-1).Info("Get keycode summary failed, will not write keycode summary into AlamedaService's status", "AlamedaService.Namespace", req.Namespace, "AlamedaService.Name", req.Name, "error", err.Error())
			reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
		} else {
			detail, err := keycodeRepository.GetKeycodeDetail("")
			if err != nil {
				r.Log.V(-1).Info("Get keycode summary failed, write empty keycode summary into AlamedaService's status", "AlamedaService.Namespace", req.Namespace, "AlamedaService.Name", req.Name, "error", err.Error())
				reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
			}
			instance.SetStatusKeycodeSummary(detail.Summary())
		}

		if err := r.Client.Update(context.Background(), instance); err != nil {
			r.Log.V(-1).Info("Update AlamedaService status failed", "AlamedaService.Namespace", req.Namespace, "AlamedaService.Name", req.Name, "error", err.Error())
			reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
		}
	}(&isAlamedaServiceOwningLock)

	// Fetch the AlamedaService instance
	err := r.Client.Get(context.TODO(), req.NamespacedName, alamedaService)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			if err := r.deleteAlamedaServiceDependencies(alamedaService); err != nil {
				r.Log.V(-1).Info("Handle AlamedaService deletion failed", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
			}
			r.Log.Info("AlamedaService not found, skip keycode reconciling", "AlamedaService.Namespace", req.Namespace, "AlamedaService.Name", req.Name)
			reconcileResult.Requeue = false
			return reconcileResult, nil
		}
		// Error reading the object - requeue the req.
		r.Log.V(-1).Info("Get AlamedaService failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
		reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
		return reconcileResult, nil
	}

	// Check if AlamedaService own lock
	ok, err := r.isAlamedaServiceOwnLock(alamedaService)
	if err != nil {
		r.Log.V(-1).Info("Check if AlamedaService is owning lock failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
		reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
		return reconcileResult, nil
	}
	isAlamedaServiceOwningLock = ok
	if !ok {
		r.Log.Info("AlamedaService is not owning lock, stop reconciling.", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
		reconcileResult = ctrl.Result{Requeue: false}
		return reconcileResult, nil
	}

	if firstRetryTime := r.getFirstRetryTime(req.NamespacedName); firstRetryTime != nil {
		now := time.Now()
		if now.Sub(*firstRetryTime) > retryLimitDuration {
			r.Log.Error(nil, "Exceeds retry limit, stop reconciing.", "AlamedaService.Namespace", req.Namespace, "AlamedaService.Name", req.Name)
			reconcileResult.Requeue = false
			return reconcileResult, nil
		}
	}

	// Get keycodeRepository
	keycodeRepository, err := r.getKeycodeRepository(alamedaService.Namespace)
	if err != nil {
		r.Log.V(-1).Info("Get licese repository failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
		alamedaService.Status.KeycodeStatus.LastErrorMessage = errors.Wrap(err, "get keycode repository instance failed").Error()
		reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
		return reconcileResult, nil
	}

	// There are two conditions to handle,
	// first, keycode is empty
	// seconde, keycode is not empty
	if alamedaService.IsCodeNumberEmpty() {
		if err := r.handleEmptyKeycode(keycodeRepository, alamedaService); err != nil {
			r.Log.V(-1).Info("Handle empty keycode failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
			alamedaService.Status.KeycodeStatus.LastErrorMessage = errors.Wrap(err, "handle empty keycode failed").Error()
			reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
			return reconcileResult, nil
		}
		alamedaService.Spec.Keycode = federatoraiv1alpha1.KeycodeSpec{}
		alamedaService.Status.KeycodeStatus = federatoraiv1alpha1.KeycodeStatus{State: federatoraiv1alpha1.KeycodeStateWaitingKeycode}
		r.Log.Info("Handle empty keycode done", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
		return reconcileResult, nil
	}

	// Check if need to reconcile keycode
	if r.needToReconcile(alamedaService) {
		r.Log.Info("Keycode not changed, skip reconciling", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
		return reconcileResult, nil
	}

	// If keycode is updated, do the update process no matter what the current state is
	if alamedaService.IsCodeNumberUpdated() {
		if err := r.handleKeycode(keycodeRepository, alamedaService); err != nil {
			r.Log.V(-1).Info("Update keycode failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
			alamedaService.Status.KeycodeStatus.LastErrorMessage = errors.Wrap(err, "update keycode failed").Error()
			reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
			return reconcileResult, nil
		}
		r.Log.Info("Update keycode done, start polling registration data", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
		reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
		return reconcileResult, nil
	}

	// Keycode is not changed, process keycode by the current state
	switch alamedaService.Status.KeycodeStatus.State {
	case federatoraiv1alpha1.KeycodeStateDefault, federatoraiv1alpha1.KeycodeStateWaitingKeycode:
		if err := r.handleKeycode(keycodeRepository, alamedaService); err != nil {
			r.Log.V(-1).Info("Handling keycode failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
			alamedaService.Status.KeycodeStatus.LastErrorMessage = errors.Wrap(err, "handle keycode failed").Error()
			reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
			return reconcileResult, nil
		}
		r.Log.Info("Handling keycode done, start polling registration data", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
		reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
		return reconcileResult, nil
	case federatoraiv1alpha1.KeycodeStatePollingRegistrationData:
		// This state will move to "federatoraiv1alpha1.KeycodeStateDone" state if the keycode detail is registered

		// Poll registration data from keycode repository
		registrationData, err := keycodeRepository.GetRegistrationData()
		if err != nil {
			r.Log.V(-1).Info("Polling registration data failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
			alamedaService.Status.KeycodeStatus.LastErrorMessage = errors.Wrap(err, "poll registration data failed").Error()
			reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
			return reconcileResult, nil
		}

		// Get keycode defailt from keycode repository
		detail, err := keycodeRepository.GetKeycodeDetail("")
		if err != nil {
			r.Log.V(-1).Info("Polling registration data failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
			alamedaService.Status.KeycodeStatus.LastErrorMessage = errors.Wrap(err, "poll registration data failed").Error()
			reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
			return reconcileResult, nil
		}
		if detail.Registered {
			alamedaService.Spec.Keycode.SignatureData = registrationData
			alamedaService.Status.KeycodeStatus = federatoraiv1alpha1.KeycodeStatus{
				CodeNumber:       alamedaService.Spec.Keycode.CodeNumber,
				RegistrationData: "",
				State:            federatoraiv1alpha1.KeycodeStateDone,
				LastErrorMessage: "",
				Summary:          "",
			}
			r.Log.Info("Keycode has been registered, move state to \"Done\"", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
		} else {
			alamedaService.Status.KeycodeStatus = federatoraiv1alpha1.KeycodeStatus{
				CodeNumber:       alamedaService.Spec.Keycode.CodeNumber,
				RegistrationData: registrationData,
				State:            federatoraiv1alpha1.KeycodeStateWaitingSignatureData,
				LastErrorMessage: "",
				Summary:          "",
			}
			r.Log.Info("Polling registration data done, waiting signature data", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
		}
		return reconcileResult, nil
	case federatoraiv1alpha1.KeycodeStateWaitingSignatureData:
		if alamedaService.Spec.Keycode.SignatureData == "" {
			r.Log.Info("Waiting signature data to be filled in, skip reconciling", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
			return reconcileResult, nil
		}
		if err := r.handleSignatureData(keycodeRepository, alamedaService); err != nil {
			r.Log.V(-1).Info("Handling signature data failed, retry reconciling keycode", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "error", err.Error())
			alamedaService.Status.KeycodeStatus.LastErrorMessage = errors.Wrap(err, "handle signature data  failed").Error()
			reconcileResult = ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}
			return reconcileResult, nil
		}
		alamedaService.Status.KeycodeStatus.LastErrorMessage = ""
		alamedaService.Status.KeycodeStatus.State = federatoraiv1alpha1.KeycodeStateDone
		r.Log.Info("Handling signature data done", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name)
		return reconcileResult, nil
	default:
		r.Log.Info("Unknown keycode state, skip reconciling", "AlamedaService.Namespace", alamedaService.Namespace, "AlamedaService.Name", alamedaService.Name, "state", alamedaService.Status.KeycodeStatus.State)
		return reconcileResult, nil
	}
}

func (r *AlamedaServiceKeycodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&federatoraiv1alpha1.AlamedaService{}).
		Complete(r)
}

// handleFirstRetryTime set/resets first retry time when requeue is true/false
func (r *AlamedaServiceKeycodeReconciler) handleFirstRetryTime(isAlamedaServiceOwningLock *bool, reconcileResult *ctrl.Result, namespacedName types.NamespacedName) {

	if isAlamedaServiceOwningLock == nil || !(*isAlamedaServiceOwningLock) {
		return
	}

	if reconcileResult.Requeue == true {
		t := time.Now()
		r.setFirstRetryTimeIfNotExist(namespacedName, &t)
	} else {
		r.setFirstRetryTime(namespacedName, nil)
	}
}

func (r *AlamedaServiceKeycodeReconciler) deleteAlamedaServiceDependencies(alamedaService *federatoraiv1alpha1.AlamedaService) error {

	datahubAddress, err := r.getDatahubAddressByNamespace(alamedaService.Namespace)
	if err != nil {
		return errors.Wrap(err, "get datahub address failed")
	}
	r.deleteDatahubClient(datahubAddress)

	r.deleteFirstRetryTime(types.NamespacedName{Namespace: alamedaService.Namespace, Name: alamedaService.Name})
	r.deleteEventChan(Namespace(alamedaService.Namespace))
	r.deleteLastReconcileTask(Namespace(alamedaService.Namespace))

	return nil
}

func (r *AlamedaServiceKeycodeReconciler) deleteDatahubClient(datahubAddr string) error {

	if client, exist := r.DatahubClientMap[Namespace(datahubAddr)]; exist {
		r.DatahubClientMapLock.Lock()
		defer r.DatahubClientMapLock.Unlock()
		if err := client.Close(); err != nil {
			return err
		}
		delete(r.DatahubClientMap, Namespace(datahubAddr))
	}

	return nil
}

func (r *AlamedaServiceKeycodeReconciler) needToReconcile(alamedaService *federatoraiv1alpha1.AlamedaService) bool {
	return !alamedaService.IsCodeNumberUpdated() &&
		alamedaService.Status.KeycodeStatus.State == federatoraiv1alpha1.KeycodeStateDone
}

func (r *AlamedaServiceKeycodeReconciler) handleEmptyKeycode(keycodeRepository repository_keycode.Interface, alamedaService *federatoraiv1alpha1.AlamedaService) error {

	if !alamedaService.IsCodeNumberUpdated() {
		return nil
	}

	details, err := keycodeRepository.ListKeycodes()
	if err != nil {
		return errors.Wrap(err, "list keycodes failed")
	}
	for _, detail := range details {
		codeNumber := detail.Keycode
		if codeNumber == "" {
			continue
		}
		if err := keycodeRepository.DeleteKeycode(codeNumber); err != nil {
			e := newLicenseEvent(
				alamedaService.Namespace,
				fmt.Sprintf(deleteKeycodeFailedMessageTemplate, codeNumber),
				r.ClusterID,
				datahubv1alpha1_event.EventLevel_EVENT_LEVEL_WARNING)
			r.addEvent(Namespace(alamedaService.Namespace), e)
			return errors.Wrap(err, "delete keycode failed")
		}
		e := newLicenseEvent(
			alamedaService.Namespace,
			fmt.Sprintf(deleteKeycodeSuccessMessageTemplate, codeNumber),
			r.ClusterID,
			datahubv1alpha1_event.EventLevel_EVENT_LEVEL_INFO)
		r.addEvent(Namespace(alamedaService.Namespace), e)
	}

	return nil
}

func (r *AlamedaServiceKeycodeReconciler) handleKeycode(keycodeRepository repository_keycode.Interface, alamedaService *federatoraiv1alpha1.AlamedaService) error {

	// Check if keycode is existing
	keycode := alamedaService.Spec.Keycode.CodeNumber
	details, err := keycodeRepository.ListKeycodes()
	if err != nil {
		return errors.Wrap(err, "list keycodes failed")
	}

	if len(details) == 0 {
		// Apply keycode to keycode repository
		if err := keycodeRepository.SendKeycode(keycode); err != nil {
			e := newLicenseEvent(
				alamedaService.Namespace,
				fmt.Sprintf(addKeycodeFailedMessageTemplate, keycode),
				r.ClusterID,
				datahubv1alpha1_event.EventLevel_EVENT_LEVEL_WARNING)
			r.addEvent(Namespace(alamedaService.Namespace), e)
			return errors.Wrap(err, "send keycode to keycode repository failed")
		}
		e := newLicenseEvent(
			alamedaService.Namespace,
			fmt.Sprintf(addKeycodeSuccessMessageTemplate, keycode),
			r.ClusterID,
			datahubv1alpha1_event.EventLevel_EVENT_LEVEL_INFO)
		r.addEvent(Namespace(alamedaService.Namespace), e)
		alamedaService.Status.KeycodeStatus.CodeNumber = alamedaService.Spec.Keycode.CodeNumber
		alamedaService.Status.KeycodeStatus.State = federatoraiv1alpha1.KeycodeStatePollingRegistrationData
	} else {
		// If keycode is special case, get the existing keycode for user
		if r.isKeycodeSpecialCase(keycode) {
			alamedaService.Spec.Keycode.CodeNumber = details[0].Keycode
			alamedaService.Spec.Keycode.SignatureData = ""
			alamedaService.Status.KeycodeStatus.CodeNumber = details[0].Keycode
			alamedaService.Status.KeycodeStatus.RegistrationData = ""
			alamedaService.Status.KeycodeStatus.State = federatoraiv1alpha1.KeycodeStatePollingRegistrationData
			alamedaService.Status.KeycodeStatus.LastErrorMessage = ""
			alamedaService.Status.KeycodeStatus.Summary = ""
		} else {

			for _, detail := range details {
				codeNumber := detail.Keycode
				if codeNumber == "" {
					continue
				}
				if err := keycodeRepository.DeleteKeycode(codeNumber); err != nil {
					e := newLicenseEvent(
						alamedaService.Namespace,
						fmt.Sprintf(deleteKeycodeFailedMessageTemplate, keycode),
						r.ClusterID,
						datahubv1alpha1_event.EventLevel_EVENT_LEVEL_WARNING)
					r.addEvent(Namespace(alamedaService.Namespace), e)
					return errors.Wrap(err, "delete keycode failed")
				}
				e := newLicenseEvent(
					alamedaService.Namespace,
					fmt.Sprintf(deleteKeycodeSuccessMessageTemplate, codeNumber),
					r.ClusterID,
					datahubv1alpha1_event.EventLevel_EVENT_LEVEL_INFO)
				r.addEvent(Namespace(alamedaService.Namespace), e)
			}

			alamedaService.Spec.Keycode.SignatureData = ""
			alamedaService.Status.KeycodeStatus.CodeNumber = ""
			alamedaService.Status.KeycodeStatus.RegistrationData = ""
			alamedaService.Status.KeycodeStatus.State = federatoraiv1alpha1.KeycodeStateWaitingKeycode
			alamedaService.Status.KeycodeStatus.LastErrorMessage = ""
			alamedaService.Status.KeycodeStatus.Summary = ""

			// Apply keycode to keycode repository
			if err := keycodeRepository.SendKeycode(keycode); err != nil {
				return errors.Wrap(err, "send keycode to keycode repository failed")
			}
			e := newLicenseEvent(
				alamedaService.Namespace,
				fmt.Sprintf(addKeycodeSuccessMessageTemplate, keycode),
				r.ClusterID,
				datahubv1alpha1_event.EventLevel_EVENT_LEVEL_INFO)
			r.addEvent(Namespace(alamedaService.Namespace), e)
			alamedaService.Status.KeycodeStatus.CodeNumber = alamedaService.Spec.Keycode.CodeNumber
			alamedaService.Status.KeycodeStatus.State = federatoraiv1alpha1.KeycodeStatePollingRegistrationData
		}
	}

	return nil
}

func (r *AlamedaServiceKeycodeReconciler) isKeycodeSpecialCase(keycode string) bool {

	for _, c := range keycodeSpecialCases {
		keycode = strings.Replace(keycode, "-", "", -1)
		if c == keycode {
			return true
		}
	}

	return false
}

func (r *AlamedaServiceKeycodeReconciler) handleSignatureData(keycodeRepository repository_keycode.Interface, alamedaService *federatoraiv1alpha1.AlamedaService) error {

	// Sending registration data to keycode repository
	err := keycodeRepository.SendSignatureData(alamedaService.Spec.Keycode.SignatureData)
	if err != nil {
		return errors.Wrap(err, "send signature data to keycode repository failed")
	}

	return nil
}

func (r *AlamedaServiceKeycodeReconciler) getKeycodeRepository(namespace string) (repository_keycode.Interface, error) {

	datahubAddress, err := r.getDatahubAddressByNamespace(namespace)
	if err != nil {
		return nil, errors.Wrap(err, "get Datahub address failed")
	}
	datahubClient := r.getOrCreateDatahubClient(datahubAddress, namespace)
	keycodeRepository := repository_keycode_datahub.NewKeycodeRepository(&datahubClient)

	return keycodeRepository, nil
}

func (r *AlamedaServiceKeycodeReconciler) getDatahubAddressByNamespace(namespace string) (string, error) {

	componentFactory := component.ComponentConfig{NameSpace: namespace}

	// Get datahub client instance
	datahubServiceAssetName := alamedaserviceparamter.GetAlamedaDatahubService()
	datahubService := componentFactory.NewService(datahubServiceAssetName)
	datahubAddress, err := util.GetServiceAddress(datahubService, "grpc")
	if err != nil {
		return "", err
	}
	return datahubAddress, nil
}

func (r *AlamedaServiceKeycodeReconciler) getOrCreateDatahubClient(datahubAddress, ns string) client_datahub.Client {

	if _, exist := r.DatahubClientMap[Namespace(datahubAddress)]; !exist {
		r.DatahubClientMapLock.Lock()
		defer r.DatahubClientMapLock.Unlock()
		datahubClientConfig := client_datahub.NewDefaultConfig(ns)
		datahubClientConfig.Address = datahubAddress
		r.DatahubClientMap[Namespace(datahubAddress)] = client_datahub.NewDatahubClient(datahubClientConfig)
	}
	return r.DatahubClientMap[Namespace(datahubAddress)]
}

func (r *AlamedaServiceKeycodeReconciler) setFirstRetryTimeIfNotExist(namespacedName types.NamespacedName, t *time.Time) {
	if r.getFirstRetryTime(namespacedName) == nil {
		r.setFirstRetryTime(namespacedName, t)
	}
}

func (r *AlamedaServiceKeycodeReconciler) setFirstRetryTime(namespacedName types.NamespacedName, t *time.Time) {

	r.FirstRetryTimeLock.Lock()
	defer r.FirstRetryTimeLock.Unlock()
	r.FirstRetryTimeCache[namespacedName] = t
}

func (r *AlamedaServiceKeycodeReconciler) getFirstRetryTime(namespacedName types.NamespacedName) *time.Time {

	r.FirstRetryTimeLock.Lock()
	defer r.FirstRetryTimeLock.Unlock()
	return r.FirstRetryTimeCache[namespacedName]
}

func (r *AlamedaServiceKeycodeReconciler) deleteFirstRetryTime(namespacedName types.NamespacedName) {

	r.FirstRetryTimeLock.Lock()
	defer r.FirstRetryTimeLock.Unlock()
	delete(r.FirstRetryTimeCache, namespacedName)
}

func (r *AlamedaServiceKeycodeReconciler) getEventChan(namespace Namespace) chan datahubv1alpha1_event.Event {

	var eventChan chan datahubv1alpha1_event.Event
	var exist bool
	if eventChan, exist = r.EventChanMap[namespace]; !exist {
		r.EventChanMapLock.Lock()
		r.EventChanMap[namespace] = make(chan datahubv1alpha1_event.Event, 100)
		r.EventChanMapLock.Unlock()
	}
	eventChan = r.EventChanMap[namespace]
	return eventChan
}

func (r *AlamedaServiceKeycodeReconciler) deleteEventChan(namespace Namespace) {

	r.EventChanMapLock.Lock()
	defer r.EventChanMapLock.Unlock()
	delete(r.EventChanMap, namespace)
}

func (r *AlamedaServiceKeycodeReconciler) addEvent(namespace Namespace, e datahubv1alpha1_event.Event) {

	var eventChan chan datahubv1alpha1_event.Event
	var exist bool
	if eventChan, exist = r.EventChanMap[namespace]; !exist {
		r.EventChanMapLock.Lock()
		r.EventChanMap[namespace] = make(chan datahubv1alpha1_event.Event, 100)
		eventChan = r.EventChanMap[namespace]
		r.EventChanMapLock.Unlock()
	}

	eventChan <- e
}

func (r *AlamedaServiceKeycodeReconciler) flushEvents(isAlamedaServiceOwningLock *bool, namespace Namespace, alamedaService *federatoraiv1alpha1.AlamedaService) error {

	if isAlamedaServiceOwningLock == nil || !(*isAlamedaServiceOwningLock) {
		return nil
	}

	r.Log.V(1).Info("Flush events...")

	datahubAddress, err := r.getDatahubAddressByNamespace(string(namespace))
	if err != nil {
		r.Log.V(-1).Info("Flush events failed: get datahub address failed %s", err.Error())
	}

	cli := r.getOrCreateDatahubClient(datahubAddress, string(namespace))

	var events []*datahubv1alpha1_event.Event
	eventChan := r.getEventChan(namespace)
Loop:
	for {
		select {
		case event := <-eventChan:
			copyEvent := event
			events = append(events, &copyEvent)
		default:
			break Loop
		}
	}

	if !r.needToflushEvents(namespace, alamedaService) {
		r.Log.V(1).Info("Need not to flush events")
		return nil
	}
	err = cli.CreateEvents(events)
	if err != nil {
		r.Log.V(-1).Info("Flush events failed: %s", "error", err.Error())
	}

	r.Log.V(1).Info("Flush events done")
	return nil
}

func (r *AlamedaServiceKeycodeReconciler) needToflushEvents(namespace Namespace, alamedaService *federatoraiv1alpha1.AlamedaService) bool {

	if alamedaService == nil || alamedaService.DeletionTimestamp != nil {
		return true
	}

	lastReconcileTask := r.getLastReconcileTask(namespace)
	if lastReconcileTask.codeNumber == alamedaService.Spec.Keycode.CodeNumber &&
		lastReconcileTask.state == alamedaService.Status.KeycodeStatus.State {
		return false
	}

	return true
}

func (r *AlamedaServiceKeycodeReconciler) setLastReconcileTask(isAlamedaServiceOwningLock *bool, namespace Namespace, alamedaService *federatoraiv1alpha1.AlamedaService) {

	if isAlamedaServiceOwningLock == nil || !(*isAlamedaServiceOwningLock) {
		return
	}

	if alamedaService == nil || alamedaService.DeletionTimestamp != nil {
		return
	}
	r.LastReconcileTaskMapLock.Lock()
	defer r.LastReconcileTaskMapLock.Unlock()
	r.LastReconcileTaskMap[namespace] = KeycodeStatus{
		codeNumber: alamedaService.Spec.Keycode.CodeNumber,
		state:      alamedaService.Status.KeycodeStatus.State,
	}
}

func (r *AlamedaServiceKeycodeReconciler) getLastReconcileTask(namespace Namespace) KeycodeStatus {
	return r.LastReconcileTaskMap[namespace]
}

func (r *AlamedaServiceKeycodeReconciler) deleteLastReconcileTask(namespace Namespace) {

	r.LastReconcileTaskMapLock.Lock()
	defer r.LastReconcileTaskMapLock.Unlock()
	delete(r.LastReconcileTaskMap, namespace)
}

func (r *AlamedaServiceKeycodeReconciler) isAlamedaServiceOwnLock(alamedaService *federatoraiv1alpha1.AlamedaService) (bool, error) {
	lock, err := controlleruitl.GetAlamedaServiceLock(context.TODO(), r.Client)
	if err != nil {
		return false, errors.Wrap(err, "get or create AlamedaService lock failed")
	}
	return controlleruitl.IsAlamedaServiceLockOwnedByAlamedaService(lock, *alamedaService), nil
}
