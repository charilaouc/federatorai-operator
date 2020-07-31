package controllers

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	autoscaling_v1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	federatoraiv1alpha1 "github.com/containers-ai/federatorai-operator/api/v1alpha1"
	federatoraioperatorcontrollerutil "github.com/containers-ai/federatorai-operator/controllers/util"
	"github.com/containers-ai/federatorai-operator/pkg/component"
	"github.com/containers-ai/federatorai-operator/pkg/lib/resourceapply"
	prom_patch "github.com/containers-ai/federatorai-operator/pkg/patch/prometheus"
	"github.com/containers-ai/federatorai-operator/pkg/processcrdspec"
	"github.com/containers-ai/federatorai-operator/pkg/processcrdspec/alamedaserviceparamter"
	"github.com/containers-ai/federatorai-operator/pkg/updateresource"
	"github.com/containers-ai/federatorai-operator/pkg/util"
	"github.com/go-logr/logr"
	routev1 "github.com/openshift/api/route/v1"
	securityv1 "github.com/openshift/api/security/v1"
	"github.com/pkg/errors"
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	admissionWebhookAnnotationKeySecretName = "secret.name"
	serviceExposureAnnotationKey            = "servicesxposures.alamedaservices.federatorai.containers.ai"

	defaultKafkaVersion              = "2.4.0"
	defaultPrometheusBearerTokenFile = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

var (
	componentConfig *component.ComponentConfig
	requeueAfter    = 3 * time.Second
)

// AlamedaServiceReconciler reconciles a AlamedaService object
type AlamedaServiceReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	// reconciledAlamedaService caches alamedaservice which has been created and reconciled once
	FirstReconcileDoneAlamedaService map[string]struct{}

	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	Apiextclient apiextension.Interface

	IsOpenshiftAPIRouteExist    bool
	IsOpenshiftAPISecurityExist bool

	PodSecurityPolicesAPIGroupVersion schema.GroupVersion
}

