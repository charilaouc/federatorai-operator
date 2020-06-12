package component

const (
	defaultImageAdmissionController       = "quay.io/prophetstor/alameda-admission-ubi:latest"
	defaultImageAIDispatcher              = "quay.io/prophetstor/alameda-ai-dispatcher:latest"
	defaultImageAIEngine                  = "quay.io/prophetstor/alameda-ai:latest"
	defaultImageAnalyzer                  = "quay.io/prophetstor/alameda-analyzer-ubi:latest"
	defaultImageDatahub                   = "quay.io/prophetstor/alameda-datahub-ubi:latest"
	defaultImageEvictioner                = "quay.io/prophetstor/alameda-evictioner-ubi:latest"
	defaultImageExecutor                  = "quay.io/prophetstor/alameda-executor-ubi:latest"
	defaultImageAlpine                    = "alpine"
	defaultImageInfluxDB                  = "influxdb:1.7-alpine"
	defaultImageNotifier                  = "quay.io/prophetstor/alameda-notifier-ubi:latest"
	defaultImageOperator                  = "quay.io/prophetstor/alameda-operator-ubi:latest"
	defaultImageRabbitMQ                  = "quay.io/prophetstor/alameda-rabbitmq:latest"
	defaultImageRecommender               = "quay.io/prophetstor/alameda-recommender-ubi:latest"
	defaultImageFedemeterAPI              = "quay.io/prophetstor/fedemeter-api-ubi:latest"
	defaultImageFederatoraiAgentGPU       = "quay.io/prophetstor/federatorai-agent-gpu:latest"
	defaultImageFederatoraiAgentPreloader = "quay.io/prophetstor/federatorai-agent-preloader:latest"
	defaultImageFederatoraiAgent          = "quay.io/prophetstor/federatorai-agent-ubi:latest"
	defaultImageFederatoraiRestAPI        = "quay.io/prophetstor/federatorai-rest-ubi:latest"
	defaultImageFedemeterInfluxdb         = "quay.io/prophetstor/fedemeter-influxdb:latest"
	defaultImageFrontend                  = "quay.io/prophetstor/federatorai-dashboard-frontend:latest"
	defaultImageBackend                   = "quay.io/prophetstor/federatorai-dashboard-backend:latest"
	defaultImageFederatoraiAgentApp       = "quay.io/prophetstor/federatorai-agent-app:latest"
)

type ImageConfig struct {
	AdmissionController       string
	AIDispatcher              string
	AIEngine                  string
	Analyzer                  string
	Datahub                   string
	Evictioner                string
	Executor                  string
	Alpine                    string
	InfluxDB                  string
	Notifier                  string
	Operator                  string
	RabbitMQ                  string
	Recommender               string
	FedemeterAPI              string
	FederatoraiAgentGPU       string
	FederatoraiAgentPreloader string
	FederatoraiAgent          string
	FederatoraiAgentApp       string
	FederatoraiRestAPI        string
	FedemeterInfluxDB         string
	DashboardFrontend         string
	DashboardBackend          string
}

// NewDefautlImageConfig returns ImageConfig with default value
func NewDefautlImageConfig() ImageConfig {
	return ImageConfig{
		AdmissionController:       defaultImageAdmissionController,
		AIDispatcher:              defaultImageAIDispatcher,
		AIEngine:                  defaultImageAIEngine,
		Analyzer:                  defaultImageAnalyzer,
		Datahub:                   defaultImageDatahub,
		Evictioner:                defaultImageEvictioner,
		Executor:                  defaultImageExecutor,
		Alpine:                    defaultImageAlpine,
		InfluxDB:                  defaultImageInfluxDB,
		Notifier:                  defaultImageNotifier,
		Operator:                  defaultImageOperator,
		RabbitMQ:                  defaultImageRabbitMQ,
		Recommender:               defaultImageRecommender,
		FedemeterAPI:              defaultImageFedemeterAPI,
		FederatoraiAgentGPU:       defaultImageFederatoraiAgentGPU,
		FederatoraiAgentPreloader: defaultImageFederatoraiAgentPreloader,
		FederatoraiAgent:          defaultImageFederatoraiAgent,
		FederatoraiAgentApp:       defaultImageFederatoraiAgentApp,
		FederatoraiRestAPI:        defaultImageFederatoraiRestAPI,
		FedemeterInfluxDB:         defaultImageFedemeterInfluxdb,
		DashboardFrontend:         defaultImageFrontend,
		DashboardBackend:          defaultImageBackend,
	}
}

