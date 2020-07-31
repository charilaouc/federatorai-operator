package entities

import (
	"time"
)

type PlanningClusterStatusApplication struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"application"`
	Measurement                  *Measurement `name:"application" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time   `json:"time"                            required:"false" column:"tag"`
	Name                         string       `json:"name"                            required:"true"  column:"tag"`
	Namespace                    string       `json:"namespace"                       required:"true"  column:"tag"`
	ClusterName                  string       `json:"cluster_name"                    required:"true"  column:"tag"`
	PlanningId                   string       `json:"planning_id"                     required:"true"  column:"tag"`
	PlanningType                 PlanningType `json:"planning_type"                   required:"true"  column:"tag"`
	Granularity                  string       `json:"granularity"                     required:"true"  column:"tag"`
	ResourceRequestCPU           float64      `json:"resource_request_cpu"            required:"false" column:"field"`
	ResourceRequestMemory        float64      `json:"resource_request_memory"         required:"false" column:"field"`
	ResourceLimitCPU             float64      `json:"resource_limit_cpu"              required:"false" column:"field"`
	ResourceLimitMemory          float64      `json:"resource_limit_memory"           required:"false" column:"field"`
	InitialResourceRequestCPU    float64      `json:"initial_resource_request_cpu"    required:"false" column:"field"`
	InitialResourceRequestMemory float64      `json:"initial_resource_request_memory" required:"false" column:"field"`
	InitialResourceLimitCPU      float64      `json:"initial_resource_limit_cpu"      required:"false" column:"field"`
	InitialResourceLimitMemory   float64      `json:"initial_resource_limit_memory"   required:"false" column:"field"`
	StartTime                    int64        `json:"start_time"                      required:"false" column:"field"`
	EndTime                      int64        `json:"end_time"                        required:"false" column:"field"`
	TotalCost                    float64      `json:"total_cost"                      required:"false" column:"field"`
	ApplyPlanningNow             bool         `json:"apply_planning_now"              required:"false" column:"field"`
}

type PlanningClusterStatusCluster struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"cluster"`
	Measurement                  *Measurement `name:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time   `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string       `json:"name"                            required:"true"  column:"tag"`
	PlanningId                   string       `json:"planning_id"                     required:"true"  column:"tag"`
	PlanningType                 PlanningType `json:"planning_type"                   required:"true"  column:"tag"`
	Granularity                  string       `json:"granularity"                     required:"true"  column:"tag"`
	ResourceRequestCPU           float64      `json:"resource_request_cpu"            required:"false" column:"field"`
	ResourceRequestMemory        float64      `json:"resource_request_memory"         required:"false" column:"field"`
	ResourceLimitCPU             float64      `json:"resource_limit_cpu"              required:"false" column:"field"`
	ResourceLimitMemory          float64      `json:"resource_limit_memory"           required:"false" column:"field"`
	InitialResourceRequestCPU    float64      `json:"initial_resource_request_cpu"    required:"false" column:"field"`
	InitialResourceRequestMemory float64      `json:"initial_resource_request_memory" required:"false" column:"field"`
	InitialResourceLimitCPU      float64      `json:"initial_resource_limit_cpu"      required:"false" column:"field"`
	InitialResourceLimitMemory   float64      `json:"initial_resource_limit_memory"   required:"false" column:"field"`
	StartTime                    int64        `json:"start_time"                      required:"false" column:"field"`
	EndTime                      int64        `json:"end_time"                        required:"false" column:"field"`
	TotalCost                    float64      `json:"total_cost"                      required:"false" column:"field"`
	ApplyPlanningNow             bool         `json:"apply_planning_now"              required:"false" column:"field"`
}

type PlanningClusterStatusContainer struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"container"`
	Measurement                  *Measurement `name:"container" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time   `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string       `json:"name"                            required:"true"  column:"tag"`
	PodName                      string       `json:"pod_name"                        required:"true"  column:"tag"`
	Namespace                    string       `json:"namespace"                       required:"true"  column:"tag"`
	ClusterName                  string       `json:"cluster_name"                    required:"true"  column:"tag"`
	PlanningId                   string       `json:"planning_id"                     required:"true"  column:"tag"`
	PlanningType                 PlanningType `json:"planning_type"                   required:"true"  column:"tag"`
	Granularity                  string       `json:"granularity"                     required:"true"  column:"tag"`
	TopControllerName            string       `json:"top_controller_name"             required:"false" column:"field"`
	TopControllerKind            Kind         `json:"top_controller_kind"             required:"false" column:"field"`
	Policy                       string       `json:"policy"                          required:"false" column:"field"`
	PolicyTime                   int64        `json:"policy_time"                     required:"false" column:"field"`
	ResourceRequestCPU           float64      `json:"resource_request_cpu"            required:"false" column:"field"`
	ResourceRequestMemory        float64      `json:"resource_request_memory"         required:"false" column:"field"`
	ResourceLimitCPU             float64      `json:"resource_limit_cpu"              required:"false" column:"field"`
	ResourceLimitMemory          float64      `json:"resource_limit_memory"           required:"false" column:"field"`
	InitialResourceRequestCPU    float64      `json:"initial_resource_request_cpu"    required:"false" column:"field"`
	InitialResourceRequestMemory float64      `json:"initial_resource_request_memory" required:"false" column:"field"`
	InitialResourceLimitCPU      float64      `json:"initial_resource_limit_cpu"      required:"false" column:"field"`
	InitialResourceLimitMemory   float64      `json:"initial_resource_limit_memory"   required:"false" column:"field"`
	StartTime                    int64        `json:"start_time"                      required:"false" column:"field"`
	EndTime                      int64        `json:"end_time"                        required:"false" column:"field"`
	PodTotalCost                 float64      `json:"pod_total_cost"                  required:"false" column:"field"`
}

