package entities

import (
	"time"
)

type PredictionKafkaTopicCurrentOffset struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"topic"`
	Measurement   *Measurement `name:"kafka_topic_partition_current_offset" metric:"current_offset" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionKafkaTopicCurrentOffsetUpperBound struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"topic"`
	Measurement   *Measurement `name:"kafka_topic_partition_current_offset_upper_bound" metric:"current_offset" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionKafkaTopicCurrentOffsetLowerBound struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"topic"`
	Measurement   *Measurement `name:"kafka_topic_partition_current_offset_lower_bound" metric:"current_offset" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionKafkaConsumerGroupCurrentOffset struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"consumer_group"`
	Measurement   *Measurement `name:"kafka_consumer_group_current_offset" metric:"current_offset" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	TopicName     string       `json:"topic_name"    required:"true"  column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionKafkaConsumerGroupCurrentOffsetUpperBound struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"consumer_group"`
	Measurement   *Measurement `name:"kafka_consumer_group_current_offset_upper_bound" metric:"current_offset" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	TopicName     string       `json:"topic_name"    required:"true"  column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionKafkaConsumerGroupCurrentOffsetLowerBound struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"consumer_group"`
	Measurement   *Measurement `name:"kafka_consumer_group_current_offset_lower_bound" metric:"current_offset" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	TopicName     string       `json:"topic_name"    required:"true"  column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusApplicationCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application" metric:"cpu_millicores_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusApplicationCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application" metric:"cpu_millicores_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusApplicationCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application" metric:"cpu_millicores_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusApplicationMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application" metric:"memory_bytes_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusApplicationMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application" metric:"memory_bytes_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusApplicationMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application" metric:"memory_bytes_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusClusterCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster"`
	Measurement   *Measurement `name:"cluster" metric:"cpu_millicores_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusClusterCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster"`
	Measurement   *Measurement `name:"cluster" metric:"cpu_millicores_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusClusterCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster"`
	Measurement   *Measurement `name:"cluster" metric:"cpu_millicores_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusClusterMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster"`
	Measurement   *Measurement `name:"cluster" metric:"memory_bytes_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusClusterMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster"`
	Measurement   *Measurement `name:"cluster" metric:"memory_bytes_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusClusterMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster"`
	Measurement   *Measurement `name:"cluster" metric:"memory_bytes_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusContainerCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container" metric:"cpu_millicores_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusContainerCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container" metric:"cpu_millicores_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusContainerCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container" metric:"cpu_millicores_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusContainerMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container" metric:"memory_bytes_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusContainerMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container" metric:"memory_bytes_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusContainerMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container"`
	Measurement   *Measurement `name:"container" metric:"memory_bytes_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	PodName       string       `json:"pod_name"      required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"     required:"true"  column:"tag"`
	NodeName      string       `json:"node_name"     required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusControllerCPU struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller"`
	Measurement    *Measurement `name:"controller" metric:"cpu_millicores_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"            required:"false" column:"tag"`
	Name           string       `json:"name"            required:"true"  column:"tag"`
	Namespace      string       `json:"namespace"       required:"true"  column:"tag"`
	ClusterName    string       `json:"cluster_name"    required:"true"  column:"tag"`
	MetricType     MetricType   `json:"metric"          required:"false" column:"tag"`
	Boundary       Boundary     `json:"kind"            required:"false" column:"tag"`
	Granularity    string       `json:"granularity"     required:"true"  column:"tag"`
	ControllerKind string       `json:"controller_kind" required:"true"  column:"tag"`
	ModelId        string       `json:"model_id"        required:"true"  column:"field"`
	PredictionId   string       `json:"prediction_id"   required:"true"  column:"field"`
	Value          float64      `json:"value"           required:"true"  column:"field"`
}

type PredictionClusterStatusControllerCPUUpperBound struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller"`
	Measurement    *Measurement `name:"controller" metric:"cpu_millicores_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"            required:"false" column:"tag"`
	Name           string       `json:"name"            required:"true"  column:"tag"`
	Namespace      string       `json:"namespace"       required:"true"  column:"tag"`
	ClusterName    string       `json:"cluster_name"    required:"true"  column:"tag"`
	MetricType     MetricType   `json:"metric"          required:"false" column:"tag"`
	Boundary       Boundary     `json:"kind"            required:"false" column:"tag"`
	Granularity    string       `json:"granularity"     required:"true"  column:"tag"`
	ControllerKind string       `json:"controller_kind" required:"true"  column:"tag"`
	ModelId        string       `json:"model_id"        required:"true"  column:"field"`
	PredictionId   string       `json:"prediction_id"   required:"true"  column:"field"`
	Value          float64      `json:"value"           required:"true"  column:"field"`
}

