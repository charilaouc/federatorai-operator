package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/containers-ai/federatorai-operator/api/v1alpha1"
	"github.com/containers-ai/federatorai-operator/pkg/consts"
	"github.com/pkg/errors"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

type GroupEnums string

const (
	Openshift_FEDEMETER_WORKER_NODE_LOWER_LIMIT = "1"
	Openshift_FEDEMETER_FILTER_TABLE            = "no_filter"
	NKS_FEDEMETER_WORKER_NODE_LOWER_LIMIT       = "2"
	NKS_FEDEMETER_FILTER_TABLE                  = "stackpoint_filter"

	FedemeterGroup GroupEnums = "alameda/fedemeter"
	AlamedaGroup   GroupEnums = "alameda"
	InfluxDBGroup  GroupEnums = "influxdb"
	//deployment name
	AlamedaaiDPN                 = "alameda-ai"
	AlamedaoperatorDPN           = "alameda-operator"
	AlamedadatahubDPN            = "alameda-datahub"
	AlamedaevictionerDPN         = "alameda-evictioner"
	AdmissioncontrollerDPN       = "admission-controller"
	AlamedarecommenderDPN        = "alameda-recommender"
	AlamedaexecutorDPN           = "alameda-executor"
	AlamedadispatcherDPN         = "alameda-ai-dispatcher"
	AlamedaRabbitMQDPN           = "alameda-rabbitmq"
	FedemeterDPN                 = "fedemeter-api"
	InfluxdbDPN                  = "alameda-influxdb"
	AlamedaweavescopeDPN         = "alameda-weave-scope-app"
	AlamedaweavescopeProbeDPN    = "alameda-weave-scope-cluster-agent"
	AlamedaanalyzerDPN           = "alameda-analyzer"
	AlamedaNotifierDPN           = "alameda-notifier"
	FederatoraiAgentDPN          = "federatorai-agent"
	FederatoraiAgentAppDPN       = "federatorai-agent-app"
	FederatoraiAgentGPUDPN       = "federatorai-agent-gpu"
	FederatoraiRestDPN           = "federatorai-rest"
	FederatoraiAgentPreloaderDPN = "federatorai-agent-preloader"
	FederatoraiFrontendDPN       = "federatorai-dashboard-frontend"
	FederatoraiBackendDPN        = "federatorai-dashboard-backend"
	//DaemonSet name
	AlamedaweavescopeAgentDS = "alameda-weave-scope-agent"
	//container name
	AlamedaaiCTN                 = "alameda-ai-engine"
	AlamedaoperatorCTN           = "alameda-operator"
	AlamedadatahubCTN            = "alameda-datahub"
	AlamedaevictionerCTN         = "alameda-evictioner"
	AdmissioncontrollerCTN       = "admission-controller"
	AlamedarecommenderCTN        = "alameda-recommender"
	AlamedaexecutorCTN           = "alameda-executor"
	AlamedadispatcherCTN         = "ai-dispatcher"
	AlamedaRabbitMQCTN           = "rabbitmq"
	FedemeterCTN                 = "fedemeter-api"
	FedemeterInfluxDBCTN         = "fedemeter-influxdb"
	GetTokenCTN                  = "gettoken"
	InfluxdbCTN                  = "influxdb"
	AlamedaweavescopeCTN         = "alameda-weave-scope-app"
	AlamedaweavescopeProbeCTN    = "alameda-weave-scope-cluster-agent"
	AlamedaweavescopeAgentCTN    = "alameda-weave-scope-agent"
	AlamedaanalyzerCTN           = "alameda-analyzer"
	AlamedaNofitierCTN           = "alameda-notifier"
	FederatoraiAgentCTN          = "federatorai-agent"
	FederatoraiAgentGPUCTN       = "federatorai-agent-gpu"
	FederatoraiAgentAppCTN       = "federatorai-agent-app"
	FederatoraiRestCTN           = "federatorai-rest"
	FederatoraiAgentPreloaderCTN = "federatorai-agent-preloader"
	FederatoraiBackendCTN        = "federatorai-dashboard-backend"
	FederatoraiFrontendCTN       = "federatorai-dashboard-frontend"

	//Statefulset name
	FedemeterInfluxDBSSN = "fedemeter-influxdb"
	//CRD NAME
	AlamedaScalerName         = "alamedascalers.autoscaling.containers.ai"
	AlamedaRecommendationName = "alamedarecommendations.autoscaling.containers.ai"

	//CRD Version
	OriAlamedaOperatorVersion = "v0.3.8"

	//AlamedaService modify Prometheus's var
	DefaultNamespace        = "federatorai"
	NamespaceServiceAccount = "serviceaccount:federatorai"

	//MountPath
	DataMountPath = "/var/lib"
	LogMountPath  = "/var/log"

	// Influxdb environment variables name
	AlamedaInfluxDBAdminUserEnvName     = "INFLUXDB_ADMIN_USER"
	AlamedaInfluxDBAdminPasswordEnvName = "INFLUXDB_ADMIN_PASSWORD"
	AlamedaInfluxDBHTTPSEnabledEnvName  = "INFLUXDB_HTTP_HTTPS_ENABLED"

	AlamedaInfluxDBAPIPort = 8086
)

