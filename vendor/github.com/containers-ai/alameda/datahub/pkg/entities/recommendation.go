package entities

import (
	"time"
)

type RecommendationKafkaConsumerGroup struct {
	DatahubEntity         `scope:"recommendation" category:"kafka" type:"consumer_group"`
	Measurement           *Measurement `name:"kafka_consumer_group" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time   `json:"time"                    required:"false" column:"tag"`
	Name                  string       `json:"name"                    required:"true"  column:"tag"`
	Namespace             string       `json:"namespace"               required:"true"  column:"tag"`
	ClusterName           string       `json:"cluster_name"            required:"true"  column:"tag"`
	ResourceK8sName       string       `json:"resource_k8s_name"       required:"true"  column:"tag"`
	ResourceK8sNamespace  string       `json:"resource_k8s_namespace"  required:"true"  column:"tag"`
	Kind                  Kind         `json:"kind"                    required:"true"  column:"tag"`
	CreateTime            int64        `json:"create_time"             required:"true"  column:"field"`
	CurrentReplicas       int32        `json:"current_replicas"        required:"true"  column:"field"`
	DesiredReplicas       int32        `json:"desired_replicas"        required:"true"  column:"field"`
	PerConsumerCapability float64      `json:"per_consumer_capability" required:"false" column:"field"`
}

type RecommendationClusterStatusApplication struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"application"`
	Measurement           *Measurement       `name:"application" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time         `json:"time"                 required:"false" column:"tag"`
	Name                  string             `json:"name"                 required:"true"  column:"tag"`
	Namespace             string             `json:"namespace"            required:"true"  column:"tag"`
	ClusterName           string             `json:"cluster_name"         required:"true"  column:"tag"`
	Type                  RecommendationType `json:"type"                 required:"true"  column:"tag"`
	Kind                  Kind               `json:"kind"                 required:"true"  column:"field"`
	CurrentReplicas       int32              `json:"current_replicas"     required:"true"  column:"field"`
	DesiredReplicas       int32              `json:"desired_replicas"     required:"true"  column:"field"`
	CreateTime            int64              `json:"create_time"          required:"true"  column:"field"`
	CurrentCPURequests    float64            `json:"current_cpu_requests" required:"false" column:"field"`
	CurrentMemoryRequests float64            `json:"current_mem_requests" required:"false" column:"field"`
	CurrentCPULimits      float64            `json:"current_cpu_limits"   required:"false" column:"field"`
	CurrentMemoryLimits   float64            `json:"current_mem_limits"   required:"false" column:"field"`
	DesiredCPULimits      float64            `json:"desired_cpu_limits"   required:"false" column:"field"`
	DesiredMemoryLimits   float64            `json:"desired_mem_limits"   required:"false" column:"field"`
	TotalCost             float64            `json:"total_cost"           required:"false" column:"field"`
}

type RecommendationClusterStatusCluster struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"cluster"`
	Measurement           *Measurement       `name:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time         `json:"time"                 required:"false" column:"tag"`
	Name                  string             `json:"name"                 required:"true"  column:"tag"`
	Type                  RecommendationType `json:"type"                 required:"true"  column:"tag"`
	Kind                  Kind               `json:"kind"                 required:"true"  column:"field"`
	CurrentReplicas       int32              `json:"current_replicas"     required:"true"  column:"field"`
	DesiredReplicas       int32              `json:"desired_replicas"     required:"true"  column:"field"`
	CreateTime            int64              `json:"create_time"          required:"true"  column:"field"`
	CurrentCPURequests    float64            `json:"current_cpu_requests" required:"false" column:"field"`
	CurrentMemoryRequests float64            `json:"current_mem_requests" required:"false" column:"field"`
	CurrentCPULimits      float64            `json:"current_cpu_limits"   required:"false" column:"field"`
	CurrentMemoryLimits   float64            `json:"current_mem_limits"   required:"false" column:"field"`
	DesiredCPULimits      float64            `json:"desired_cpu_limits"   required:"false" column:"field"`
	DesiredMemoryLimits   float64            `json:"desired_mem_limits"   required:"false" column:"field"`
	TotalCost             float64            `json:"total_cost"           required:"false" column:"field"`
}

type RecommendationClusterStatusContainerLimit struct {
	DatahubEntity              `scope:"recommendation" category:"cluster_status" type:"container"`
	Measurement                *Measurement `name:"container" metric:"undefined" boundary:"undefined" quota:"limit" ts:"true"`
	Time                       *time.Time   `json:"time"                          required:"false" column:"tag"`
	Name                       string       `json:"name"                          required:"true"  column:"tag"`
	Namespace                  string       `json:"namespace"                     required:"true"  column:"tag"`
	ClusterName                string       `json:"cluster_name"                  required:"true"  column:"tag"`
	PodName                    string       `json:"pod_name"                      required:"true"  column:"tag"`
	Granularity                string       `json:"granularity"                   required:"true"  column:"tag"`
	TopControllerName          string       `json:"top_controller_name"           required:"true"  column:"field"`
	TopControllerKind          Kind         `json:"top_controller_kind"           required:"true"  column:"field"`
	Policy                     string       `json:"policy"                        required:"true"  column:"field"`
	PolicyTime                 int64        `json:"policy_time"                   required:"true"  column:"field"`
	PodTotalCost               float64      `json:"pod_total_cost"                required:"true"  column:"field"`
	StartTime                  int64        `json:"start_time"                    required:"true"  column:"field"`
	EndTime                    int64        `json:"end_time"                      required:"true"  column:"field"`
	ResourceLimitCPU           float64      `json:"resource_limit_cpu"            required:"true"  column:"field"`
	ResourceLimitMemory        float64      `json:"resource_limit_memory"         required:"true"  column:"field"`
	InitialResourceLimitCPU    float64      `json:"initial_resource_limit_cpu"    required:"true"  column:"field"`
	InitialResourceLimitMemory float64      `json:"initial_resource_limit_memory" required:"true"  column:"field"`
}

type RecommendationClusterStatusContainerRequest struct {
	DatahubEntity                `scope:"recommendation" category:"cluster_status" type:"container"`
	Measurement                  *Measurement `name:"container" metric:"undefined" boundary:"undefined" quota:"request" ts:"true"`
	Time                         *time.Time   `json:"time"                            required:"false" column:"tag"`
	Name                         string       `json:"name"                            required:"true"  column:"tag"`
	Namespace                    string       `json:"namespace"                       required:"true"  column:"tag"`
	ClusterName                  string       `json:"cluster_name"                    required:"true"  column:"tag"`
	PodName                      string       `json:"pod_name"                        required:"true"  column:"tag"`
	Granularity                  string       `json:"granularity"                     required:"true"  column:"tag"`
	TopControllerName            string       `json:"top_controller_name"             required:"true"  column:"field"`
	TopControllerKind            Kind         `json:"top_controller_kind"             required:"true"  column:"field"`
	Policy                       string       `json:"policy"                          required:"true"  column:"field"`
	PolicyTime                   int64        `json:"policy_time"                     required:"true"  column:"field"`
	PodTotalCost                 float64      `json:"pod_total_cost"                  required:"true"  column:"field"`
	StartTime                    int64        `json:"start_time"                      required:"true"  column:"field"`
	EndTime                      int64        `json:"end_time"                        required:"true"  column:"field"`
	ResourceRequestCPU           float64      `json:"resource_request_cpu"            required:"true"  column:"field"`
	ResourceRequestMemory        float64      `json:"resource_request_memory"         required:"true"  column:"field"`
	InitialResourceRequestCPU    float64      `json:"initial_resource_request_cpu"    required:"true"  column:"field"`
	InitialResourceRequestMemory float64      `json:"initial_resource_request_memory" required:"true"  column:"field"`
}

type RecommendationClusterStatusController struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"controller"`
	Measurement           *Measurement       `name:"controller" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time         `json:"time"                 required:"false" column:"tag"`
	Name                  string             `json:"name"                 required:"true"  column:"tag"`
	Namespace             string             `json:"namespace"            required:"true"  column:"tag"`
	ClusterName           string             `json:"cluster_name"         required:"true"  column:"tag"`
	Type                  RecommendationType `json:"type"                 required:"true"  column:"tag"`
	Kind                  Kind               `json:"kind"                 required:"true"  column:"field"`
	CurrentReplicas       int32              `json:"current_replicas"     required:"true"  column:"field"`
	DesiredReplicas       int32              `json:"desired_replicas"     required:"true"  column:"field"`
	CreateTime            int64              `json:"create_time"          required:"true"  column:"field"`
	CurrentCPURequests    float64            `json:"current_cpu_requests" required:"false" column:"field"`
	CurrentMemoryRequests float64            `json:"current_mem_requests" required:"false" column:"field"`
	CurrentCPULimits      float64            `json:"current_cpu_limits"   required:"false" column:"field"`
	CurrentMemoryLimits   float64            `json:"current_mem_limits"   required:"false" column:"field"`
	DesiredCPULimits      float64            `json:"desired_cpu_limits"   required:"false" column:"field"`
	DesiredMemoryLimits   float64            `json:"desired_mem_limits"   required:"false" column:"field"`
	TotalCost             float64            `json:"total_cost"           required:"false" column:"field"`
}

type RecommendationClusterStatusNamespace struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"namespace"`
	Measurement           *Measurement       `name:"namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time         `json:"time"                 required:"false" column:"tag"`
	Name                  string             `json:"name"                 required:"true"  column:"tag"`
	ClusterName           string             `json:"cluster_name"         required:"true"  column:"tag"`
	Type                  RecommendationType `json:"type"                 required:"true"  column:"tag"`
	Kind                  Kind               `json:"kind"                 required:"true"  column:"field"`
	CurrentReplicas       int32              `json:"current_replicas"     required:"true"  column:"field"`
	DesiredReplicas       int32              `json:"desired_replicas"     required:"true"  column:"field"`
	CreateTime            int64              `json:"create_time"          required:"true"  column:"field"`
	CurrentCPURequests    float64            `json:"current_cpu_requests" required:"false" column:"field"`
	CurrentMemoryRequests float64            `json:"current_mem_requests" required:"false" column:"field"`
	CurrentCPULimits      float64            `json:"current_cpu_limits"   required:"false" column:"field"`
	CurrentMemoryLimits   float64            `json:"current_mem_limits"   required:"false" column:"field"`
	DesiredCPULimits      float64            `json:"desired_cpu_limits"   required:"false" column:"field"`
	DesiredMemoryLimits   float64            `json:"desired_mem_limits"   required:"false" column:"field"`
	TotalCost             float64            `json:"total_cost"           required:"false" column:"field"`
}

type RecommendationClusterStatusNode struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"node"`
	Measurement           *Measurement       `name:"node" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time         `json:"time"                 required:"false" column:"tag"`
	Name                  string             `json:"name"                 required:"true"  column:"tag"`
	ClusterName           string             `json:"cluster_name"         required:"true"  column:"tag"`
	Type                  RecommendationType `json:"type"                 required:"true"  column:"tag"`
	Kind                  Kind               `json:"kind"                 required:"true"  column:"field"`
	CurrentReplicas       int32              `json:"current_replicas"     required:"true"  column:"field"`
	DesiredReplicas       int32              `json:"desired_replicas"     required:"true"  column:"field"`
	CreateTime            int64              `json:"create_time"          required:"true"  column:"field"`
	CurrentCPURequests    float64            `json:"current_cpu_requests" required:"false" column:"field"`
	CurrentMemoryRequests float64            `json:"current_mem_requests" required:"false" column:"field"`
	CurrentCPULimits      float64            `json:"current_cpu_limits"   required:"false" column:"field"`
	CurrentMemoryLimits   float64            `json:"current_mem_limits"   required:"false" column:"field"`
	DesiredCPULimits      float64            `json:"desired_cpu_limits"   required:"false" column:"field"`
	DesiredMemoryLimits   float64            `json:"desired_mem_limits"   required:"false" column:"field"`
	TotalCost             float64            `json:"total_cost"           required:"false" column:"field"`
}
