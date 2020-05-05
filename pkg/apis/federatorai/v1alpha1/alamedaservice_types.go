package v1alpha1

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Platform = string

const (
	PlatformOpenshift3_9 Platform = "openshift3.9"
)

type TLSSpec struct {
	Enabled            bool `json:"enabled,omitempty"`
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

type SASLSpec struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type KafkaSpec struct {
	BrokerAddresses []string `json:"brokerAddresses,omitempty"`
	// Version         string   `json:"version,omitempty"`
	TLS  TLSSpec  `json:"tls,omitempty"`
	SASL SASLSpec `json:"sasl,omitempty"`
}

type NginxSpec struct {
	Enabled bool `json:"enabled"`
}

type ClusterAutoScalerSpec struct {
	EnableExecution bool   `json:"enableExecution,omitempty"`
	ForeseeTime     *int32 `json:"foreseeTime,omitempty"`
	MaxNodes        *int32 `json:"maxNodes,omitempty"`
	MaxCPU          *int32 `json:"maxCPU,omitempty"`
	MaxMem          *int32 `json:"maxMem,omitempty"`
}

// AlamedaServiceSpec defines the desired state of AlamedaService
// +k8s:openapi-gen=true
type AlamedaServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	// +kubebuilder:validation:Enum=openshift3.9
	Platform        Platform `json:"platform,omitempty"`
	EnableExecution bool     `json:"enableExecution"`
	EnableGUI       bool     `json:"enableGui"`
	EnableVPA       *bool    `json:"enableVPA"`
	EnableGPU       *bool    `json:"enableGPU"`
	// +nullable
	EnableDispatcher         *bool         `json:"enableDispatcher,omitempty"`
	EnablePreloader          bool          `json:"enablePreloader,omitempty"`
	SelfDriving              bool          `json:"selfDriving"`
	AutoPatchPrometheusRules bool          `json:"autoPatchPrometheusRules"`
	Version                  string        `json:"version"`
	ImageLocation            string        `json:"imageLocation,omitempty"`
	PrometheusService        string        `json:"prometheusService"`
	Storages                 []StorageSpec `json:"storages"`
	// +nullable
	ServiceExposures  []ServiceExposureSpec `json:"serviceExposures,omitempty"`
	EnableWeavescope  bool                  `json:"enableWeavescope,omitempty"`
	Keycode           KeycodeSpec           `json:"keycode"`
	Kafka             KafkaSpec             `json:"kafka"`
	Nginx             NginxSpec             `json:"nginx"`
	ClusterAutoScaler ClusterAutoScalerSpec `json:"clusterAutoScaler,omitempty"`

	//Component Section Schema
	InfluxdbSectionSet                  AlamedaComponentSpec    `json:"alamedaInfluxdb,omitempty"`
	GrafanaSectionSet                   AlamedaComponentSpec    `json:"alamedaGrafana,omitempty"`
	AlamedaAISectionSet                 AlamedaComponentSpec    `json:"alamedaAi,omitempty"`
	AlamedaOperatorSectionSet           AlamedaComponentSpec    `json:"alamedaOperator,omitempty"`
	AlamedaDatahubSectionSet            AlamedaComponentSpec    `json:"alamedaDatahub,omitempty"`
	AlamedaEvictionerSectionSet         AlamedaComponentSpec    `json:"alamedaEvictioner,omitempty"`
	AdmissionControllerSectionSet       AlamedaComponentSpec    `json:"alamedaAdmissionController,omitempty"`
	AlamedaRecommenderSectionSet        AlamedaComponentSpec    `json:"alamedaRecommender,omitempty"`
	AlamedaExecutorSectionSet           AlamedaComponentSpec    `json:"alamedaExecutor,omitempty"`
	AlamedaFedemeterSectionSet          AlamedaComponentSpec    `json:"fedemeter,omitempty"`
	AlamedaFedemeterInfluxDBSectionSet  AlamedaComponentSpec    `json:"fedemeterInfluxdb,omitempty"`
	AlamedaWeavescopeSectionSet         AlamedaComponentSpec    `json:"alameda-weavescope,omitempty"`
	AlamedaDispatcherSectionSet         AlamedaComponentSpec    `json:"alameda-dispatcher,omitempty"`
	AlamedaRabbitMQSectionSet           AlamedaComponentSpec    `json:"alamedaRabbitMQ,omitempty"`
	AlamedaAnalyzerSectionSet           AlamedaComponentSpec    `json:"alameda-analyzer,omitempty"`
	AlamedaNotifierSectionSet           AlamedaComponentSpec    `json:"alamedaNotifier,omitempty"`
	FederatoraiAgentSectionSet          AlamedaComponentSpec    `json:"federatoraiAgent,omitempty"`
	FederatoraiAgentGPUSectionSet       FederatoraiAgentGPUSpec `json:"federatoraiAgentGPU,omitempty"`
	FederatoraiRestSectionSet           AlamedaComponentSpec    `json:"federatoraiRest,omitempty"`
	FederatoraiAgentPreloaderSectionSet AlamedaComponentSpec    `json:"federatoraiAgentPreloader,omitempty"`
	FederatoraiFrontendSectionSet       AlamedaComponentSpec    `json:"federatoraiFrontend,omitempty"`
	FederatoraiBackendSectionSet        AlamedaComponentSpec    `json:"federatoraiBackend,omitempty"`
	FederatoraiAgentAppSectionSet       AlamedaComponentSpec    `json:"federatoraiAgentApp,omitempty"`
}