type PlanningClusterStatusController struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"controller"`
	Measurement                  *Measurement `name:"controller" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time   `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string       `json:"name"                            required:"true"  column:"tag"`
	Namespace                    string       `json:"namespace"                       required:"true"  column:"tag"`
	ClusterName                  string       `json:"cluster_name"                    required:"true"  column:"tag"`
	PlanningId                   string       `json:"planning_id"                     required:"true"  column:"tag"`
	PlanningType                 PlanningType `json:"planning_type"                   required:"true"  column:"tag"`
	Kind                         Kind         `json:"kind"                            required:"true"  column:"tag"`
	Granularity                  string       `json:"granularity"                     required:"true"  column:"tag"`
	ResourceRequestCPU           float64      `json:"resource_request_cpu"            required:"false" column:"field"`
	ResourceRequestMemory        float64      `json:"resource_request_memory"         required:"false" column:"field"`
	ResourceLimitCPU             float64      `json:"resource_limit_cpu"              required:"false" column:"field"`
	ResourceLimitMemory          float64      `json:"resource_limit_memory"           required:"false" column:"field"`
	InitialResourceRequestCPU    float64      `json:"initial_resource_request_cpu"    required:"false" column:"field"`
	InitialResourceRequestMemory float64      `json:"initial_resource_request_memory" required:"false" column:"field"`
	InitialResourceLimitCPU      float64      `json:"initial_resource_limit_cpu"      required:"false" column:"field"`
	InitialResourceLimitMemory   float64      `json:"initial_resource_limit_memory"   required:"false" column:"field"`
	StartTime                    int64        `json:"start_time"                      required:"false" column:"field"`
	EndTime                      int64        `json:"end_time"                        required:"false" column:"field"`
	TotalCost                    float64      `json:"total_cost"                      required:"false" column:"field"`
	ApplyPlanningNow             bool         `json:"apply_planning_now"              required:"false" column:"field"`
}

type PlanningClusterStatusNamespace struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"namespace"`
	Measurement                  *Measurement `name:"namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time   `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string       `json:"name"                            required:"true"  column:"tag"`
	ClusterName                  string       `json:"cluster_name"                    required:"true"  column:"tag"`
	PlanningId                   string       `json:"planning_id"                     required:"true"  column:"tag"`
	PlanningType                 PlanningType `json:"planning_type"                   required:"true"  column:"tag"`
	Granularity                  string       `json:"granularity"                     required:"true"  column:"tag"`
	ResourceRequestCPU           float64      `json:"resource_request_cpu"            required:"false" column:"field"`
	ResourceRequestMemory        float64      `json:"resource_request_memory"         required:"false" column:"field"`
	ResourceLimitCPU             float64      `json:"resource_limit_cpu"              required:"false" column:"field"`
	ResourceLimitMemory          float64      `json:"resource_limit_memory"           required:"false" column:"field"`
	InitialResourceRequestCPU    float64      `json:"initial_resource_request_cpu"    required:"false" column:"field"`
	InitialResourceRequestMemory float64      `json:"initial_resource_request_memory" required:"false" column:"field"`
	InitialResourceLimitCPU      float64      `json:"initial_resource_limit_cpu"      required:"false" column:"field"`
	InitialResourceLimitMemory   float64      `json:"initial_resource_limit_memory"   required:"false" column:"field"`
	StartTime                    int64        `json:"start_time"                      required:"false" column:"field"`
	EndTime                      int64        `json:"end_time"                        required:"false" column:"field"`
	TotalCost                    float64      `json:"total_cost"                      required:"false" column:"field"`
	ApplyPlanningNow             bool         `json:"apply_planning_now"              required:"false" column:"field"`
}

type PlanningClusterStatusNode struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"node"`
	Measurement                  *Measurement `name:"node" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time   `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string       `json:"name"                            required:"true"  column:"tag"`
	ClusterName                  string       `json:"cluster_name"                    required:"true"  column:"tag"`
	PlanningId                   string       `json:"planning_id"                     required:"true"  column:"tag"`
	PlanningType                 PlanningType `json:"planning_type"                   required:"true"  column:"tag"`
	Granularity                  string       `json:"granularity"                     required:"true"  column:"tag"`
	ResourceRequestCPU           float64      `json:"resource_request_cpu"            required:"false" column:"field"`
	ResourceRequestMemory        float64      `json:"resource_request_memory"         required:"false" column:"field"`
	ResourceLimitCPU             float64      `json:"resource_limit_cpu"              required:"false" column:"field"`
	ResourceLimitMemory          float64      `json:"resource_limit_memory"           required:"false" column:"field"`
	InitialResourceRequestCPU    float64      `json:"initial_resource_request_cpu"    required:"false" column:"field"`
	InitialResourceRequestMemory float64      `json:"initial_resource_request_memory" required:"false" column:"field"`
	InitialResourceLimitCPU      float64      `json:"initial_resource_limit_cpu"      required:"false" column:"field"`
	InitialResourceLimitMemory   float64      `json:"initial_resource_limit_memory"   required:"false" column:"field"`
	StartTime                    int64        `json:"start_time"                      required:"false" column:"field"`
	EndTime                      int64        `json:"end_time"                        required:"false" column:"field"`
	TotalCost                    float64      `json:"total_cost"                      required:"false" column:"field"`
	ApplyPlanningNow             bool         `json:"apply_planning_now"              required:"false" column:"field"`
}