var (
	//if disable resource protection
	Disable_operand_resource_protection = "false"
	log                                 = logf.Log.WithName("controller_alamedaservice")
	//AlamedaScaler version
	AlamedaScalerVersion        = []string{"v1", "v2"}
	V1scalerOperatorVersionList = []string{
		"v0.3.6", "v0.3.7", "v0.3.8", "v0.3.9", "v0.3.10", "v0.3.11", "v0.3.12",
	}
)

// GetServiceAddress returns address combining dns name with port number base on port name
func GetServiceAddress(svc *corev1.Service, portName string) (string, error) {

	portNum := int32(0)
	exist := false
	for _, port := range svc.Spec.Ports {
		if port.Name == portName {
			portNum = port.Port
			exist = true
			break
		}
	}
	if !exist {
		return "", errors.New("port name does not exist")
	}

	namespace := svc.Namespace
	name := svc.Name
	return fmt.Sprintf("%s.%s.svc:%d", name, namespace, portNum), nil
}

// GetServiceDNS returns service dns
func GetServiceDNS(svc *corev1.Service) string {
	namespace := svc.Namespace
	name := svc.Name
	return fmt.Sprintf("%s.%s.svc", name, namespace)
}

func SetBootStrapImageStruct(
	dep *appsv1.Deployment,
	componentspec v1alpha1.AlamedaComponentSpec,
	ctn string) {
	for index, value := range dep.Spec.Template.Spec.InitContainers {
		if value.Name == ctn {
			if componentspec.BootStrapContainer.Image != "" ||
				componentspec.BootStrapContainer.Version != "" {
				image := fmt.Sprintf("%s:%s",
					componentspec.BootStrapContainer.Image,
					componentspec.BootStrapContainer.Version)
				dep.Spec.Template.Spec.InitContainers[index].Image = image
			}
			dep.Spec.Template.Spec.InitContainers[index].ImagePullPolicy =
				componentspec.BootStrapContainer.ImagePullPolicy
		}
	}
}

//if user section schema set pullpolicy then AlamedaService set Containers image's pullpolicy
func SetImagePullPolicy(
	dep interface{}, ctn string, imagePullPolicy corev1.PullPolicy) {
	rtObj := reflect.ValueOf(dep).Elem()
	containers := rtObj.FieldByName("Spec").FieldByName(
		"Template").FieldByName("Spec").FieldByName("Containers")
	for i := 0; i < reflect.ValueOf(containers.Interface()).Len(); i++ {
		ctName := containers.Index(i).FieldByName("Name").String()
		if ctName == ctn {
			containers.Index(i).FieldByName("ImagePullPolicy").SetString(
				string(imagePullPolicy))
			log.V(1).Info("Set ImagePullPolicy", "Kind",
				rtObj.FieldByName("Kind"), "Container Name",
				ctName, "Image Pull Policy", imagePullPolicy)
		}
	}
}

func getVolumeIndex(dep interface{}, suffix string) int {
	rtObj := reflect.ValueOf(dep).Elem()
	volumes := rtObj.FieldByName("Spec").FieldByName(
		"Template").FieldByName("Spec").FieldByName("Volumes")
	for i := 0; i < reflect.ValueOf(volumes.Interface()).Len(); i++ {
		volName := volumes.Index(i).FieldByName("Name").String()
		if strings.HasSuffix(volName, suffix) {
			return i
		}
	}
	return -1
}

//if user set ephemeral then AlamedaService set Deployment VolumeSource is EmptyDir
func setEmptyDir(dep interface{}, index int, size string) {
	rtObj := reflect.ValueOf(dep).Elem()
	rtName := rtObj.FieldByName("Name").String()
	rtKind := rtObj.FieldByName("Kind").String()
	volumes := rtObj.FieldByName("Spec").FieldByName(
		"Template").FieldByName("Spec").FieldByName("Volumes")
	if size != "" {
		quantity := resource.MustParse(size)
		emptydir := &corev1.EmptyDirVolumeSource{SizeLimit: &quantity}
		vs := corev1.VolumeSource{EmptyDir: emptydir}
		volumes.Index(index).FieldByName("VolumeSource").Set(reflect.ValueOf(vs))
	} else {
		vs := corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}
		volumes.Index(index).FieldByName("VolumeSource").Set(reflect.ValueOf(vs))
	}
	log.V(1).Info("SetVolumeSourceEmptyDir", "Kind", rtKind, "Name", rtName)
}

