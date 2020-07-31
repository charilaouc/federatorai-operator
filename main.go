package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	autoscaling_v1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	alamedautilsk8s "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	datahubv1alpha1_event "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events"
	federatoraiv1alpha1 "github.com/containers-ai/federatorai-operator/api/v1alpha1"
	assets "github.com/containers-ai/federatorai-operator/assets"
	"github.com/containers-ai/federatorai-operator/controllers"
	assetsBin "github.com/containers-ai/federatorai-operator/pkg/assets"
	client_datahub "github.com/containers-ai/federatorai-operator/pkg/client/datahub"
	"github.com/containers-ai/federatorai-operator/pkg/component"
	"github.com/containers-ai/federatorai-operator/pkg/consts"
	"github.com/containers-ai/federatorai-operator/pkg/lib/resourceread"
	fedOperatorLog "github.com/containers-ai/federatorai-operator/pkg/log"
	patchprom "github.com/containers-ai/federatorai-operator/pkg/patch/prometheus"
	alamedaserviceparamter "github.com/containers-ai/federatorai-operator/pkg/processcrdspec/alamedaserviceparamter"
	"github.com/containers-ai/federatorai-operator/pkg/protocol/grpc"
	upgradeinfluxdb "github.com/containers-ai/federatorai-operator/pkg/upgrader/influxdb"
	"github.com/containers-ai/federatorai-operator/pkg/util"
	"github.com/containers-ai/federatorai-operator/pkg/version"
	prom_op_api "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/openshift/api/route"
	routev1 "github.com/openshift/api/route/v1"
	"github.com/openshift/api/security"
	securityv1 "github.com/openshift/api/security/v1"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	apiextensionv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	k8sapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/scale/scheme/extensionsv1beta1"
	apiregistrationv1beta1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	scheme   = k8sruntime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(routev1.AddToScheme(scheme))
	utilruntime.Must(securityv1.AddToScheme(scheme))
	utilruntime.Must(autoscaling_v1alpha1.AddToScheme(scheme))
	utilruntime.Must(apiregistrationv1beta1.AddToScheme(scheme))
	utilruntime.Must(prom_op_api.AddToScheme(scheme))
	utilruntime.Must(apiextensionv1beta1.AddToScheme(scheme))
	utilruntime.Must(federatoraiv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

var (
	configurationFilePath string

	fedOperatorConfig Config

	log = logf.Log.WithName("manager")

	watchNamespace = ""

	registerdAPIResources = make(map[string]bool)
)

func initConfiguration() {

	fedOperatorConfig = NewDefaultConfig()

	initViperSetting()
	mergeViperValueWithDefaultConfig()
}

func initViperSetting() {

	viper.SetEnvPrefix(consts.EnvVarPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(strings.Split(consts.EnvReplacerOldNew, ";")...))
	viper.AllowEmptyEnv(consts.AllowEmptyEnv)
}

func mergeViperValueWithDefaultConfig() {

	viper.SetConfigFile(configurationFilePath)

	if err := viper.ReadInConfig(); err != nil {
		panic(errors.New("Read configuration file failed: " + err.Error()))
	}

	if err := viper.Unmarshal(&fedOperatorConfig); err != nil {
		panic(errors.New("Unmarshal configuration failed: " + err.Error()))
	}
}

func initLogger() {

	logPaths := viper.GetStringSlice("log.outputPaths")
	if len(logPaths) == 0 {
		fedOperatorConfig.Log.AppendOutput(consts.DefaultLogOutputPath)
	} else {
		fedOperatorConfig.Log.OutputPaths = []string{}
		for _, logPath := range logPaths {
			fedOperatorConfig.Log.AppendOutput(logPath)
		}
	}
	logger, err := fedOperatorLog.NewZaprLogger(fedOperatorConfig.Log)
	if err != nil {
		panic(err)
	}
	logf.SetLogger(logger)

	grpcLogPaths := viper.GetStringSlice("grpc.log.outputPaths")
	if len(grpcLogPaths) == 0 {
		fedOperatorConfig.GRPC.Log.AppendOutput(consts.DefaultLogOutputPath)
	} else {
		fedOperatorConfig.GRPC.Log.OutputPaths = []string{}
		for _, grpcLogPath := range grpcLogPaths {
			fedOperatorConfig.GRPC.Log.AppendOutput(grpcLogPath)
		}
	}

	grpcLogger, err := fedOperatorLog.NewZapLogger(fedOperatorConfig.GRPC.Log)
	if err != nil {
		panic(err)
	}
	grpc.SetGRPCLogger(grpcLogger)
}

func printVersion() {
	log.Info(fmt.Sprintf("Go Version: %s", runtime.Version()))
	log.Info(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
	log.Info(fmt.Sprintf("Federatorai Operator Version: %v", version.String))
}

func printConfiguration() {
	if b, err := json.MarshalIndent(fedOperatorConfig, "", "    "); err != nil {
		panic(err.Error())
	} else {
		log.Info(fmt.Sprintf("%+v", string(b)))
	}
}

func initProcesses() {
	initConfiguration()
	initLogger()
	printVersion()
	printConfiguration()
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))
}