type AlamedaComponentSpec struct {
	Image           string            `json:"image"`
	Version         string            `json:"version"`
	ImagePullPolicy corev1.PullPolicy `json:"imagepullpolicy"`
	// +nullable
	Storages           []StorageSpec `json:"storages"`
	BootStrapContainer Imagestruct   `json:"bootstrap"`
	// +nullable
	EnvVars []corev1.EnvVar `json:"env"`
}

type FederatoraiAgentGPUSpec struct {
	AlamedaComponentSpec `json:",inline"`
	// +nullable
	Prometheus *PrometheusConfig `json:"prometheus"`
	// +nullable
	InfluxDB *InfluxDBConfig `json:"influxDB"`
}

type PrometheusConfig struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type InfluxDBConfig struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Imagestruct struct {
	Image           string            `json:"image"`
	Version         string            `json:"version"`
	ImagePullPolicy corev1.PullPolicy `json:"imagepullpolicy"`
}
type Usage string
type Type string

const (
	Empty     Usage = ""
	Log       Usage = "log"
	Data      Usage = "data"
	PVC       Type  = "pvc"
	Ephemeral Type  = "ephemeral"
)

var (
	PvcUsage = []Usage{Data, Log}
)

type StorageSpec struct {
	Type  Type   `json:"type"`
	Usage Usage  `json:"usage"`
	Size  string `json:"size,omitempty"`
	// +nullable
	Class       *string                           `json:"class,omitempty"`
	AccessModes corev1.PersistentVolumeAccessMode `json:"accessMode,omitempty"`
}

//check StorageStruct
func (storageStruct StorageSpec) StorageIsEmpty() bool {
	if storageStruct.Size != "" && storageStruct.Type == PVC {
		return false
	}
	return true
}

// ServiceExposureType defines the type of the service to be exposed
type ServiceExposureType = string

var (
	// ServiceExposureTypeNodePort represents NodePort type
	ServiceExposureTypeNodePort ServiceExposureType = "NodePort"
)

// ServiceExposureSpec defines the service to be exposed
type ServiceExposureSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:Enum=NodePort
	Type     ServiceExposureType `json:"type"`
	NodePort *NodePortSpec       `json:"nodePort,omitempty"`
}

// NodePortSpec defines the ports to be proxied from node to service
type NodePortSpec struct {
	Ports []PortSpec `json:"ports"`
}

// PortSpec defines the service port
type PortSpec struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	NodePort int32 `json:"nodePort"`
}

// KeycodeSpec contains data for keycode check
type KeycodeSpec struct {
	// CodeNumber provides user to apply keycode to Federator.ai
	CodeNumber string `json:"codeNumber"`
	// SignatureData provides user to apply signature data which is download from ProphetStor to Federator.ai
	SignatureData string `json:"signatureData,omitempty"`
}

// KeycodeState defines type of keycode processing state
type KeycodeState string

var (
	// KeycodeStateDefault represents default state
	KeycodeStateDefault KeycodeState
	// KeycodeStateWaitingKeycode represents state in waiting keycode to be filled in
	KeycodeStateWaitingKeycode KeycodeState = "WaitingKeycode"
	// KeycodeStatePollingRegistrationData represents in poll registration data state
	KeycodeStatePollingRegistrationData KeycodeState = "PollingRegistrationData"
	// KeycodeStateWaitingSignatureData represents state waiting user fill in signature data
	KeycodeStateWaitingSignatureData KeycodeState = "WaitingSignatureData"
	// KeycodeStateDone represents state waiting user fill in signature data
	KeycodeStateDone KeycodeState = "Done"
)

// KeycodeStatus contains current keycode information
type KeycodeStatus struct {
	// CodeNumber represents the last keycode user successfully applied
	CodeNumber string `json:"codeNumber"`
	// RegistrationData contains data that user need to send to ProphetStor to activate keycode
	RegistrationData string `json:"registrationData"`
	// State represents the state of keycode processing
	State KeycodeState `json:"state"`
	// LastErrorMessage stores the error message that happend when Federatorai-Operator handled keycode
	LastErrorMessage string `json:"lastErrorMessage"`
	// Summary stores the summary of the keycode
	Summary string `json:"summary"`
}