//if user set pvc then AlamedaService set Deployment VolumeSource is PersistentVolumeClaim
func setVolumeSource(dep interface{}, index int, claimName string) {
	rtObj := reflect.ValueOf(dep).Elem()
	rtName := rtObj.FieldByName("Name").String()
	rtKind := rtObj.FieldByName("Kind").String()
	volumes := rtObj.FieldByName("Spec").FieldByName(
		"Template").FieldByName("Spec").FieldByName("Volumes")
	pvcs := &corev1.PersistentVolumeClaimVolumeSource{ClaimName: claimName}
	vs := corev1.VolumeSource{PersistentVolumeClaim: pvcs}
	volumes.Index(index).FieldByName("VolumeSource").Set(reflect.ValueOf(vs))
	log.V(1).Info("SetVolumeSource", "Kind", rtKind, "Name", rtName, "PVCs", pvcs)
}

//if user set pvc then AlamedaService set pvc to Deployment's VolumeSource
func SetStorageToVolumeSource(dep interface{},
	storagestructs []v1alpha1.StorageSpec,
	volumeName string, group GroupEnums) {
	logSuffix := consts.LogSuffix
	dataSuffix := consts.DataSuffix
	for _, v := range storagestructs {
		if !v.StorageIsEmpty() {
			if index := getVolumeIndex(dep, logSuffix); index != -1 &&
				v.Usage == v1alpha1.Log {
				setVolumeSource(dep, index, strings.Replace(
					volumeName, "type", string(v1alpha1.Log), -1))
			}
			if index := getVolumeIndex(dep, dataSuffix); index != -1 &&
				v.Usage == v1alpha1.Data {
				setVolumeSource(dep, index, strings.Replace(
					volumeName, "type", string(v1alpha1.Data), -1))
			}
			if v.Usage == v1alpha1.Empty && group == AlamedaGroup {
				if index := getVolumeIndex(dep, logSuffix); index != -1 {
					setVolumeSource(dep, index, strings.Replace(
						volumeName, "type", string(v1alpha1.Log), -1))
				}
				if index := getVolumeIndex(dep, dataSuffix); index != -1 {
					setVolumeSource(dep, index, strings.Replace(
						volumeName, "type", string(v1alpha1.Data), -1))
				}
			} else if v.Usage == v1alpha1.Empty && group != AlamedaGroup {
				if index := getVolumeIndex(dep, dataSuffix); index != -1 {
					setVolumeSource(dep, index, strings.Replace(
						volumeName, "type", string(v1alpha1.Data), -1))
				}
			}
		}
		if v.Type == v1alpha1.Ephemeral {
			if index := getVolumeIndex(dep, logSuffix); index != -1 &&
				v.Usage == v1alpha1.Log {
				setEmptyDir(dep, index, v.Size)
			}
			if index := getVolumeIndex(dep, dataSuffix); index != -1 &&
				v.Usage == v1alpha1.Data {
				setEmptyDir(dep, index, v.Size)
			}
		}
	}
}

func setMountPath(dep interface{}, volumeName string,
	mountPath string, ctn string, group GroupEnums) {
	rtObj := reflect.ValueOf(dep).Elem()
	containers := rtObj.FieldByName("Spec").FieldByName(
		"Template").FieldByName("Spec").FieldByName("Containers")
	for i := 0; i < reflect.ValueOf(containers.Interface()).Len(); i++ {
		ctName := containers.Index(i).FieldByName("Name").String()
		if ctName == ctn {
			volMnts := containers.Index(i).FieldByName("VolumeMounts")
			for v := 0; v < reflect.ValueOf(volMnts.Interface()).Len(); v++ {
				if volMnts.Index(v).FieldByName("Name").String() == volumeName {
					return
				}
			}
			vm := corev1.VolumeMount{
				Name:      volumeName,
				MountPath: mountPath,
			}
			if group != AlamedaGroup {
				vm.SubPath = string(group)
			}
			containers.Index(i).FieldByName("VolumeMounts").Set(reflect.ValueOf(
				append([]corev1.VolumeMount{vm}, containers.Index(i).FieldByName(
					"VolumeMounts").Interface().([]corev1.VolumeMount)...)))
		}
	}
}