func main() {
	restCfg := ctrl.GetConfigOrDie()
	cli, err := client.New(restCfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		panic(err)
	}

	upgradeCmd := flag.NewFlagSet("upgrade", flag.ExitOnError)
	promCheckCmd := flag.NewFlagSet("prom_check", flag.ExitOnError)
	promApplyCmd := flag.NewFlagSet("prom_apply", flag.ExitOnError)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "upgrade":
			upgradeType := upgradeCmd.String("type", "influxdb", "type")
			upgradeTimeout := upgradeCmd.String("timeout", "300s", "timeout")
			cfgPath := upgradeCmd.String("config", consts.DefaultConfigPath, "File path to federatorai-operator coniguration")
			upgradeCmd.Parse(os.Args[2:])
			configurationFilePath = *cfgPath

			initProcesses()
			if *upgradeType == "influxdb" {
				log.Info("Start upgrading influxdb")
				if err := upgradeinfluxdb.InfluxdbUpgrade(*upgradeTimeout, cli); err != nil {
					log.Error(err, "")
					os.Exit(1)
				}
				log.Info("Upgrade influxdb successfully")
			}

			os.Exit(0)
		case "prom_check":
			cfgPath := promCheckCmd.String("config", consts.DefaultConfigPath, "File path to federatorai-operator coniguration")
			promCheckCmd.Parse(os.Args[2:])
			configurationFilePath = *cfgPath

			initProcesses()
			log.Info("Start checking prometheus rules")
			if err := patchprom.PromCheck(cli); err != nil {
				log.Error(err, "")
				os.Exit(1)
			}
			log.Info("Check prometheus rules successfully")
			os.Exit(0)
		case "prom_apply":
			cfgPath := promApplyCmd.String("config", consts.DefaultConfigPath, "File path to federatorai-operator coniguration")
			promApplyCmd.Parse(os.Args[2:])
			configurationFilePath = *cfgPath

			initProcesses()
			log.Info("Start applying prometheus rules")
			if err := patchprom.PromApply(cli); err != nil {
				log.Error(err, "")
				os.Exit(1)
			}
			log.Info("Apply prometheus rules successfully")
			os.Exit(0)
		default:
		}
	}

	var metricsAddr string
	var enableLeaderElection bool
	//flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.StringVar(&configurationFilePath, "config", consts.DefaultConfigPath, "File path to federatorai-operator coniguration")
	flag.Parse()

	initProcesses()
	metricsAddr = fmt.Sprintf("%s:%d", viper.GetString("metrics.host"), viper.GetInt32("metrics.port"))
	mgr, err := ctrl.NewManager(restCfg, ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "0c08aad9.containers.ai",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	clusterID, err := alamedautilsk8s.GetClusterUID(cli)
	if err != nil {
		setupLog.Error(err, "unable to get cluster Id")
	}
	if err = (&controllers.AlamedaServiceKeycodeReconciler{
		Client:                   mgr.GetClient(),
		Log:                      ctrl.Log.WithName("controllers").WithName("AlamedaServiceKeycode"),
		Scheme:                   mgr.GetScheme(),
		DatahubClientMap:         make(map[controllers.Namespace]client_datahub.Client),
		DatahubClientMapLock:     sync.Mutex{},
		FirstRetryTimeCache:      make(map[types.NamespacedName]*time.Time),
		FirstRetryTimeLock:       sync.Mutex{},
		EventChanMap:             make(map[controllers.Namespace]chan datahubv1alpha1_event.Event),
		EventChanMapLock:         sync.Mutex{},
		LastReconcileTaskMap:     make(map[controllers.Namespace]controllers.KeycodeStatus),
		LastReconcileTaskMapLock: sync.Mutex{},
		ClusterID:                clusterID,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AlamedaServiceKeycode")
		os.Exit(1)
	}

	hasOpenshiftAPIRoute, err := util.ServerHasAPIGroup(route.GroupName)
	if err != nil {
		panic(err)
	}
	hasOpenshiftAPISecurity, err := util.ServerHasAPIGroup(security.GroupName)
	if err != nil {
		panic(err)
	}
	var podSecurityPolicesAPIGroupVersion schema.GroupVersion
	hasPSPInExtensionV1beta1, err := util.ServerHasResourceInAPIGroupVersion(
		"podsecuritypolicies", extensionsv1beta1.SchemeGroupVersion.String())
	if err != nil {
		panic(err)
	} else if hasPSPInExtensionV1beta1 {
		podSecurityPolicesAPIGroupVersion = extensionsv1beta1.SchemeGroupVersion
	}
	hasPSPInPolicyV1beta1, err := util.ServerHasResourceInAPIGroupVersion(
		"podsecuritypolicies", policyv1beta1.SchemeGroupVersion.String())
	if err != nil {
		panic(err)
	} else if hasPSPInPolicyV1beta1 {
		podSecurityPolicesAPIGroupVersion = policyv1beta1.SchemeGroupVersion
	}
	if err = (&controllers.AlamedaServiceReconciler{
		Client:                            mgr.GetClient(),
		Log:                               ctrl.Log.WithName("controllers").WithName("AlamedaService"),
		Scheme:                            mgr.GetScheme(),
		FirstReconcileDoneAlamedaService:  make(map[string]struct{}),
		Apiextclient:                      clientset.NewForConfigOrDie(mgr.GetConfig()),
		IsOpenshiftAPIRouteExist:          hasOpenshiftAPIRoute,
		IsOpenshiftAPISecurityExist:       hasOpenshiftAPISecurity,
		PodSecurityPolicesAPIGroupVersion: podSecurityPolicesAPIGroupVersion,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AlamedaService")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	if err := createCustomeResourceDefinitions(restCfg); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	if err := waitCRDReady(restCfg); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Setup requirements before starts the manager
	if err := setupRequirements(restCfg); err != nil {
		log.Error(err, "setup requirements failed")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

}

func setupRequirements(clientConfig *rest.Config) error {

	if err := createConfigMaps(clientConfig); err != nil {
		return errors.Wrapf(err, "create configMaps failed")
	}

	return nil
}

// In order to let manager has the scheme definition of the crds,
// it should install those crds in to the cluster before the instance of manager is created,
func createCustomeResourceDefinitions(clientConfig *rest.Config) error {

	apiExtensionsClientset, err := clientset.NewForConfig(clientConfig)
	if err != nil {
		return errors.Errorf("create k8s clientset failed: %s", err.Error())
	}

	assets := alamedaserviceparamter.GetCustomResourceDefinitions()
	for _, asset := range assets {

		// Since CRDs will be update later, we can use empty ComponentConfig here.
		cc := component.ComponentConfig{}
		crd := cc.NewCustomResourceDefinition(asset)
		addCRDToRegisterdAPIResources(crd)
		//use apiExtensionsClientset.ApiextensionsV1() in k8s 1.19 or later
		_, err = apiExtensionsClientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(context.TODO(), crd, metav1.CreateOptions{})
		if err != nil && k8sapierrors.IsAlreadyExists(err) {
			log.Info("CustomResourceDefinition is existing in cluster, will not create or update it.", "CustomResourceDefinition name", crd.Name)
			continue
		} else if err != nil {
			return errors.Wrapf(err, "create CustomResourceDefinition(%s) failed", crd.Name)
		}
	}
	return nil
}

func addCRDToRegisterdAPIResources(crd *apiextensionv1beta1.CustomResourceDefinition) {

	group := crd.Spec.Group
	kind := crd.Spec.Names.Kind

	for _, crdVersion := range crd.Spec.Versions {
		if crdVersion.Storage {
			version := crdVersion.Name
			groupVersionKind := fmt.Sprintf("%s/%s/%s", group, version, kind)
			addAPIResource(groupVersionKind, registerdAPIResources)
			return
		}
	}

	if crd.Spec.Version != "" {
		version := crd.Spec.Version
		groupVersionKind := fmt.Sprintf("%s/%s/%s", group, version, kind)
		addAPIResource(groupVersionKind, registerdAPIResources)
	}
}

func addAPIResource(groupVersionKind string, gvkMap map[string]bool) {
	gvkMap[groupVersionKind] = true
}

func deleteAPIResource(groupVersionKind string, gvkMap map[string]bool) {
	delete(gvkMap, groupVersionKind)
}

func createConfigMaps(clientConfig *rest.Config) error {

	cli, err := client.New(clientConfig, client.Options{})
	if err != nil {
		return errors.Errorf("new k8s client failed: %s", err.Error())
	}

	ctx := context.TODO()
	files := assets.GetRequiredConfigMaps()
	for _, file := range files {
		fileBytes, err := assetsBin.Asset(file)
		if err != nil {
			return errors.Errorf("get asset's bytes failed: %s", err.Error())
		}
		configMap := resourceread.ReadConfigMapV1(fileBytes)
		err = cli.Create(ctx, configMap)
		if err != nil && !k8sapierrors.IsAlreadyExists(err) {
			return errors.Errorf("create configMap %s/%s failed: %s", configMap.Namespace, configMap.Name, err.Error())
		}
	}

	return nil
}

func waitCRDReady(clientConfig *rest.Config) error {

	waitInterval := 500 * time.Millisecond
	if err := wait.Poll(waitInterval, 30*time.Second, func() (bool, error) {
		apiExtensionsClientset, err := clientset.NewForConfig(clientConfig)
		if err != nil {
			log.V(-1).Info("Create k8s clientset failed, will retry", "msg", err.Error())
			return false, nil
		}

		_, apiList, err := apiExtensionsClientset.DiscoveryClient.ServerGroupsAndResources()
		if err != nil {
			log.V(-1).Info("Get k8s ServerGroupsAndResources failed, will retry", "msg", err.Error())
		}

		for _, apiResourceList := range apiList {
			for _, apiResource := range apiResourceList.APIResources {
				groupVersion := apiResourceList.GroupVersion // fmt.Sprintf("%s/%s",group,version)
				kind := apiResource.Kind
				groupVersionKind := fmt.Sprintf("%s/%s", groupVersion, kind)
				deleteAPIResource(groupVersionKind, registerdAPIResources)
			}
		}

		ok := len(registerdAPIResources) == 0
		if !ok {
			log.V(-1).Info("Server does not have required apiResources, will retry fetching")
		}
		return ok, nil

	}); err != nil {
		return err
	}

	return nil
}
