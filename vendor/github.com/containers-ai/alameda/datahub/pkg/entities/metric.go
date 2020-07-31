package entities

import (
	"time"
)

type MetricKafkaTopicCurrentOffset struct {
	DatahubEntity `scope:"metric" category:"kafka" type:"topic"`
	Measurement   *Measurement `name:"kafka_topic_partition_current_offset" metric:"current_offset" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"    required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricKafkaConsumerGroupCurrentOffset struct {
	DatahubEntity `scope:"metric" category:"kafka" type:"consumer_group"`
	Measurement   *Measurement `name:"kafka_consumer_group_current_offset" metric:"current_offset" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"    required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	TopicName     string       `json:"topic_name"   required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricKafkaConsumerGroupLag struct {
	DatahubEntity `scope:"metric" category:"kafka" type:"consumer_group"`
	Measurement   *Measurement `name:"kafka_consumer_group_lag" metric:"lag" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"    required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	TopicName     string       `json:"topic_name"   required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusApplicationCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"    required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusApplicationMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"    required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusClusterCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"cluster"`
	Measurement   *Measurement `name:"cluster_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"  required:"false" column:"tag"`
	Name          string       `json:"name"  required:"true"  column:"tag"`
	Uid           string       `json:"uid"   required:"true"  column:"tag"`
	Value         float64      `json:"value" required:"true"  column:"field"`
}

type MetricClusterStatusClusterMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"cluster"`
	Measurement   *Measurement `name:"cluster_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"  required:"false" column:"tag"`
	Name          string       `json:"name"  required:"true"  column:"tag"`
	Uid           string       `json:"uid"   required:"true"  column:"tag"`
	Value         float64      `json:"value" required:"true"  column:"field"`
}

type MetricClusterStatusContainerCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	RateRange     string       `json:"rate_range"    required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type MetricClusterStatusContainerMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	RateRange     string       `json:"rate_range"    required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

// Resource requests cpu millicores
type MetricClusterStatusContainerRscReqCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container_resource_requests_cpu_millicores" metric:"cpu_millicores_total" boundary:"undefined" quota:"request" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

// Resource limits cpu millicores
type MetricClusterStatusContainerRscLimitCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container_resource_limits_cpu_millicores" metric:"cpu_millicores_total" boundary:"undefined" quota:"limit" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

// Resource requests memory bytes
type MetricClusterStatusContainerRscReqMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container_resource_requests_memory_bytes" metric:"memory_bytes_total" boundary:"undefined" quota:"request" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

// Resource limits memory bytes
type MetricClusterStatusContainerRscLimitMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container_resource_limits_memory_bytes" metric:"memory_bytes_total" boundary:"undefined" quota:"limit" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type MetricClusterStatusContainerRestartsTotal struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container_restarts_total" metric:"restarts_total" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type MetricClusterStatusControllerCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"controller"`
	Measurement   *Measurement `name:"controller_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"    required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Kind          Kind         `json:"kind"         required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusControllerMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"controller"`
	Measurement   *Measurement `name:"controller_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"    required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Kind          Kind         `json:"kind"         required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusNamespaceCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true" column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true" column:"tag"`
	Uid           string       `json:"uid"          required:"true" column:"tag"`
	Value         float64      `json:"value"        required:"true" column:"field"`
}

type MetricClusterStatusNamespaceMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusNodeCPUAllocatable struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_cpu_cores_allocatable" metric:"cpu_cores_alloc" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusNodeCPUTotal struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_cpu_millicores_total" metric:"cpu_millicores_total" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusNodeCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusNodeMemoryTotal struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_memory_bytes_total" metric:"memory_bytes_total" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusNodeMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

// Filesystem bytes usage percentage
type MetricClusterStatusNodeFSPCT struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_fs_bytes_usage_pct" metric:"fs_bytes_usage_pct" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Device        string       `json:"device"       required:"false" column:"field"`
	FSType        string       `json:"fs_type"      required:"false" column:"field"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

// Disk I/O utilization
type MetricClusterStatusNodeDiskIOUtil struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_disk_io_util" metric:"disk_io_util" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusNodePodPhaseCount struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_pod_phase_count" metric:"number" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Pending       int64        `json:"pending"      required:"true"  column:"field"`
	Running       int64        `json:"running"      required:"true"  column:"field"`
	Succeeded     int64        `json:"succeeded"    required:"true"  column:"field"`
	Failed        int64        `json:"failed"       required:"true"  column:"field"`
	Unknown       int64        `json:"unknown"      required:"true"  column:"field"`
}

type MetricClusterStatusNodeUnschedulable struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node_unschedulable" metric:"unschedulable" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         float64      `json:"value"        required:"true"  column:"field"`
}

type MetricClusterStatusServiceHealth struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"service"`
	Measurement   *Measurement `name:"service_health" metric:"health" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	Job           string       `json:"job"           required:"true"  column:"tag"`
	Value         int32        `json:"value"         required:"true"  column:"field"` // 1: healthy, 0: failed
}

type MetricTopContainerCPUUsagePCT struct {
	DatahubEntity `scope:"metric" category:"top" type:"container"`
	Measurement   *Measurement `name:"top_container_cpu_millicores_usage_pct" metric:"cpu_millicores_usage_pct" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type MetricTopContainerMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"top" type:"container"`
	Measurement   *Measurement `name:"top_container_memory_bytes_usage" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	PodNamespace  string       `json:"pod_namespace" required:"true"  column:"tag"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}