// SetAdmissionController sets image to imageConfig
func (i *ImageConfig) SetAdmissionController(image string) {
	i.AdmissionController = image
}

// SetAIDispatcher sets image to imageConfig
func (i *ImageConfig) SetAIDispatcher(image string) {
	i.AIDispatcher = image
}

// SetAIEngine sets image to imageConfig
func (i *ImageConfig) SetAIEngine(image string) {
	i.AIEngine = image
}

// SetAnalyzer sets image to imageConfig
func (i *ImageConfig) SetAnalyzer(image string) {
	i.Analyzer = image
}

// SetDatahub sets image to imageConfig
func (i *ImageConfig) SetDatahub(image string) {
	i.Datahub = image
}

// SetEvictioner sets image to imageConfig
func (i *ImageConfig) SetEvictioner(image string) {
	i.Evictioner = image
}

// SetExecutor sets image to imageConfig
func (i *ImageConfig) SetExecutor(image string) {
	i.Executor = image
}

// SetAlpine sets image to imageConfig
func (i *ImageConfig) SetAlpine(image string) {
	i.Alpine = image
}

// SetInfluxdb sets image to imageConfig
func (i *ImageConfig) SetInfluxdb(image string) {
	i.InfluxDB = image
}

// SetNotifier sets image to imageConfig
func (i *ImageConfig) SetNotifier(image string) {
	i.Notifier = image
}

// SetOperator sets image to imageConfig
func (i *ImageConfig) SetOperator(image string) {
	i.Operator = image
}

// SetRabbitMQ sets image to imageConfig
func (i *ImageConfig) SetRabbitMQ(image string) {
	i.RabbitMQ = image
}

// SetRecommender sets image to imageConfig
func (i *ImageConfig) SetRecommender(image string) {
	i.Recommender = image
}

// SetFedemeterAPI sets image to imageConfig
func (i *ImageConfig) SetFedemeterAPI(image string) {
	i.FedemeterAPI = image
}

// SetFederatoraiAgentGPU sets image to imageConfig
func (i *ImageConfig) SetFederatoraiAgentGPU(image string) {
	i.FederatoraiAgentGPU = image
}

// SetFederatoraiAgentApp sets image to imageConfig
func (i *ImageConfig) SetFederatoraiAgentApp(image string) {
	i.FederatoraiAgentApp = image
}

// SetFederatoraiAgentPreloader sets image to imageConfig
func (i *ImageConfig) SetFederatoraiAgentPreloader(image string) {
	i.FederatoraiAgentPreloader = image
}

// SetFederatoraiAgent sets image to imageConfig
func (i *ImageConfig) SetFederatoraiAgent(image string) {
	i.FederatoraiAgent = image
}

// SetFederatoraiRestAPI sets image to imageConfig
func (i *ImageConfig) SetFederatoraiRestAPI(image string) {
	i.FederatoraiRestAPI = image
}

// SetFedemeterInfluxdb sets image to imageConfig
func (i *ImageConfig) SetFedemeterInfluxdb(image string) {
	i.FedemeterInfluxDB = image
}

// SetFrontEnd sets image to imageConfig
func (i *ImageConfig) SetFrontend(image string) {
	i.DashboardFrontend = image
}

// SetBackEnd sets image to imageConfig
func (i *ImageConfig) SetBackend(image string) {
	i.DashboardBackend = image
}
