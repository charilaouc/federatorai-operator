package entities

type Kind string

const (
	KindUndefined    Kind = "KIND_UNDEFINED"
	Deployment       Kind = "DEPLOYMENT"
	DeploymentConfig Kind = "DEPLOYMENTCONFIG"
	StatefulSet      Kind = "STATEFULSET"
	AlamedaScaler    Kind = "ALAMEDASCALER"
)

type ScalingTool string

const (
	ScalingToolUndefined ScalingTool = "SCALING_TOOL_UNDEFINED"
	None                 ScalingTool = "NONE"
	VPA                  ScalingTool = "VPA"
	HPA                  ScalingTool = "HPA"
)

type Policy string

const (
	PolicyUndefined Policy = "RECOMMENDATION_POLICY_UNDEFINED"
	Stable          Policy = "STABLE"
	Compact         Policy = "COMPACT"
)

type Boundary string

const (
	BoundaryUndefined Boundary = "undefined"
	Raw               Boundary = "raw"
	UpperBound        Boundary = "upper_bound"
	LowerBound        Boundary = "lower_bound"
)

type MetricType string

const (
	CPUMilliCoresUsage MetricType = "cpu_usage_seconds_percentage"
	MemoryBytesUsage   MetricType = "memory_usage_bytes"
)

type PlanningType string

const (
	PlanningUndefined PlanningType = "PT_UNDEFINED"
	PTRecommendation  PlanningType = "PT_RECOMMENDATION"
	PTPlanning        PlanningType = "PT_PLANNING"
)

type RecommendationType string

const (
	RecommendationUndefined RecommendationType = "CRT_UNDEFINED"
	Primitive               RecommendationType = "PRIMITIVE"
	K8s                     RecommendationType = "K8S"
)

type PodPhase string

const (
	PodPhaseUndefined PodPhase = "POD_PHASE_UNDEFINED"
	Pending           PodPhase = "PENDING"
	Running           PodPhase = "RUNNING"
	Succeeded         PodPhase = "SUCCEEDED"
	Failed            PodPhase = "FAILED"
	Unknown           PodPhase = "UNKNOWN"
	Completed         PodPhase = "COMPLETED"
	CrashLoopBackOff  PodPhase = "CRASHLOOPBACKOFF"
)