type PredictionClusterStatusControllerCPULowerBound struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller"`
	Measurement    *Measurement `name:"controller" metric:"cpu_millicores_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"            required:"false" column:"tag"`
	Name           string       `json:"name"            required:"true"  column:"tag"`
	Namespace      string       `json:"namespace"       required:"true"  column:"tag"`
	ClusterName    string       `json:"cluster_name"    required:"true"  column:"tag"`
	MetricType     MetricType   `json:"metric"          required:"false" column:"tag"`
	Boundary       Boundary     `json:"kind"            required:"false" column:"tag"`
	Granularity    string       `json:"granularity"     required:"true"  column:"tag"`
	ControllerKind string       `json:"controller_kind" required:"true"  column:"tag"`
	ModelId        string       `json:"model_id"        required:"true"  column:"field"`
	PredictionId   string       `json:"prediction_id"   required:"true"  column:"field"`
	Value          float64      `json:"value"           required:"true"  column:"field"`
}

type PredictionClusterStatusControllerMemory struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller"`
	Measurement    *Measurement `name:"controller" metric:"memory_bytes_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"            required:"false" column:"tag"`
	Name           string       `json:"name"            required:"true"  column:"tag"`
	Namespace      string       `json:"namespace"       required:"true"  column:"tag"`
	ClusterName    string       `json:"cluster_name"    required:"true"  column:"tag"`
	MetricType     MetricType   `json:"metric"          required:"false" column:"tag"`
	Boundary       Boundary     `json:"kind"            required:"false" column:"tag"`
	Granularity    string       `json:"granularity"     required:"true"  column:"tag"`
	ControllerKind string       `json:"controller_kind" required:"true"  column:"tag"`
	ModelId        string       `json:"model_id"        required:"true"  column:"field"`
	PredictionId   string       `json:"prediction_id"   required:"true"  column:"field"`
	Value          float64      `json:"value"           required:"true"  column:"field"`
}

type PredictionClusterStatusControllerMemoryUpperBound struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller"`
	Measurement    *Measurement `name:"controller" metric:"memory_bytes_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"            required:"false" column:"tag"`
	Name           string       `json:"name"            required:"true"  column:"tag"`
	Namespace      string       `json:"namespace"       required:"true"  column:"tag"`
	ClusterName    string       `json:"cluster_name"    required:"true"  column:"tag"`
	MetricType     MetricType   `json:"metric"          required:"false" column:"tag"`
	Boundary       Boundary     `json:"kind"            required:"false" column:"tag"`
	Granularity    string       `json:"granularity"     required:"true"  column:"tag"`
	ControllerKind string       `json:"controller_kind" required:"true"  column:"tag"`
	ModelId        string       `json:"model_id"        required:"true"  column:"field"`
	PredictionId   string       `json:"prediction_id"   required:"true"  column:"field"`
	Value          float64      `json:"value"           required:"true"  column:"field"`
}

type PredictionClusterStatusControllerMemoryLowerBound struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller"`
	Measurement    *Measurement `name:"controller" metric:"memory_bytes_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"            required:"false" column:"tag"`
	Name           string       `json:"name"            required:"true"  column:"tag"`
	Namespace      string       `json:"namespace"       required:"true"  column:"tag"`
	ClusterName    string       `json:"cluster_name"    required:"true"  column:"tag"`
	MetricType     MetricType   `json:"metric"          required:"false" column:"tag"`
	Boundary       Boundary     `json:"kind"            required:"false" column:"tag"`
	Granularity    string       `json:"granularity"     required:"true"  column:"tag"`
	ControllerKind string       `json:"controller_kind" required:"true"  column:"tag"`
	ModelId        string       `json:"model_id"        required:"true"  column:"field"`
	PredictionId   string       `json:"prediction_id"   required:"true"  column:"field"`
	Value          float64      `json:"value"           required:"true"  column:"field"`
}

type PredictionClusterStatusNamespaceCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace" metric:"cpu_millicores_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNamespaceCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace" metric:"cpu_millicores_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNamespaceCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace" metric:"cpu_millicores_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNamespaceMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace" metric:"memory_bytes_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNamespaceMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace" metric:"memory_bytes_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNamespaceMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace" metric:"memory_bytes_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNodeCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node" metric:"cpu_millicores_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	IsScheduled   string       `json:"is_scheduled"  required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNodeCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node" metric:"cpu_millicores_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	IsScheduled   string       `json:"is_scheduled"  required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNodeCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node" metric:"cpu_millicores_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	IsScheduled   string       `json:"is_scheduled"  required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNodeMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node" metric:"memory_bytes_usage" boundary:"raw" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	IsScheduled   string       `json:"is_scheduled"  required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNodeMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node" metric:"memory_bytes_usage" boundary:"upper_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	IsScheduled   string       `json:"is_scheduled"  required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}

type PredictionClusterStatusNodeMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node"`
	Measurement   *Measurement `name:"node" metric:"memory_bytes_usage" boundary:"lower_bound" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"          required:"false" column:"tag"`
	Name          string       `json:"name"          required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name"  required:"true"  column:"tag"`
	MetricType    MetricType   `json:"metric"        required:"false" column:"tag"`
	Boundary      Boundary     `json:"kind"          required:"false" column:"tag"`
	Granularity   string       `json:"granularity"   required:"true"  column:"tag"`
	IsScheduled   string       `json:"is_scheduled"  required:"true"  column:"tag"`
	ModelId       string       `json:"model_id"      required:"true"  column:"field"`
	PredictionId  string       `json:"prediction_id" required:"true"  column:"field"`
	Value         float64      `json:"value"         required:"true"  column:"field"`
}