//if user set pvc then AlamedaService set pvc to Deployment's MountPath
func SetStorageToMountPath(
	dep interface{}, storagestructs []v1alpha1.StorageSpec,
	ctn string, volumeName string, group GroupEnums) {
	for _, v := range storagestructs {
		if v.Type == v1alpha1.Ephemeral || v.Type == v1alpha1.PVC {
			if v.Usage == v1alpha1.Data {
				setMountPath(dep, strings.Replace(volumeName, "type",
					string(v1alpha1.Data), -1),
					fmt.Sprintf("%s/%s", DataMountPath, group), ctn, group)
			} else if v.Usage == v1alpha1.Log {
				setMountPath(dep, strings.Replace(volumeName, "type",
					string(v1alpha1.Log), -1),
					fmt.Sprintf("%s/%s", LogMountPath, group), ctn, group)
			} else if v.Usage == v1alpha1.Empty && group == AlamedaGroup {
				setMountPath(dep, strings.Replace(volumeName, "type",
					string(v1alpha1.Data), -1),
					fmt.Sprintf("%s/%s", DataMountPath, group), ctn, group)
				setMountPath(dep, strings.Replace(volumeName, "type",
					string(v1alpha1.Log), -1),
					fmt.Sprintf("%s/%s", LogMountPath, group), ctn, group)
			} else if v.Usage == v1alpha1.Empty && group != AlamedaGroup {
				// if not alameda component's then only set data
				setMountPath(dep, strings.Replace(volumeName, "type",
					string(v1alpha1.Data), -1),
					fmt.Sprintf("%s/%s", DataMountPath, group), ctn, group)
			}
		}
	}
}

func setPVCSpec(pvc *corev1.PersistentVolumeClaim, value v1alpha1.StorageSpec) {
	if value.AccessModes != "" {
		pvc.Spec.AccessModes = append(pvc.Spec.AccessModes, value.AccessModes)
	}
	if value.Size != "" {
		pvc.Spec.Resources.Requests[corev1.ResourceStorage] =
			resource.MustParse(value.Size)
	}
	if value.Class != nil {
		pvc.Spec.StorageClassName = value.Class
	}
}

//if user set pvc then AlamedaService set PersistentVolumeClaimSpec
func SetStorageToPersistentVolumeClaimSpec(
	pvc *corev1.PersistentVolumeClaim,
	storagestructs []v1alpha1.StorageSpec,
	pvctype v1alpha1.Usage) {
	for k, v := range storagestructs {
		if v.Usage == pvctype || v.Usage == v1alpha1.Empty {
			setPVCSpec(pvc, storagestructs[k])
		}
	}
}

func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// StringSliceDelete removes elements in slice2 from slice1
func StringSliceDelete(slice1 []string, slice2 []string) []string {

	var diff []string

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}

func ServerHasAPIGroup(apiGroupName string) (bool, error) {

	config, err := config.GetConfig()
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return false, err
	}
	apiGroups, err := k8sClient.ServerGroups()
	if err != nil {
		return false, err
	}
	for _, apiGroup := range apiGroups.Groups {
		if apiGroup.Name == apiGroupName {
			return true, nil
		}
	}
	return false, nil
}

func ServerHasResourceInAPIGroupVersion(
	resourceName, apiGroupVersion string) (bool, error) {

	config, err := config.GetConfig()
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return false, err
	}
	resourceList, err :=
		k8sClient.ServerResourcesForGroupVersion(apiGroupVersion)
	if err != nil {
		return false, err
	}
	for _, resource := range resourceList.APIResources {
		if resource.Name == resourceName {
			return true, nil
		}
	}
	return false, nil
}

func SetResourcesForContainers(
	obj interface{}, targetResources corev1.ResourceRequirements,
	isInitContainer bool) {
	rtObj := reflect.ValueOf(obj).Elem()
	ctType := "Containers"
	if isInitContainer {
		ctType = "InitContainers"
	}
	containers := rtObj.FieldByName("Spec").FieldByName(
		"Template").FieldByName("Spec").FieldByName(ctType)
	for i := 0; i < reflect.ValueOf(containers.Interface()).Len(); i++ {
		resources := containers.Index(i).FieldByName("Resources")
		if resources.FieldByName("Requests").IsNil() {
			resources.FieldByName("Requests").Set(
				reflect.ValueOf(corev1.ResourceList{}))
		}
		if resources.FieldByName("Limits").IsNil() {
			resources.FieldByName("Limits").Set(
				reflect.ValueOf(corev1.ResourceList{}))
		}

		for k, v := range targetResources.Requests {
			resources.FieldByName("Requests").SetMapIndex(
				reflect.ValueOf(k), reflect.ValueOf(v))
		}
		for k, v := range targetResources.Limits {
			resources.FieldByName("Limits").SetMapIndex(
				reflect.ValueOf(k), reflect.ValueOf(v))
		}
	}
}