// Reconcile reads that state of the cluster for a AlamedaService object and makes changes based on the state read
// and what is in the AlamedaService.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *AlamedaServiceReconciler) Reconcile(
	req ctrl.Request) (ctrl.Result, error) {

	instance := &federatoraiv1alpha1.AlamedaService{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil && !k8sErrors.IsNotFound(err) {
		r.Log.V(-1).Info("Get AlamedaService failed, retry reconciling.",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	} else if k8sErrors.IsNotFound(err) {
		// Request object not found, could have been deleted after reconcile req.
		r.Log.Info("Handing AlamedaService deletion.", "AlamedaService.Namespace",
			req.Namespace, "AlamedaService.Name", req.Name)
		if err := r.handleAlamedaServiceDeletion(req); err != nil {
			r.Log.V(-1).Info("Handle AlamedaService deletion failed, retry reconciling.",
				"AlamedaService.Namespace", req.Namespace, "AlamedaService.Name",
				req.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		r.Log.Info("Handle AlamedaService deletion done.", "AlamedaService.Namespace",
			req.Namespace, "AlamedaService.Name", req.Name)
		return ctrl.Result{}, nil
	}

	r.Log.Info("Reconciling AlamedaService.", "AlamedaService.Namespace",
		instance.Namespace, "AlamedaService.Name", instance.Name)
	r.InitAlamedaService(instance)

	clusterRoleGC, err := util.GetOrCreateGCClusterRole(r.Client)
	if err != nil {
		r.Log.V(-1).Info("get clusterRole GC failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	// Check if AlamedaService need to reconcile, currently only reconcile one AlamedaService in one cluster
	isNeedToBeReconciled, err := r.isNeedToBeReconciled(instance, clusterRoleGC)
	if err != nil {
		r.Log.V(-1).Info("check if AlamedaService needs to reconcile failed, retry reconciling",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if !isNeedToBeReconciled {
		r.Log.Info("AlamedaService does not need to be reconcile", "AlamedaService.Namespace",
			instance.Namespace, "AlamedaService.Name", instance.Name)
		err := r.updateAlamedaServiceActivation(instance, false)
		if err != nil {
			r.Log.V(-1).Info("Update AlamedaService activation failed, retry reconciling",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
		return ctrl.Result{}, nil
	} else {
		if err := r.updateAlamedaServiceActivation(instance, true); err != nil {
			r.Log.V(-1).Info("Update AlamedaService activation failed, retry reconciling",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}

	hasGCOwner := false
	for _, or := range instance.GetOwnerReferences() {
		if strings.ToLower(or.Kind) == strings.ToLower("ClusterRole") {
			hasGCOwner = true
			break
		}
	}
	if !hasGCOwner {
		tmpInstance := &federatoraiv1alpha1.AlamedaService{}
		if err = r.Client.Get(
			context.TODO(), req.NamespacedName, tmpInstance); err != nil {
			r.Log.V(-1).Info("get latest alamedaservice failed for setting clusterrole gc",
				"AlamedaService.Namespace", instance.Namespace,
				"AlamedaService.Name", instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
		if err := controllerutil.SetControllerReference(
			clusterRoleGC, tmpInstance, r.Scheme); err != nil {
			r.Log.V(-1).Info("set clusterrole gc for alamedaservice failed",
				"AlamedaService.Namespace", instance.Namespace,
				"AlamedaService.Name", instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}

		if err := r.Client.Update(context.Background(), tmpInstance); err != nil {
			r.Log.V(-1).Info("update alamedaservice for clusterrole gc failed",
				"AlamedaService.Namespace", instance.Namespace,
				"AlamedaService.Name", instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
		return ctrl.Result{}, nil
	}

	isFirstReconciled := r.isAlamedaServiceFirstReconciledDone(*instance)
	hasSpecBeenChanged, _ := r.checkAlamedaServiceSpecIsChange(
		instance, req.NamespacedName)
	if !hasSpecBeenChanged &&
		util.Disable_operand_resource_protection == "true" && !isFirstReconciled {
		r.Log.Info("AlamedaService spec is not changed, skip reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace,
			"AlamedaService.Name", instance.Name)
		return ctrl.Result{}, nil
	}

	asp := alamedaserviceparamter.NewAlamedaServiceParamter(instance)
	ns, err := r.getNamespace(req.Namespace)
	if err != nil {
		r.Log.V(-1).Info("Get Namespace failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	componentConfig, err = r.newComponentConfig(ns, *instance, *asp)
	if err != nil {
		r.Log.V(-1).Info("New ComponentConfig failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	resource := r.removeUnsupportedResource(*asp.GetInstallResource())
	installResource := &resource

	if asp.AutoPatchPrometheusRules {
		if ok, missingRulesMap, err := prom_patch.RulesCheck(
			prom_patch.GetK8SClient()); err != nil {
			r.Log.V(-1).Info("check prometheusrules failed, retry reconciling AlamedaService",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
		} else if !ok {
			if err = prom_patch.PatchMissingRules(
				prom_patch.GetK8SClient(), missingRulesMap); err != nil {
				r.Log.V(-1).Info("patch prometheusrules failed, retry reconciling AlamedaService",
					"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
					instance.Name, "msg", err.Error())
				return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
			}
		}
		/*
			if ok, err := prom_patch.RelabelingCheck(prom_patch.GetK8SClient()); err != nil {
			r.Log.V(-1).Info("check relabelings failed, retry reconciling AlamedaService",
					"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
					instance.Name, "msg", err.Error())
				return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
			} else if !ok {
				if err = prom_patch.PatchRelabelings(prom_patch.GetK8SClient()); err != nil {
				r.Log.V(-1).Info("patch relabeling failed, retry reconciling AlamedaService",
						"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
						instance.Name, "msg", err.Error())
					return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
				}
			}
		*/
	}

	if err = r.syncCustomResourceDefinition(
		instance, clusterRoleGC, asp, installResource); err != nil {
		r.Log.Error(err, "create crd failed")
	}
	if err := r.syncPodSecurityPolicy(
		instance, clusterRoleGC, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync podSecurityPolicy failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncSecurityContextConstraints(
		instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync securityContextConstraint failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	if err := r.syncClusterRole(
		instance, clusterRoleGC, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync clusterRole failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace,
			"AlamedaService.Name", instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.createServiceAccount(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Create serviceAccount failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	if err := r.syncClusterRoleBinding(
		instance, clusterRoleGC, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync clusterRoleBinding failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	if err := r.syncRole(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync Role failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncRoleBinding(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync RoleBinding failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.createSecret(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("create secret failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.createPersistentVolumeClaim(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("create PersistentVolumeClaim failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncConfigMap(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync configMap failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncService(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync service failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncServiceExposure(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync service exposure failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncMutatingWebhookConfiguration(
		instance, clusterRoleGC, asp, installResource); err != nil {
		r.Log.V(-1).Info(
			"create MutatingWebhookConfiguration failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncDeployment(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync deployment failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncValidatingWebhookConfiguration(
		instance, clusterRoleGC, asp, installResource); err != nil {
		r.Log.V(-1).Info(
			"Sync ValidatingWebhookConfiguration failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncStatefulSet(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync statefulset failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.syncRoute(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync route failed, retry reconciling.", "AlamedaService.Namespace",
			instance.Namespace, "AlamedaService.Name", instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	if err := r.syncDaemonSet(instance, asp, installResource); err != nil {
		r.Log.V(-1).Info("Sync DaemonSet failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.createAlamedaNotificationChannels(clusterRoleGC, installResource); err != nil {
		r.Log.V(-1).Info(
			"create AlamedaNotificationChannels failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
	}
	if err := r.createAlamedaNotificationTopics(clusterRoleGC, installResource); err != nil {
		r.Log.V(-1).Info(
			"create AlamedaNotificationTopic failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
	}

	//Uninstall Execution Component
	if !asp.EnableExecution {
		r.Log.Info("EnableExecution has been changed to false")
		excutionResource := alamedaserviceparamter.GetExcutionResource()
		if err := r.uninstallResource(*excutionResource); err != nil {
			r.Log.V(-1).Info("Uninstall execution resources failed, retry reconciling AlamedaService",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}

	//Uninstall dispatcher Component
	if !asp.EnableDispatcher {
		resource := r.removeUnsupportedResource(*alamedaserviceparamter.GetDispatcherResource())
		if err := r.uninstallResource(resource); err != nil {
			r.Log.V(-1).Info("Uninstall dispatcher resources failed, retry reconciling AlamedaService",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}
	if !asp.EnablePreloader {
		resource := r.removeUnsupportedResource(*alamedaserviceparamter.GetPreloaderResource())
		if err := r.uninstallResource(resource); err != nil {
			r.Log.V(-1).Info("Uninstall preloader resources failed, retry reconciling AlamedaService",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}
	//Uninstall weavescope components
	if !asp.EnableWeavescope {
		resource := r.removeUnsupportedResource(alamedaserviceparamter.GetWeavescopeResource())
		if err := r.uninstallResource(resource); err != nil {
			r.Log.V(-1).Info("Uninstall weavescope resources failed, retry reconciling AlamedaService",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}
	//Uninstall vpa components
	if !asp.EnableVPA {
		resource := r.removeUnsupportedResource(alamedaserviceparamter.GetVPAResource())
		if err := r.uninstallResource(resource); err != nil {
			r.Log.V(-1).Info("Uninstall vpa resources failed, retry reconciling AlamedaService",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}
	//Uninstall vpa components
	if !asp.EnableGPU {
		resource := r.removeUnsupportedResource(alamedaserviceparamter.GetFederatoraiAgentGPU())
		if err := r.uninstallResource(resource); err != nil {
			r.Log.V(-1).Info(
				"Uninstall federatorai-agent-gpu resources failed, retry reconciling AlamedaService",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}

	//Uninstall PersistentVolumeClaim Source
	pvcResource := asp.GetUninstallPersistentVolumeClaimSource()
	if err := r.uninstallPersistentVolumeClaim(instance, pvcResource); err != nil {
		r.Log.V(-1).Info("retry reconciling AlamedaService", "AlamedaService.Namespace",
			instance.Namespace, "AlamedaService.Name", instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if !asp.SelfDriving {
		r.Log.Info("selfDriving has been changed to false")
		selfDrivingResource := alamedaserviceparamter.GetSelfDrivingRsource()
		if err := r.uninstallScalerforAlameda(instance, selfDrivingResource); err != nil {
			r.Log.V(-1).Info("retry reconciling AlamedaService", "AlamedaService.Namespace",
				instance.Namespace, "AlamedaService.Name", instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	} else { //install Alameda Scaler
		if err := r.createScalerforAlameda(instance, asp, installResource); err != nil {
			r.Log.V(-1).Info("create scaler for alameda failed, retry reconciling AlamedaService",
				"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
				instance.Name, "msg", err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}

	if err = r.updateAlamedaService(instance, req.NamespacedName, asp); err != nil {
		r.Log.Error(err, "Update AlamedaService failed, retry reconciling AlamedaService",
			"AlamedaService.Namespace", instance.Namespace, "AlamedaService.Name",
			instance.Name, "msg", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	r.Log.Info("Reconciling done.", "AlamedaService.Namespace", instance.Namespace,
		"AlamedaService.Name", instance.Name)
	id := fmt.Sprintf(`%s/%s`, instance.GetNamespace(), instance.GetName())
	r.FirstReconcileDoneAlamedaService[id] = struct{}{}

	return ctrl.Result{}, nil
}

func (r *AlamedaServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	prom_patch.InitK8SClient()

	util.Disable_operand_resource_protection =
		os.Getenv("DISABLE_OPERAND_RESOURCE_PROTECTION")
	if util.Disable_operand_resource_protection != "true" {
		return ctrl.NewControllerManagedBy(mgr).
			For(&federatoraiv1alpha1.AlamedaService{}).Watches(&source.Kind{
			Type: &appsv1.Deployment{},
		}, &handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &federatoraiv1alpha1.AlamedaService{},
		}).Watches(&source.Kind{
			Type: &corev1.ConfigMap{},
		}, &handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &federatoraiv1alpha1.AlamedaService{},
		}).Watches(&source.Kind{
			Type: &corev1.Service{},
		}, &handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &federatoraiv1alpha1.AlamedaService{},
		}).Complete(r)
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&federatoraiv1alpha1.AlamedaService{}).
		Complete(r)
}

func (r *AlamedaServiceReconciler) handleAlamedaServiceDeletion(
	req ctrl.Request) error {

	var err error

	id := fmt.Sprintf(`%s/%s`, req.Namespace, req.Name)
	delete(r.FirstReconcileDoneAlamedaService, id)

	// Before handling, check if the AlamedaService owns the lock
	lock, err := federatoraioperatorcontrollerutil.GetAlamedaServiceLock(
		context.TODO(), r.Client)
	if err != nil && !k8sErrors.IsNotFound(err) {
		return errors.Wrap(err, "get AlamedaService lock failed")
	} else if k8sErrors.IsNotFound(err) {
		return nil
	} else if !federatoraioperatorcontrollerutil.IsAlamedaServiceLockOwnedByAlamedaService(
		lock, federatoraiv1alpha1.AlamedaService{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: req.Namespace, Name: req.Name,
			},
		}) {
		return nil
	}

	// Deletion of AlamedaService lock must in the last step
	defer func() {
		if err := r.deleteAlamedaServiceLock(context.TODO()); err != nil {
			err = errors.Wrap(err, "delete AlamedaService lock failed")
		}
	}()

	gcSecret, err := util.GetGCClusterRole(context.TODO(), r.Client)
	if err != nil {
		err = errors.Wrap(err, "get gc secret failed")
		return err
	}
	if err := r.Client.Delete(context.TODO(), &gcSecret); err != nil {
		err = errors.Wrap(err, "delete gc secret failed")
		return err
	}

	return err
}

func (r *AlamedaServiceReconciler) InitAlamedaService(
	alamedaService *federatoraiv1alpha1.AlamedaService) {
	if alamedaService.Spec.EnableDispatcher == nil {
		enableTrue := true
		alamedaService.Spec.EnableDispatcher = &enableTrue
	}
}

func (r *AlamedaServiceReconciler) getNamespace(namespaceName string) (
	corev1.Namespace, error) {
	namespace := corev1.Namespace{}
	if err := r.Client.Get(context.TODO(), client.ObjectKey{
		Name: namespaceName,
	}, &namespace); err != nil {
		return namespace, errors.Errorf("get namespace %s failed: %s",
			namespaceName, err.Error())
	}
	return namespace, nil
}

func (r *AlamedaServiceReconciler) newComponentConfig(
	namespace corev1.Namespace,
	alamedaService federatoraiv1alpha1.AlamedaService,
	asp alamedaserviceparamter.AlamedaServiceParamter) (
	*component.ComponentConfig, error) {

	imageConfig := newDefautlImageConfig()
	imageConfig =
		setImageConfigWithAlamedaServiceParameterGlobalConfiguration(imageConfig, asp)
	imageConfig = setImageConfigWithEnv(imageConfig)
	imageConfig = setImageConfigWithAlamedaServiceParameter(imageConfig, asp)

	podTemplateConfig := component.NewDefaultPodTemplateConfig(namespace)

	prometheusConfig := component.PrometheusConfig{
		Address:         alamedaService.Spec.PrometheusService,
		BearerTokenFile: defaultPrometheusBearerTokenFile,
		TLS: component.TLSConfig{
			InsecureSkipVerify: true,
		},
	}
	prometheusURL, err := url.Parse(alamedaService.Spec.PrometheusService)
	if err != nil {
		return nil, errors.Wrap(err, "parse Prometheus url failed")
	} else {
		prometheusConfig.Host = prometheusURL.Hostname()
		if prometheusURL.Port() != "" {
			prometheusConfig.Port = prometheusURL.Port()
		} else {
			if prometheusURL.Scheme == "http" {
				prometheusConfig.Port = "80"
			} else if prometheusURL.Scheme == "https" {
				prometheusConfig.Port = "443"
			}
		}
		prometheusConfig.Protocol = prometheusURL.Scheme
	}

	enabled := false
	if len(asp.Kafka.BrokerAddresses) > 0 {
		enabled = true
	}
	kafka := component.KafkaConfig{
		Enabled:         enabled,
		BrokerAddresses: asp.Kafka.BrokerAddresses,
		Version:         defaultKafkaVersion,
		SASL: component.SASLConfig{
			Enabled: asp.Kafka.SASL.Enabled,
			BasicAuth: component.BasicAuth{
				Username: asp.Kafka.SASL.Username,
				Password: asp.Kafka.SASL.Password,
			},
		},
		TLS: component.TLSConfig{
			Enabled:            asp.Kafka.TLS.Enabled,
			InsecureSkipVerify: asp.Kafka.TLS.InsecureSkipVerify,
		},
	}

	execition := component.ExecutionConfig{}
	if asp.EnableVPA {
		execition.EnabledVPA = true
	}

	faiAgentGPU := component.FederatoraiAgentGPUConfig{}
	faiAgentGPUSectionSet := alamedaService.Spec.FederatoraiAgentGPUSectionSet
	faiAgentGPU.Enabled = asp.EnableGPU
	if faiAgentGPUSectionSet.Prometheus != nil {
		faiAgentGPU.Datasource.Prometheus.Address =
			faiAgentGPUSectionSet.Prometheus.Address
		faiAgentGPU.Datasource.Prometheus.BasicAuth.Username =
			faiAgentGPUSectionSet.Prometheus.Username
		faiAgentGPU.Datasource.Prometheus.BasicAuth.Password =
			faiAgentGPUSectionSet.Prometheus.Password
	}
	if faiAgentGPUSectionSet.InfluxDB != nil {
		faiAgentGPU.Datasource.InfluxDB.Address =
			faiAgentGPUSectionSet.InfluxDB.Address
		faiAgentGPU.Datasource.InfluxDB.BasicAuth.Username =
			faiAgentGPUSectionSet.InfluxDB.Username
		faiAgentGPU.Datasource.InfluxDB.BasicAuth.Password =
			faiAgentGPUSectionSet.InfluxDB.Password
	}

	aiDispatcher := component.AIDispatcherConfig{}
	aiDispatcher.Enabled = asp.EnableDispatcher

	nginxHPA := component.NginxConfig{}
	nginxHPA.Enabled = asp.Nginx.Enabled

	caConfig := component.ClusterAutoScalerConfig{}
	caConfig.EnableExecution = *asp.ClusterAutoScaler.EnableExecution

	clusterType := resourceapply.CheckClusterType(r.Apiextclient.ApiextensionsV1beta1())
	componentConfg := component.NewComponentConfig(podTemplateConfig, alamedaService,
		component.WithNamespace(namespace.Name),
		component.WithImageConfig(imageConfig),
		component.WithPodSecurityPolicyGroup(r.PodSecurityPolicesAPIGroupVersion.Group),
		component.WithPodSecurityPolicyVersion(r.PodSecurityPolicesAPIGroupVersion.Version),
		component.WithPrometheusConfig(prometheusConfig),
		component.WithKafkaConfig(kafka),
		component.WithFedermeterConfig(clusterType),
		component.WithExecutionConfig(execition),
		component.WithFederatoraiAgentGPUConfig(faiAgentGPU),
		component.WithAIDispatcherConfig(aiDispatcher),
		component.WithNginxConfig(nginxHPA),
		component.WithClusterAutoScalerConfig(caConfig),
		component.WithNodeSelector(asp.NodeSelector),
	)
	return componentConfg, nil
}

func (r *AlamedaServiceReconciler) createScalerforAlameda(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.AlamedaScalerList {
		resourceScaler := componentConfig.NewAlamedaScaler(fileString)
		if err := controllerutil.SetControllerReference(
			instance, resourceScaler, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceScaler SetControllerReference: %s", err.Error())
		}
		foundScaler := &autoscaling_v1alpha1.AlamedaScaler{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceScaler.Name, Namespace: resourceScaler.Namespace,
		}, foundScaler)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource Scaler... ",
				"resourceScaler.Name", resourceScaler.Name)
			err = r.Client.Create(context.TODO(), resourceScaler)
			if err != nil {
				return errors.Errorf("create Scaler %s/%s failed: %s",
					resourceScaler.Namespace, resourceScaler.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource Scaler",
				"resourceScaler.Name", resourceScaler.Name)
		} else if err != nil {
			return errors.Errorf("get Scaler %s/%s failed: %s",
				resourceScaler.Namespace, resourceScaler.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncCustomResourceDefinition(
	instance *federatoraiv1alpha1.AlamedaService,
	gcIns *rbacv1.ClusterRole, asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.CustomResourceDefinitionList {
		crd := componentConfig.NewCustomResourceDefinition(fileString)
		_, err := resourceapply.ApplyCustomResourceDefinition(
			r.Apiextclient.ApiextensionsV1beta1(), gcIns, r.Scheme, crd, asp)
		if err != nil {
			return errors.Wrapf(err,
				"syncCustomResourceDefinition faild: CustomResourceDefinition.Name: %s",
				crd.Name)
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallCustomResourceDefinition(
	resource *alamedaserviceparamter.Resource) {
	for _, fileString := range resource.CustomResourceDefinitionList {
		crd := componentConfig.NewCustomResourceDefinition(fileString)
		_, _, _ = resourceapply.DeleteCustomResourceDefinition(
			r.Apiextclient.ApiextensionsV1beta1(), crd)
	}
}

func (r *AlamedaServiceReconciler) syncClusterRoleBinding(
	instance *federatoraiv1alpha1.AlamedaService,
	gcIns *rbacv1.ClusterRole,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.ClusterRoleBindingList {
		resourceCRB := componentConfig.NewClusterRoleBinding(FileStr)
		//cluster-scoped resource must not have a namespace-scoped owner
		if err := controllerutil.SetControllerReference(
			gcIns, resourceCRB, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceCRB SetControllerReference: %s", err.Error())
		}

		foundCRB := &rbacv1.ClusterRoleBinding{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceCRB.Name,
		}, foundCRB)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource ClusterRoleBinding... ",
				"resourceCRB.Name", resourceCRB.Name)
			err = r.Client.Create(context.TODO(), resourceCRB)
			if err != nil {
				return errors.Errorf("create clusterRoleBinding %s/%s failed: %s",
					resourceCRB.Namespace, resourceCRB.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource ClusterRoleBinding",
				"resourceCRB.Name", resourceCRB.Name)
		} else if err != nil {
			return errors.Errorf("get clusterRoleBinding %s/%s failed: %s",
				resourceCRB.Namespace, resourceCRB.Name, err.Error())
		} else {
			err = r.Client.Update(context.TODO(), resourceCRB)
			if err != nil {
				return errors.Errorf("Update clusterRoleBinding %s/%s failed: %s",
					resourceCRB.Namespace, resourceCRB.Name, err.Error())
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) createAlamedaNotificationChannels(
	owner metav1.Object, resource *alamedaserviceparamter.Resource) error {
	for _, file := range resource.AlamedaNotificationChannelList {
		src, err := componentConfig.NewAlamedaNotificationChannel(file)
		if err != nil {
			return errors.Errorf(
				"get AlamedaNotificationChannel failed: file: %s, error: %s",
				file, err.Error())
		}
		if err := controllerutil.SetControllerReference(
			owner, src, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail AlamedaNotificationChannel SetControllerReference: %s",
				err.Error())
		}
		err = r.Client.Create(context.TODO(), src)
		if err != nil && !k8sErrors.IsAlreadyExists(err) {
			return errors.Errorf(
				"create AlamedaNotificationChannel %s failed: %s",
				src.GetName(), err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) createAlamedaNotificationTopics(
	owner metav1.Object, resource *alamedaserviceparamter.Resource) error {
	for _, file := range resource.AlamedaNotificationTopic {
		src, err := componentConfig.NewAlamedaNotificationTopic(file)
		if err != nil {
			return errors.Errorf(
				"get AlamedaNotificationTopic failed: file: %s, error: %s",
				file, err.Error())
		}
		if err := controllerutil.SetControllerReference(
			owner, src, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail AlamedaNotificationTopic SetControllerReference: %s",
				err.Error())
		}
		err = r.Client.Create(context.TODO(), src)
		if err != nil && !k8sErrors.IsAlreadyExists(err) {
			return errors.Errorf(
				"create AlamedaNotificationTopic %s failed: %s",
				src.GetName(), err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncPodSecurityPolicy(
	instance *federatoraiv1alpha1.AlamedaService,
	gcIns *rbacv1.ClusterRole,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {

	var psp runtime.Object
	switch r.PodSecurityPolicesAPIGroupVersion {
	case policyv1beta1.SchemeGroupVersion:
		psp = &policyv1beta1.PodSecurityPolicy{}
	case extensionsv1beta1.SchemeGroupVersion:
		psp = &extensionsv1beta1.PodSecurityPolicy{}
	default:
		return errors.Errorf(
			`not supported apiGroup "%s" for Kind:"PodSecurityPolicy"`,
			r.PodSecurityPolicesAPIGroupVersion)
	}

	for _, FileStr := range resource.PodSecurityPolicyList {
		var resourceMeta metav1.Object
		resourcePSP, err := componentConfig.NewPodSecurityPolicy(FileStr)
		switch v := resourcePSP.(type) {
		case *policyv1beta1.PodSecurityPolicy:
			resourceMeta = v
			if err := controllerutil.SetControllerReference(
				gcIns, v, r.Scheme); err != nil {
				return errors.Errorf(
					"Fail resourcePSP SetControllerReference: %s", err.Error())
			}
		case *extensionsv1beta1.PodSecurityPolicy:
			resourceMeta = v
			if err := controllerutil.SetControllerReference(
				gcIns, v, r.Scheme); err != nil {
				return errors.Errorf(
					"Fail resourcePSP SetControllerReference: %s", err.Error())
			}
		default:
			return errors.Errorf(
				`not supported type "%T" for Kind:"PodSecurityPolicy"`, resourcePSP)
		}
		err = r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceMeta.GetName(),
		}, psp)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource PodSecurityPolicy... ",
				"resourcePSP.Name", resourceMeta.GetName())
			err = r.Client.Create(context.TODO(), resourcePSP)
			if err != nil {
				return errors.Errorf("create PodSecurityPolicy %s/%s failed: %s",
					resourceMeta.GetNamespace(), resourceMeta.GetName(), err.Error())
			}
			r.Log.Info("Successfully Creating Resource PodSecurityPolicy",
				"resourcePSP.Name", resourceMeta.GetName())
		} else if err != nil {
			return errors.Errorf("get PodSecurityPolicy %s/%s failed: %s",
				resourceMeta.GetNamespace(), resourceMeta.GetName(), err.Error())
		} else {
			err = r.Client.Update(context.TODO(), resourcePSP)
			if err != nil {
				return errors.Errorf("Update PodSecurityPolicy %s/%s failed: %s",
					resourceMeta.GetNamespace(), resourceMeta.GetName(), err.Error())
			}
		}
	}

	return nil
}

func (r *AlamedaServiceReconciler) syncSecurityContextConstraints(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.SecurityContextConstraintsList {
		resourceSCC := componentConfig.NewSecurityContextConstraints(FileStr)
		if err := controllerutil.SetControllerReference(
			instance, resourceSCC, r.Scheme); err != nil {
			return errors.Errorf("Fail resourceSCC SetControllerReference: %s", err.Error())
		}
		//process resource SecurityContextConstraints according to AlamedaService CR
		resourceSCC =
			processcrdspec.ParamterToSecurityContextConstraints(resourceSCC, asp)
		foundSCC := &securityv1.SecurityContextConstraints{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceSCC.Name,
		}, foundSCC)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource SecurityContextConstraints... ",
				"resourceSCC.Name", resourceSCC.Name)
			err = r.Client.Create(context.TODO(), resourceSCC)
			if err != nil {
				return errors.Errorf("create SecurityContextConstraints %s/%s failed: %s",
					resourceSCC.Namespace, resourceSCC.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource SecurityContextConstraints",
				"resourceSCC.Name", resourceSCC.Name)
		} else if err != nil {
			return errors.Errorf("get SecurityContextConstraints %s/%s failed: %s",
				resourceSCC.Namespace, resourceSCC.Name, err.Error())
		} else {
			err = r.Client.Update(context.TODO(), resourceSCC)
			if err != nil {
				return errors.Errorf("Update SecurityContextConstraints %s/%s failed: %s",
					resourceSCC.Namespace, resourceSCC.Name, err.Error())
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncDaemonSet(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.DaemonSetList {
		resourceDS := componentConfig.NewDaemonSet(FileStr)
		if err := controllerutil.SetControllerReference(
			instance, resourceDS, r.Scheme); err != nil {
			return errors.Errorf("Fail resourceDS SetControllerReference: %s", err.Error())
		}
		//process resource DaemonSet according to AlamedaService CR
		resourceDS = processcrdspec.ParamterToDaemonSet(resourceDS, asp)
		if err := r.patchConfigMapResourceVersionIntoPodTemplateSpecLabel(
			resourceDS.Namespace, &resourceDS.Spec.Template); err != nil {
			return errors.Wrap(err,
				"patch resourceVersion of mounted configMaps into PodTemplateSpec failed")
		}
		foundDS := &appsv1.DaemonSet{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceDS.Name, Namespace: resourceDS.Namespace,
		}, foundDS)

		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource DaemonSet... ",
				"resourceDS.Namespace", resourceDS.Namespace,
				"resourceDS.Name", resourceDS.Name)
			err = r.Client.Create(context.TODO(), resourceDS)
			if err != nil {
				return errors.Errorf("create DaemonSet %s/%s failed: %s",
					resourceDS.Namespace, resourceDS.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource DaemonSet",
				"resourceDS.Namespace", resourceDS.Namespace,
				"resourceDS.Name", resourceDS.Name)
		} else if err != nil {
			return errors.Errorf("get DaemonSet %s/%s failed: %s",
				resourceDS.Namespace, resourceDS.Name, err.Error())
		} else {
			if updateresource.MisMatchResourceDaemonSet(foundDS, resourceDS) {
				r.Log.Info("Update Resource DaemonSet:",
					"foundDS.Name", foundDS.Name)
				err = r.Client.Delete(context.TODO(), foundDS)
				if err != nil {
					return errors.Errorf("delete DaemonSet %s/%s failed: %s",
						foundDS.Namespace, foundDS.Name, err.Error())
				}
				err = r.Client.Create(context.TODO(), resourceDS)
				if err != nil {
					return errors.Errorf("create DaemonSet %s/%s failed: %s",
						foundDS.Namespace, foundDS.Name, err.Error())
				}
				r.Log.Info("Successfully Update Resource DaemonSet",
					"resourceDS.Namespace", resourceDS.Namespace,
					"resourceDS.Name", resourceDS.Name)
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncClusterRole(
	instance *federatoraiv1alpha1.AlamedaService,
	gcIns *rbacv1.ClusterRole, asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.ClusterRoleList {
		resourceCR := componentConfig.NewClusterRole(FileStr)
		//cluster-scoped resource must not have a namespace-scoped owner
		if err := controllerutil.SetControllerReference(
			gcIns, resourceCR, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceCR SetControllerReference: %s", err.Error())
		}

		foundCR := &rbacv1.ClusterRole{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceCR.Name,
		}, foundCR)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource ClusterRole... ",
				"resourceCR.Name", resourceCR.Name)
			err = r.Client.Create(context.TODO(), resourceCR)
			if err != nil {
				return errors.Errorf("create clusterRole %s/%s failed: %s",
					resourceCR.Namespace, resourceCR.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource ClusterRole",
				"resourceCR.Name", resourceCR.Name)
		} else if err != nil {
			return errors.Errorf("get clusterRole %s/%s failed: %s",
				resourceCR.Namespace, resourceCR.Name, err.Error())
		} else {
			err = r.Client.Update(context.TODO(), resourceCR)
			if err != nil {
				return errors.Errorf("Update clusterRole %s/%s failed: %s",
					resourceCR.Namespace, resourceCR.Name, err.Error())
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) createServiceAccount(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.ServiceAccountList {
		resourceSA := componentConfig.NewServiceAccount(FileStr)
		if err := controllerutil.SetControllerReference(
			instance, resourceSA, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceSA SetControllerReference: %s", err.Error())
		}
		foundSA := &corev1.ServiceAccount{}

		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceSA.Name, Namespace: resourceSA.Namespace,
		}, foundSA)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource ServiceAccount... ",
				"resourceSA.Name", resourceSA.Name)
			err = r.Client.Create(context.TODO(), resourceSA)
			if err != nil {
				return errors.Errorf("create serviceAccount %s/%s failed: %s",
					resourceSA.Namespace, resourceSA.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource ServiceAccount",
				"resourceSA.Name", resourceSA.Name)
		} else if err != nil {
			return errors.Errorf("get serviceAccount %s/%s failed: %s",
				resourceSA.Namespace, resourceSA.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncRole(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.RoleList {
		resourceCR := componentConfig.NewRole(FileStr)
		if err := controllerutil.SetControllerReference(
			instance, resourceCR, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceCR SetControllerReference: %s", err.Error())
		}
		foundCR := &rbacv1.Role{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Namespace: resourceCR.Namespace, Name: resourceCR.Name,
		}, foundCR)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource Role... ", "resourceCR.Namespace",
				resourceCR.Namespace, "resourceCR.Name", resourceCR.Name)
			err = r.Client.Create(context.TODO(), resourceCR)
			if err != nil {
				return errors.Errorf("create Role %s/%s failed: %s",
					resourceCR.Namespace, resourceCR.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource Role",
				"resourceCR.Namespace", resourceCR.Namespace,
				"resourceCR.Name", resourceCR.Name)
		} else if err != nil {
			return errors.Errorf("get Role %s/%s failed: %s",
				resourceCR.Namespace, resourceCR.Name, err.Error())
		} else {
			err = r.Client.Update(context.TODO(), resourceCR)
			if err != nil {
				return errors.Errorf("Update Role %s/%s failed: %s",
					resourceCR.Namespace, resourceCR.Name, err.Error())
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncRoleBinding(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.RoleBindingList {
		resourceCR := componentConfig.NewRoleBinding(FileStr)
		if err := controllerutil.SetControllerReference(
			instance, resourceCR, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceCR SetControllerReference: %s", err.Error())
		}
		foundCR := &rbacv1.RoleBinding{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Namespace: resourceCR.Namespace, Name: resourceCR.Name,
		}, foundCR)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource RoleBinding... ",
				"resourceCR.Namespace", resourceCR.Namespace,
				"resourceCR.Name", resourceCR.Name)
			err = r.Client.Create(context.TODO(), resourceCR)
			if err != nil {
				return errors.Errorf("create RoleBinding %s/%s failed: %s",
					resourceCR.Namespace, resourceCR.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource RoleBinding",
				"resourceCR.Namespace", resourceCR.Namespace,
				"resourceCR.Name", resourceCR.Name)
		} else if err != nil {
			return errors.Errorf("get RoleBinding %s/%s failed: %s",
				resourceCR.Namespace, resourceCR.Name, err.Error())
		} else {
			err = r.Client.Update(context.TODO(), resourceCR)
			if err != nil {
				return errors.Errorf("Update RoleBinding %s/%s failed: %s",
					resourceCR.Namespace, resourceCR.Name, err.Error())
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) createPersistentVolumeClaim(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.PersistentVolumeClaimList {
		resourcePVC := componentConfig.NewPersistentVolumeClaim(FileStr)
		//process resource configmap into desire configmap
		resourcePVC =
			processcrdspec.ParamterToPersistentVolumeClaim(resourcePVC, asp)
		if err := controllerutil.SetControllerReference(
			instance, resourcePVC, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourcePVC SetControllerReference: %s", err.Error())
		}
		foundPVC := &corev1.PersistentVolumeClaim{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourcePVC.Name, Namespace: resourcePVC.Namespace,
		}, foundPVC)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource PersistentVolumeClaim... ",
				"resourcePVC.Name", resourcePVC.Name)
			err = r.Client.Create(context.TODO(), resourcePVC)
			if err != nil {
				return errors.Errorf("create PersistentVolumeClaim %s/%s failed: %s",
					resourcePVC.Namespace, resourcePVC.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource PersistentVolumeClaim",
				"resourcePVC.Name", resourcePVC.Name)
		} else if err != nil {
			return errors.Errorf("get PersistentVolumeClaim %s/%s failed: %s",
				resourcePVC.Namespace, resourcePVC.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) createSecret(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {

	secret, err := componentConfig.NewAdmissionControllerSecret()
	if err != nil {
		return errors.Errorf("build AdmissionController secret failed: %s", err.Error())
	}
	if err := controllerutil.SetControllerReference(
		instance, secret, r.Scheme); err != nil {
		return errors.Errorf("set controller reference to secret %s/%s failed: %s",
			secret.Namespace, secret.Name, err.Error())
	}
	err = r.Client.Create(context.TODO(), secret)
	if err != nil && k8sErrors.IsAlreadyExists(err) {
	} else if err != nil {
		return errors.Errorf("create secret %s/%s failed: %s",
			secret.Namespace, secret.Name, err.Error())
	}
	secret, err = componentConfig.NewInfluxDBSecret()
	if err != nil {
		return errors.Errorf("build InfluxDB secret failed: %s", err.Error())
	}
	if err := controllerutil.SetControllerReference(
		instance, secret, r.Scheme); err != nil {
		return errors.Errorf("set controller reference to secret %s/%s failed: %s",
			secret.Namespace, secret.Name, err.Error())
	}
	err = r.Client.Create(context.TODO(), secret)
	if err != nil && k8sErrors.IsAlreadyExists(err) {
	} else if err != nil {
		return errors.Errorf("create secret %s/%s failed: %s",
			secret.Namespace, secret.Name, err.Error())
	}
	secret, err = componentConfig.NewfedemeterSecret()
	if err != nil {
		return errors.Errorf("build Fedemeter secret failed: %s", err.Error())
	}
	if err := controllerutil.SetControllerReference(
		instance, secret, r.Scheme); err != nil {
		return errors.Errorf("set controller reference to secret %s/%s failed: %s",
			secret.Namespace, secret.Name, err.Error())
	}
	err = r.Client.Create(context.TODO(), secret)
	if err != nil && k8sErrors.IsAlreadyExists(err) {
	} else if err != nil {
		return errors.Errorf("create secret %s/%s failed: %s",
			secret.Namespace, secret.Name, err.Error())
	}

	notifierWebhookServiceAsset :=
		alamedaserviceparamter.GetAlamedaNotifierWebhookService()
	notifierWebhookService := componentConfig.NewService(notifierWebhookServiceAsset)
	notifierWebhookServiceAddress := util.GetServiceDNS(notifierWebhookService)
	notifierWebhookServiceCertSecretAsset :=
		alamedaserviceparamter.GetAlamedaNotifierWebhookServerCertSecret()
	notifierWebhookServiceSecret, err := componentConfig.NewTLSSecret(
		notifierWebhookServiceCertSecretAsset, notifierWebhookServiceAddress)
	if err != nil {
		return errors.Errorf("build secret failed: %s", err.Error())
	}
	if err := controllerutil.SetControllerReference(
		instance, notifierWebhookServiceSecret, r.Scheme); err != nil {
		return errors.Errorf("set controller reference to secret %s/%s failed: %s",
			notifierWebhookServiceSecret.Namespace,
			notifierWebhookServiceSecret.Name, err.Error())
	}
	err = r.Client.Create(context.TODO(), notifierWebhookServiceSecret)
	if err != nil && k8sErrors.IsAlreadyExists(err) {
	} else if err != nil {
		return errors.Errorf("get secret %s/%s failed: %s",
			notifierWebhookServiceSecret.Namespace,
			notifierWebhookServiceSecret.Name, err.Error())
	}

	operatorWebhookServiceAsset :=
		alamedaserviceparamter.GetAlamedaOperatorWebhookService()
	operatorWebhookService :=
		componentConfig.NewService(operatorWebhookServiceAsset)
	operatorWebhookServiceAddress := util.GetServiceDNS(operatorWebhookService)
	operatorWebhookServiceCertSecretAsset :=
		alamedaserviceparamter.GetAlamedaOperatorWebhookServerCertSecret()
	operatorWebhookServiceSecret, err :=
		componentConfig.NewTLSSecret(operatorWebhookServiceCertSecretAsset,
			operatorWebhookServiceAddress)
	if err != nil {
		return errors.Errorf("build secret failed: %s", err.Error())
	}
	if err := controllerutil.SetControllerReference(
		instance, operatorWebhookServiceSecret, r.Scheme); err != nil {
		return errors.Errorf("set controller reference to secret %s/%s failed: %s",
			operatorWebhookServiceSecret.Namespace,
			operatorWebhookServiceSecret.Name, err.Error())
	}
	err = r.Client.Create(context.TODO(), operatorWebhookServiceSecret)
	if err != nil && k8sErrors.IsAlreadyExists(err) {
	} else if err != nil {
		return errors.Errorf("get secret %s/%s failed: %s",
			operatorWebhookServiceSecret.Namespace,
			operatorWebhookServiceSecret.Name, err.Error())
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncConfigMap(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.ConfigMapList {
		resourceCM := componentConfig.NewConfigMap(fileString)
		if err := controllerutil.SetControllerReference(
			instance, resourceCM, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceCM SetControllerReference: %s", err.Error())
		}
		//process resource configmap into desire configmap
		foundCM := &corev1.ConfigMap{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceCM.Name, Namespace: resourceCM.Namespace,
		}, foundCM)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource ConfigMap... ",
				"resourceCM.Name", resourceCM.Name)
			err = r.Client.Create(context.TODO(), resourceCM)
			if err != nil {
				return errors.Errorf("create configMap %s/%s failed: %s",
					resourceCM.Namespace, resourceCM.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource ConfigMap",
				"resourceCM.Name", resourceCM.Name)
		} else if err != nil {
			return errors.Errorf("get configMap %s/%s failed: %s",
				resourceCM.Namespace, resourceCM.Name, err.Error())
		} else {
			if updateresource.MisMatchResourceConfigMap(foundCM, resourceCM) {
				r.Log.Info("Update Resource Service:", "foundCM.Name", foundCM.Name)
				err = r.Client.Update(context.TODO(), foundCM)
				if err != nil {
					return errors.Errorf("update configMap %s/%s failed: %s",
						foundCM.Namespace, foundCM.Name, err.Error())
				}
				r.Log.Info("Successfully Update Resource ConfigMap",
					"resourceCM.Name", foundCM.Name)
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncService(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.ServiceList {
		resourceSV := componentConfig.NewService(fileString)
		if err := controllerutil.SetControllerReference(
			instance, resourceSV, r.Scheme); err != nil {
			return errors.Errorf("Fail resourceSV SetControllerReference: %s", err.Error())
		}
		foundSV := &corev1.Service{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceSV.Name, Namespace: resourceSV.Namespace,
		}, foundSV)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource Service... ",
				"resourceSV.Name", resourceSV.Name)
			err = r.Client.Create(context.TODO(), resourceSV)
			if err != nil {
				return errors.Errorf("create service %s/%s failed: %s",
					resourceSV.Namespace, resourceSV.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource Service", "resourceSV.Name",
				resourceSV.Name)
		} else if err != nil {
			return errors.Errorf("get service %s/%s failed: %s",
				resourceSV.Namespace, resourceSV.Name, err.Error())
		} else {
			if updateresource.MisMatchResourceService(foundSV, resourceSV) {
				r.Log.Info("Update Resource Service:", "foundSV.Name", foundSV.Name)
				err = r.Client.Delete(context.TODO(), foundSV)
				if err != nil {
					return errors.Errorf("delete service %s/%s failed: %s",
						foundSV.Namespace, foundSV.Name, err.Error())
				}
				err = r.Client.Create(context.TODO(), resourceSV)
				if err != nil {
					return errors.Errorf("create service %s/%s failed: %s",
						foundSV.Namespace, foundSV.Name, err.Error())
				}
				r.Log.Info("Successfully Update Resource Service",
					"resourceSV.Name", foundSV.Name)
			}
		}
	}
	return nil
}

// syncServiceExposure synchornize AlamedaService.Spec.ServiceExposures with current services in type NodePort.
func (r *AlamedaServiceReconciler) syncServiceExposure(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {

	// prepare services need to be created by service exposures
	serviceExposureMap := make(map[string]federatoraiv1alpha1.ServiceExposureSpec)
	for _, serviceExposure := range instance.Spec.ServiceExposures {
		serviceExposureMap[serviceExposure.Name] = serviceExposure
	}
	servicesNeedToBeCreatedMap := make(map[string]corev1.Service)
	for _, fileString := range resource.ServiceList {
		resource := componentConfig.NewService(fileString)
		serviceExposure, exist := serviceExposureMap[resource.Name]
		if !exist {
			continue
		}
		want, err := r.newServiceByServiceExposure(*resource, serviceExposure)
		if err != nil {
			return errors.Wrap(err, "apply service exposure to service failed")
		}
		servicesNeedToBeCreatedMap[fmt.Sprintf("%s/%s", want.Namespace, want.Name)] = want
	}

	// create or update services
	for _, service := range servicesNeedToBeCreatedMap {

		if err := controllerutil.SetControllerReference(
			instance, &service, r.Scheme); err != nil {
			return errors.Wrapf(err, "set controller reference to Service(%s/%s) failed",
				service.Namespace, service.Name)
		}

		found := corev1.Service{}
		err := r.Client.Get(context.TODO(), client.ObjectKey{
			Namespace: service.Namespace, Name: service.Name,
		}, &found)
		if err != nil && !k8sErrors.IsNotFound(err) {
			return errors.Wrapf(err, "get Service(%s/%s) failed",
				service.Namespace, service.Name)
		} else if k8sErrors.IsNotFound(err) {
			if err := r.Client.Create(context.TODO(), &service); err != nil {
				return errors.Wrapf(err, "create Service(%s/%s) failed",
					service.Namespace, service.Name)
			}
		} else {
			// update
			service.ResourceVersion = found.ResourceVersion
			service.Spec.ClusterIP = found.Spec.ClusterIP
			if err := r.Client.Update(context.TODO(), &service); err != nil {
				return errors.Wrapf(err, "update Service(%s/%s) failed",
					service.Namespace, service.Name)
			}
		}
	}

	// deletes services created by service exposures but not in current exist list
	serviceList := corev1.ServiceList{}
	listOpt := client.ListOptions{}
	client.MatchingLabels{serviceExposureAnnotationKey: ""}.ApplyToList(&listOpt)
	err := r.Client.List(context.TODO(), &serviceList, &listOpt)
	if err != nil {
		return errors.Wrap(err, "list service failed")
	}
	for _, service := range serviceList.Items {
		if _, exist := servicesNeedToBeCreatedMap[fmt.Sprintf(
			"%s/%s", service.Namespace, service.Name)]; !exist {
			if err := r.Client.Delete(context.TODO(), &service); err != nil {
				return errors.Wrapf(err, "delete Service(%s/%s) failed",
					service.Namespace, service.Name)
			}
		}
	}

	return nil
}

func (r *AlamedaServiceReconciler) newServiceByServiceExposure(svc corev1.Service,
	svcExposure federatoraiv1alpha1.ServiceExposureSpec) (corev1.Service, error) {

	if svc.Name != svcExposure.Name {
		return corev1.Service{},
			errors.New("service name must be equal to service exposure name")
	}

	// add service exposure label to service
	labels := svc.GetLabels()
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[serviceExposureAnnotationKey] = ""
	svc.SetLabels(labels)

	switch svcExposure.Type {
	case federatoraiv1alpha1.ServiceExposureTypeNodePort:
		return r.newNodePortServiceByServiceExposure(svc, svcExposure), nil
	default:
		return corev1.Service{},
			errors.Errorf(`not supported service exposure type(%s)`,
				svcExposure.Type)
	}

}

func (r *AlamedaServiceReconciler) newNodePortServiceByServiceExposure(
	svc corev1.Service,
	svcExposure federatoraiv1alpha1.ServiceExposureSpec) corev1.Service {

	if svcExposure.NodePort == nil {
		return svc
	}

	svc.Name = fmt.Sprintf(`%s-node-port`, svc.Name)
	svc.Spec.Type = corev1.ServiceTypeNodePort

	portMap := make(map[int32]federatoraiv1alpha1.PortSpec)
	for _, port := range svcExposure.NodePort.Ports {
		portMap[port.Port] = port
	}
	newPorts := make([]corev1.ServicePort, 0, len(portMap))
	for _, port := range svc.Spec.Ports {
		portSpec, exist := portMap[port.Port]
		if !exist {
			continue
		}
		port.NodePort = portSpec.NodePort
		newPorts = append(newPorts, port)
		delete(portMap, port.Port)
	}
	for _, portSpec := range portMap {
		newPorts = append(newPorts, corev1.ServicePort{
			Port:     portSpec.Port,
			NodePort: portSpec.NodePort,
			Name:     fmt.Sprintf("port-%d", portSpec.Port),
		})
	}
	svc.Spec.Ports = newPorts
	return svc
}

func (r *AlamedaServiceReconciler) getSecret(namespace, name string) (corev1.Secret, error) {

	secret := corev1.Secret{}
	err := r.Client.Get(context.TODO(), client.ObjectKey{Namespace: namespace, Name: name}, &secret)
	if err != nil {
		return secret, errors.Errorf("get secret (%s/%s) failed", namespace, name)
	}

	return secret, nil
}

func (r *AlamedaServiceReconciler) syncMutatingWebhookConfiguration(
	instance *federatoraiv1alpha1.AlamedaService,
	gcIns *rbacv1.ClusterRole, asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.MutatingWebhookConfigurationList {
		mutatingWebhookConfiguration, err :=
			componentConfig.NewMutatingWebhookConfiguration(fileString)
		if err != nil {
			return errors.Wrap(err, "new MutatingWebhookConfiguration failed")
		}

		//cluster-scoped resource must not have a namespace-scoped owner
		if err := controllerutil.SetControllerReference(
			gcIns, mutatingWebhookConfiguration, r.Scheme); err != nil {
			return errors.Errorf("Fail MutatingWebhookConfiguration SetControllerReference: %s",
				err.Error())
		}

		secretName, exist :=
			mutatingWebhookConfiguration.ObjectMeta.Annotations[admissionWebhookAnnotationKeySecretName]
		if !exist {
			return errors.Errorf(`annotation key("%s") is empty`,
				admissionWebhookAnnotationKeySecretName)
		}

		secret, err := r.getSecret(instance.Namespace, secretName)
		if err != nil {
			return errors.Errorf("get secret failed: %s", err.Error())
		}
		caCert := secret.Data["ca.crt"]
		for i := range mutatingWebhookConfiguration.Webhooks {
			mutatingWebhookConfiguration.Webhooks[i].ClientConfig.CABundle = caCert
		}

		instance := admissionregistrationv1beta1.MutatingWebhookConfiguration{}
		err = r.Client.Get(context.TODO(), types.NamespacedName{
			Name: mutatingWebhookConfiguration.Name,
		}, &instance)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating MutatingWebhookConfiguration...", "name",
				mutatingWebhookConfiguration.Name)
			err = r.Client.Create(context.TODO(), mutatingWebhookConfiguration)
			if err != nil && !k8sErrors.IsAlreadyExists(err) {
				return errors.Wrapf(err, `create MutatingWebhookConfiguration("%s") failed`,
					mutatingWebhookConfiguration.Name)
			}
		} else if err != nil {
			return errors.Wrapf(err, `get MutatingWebhookConfiguration("%s") failed`,
				mutatingWebhookConfiguration.Name)
		} else {
			copyInstance := admissionregistrationv1beta1.MutatingWebhookConfiguration{}
			instance.DeepCopyInto(&copyInstance)

			copyInstance.Webhooks = mutatingWebhookConfiguration.Webhooks
			r.Log.Info("Updating MutatingWebhookConfiguration", "name",
				mutatingWebhookConfiguration.Name)
			err = r.Client.Update(context.TODO(), &copyInstance)
			if err != nil {
				return errors.Wrapf(err, `update MutatingWebhookConfiguration("%s")`,
					mutatingWebhookConfiguration.Name)
			}
			r.Log.Info("Updating MutatingWebhookConfiguration done.", "name",
				mutatingWebhookConfiguration.Name)
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncValidatingWebhookConfiguration(
	instance *federatoraiv1alpha1.AlamedaService,
	gcIns *rbacv1.ClusterRole, asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.ValidatingWebhookConfigurationList {
		validatingWebhookConfiguration, err :=
			componentConfig.NewValidatingWebhookConfiguration(fileString)
		if err != nil {
			return errors.Wrap(err, "new ValidatingWebhookConfigurationList failed")
		}
		//cluster-scoped resource must not have a namespace-scoped owner
		if err := controllerutil.SetControllerReference(
			gcIns, validatingWebhookConfiguration, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail ValidatingWebhookConfiguration SetControllerReference: %s", err.Error())
		}

		secretName, exist :=
			validatingWebhookConfiguration.ObjectMeta.Annotations[admissionWebhookAnnotationKeySecretName]
		if !exist {
			return errors.Errorf(`annotation key("%s") is empty`,
				admissionWebhookAnnotationKeySecretName)
		}

		secret, err := r.getSecret(instance.Namespace, secretName)
		if err != nil {
			return errors.Errorf("get secret failed: %s", err.Error())
		}
		caCert := secret.Data["ca.crt"]
		for i := range validatingWebhookConfiguration.Webhooks {
			validatingWebhookConfiguration.Webhooks[i].ClientConfig.CABundle = caCert
		}

		instance := admissionregistrationv1beta1.ValidatingWebhookConfiguration{}
		err = r.Client.Get(context.TODO(), types.NamespacedName{
			Name: validatingWebhookConfiguration.Name,
		}, &instance)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating ValidatingWebhookConfiguration...", "name",
				validatingWebhookConfiguration.Name)
			err = r.Client.Create(context.TODO(), validatingWebhookConfiguration)
			if err != nil && !k8sErrors.IsAlreadyExists(err) {
				return errors.Wrapf(err,
					`create ValidatingWebhookConfiguration("%s") failed`,
					validatingWebhookConfiguration.Name)
			}
		} else if err != nil {
			return errors.Wrapf(err, `get ValidatingWebhookConfiguration("%s") failed`,
				validatingWebhookConfiguration.Name)
		} else {
			copyInstance := admissionregistrationv1beta1.ValidatingWebhookConfiguration{}
			instance.DeepCopyInto(&copyInstance)

			copyInstance.Webhooks = validatingWebhookConfiguration.Webhooks
			r.Log.Info("Updating ValidatingWebhookConfiguration", "name",
				validatingWebhookConfiguration.Name)
			err = r.Client.Update(context.TODO(), &copyInstance)
			if err != nil {
				return errors.Wrapf(err, `update ValidatingWebhookConfiguration("%s")`,
					validatingWebhookConfiguration.Name)
			}
			r.Log.Info("Updating ValidatingWebhookConfiguration done.", "name",
				validatingWebhookConfiguration.Name)
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncDeployment(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.DeploymentList {
		resourceDep := componentConfig.NewDeployment(fileString)
		if err := controllerutil.SetControllerReference(
			instance, resourceDep, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceDep SetControllerReference: %s", err.Error())
		}
		//process resource deployment into desire deployment
		resourceDep = processcrdspec.ParamterToDeployment(resourceDep, asp)
		if err := r.patchConfigMapResourceVersionIntoPodTemplateSpecLabel(
			resourceDep.Namespace, &resourceDep.Spec.Template); err != nil {
			return errors.Wrap(err,
				"patch resourceVersion of mounted configMaps into PodTemplateSpec failed")
		}

		foundDep := &appsv1.Deployment{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceDep.Name, Namespace: resourceDep.Namespace,
		}, foundDep)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource Deployment... ",
				"resourceDep.Name", resourceDep.Name)
			err = r.Client.Create(context.TODO(), resourceDep)
			if err != nil {
				return errors.Errorf("create deployment %s/%s failed: %s",
					resourceDep.Namespace, resourceDep.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource Deployment",
				"resourceDep.Name", resourceDep.Name)
			continue
		} else if err != nil {
			return errors.Errorf("get deployment %s/%s failed: %s",
				resourceDep.Namespace, resourceDep.Name, err.Error())
		} else {
			if updateresource.MisMatchResourceDeployment(
				foundDep, resourceDep) {
				r.Log.Info("Update Resource Deployment:",
					"resourceDep.Name", foundDep.Name)
				err = r.Client.Update(context.TODO(), foundDep)
				if err != nil {
					return errors.Errorf("update deployment %s/%s failed: %s",
						foundDep.Namespace, foundDep.Name, err.Error())
				}
				r.Log.Info("Successfully Update Resource Deployment",
					"resourceDep.Name", foundDep.Name)
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncRoute(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.RouteList {
		resourceRT := componentConfig.NewRoute(FileStr)
		if err := controllerutil.SetControllerReference(
			instance, resourceRT, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceRT SetControllerReference: %s", err.Error())
		}
		foundRT := &routev1.Route{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceRT.Name, Namespace: resourceRT.Namespace,
		}, foundRT)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource Route... ",
				"resourceRT.Name", resourceRT.Name)
			err = r.Client.Create(context.TODO(), resourceRT)
			if err != nil {
				return errors.Errorf("create route %s/%s failed: %s",
					resourceRT.Namespace, resourceRT.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource route",
				"resourceRT.Name", resourceRT.Name)
		} else if err != nil {
			return errors.Errorf("get route %s/%s failed: %s",
				resourceRT.Namespace, resourceRT.Name, err.Error())
		} else {
			if foundRT.Spec.TLS == nil {
				foundRT.Spec.TLS = &routev1.TLSConfig{}
			}
			foundRT.Spec.TLS.InsecureEdgeTerminationPolicy =
				resourceRT.Spec.TLS.InsecureEdgeTerminationPolicy
			foundRT.Spec.TLS.Termination = resourceRT.Spec.TLS.Termination
			foundRT.Spec.Port = resourceRT.Spec.Port
			if err := r.Client.Update(context.TODO(), foundRT); err != nil {
				return errors.Wrapf(err, "update Route(%s/%s) failed",
					foundRT.Namespace, foundRT.Name)
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) syncStatefulSet(
	instance *federatoraiv1alpha1.AlamedaService,
	asp *alamedaserviceparamter.AlamedaServiceParamter,
	resource *alamedaserviceparamter.Resource) error {
	for _, FileStr := range resource.StatefulSetList {
		resourceSS := componentConfig.NewStatefulSet(FileStr)
		if err := controllerutil.SetControllerReference(
			instance, resourceSS, r.Scheme); err != nil {
			return errors.Errorf(
				"Fail resourceSS SetControllerReference: %s", err.Error())
		}
		resourceSS = processcrdspec.ParamterToStatefulset(resourceSS, asp)
		if err := r.patchConfigMapResourceVersionIntoPodTemplateSpecLabel(
			resourceSS.Namespace, &resourceSS.Spec.Template); err != nil {
			return errors.Wrap(err,
				"patch resourceVersion of mounted configMaps into PodTemplateSpec failed")
		}
		foundSS := &appsv1.StatefulSet{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceSS.Name, Namespace: resourceSS.Namespace,
		}, foundSS)
		if err != nil && k8sErrors.IsNotFound(err) {
			r.Log.Info("Creating a new Resource Route... ",
				"resourceSS.Name", resourceSS.Name)
			err = r.Client.Create(context.TODO(), resourceSS)
			if err != nil {
				return errors.Errorf("create route %s/%s failed: %s",
					resourceSS.Namespace, resourceSS.Name, err.Error())
			}
			r.Log.Info("Successfully Creating Resource route",
				"resourceSS.Name", resourceSS.Name)
		} else if err != nil {
			return errors.Errorf("get route %s/%s failed: %s",
				resourceSS.Namespace, resourceSS.Name, err.Error())
		} else {
			r.Log.Info("Update Resource StatefulSet:",
				"resourceSS.Name", resourceSS.Name)
			if updateresource.MisMatchResourceStatefulSet(foundSS, resourceSS) {
				r.Log.Info("Update Resource StatefulSet:", "name", foundSS.Name)
				err = r.Client.Update(context.TODO(), foundSS)
				if err != nil {
					return errors.Errorf("update statefulSet %s/%s failed: %s",
						foundSS.Namespace, foundSS.Name, err.Error())
				}
				r.Log.Info("Successfully Update Resource StatefulSet",
					"name", foundSS.Name)
			}
			r.Log.Info("Updating Resource StatefulSet", "name", resourceSS.Name)
			err = r.Client.Update(context.TODO(), resourceSS)
			if err != nil {
				return errors.Errorf("update StatefulSet %s/%s failed: %s",
					resourceSS.Namespace, resourceSS.Name, err.Error())
			}
			r.Log.Info("Successfully Update Resource StatefulSet",
				"resourceSS.Name", resourceSS.Name)
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallStatefulSet(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.StatefulSetList {
		resourceSS := componentConfig.NewStatefulSet(fileString)
		err := r.Client.Delete(context.TODO(), resourceSS)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete statefulset %s/%s failed: %s",
				resourceSS.Namespace, resourceSS.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallRoute(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.RouteList {
		resourceRT := componentConfig.NewRoute(fileString)
		err := r.Client.Delete(context.TODO(), resourceRT)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete route %s/%s failed: %s",
				resourceRT.Namespace, resourceRT.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallDeployment(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.DeploymentList {
		resourceDep := componentConfig.NewDeployment(fileString)
		err := r.Client.Delete(context.TODO(), resourceDep)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete deployment %s/%s failed: %s",
				resourceDep.Namespace, resourceDep.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallService(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.ServiceList {
		resourceSVC := componentConfig.NewService(fileString)
		err := r.Client.Delete(context.TODO(), resourceSVC)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete service %s/%s failed: %s",
				resourceSVC.Namespace, resourceSVC.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallConfigMap(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.ConfigMapList {
		resourceCM := componentConfig.NewConfigMap(fileString)
		err := r.Client.Delete(context.TODO(), resourceCM)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete comfigMap %s/%s failed: %s",
				resourceCM.Namespace, resourceCM.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallSecret(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.SecretList {
		resourceSec, _ := componentConfig.NewSecret(fileString)
		err := r.Client.Delete(context.TODO(), resourceSec)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete secret %s/%s failed: %s",
				resourceSec.Namespace, resourceSec.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallServiceAccount(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.ServiceAccountList {
		resourceSA := componentConfig.NewServiceAccount(fileString)
		err := r.Client.Delete(context.TODO(), resourceSA)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete serviceAccount %s/%s failed: %s",
				resourceSA.Namespace, resourceSA.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallClusterRole(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.ClusterRoleList {
		resourceCR := componentConfig.NewClusterRole(fileString)
		err := r.Client.Delete(context.TODO(), resourceCR)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete clusterRole %s/%s failed: %s",
				resourceCR.Namespace, resourceCR.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallClusterRoleBinding(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.ClusterRoleBindingList {
		resourceCRB := componentConfig.NewClusterRoleBinding(fileString)
		err := r.Client.Delete(context.TODO(), resourceCRB)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete clusterRoleBinding %s/%s failed: %s",
				resourceCRB.Namespace, resourceCRB.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallAlamedaScaler(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.AlamedaScalerList {
		resourceScaler := componentConfig.NewAlamedaScaler(fileString)
		err := r.Client.Delete(context.TODO(), resourceScaler)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete resourceScaler %s/%s failed: %s",
				resourceScaler.Namespace, resourceScaler.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallMutatingWebhookConfiguration(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.MutatingWebhookConfigurationList {
		mutatingWebhookConfiguration, err :=
			componentConfig.NewMutatingWebhookConfiguration(fileString)
		if err != nil {
			return errors.Wrap(err, "new MutatingWebhookConfiguration failed")
		}
		err = r.Client.Delete(context.TODO(), mutatingWebhookConfiguration)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete MutatingWebhookConfiguratio %s failed: %s",
				mutatingWebhookConfiguration.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallValidatingWebhookConfiguration(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.AlamedaScalerList {
		validatingWebhookConfiguration, err :=
			componentConfig.NewValidatingWebhookConfiguration(fileString)
		if err != nil {
			return errors.Wrap(err, "new ValidatingWebhookConfiguration failed")
		}
		err = r.Client.Delete(context.TODO(), validatingWebhookConfiguration)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			return errors.Errorf("delete ValidatingWebhookConfiguration %s failed: %s",
				validatingWebhookConfiguration.Name, err.Error())
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallScalerforAlameda(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	if err := r.uninstallAlamedaScaler(instance, resource); err != nil {
		return errors.Wrapf(err, "uninstall selfDriving scaler failed")
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallPersistentVolumeClaim(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.PersistentVolumeClaimList {
		resourcePVC := componentConfig.NewPersistentVolumeClaim(fileString)
		foundPVC := &corev1.PersistentVolumeClaim{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourcePVC.Name, Namespace: resourcePVC.Namespace,
		}, foundPVC)
		if err != nil && k8sErrors.IsNotFound(err) {
			continue
		} else if err != nil {
			return errors.Errorf("get PersistentVolumeClaim %s/%s failed: %s",
				resourcePVC.Namespace, resourcePVC.Name, err.Error())
		} else {
			err := r.Client.Delete(context.TODO(), resourcePVC)
			if err != nil && k8sErrors.IsNotFound(err) {
				return nil
			} else if err != nil {
				return errors.Errorf("delete PersistentVolumeClaim %s/%s failed: %s",
					resourcePVC.Namespace, resourcePVC.Name, err.Error())
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallDaemonSet(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.DaemonSetList {
		resourceDaemonSet := componentConfig.NewDaemonSet(fileString)
		foundDaemonSet := &appsv1.DaemonSet{}
		err := r.Client.Get(context.TODO(), types.NamespacedName{
			Name: resourceDaemonSet.Name, Namespace: resourceDaemonSet.Namespace,
		}, foundDaemonSet)
		if err != nil && k8sErrors.IsNotFound(err) {
			continue
		} else if err != nil {
			return errors.Errorf("get DaemonSet %s/%s failed: %s",
				resourceDaemonSet.Namespace, resourceDaemonSet.Name, err.Error())
		} else {
			err := r.Client.Delete(context.TODO(), resourceDaemonSet)
			if err != nil && k8sErrors.IsNotFound(err) {
				return nil
			} else if err != nil {
				return errors.Errorf("delete DaemonSet %s/%s failed: %s",
					resourceDaemonSet.Namespace, resourceDaemonSet.Name, err.Error())
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallPodSecurityPolicy(
	instance *federatoraiv1alpha1.AlamedaService,
	resource *alamedaserviceparamter.Resource) error {
	for _, fileString := range resource.PodSecurityPolicyList {
		psp, err := componentConfig.NewPodSecurityPolicy(fileString)
		if err != nil {
			return err
		}
		err = r.Client.Delete(context.TODO(), psp)
		if err != nil && k8sErrors.IsNotFound(err) {
			return nil
		} else if err != nil {
			switch psp := psp.(type) {
			case (*policyv1beta1.PodSecurityPolicy):
				return errors.Errorf("delete PodSecurityPolicy %s/%s failed: %s",
					psp.GetNamespace(), psp.GetName(), err.Error())
			case (*extensionsv1beta1.PodSecurityPolicy):
				return errors.Errorf("delete PodSecurityPolicy %s/%s failed: %s",
					psp.GetNamespace(), psp.GetName(), err.Error())
			default:
				return errors.Errorf(`not supported type %T`, psp)
			}
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) uninstallResource(
	resource alamedaserviceparamter.Resource) error {

	if err := r.uninstallStatefulSet(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall StatefulSet failed")
	}

	if err := r.uninstallRoute(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall Route failed")
	}

	if err := r.uninstallDeployment(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall Deployment failed")
	}

	if err := r.uninstallService(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall Service failed")
	}

	if err := r.uninstallConfigMap(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall ConfigMap failed")
	}

	if err := r.uninstallSecret(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall Secret failed")
	}

	if err := r.uninstallServiceAccount(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall ServiceAccount failed")
	}

	if err := r.uninstallClusterRole(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall ClusterRole failed")
	}

	if err := r.uninstallClusterRoleBinding(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall ClusterRoleBinding failed")
	}

	if err := r.uninstallAlamedaScaler(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall AlamedaScaler failed")
	}

	if err := r.uninstallPersistentVolumeClaim(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall PersistentVolumeClaim failed")
	}

	if err := r.uninstallDaemonSet(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall DaemonSet failed")
	}

	if err := r.uninstallPodSecurityPolicy(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall PodSecurityPolicy failed")
	}

	if err := r.uninstallMutatingWebhookConfiguration(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall MutatingWebhookConfiguration failed")
	}

	if err := r.uninstallValidatingWebhookConfiguration(nil, &resource); err != nil {
		return errors.Wrap(err, "uninstall ValidatingWebhookConfiguration failed")
	}

	return nil
}

func (r *AlamedaServiceReconciler) isNeedToBeReconciled(
	alamedaService *federatoraiv1alpha1.AlamedaService,
	gcIns *rbacv1.ClusterRole) (bool, error) {
	lock, err := r.getOrCreateAlamedaServiceLock(
		context.TODO(), *alamedaService, gcIns)
	if err != nil {
		return false, errors.Wrap(err, "get or create AlamedaService lock failed")
	}
	return federatoraioperatorcontrollerutil.IsAlamedaServiceLockOwnedByAlamedaService(
		lock, *alamedaService), nil
}

func (r *AlamedaServiceReconciler) getOrCreateAlamedaServiceLock(ctx context.Context,
	alamedaService federatoraiv1alpha1.AlamedaService,
	gcIns *rbacv1.ClusterRole) (rbacv1.ClusterRole, error) {
	lock, err := federatoraioperatorcontrollerutil.GetAlamedaServiceLock(ctx, r.Client)
	if err != nil && !k8sErrors.IsNotFound(err) {
		return lock, errors.Wrap(err, "get ClusterRole failed")
	} else if k8sErrors.IsNotFound(err) {
		lock = rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{
				Name: federatoraioperatorcontrollerutil.GetAlamedaServiceLockName(),
				Annotations: map[string]string{
					federatoraioperatorcontrollerutil.GetAlamedaServiceLockAnnotationKey(): fmt.Sprintf(
						"%s/%s", alamedaService.Namespace, alamedaService.Name),
				},
			},
		}
		if err := controllerutil.SetControllerReference(gcIns, &lock, r.Scheme); err != nil {
			return lock, errors.Errorf(
				"Fail alamedaservice lock SetControllerReference: %s", err.Error())
		}
		if err := r.Client.Create(ctx, &lock); err != nil {
			return lock, errors.Wrap(err, "create ClusterRole failed")
		}
	}
	return lock, nil
}

func (r *AlamedaServiceReconciler) deleteAlamedaServiceLock(ctx context.Context) error {
	lock := rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: federatoraioperatorcontrollerutil.GetAlamedaServiceLockName(),
		},
	}
	if err := r.Client.Delete(ctx, &lock); err != nil {
		return errors.Wrap(err, "delete ClusterRole failed")
	}
	return nil
}

func (r *AlamedaServiceReconciler) updateAlamedaServiceActivation(
	alamedaService *federatoraiv1alpha1.AlamedaService, active bool) error {
	copyAlamedaService := &federatoraiv1alpha1.AlamedaService{}
	r.Client.Get(context.TODO(), client.ObjectKey{
		Namespace: alamedaService.Namespace, Name: alamedaService.Name,
	}, copyAlamedaService)
	if active {
		copyAlamedaService.Status.Conditions =
			[]federatoraiv1alpha1.AlamedaServiceStatusCondition{
				{
					Paused: !active,
				},
			}
	} else {
		copyAlamedaService.Status.Conditions =
			[]federatoraiv1alpha1.AlamedaServiceStatusCondition{
				{
					Paused:  !active,
					Message: "Other AlamedaService is active.",
				},
			}
	}
	if err := r.Client.Update(context.Background(), copyAlamedaService); err != nil {
		return errors.Errorf("update AlamedaService active failed: %s", err.Error())
	}
	return nil
}

func (r *AlamedaServiceReconciler) updateAlamedaService(
	alamedaService *federatoraiv1alpha1.AlamedaService,
	namespaceName client.ObjectKey,
	asp *alamedaserviceparamter.AlamedaServiceParamter) error {
	if err := r.updateAlamedaServiceStatus(
		alamedaService, namespaceName, asp); err != nil {
		return err
	}
	if err := r.updateAlamedaServiceAnnotations(
		alamedaService, namespaceName); err != nil {
		return err
	}
	return nil
}

func (r *AlamedaServiceReconciler) updateAlamedaServiceStatus(
	alamedaService *federatoraiv1alpha1.AlamedaService, namespaceName client.ObjectKey,
	asp *alamedaserviceparamter.AlamedaServiceParamter) error {
	copyAlamedaService := alamedaService.DeepCopy()
	if err := r.Client.Get(
		context.TODO(), namespaceName, copyAlamedaService); err != nil {
		return errors.Errorf("get AlamedaService failed: %s", err.Error())
	}
	r.InitAlamedaService(copyAlamedaService)
	copyAlamedaService.Status.CRDVersion = asp.CurrentCRDVersion
	if err := r.Client.Update(context.Background(), copyAlamedaService); err != nil {
		return errors.Errorf("update AlamedaService Status failed: %s", err.Error())
	}
	r.Log.Info("Update AlamedaService Status Successfully",
		"resource.Name", copyAlamedaService.Name)
	return nil
}

func (r *AlamedaServiceReconciler) updateAlamedaServiceAnnotations(
	alamedaService *federatoraiv1alpha1.AlamedaService,
	namespaceName client.ObjectKey) error {
	copyAlamedaService := alamedaService.DeepCopy()
	if err := r.Client.Get(
		context.TODO(), namespaceName, copyAlamedaService); err != nil {
		return errors.Errorf("get AlamedaService failed: %s", err.Error())
	}
	r.InitAlamedaService(copyAlamedaService)
	jsonSpec, err := copyAlamedaService.GetSpecAnnotationWithoutKeycode()
	if err != nil {
		return errors.Errorf(
			"get AlamedaService spec annotation without keycode failed: %s", err.Error())
	}
	if copyAlamedaService.Annotations != nil {
		copyAlamedaService.Annotations["previousAlamedaServiceSpec"] = jsonSpec
	} else {
		annotations := make(map[string]string)
		annotations["previousAlamedaServiceSpec"] = jsonSpec
		copyAlamedaService.Annotations = annotations
	}
	if err := r.Client.Update(
		context.Background(), copyAlamedaService); err != nil {
		return errors.Errorf(
			"update AlamedaService Annotations failed: %s", err.Error())
	}
	r.Log.Info("Update AlamedaService Annotations Successfully",
		"resource.Name", copyAlamedaService.Name)
	return nil
}

func (r *AlamedaServiceReconciler) checkAlamedaServiceSpecIsChange(
	alamedaService *federatoraiv1alpha1.AlamedaService,
	namespaceName client.ObjectKey) (bool, error) {
	jsonSpec, err := alamedaService.GetSpecAnnotationWithoutKeycode()
	if err != nil {
		return false, errors.Errorf(
			"get AlamedaService spec annotation without keycode failed: %s", err.Error())
	}
	currentAlamedaServiceSpec := jsonSpec
	previousAlamedaServiceSpec := alamedaService.Annotations["previousAlamedaServiceSpec"]
	if currentAlamedaServiceSpec == previousAlamedaServiceSpec {
		return false, nil
	}
	return true, nil
}

func (r *AlamedaServiceReconciler) isAlamedaServiceFirstReconciledDone(
	alamedaService federatoraiv1alpha1.AlamedaService) bool {
	id := fmt.Sprintf(
		`%s/%s`, alamedaService.GetNamespace(), alamedaService.GetName())
	_, exist := r.FirstReconcileDoneAlamedaService[id]
	return !exist
}

func (r *AlamedaServiceReconciler) deleteDeploymentWhenModifyConfigMapOrService(
	dep *appsv1.Deployment) error {
	err := r.Client.Delete(context.TODO(), dep)
	if err != nil {
		return err
	}
	return nil
}

func (r *AlamedaServiceReconciler) patchConfigMapResourceVersionIntoPodTemplateSpecLabel(
	namespace string, podTemplateSpec *corev1.PodTemplateSpec) error {

	mountedConfigMapKey :=
		"configmaps.volumes.federator.ai/name-resourceversion"
	mountedConfigMapValueFormat := "%s-%s"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, volume := range podTemplateSpec.Spec.Volumes {
		if volume.ConfigMap != nil {
			configMap := corev1.ConfigMap{}
			err := r.Client.Get(ctx, client.ObjectKey{
				Namespace: namespace, Name: volume.ConfigMap.Name,
			}, &configMap)
			if err != nil {
				return errors.Errorf("get ConfigMap failed: %s", err.Error())
			}
			labels := podTemplateSpec.Labels
			if labels == nil {
				labels = make(map[string]string)
			}
			key := mountedConfigMapKey
			labels[key] = fmt.Sprintf(mountedConfigMapValueFormat,
				configMap.Name, configMap.ResourceVersion)
			podTemplateSpec.Labels = labels
		}
	}
	return nil
}

func (r *AlamedaServiceReconciler) removeUnsupportedResource(
	resource alamedaserviceparamter.Resource) alamedaserviceparamter.Resource {

	if !r.IsOpenshiftAPIRouteExist {
		resource.RouteList = nil
	}

	if !r.IsOpenshiftAPISecurityExist {
		resource.SecurityContextConstraintsList = nil
	}

	return resource
}