// AlamedaServiceStatus defines the observed state of AlamedaService
// +k8s:openapi-gen=true
type AlamedaServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	CRDVersion AlamedaServiceStatusCRDVersion `json:"crdversion"`
	// +nullable
	Conditions    []AlamedaServiceStatusCondition `json:"conditions"`
	KeycodeStatus KeycodeStatus                   `json:"keycodeStatus"`
}

type AlamedaServiceStatusCRDVersion struct {

	// Represents whether any actions on the underlaying managed objects are
	// being performed. Only delete actions will be performed.
	ChangeVersion bool   `json:"-"`
	ScalerVersion string `json:"scalerversion"`
	CRDName       string `json:"crdname"`
}

type AlamedaServiceStatusCondition struct {

	// Represents whether any actions on the underlaying managed objects are
	// being performed. Only delete actions will be performed.
	Paused  bool   `json:"paused"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlamedaService is the Schema for the alamedaservices API
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="Execution",type="boolean",JSONPath=".spec.enableExecution",description="The enable of execution"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.version",description="The type of alameda service"
// +kubebuilder:printcolumn:name="PROMETHEUS",type="string",JSONPath=".spec.prometheusService",description="The URL of prometheus"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The time of creation"
type AlamedaService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlamedaServiceSpec   `json:"spec,omitempty"`
	Status AlamedaServiceStatus `json:"status,omitempty"`
}

// SetDefaultValue sets default value into AlamedaServie
func (as *AlamedaService) SetDefaultValue() {
	if as == nil {
		as = &AlamedaService{}
	}

	if as.Spec.EnableDispatcher == nil {
		enable := true
		as.Spec.EnableDispatcher = &enable
	}

	if as.Spec.EnableVPA == nil {
		enable := true
		as.Spec.EnableVPA = &enable
	}

	if as.Spec.EnableGPU == nil {
		enable := false
		as.Spec.EnableGPU = &enable
	}
}

// IsCodeNumberEmpty returns true if keycode is empty
func (as *AlamedaService) IsCodeNumberEmpty() bool {

	if as.Spec.Keycode.CodeNumber == "" {
		return true
	}

	return false
}

// IsCodeNumberUpdated returns true if current keycode is not equal to previous keycode
func (as *AlamedaService) IsCodeNumberUpdated() bool {

	if as.Spec.Keycode.CodeNumber != as.Status.KeycodeStatus.CodeNumber {
		return true
	}

	return false
}

// SetCRDVersion sets crdVersion into AlamedaService's status
func (as *AlamedaService) SetCRDVersion(crdVer AlamedaServiceStatusCRDVersion) {
	as.Status.CRDVersion = crdVer
}

// SetStatusCodeNumber sets codeNumber into AlamedaService's status
func (as *AlamedaService) SetStatusCodeNumber(codeNumber string) {
	as.Status.KeycodeStatus.CodeNumber = codeNumber
}

// SetStatusKeycode sets keycode status into AlamedaService's status
func (as *AlamedaService) SetStatusKeycode(status KeycodeStatus) {
	as.Status.KeycodeStatus = status
}

// SetStatusRegistrationData sets registration data into AlamedaService's status
func (as *AlamedaService) SetStatusRegistrationData(registrationData string) {
	as.Status.KeycodeStatus.RegistrationData = registrationData
}

// SetStatusKeycodeState sets registration data into AlamedaService's status
func (as *AlamedaService) SetStatusKeycodeState(state KeycodeState) {
	as.Status.KeycodeStatus.State = state
}

// SetStatusKeycodeLastErrorMessage sets last error message into AlamedaService's keycode status
func (as *AlamedaService) SetStatusKeycodeLastErrorMessage(msg string) {
	as.Status.KeycodeStatus.LastErrorMessage = msg
}

// SetStatusKeycodeSummary sets keycode summary into AlamedaService's status
func (as *AlamedaService) SetStatusKeycodeSummary(summary string) {
	as.Status.KeycodeStatus.Summary = summary
}

// GetSpecAnnotationWithoutKeycode sets keycode summary into AlamedaService's status
func (as AlamedaService) GetSpecAnnotationWithoutKeycode() (string, error) {
	as.Spec.Keycode = KeycodeSpec{}
	jsonSpec, err := json.Marshal(as.Spec)
	if err != nil {
		return "", err
	}
	return string(jsonSpec), nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlamedaServiceList contains a list of AlamedaService
type AlamedaServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlamedaService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlamedaService{}, &AlamedaServiceList{})
}
