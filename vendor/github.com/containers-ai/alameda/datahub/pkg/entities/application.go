package entities

import (
	"time"
)

type ApplicationKafkaTopic struct {
	DatahubEntity          `scope:"application" category:"kafka" type:"topic"`
	Measurement            *Measurement `name:"kafka_topic" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                   *time.Time   `json:"time"                     required:"false" column:"tag"`
	Name                   string       `json:"name"                     required:"true"  column:"tag"`
	Namespace              string       `json:"namespace"                required:"true"  column:"tag"`
	ClusterName            string       `json:"cluster_name"             required:"true"  column:"tag"`
	AlamedaScalerName      string       `json:"alameda_scaler_name"      required:"true"  column:"tag"`
	AlamedaScalerNamespace string       `json:"alameda_scaler_namespace" required:"true"  column:"tag"`
	Dummy                  string       `json:"dummy"                    required:"true"  column:"field"`
}

type ApplicationKafkaConsumerGroup struct {
	DatahubEntity            `scope:"application" category:"kafka" type:"consumer_group"`
	Measurement              *Measurement `name:"kafka_consumer_group" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                     *time.Time   `json:"time"                        required:"false" column:"tag"`
	Name                     string       `json:"name"                        required:"true"  column:"tag"`
	Namespace                string       `json:"namespace"                   required:"true"  column:"tag"`
	ClusterName              string       `json:"cluster_name"                required:"true"  column:"tag"`
	TopicName                string       `json:"topic_name"                  required:"true"  column:"tag"`
	AlamedaScalerName        string       `json:"alameda_scaler_name"         required:"true"  column:"tag"`
	AlamedaScalerNamespace   string       `json:"alameda_scaler_namespace"    required:"true"  column:"tag"`
	AlamedaScalerScalingTool ScalingTool  `json:"alameda_scaler_scaling_tool" required:"true"  column:"tag"`
	ResourceK8sName          string       `json:"resource_k8s_name"           required:"true"  column:"tag"`
	ResourceK8sNamespace     string       `json:"resource_k8s_namespace"      required:"true"  column:"tag"`
	ResourceK8sKind          Kind         `json:"resource_k8s_kind"           required:"true"  column:"tag"`
	ResourceK8sSpecReplicas  int32        `json:"resource_k8s_spec_replicas"  required:"true"  column:"field"`
	ResourceK8sReplicas      int32        `json:"resource_k8s_replicas"       required:"false" column:"field"`
	ResourceK8sMinReplicas   int32        `json:"resource_k8s_min_replicas"   required:"false" column:"field"`
	ResourceK8sMaxReplicas   int32        `json:"resource_k8s_max_replicas"   required:"false" column:"field"`
	ResourceCustomName       string       `json:"resource_custom_name"        required:"false" column:"field"`
	Policy                   Policy       `json:"policy"                      required:"false" column:"field"`
	EnableExecution          bool         `json:"enable_execution"            required:"false" column:"field"`
	ResourceCPULimit         string       `json:"resource_cpu_limit"          required:"false" column:"field"`
	ResourceCPURequest       string       `json:"resource_cpu_request"        required:"false" column:"field"`
	ResourceMemoryLimit      string       `json:"resource_memory_limit"       required:"false" column:"field"`
	ResourceMemoryRequest    string       `json:"resource_memory_request"     required:"false" column:"field"`
	VolumesSize              string       `json:"volumes_size"                required:"false" column:"field"`
	VolumesPvcSize           string       `json:"volumes_pvc_size"            required:"false" column:"field"`
}
