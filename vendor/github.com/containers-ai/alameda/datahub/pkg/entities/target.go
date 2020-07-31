package entities

import (
	"time"
)

type TargetClusterStatusCluster struct {
	DatahubEntity          `scope:"target" category:"cluster_status" type:"cluster"`
	Measurement            *Measurement `name:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                   *time.Time   `json:"time"                     required:"false" column:"tag"`
	Name                   string       `json:"name"                     required:"true"  column:"tag"`
	AlamedaScalerName      string       `json:"alameda_scaler_name"      required:"true"  column:"tag"`
	AlamedaScalerNamespace string       `json:"alameda_scaler_namespace" required:"true"  column:"tag"`
	RawSpec                string       `json:"raw_spec"                 required:"true"  column:"field"`
}

type TargetClusterStatusController struct {
	DatahubEntity            `scope:"target" category:"cluster_status" type:"controller"`
	Measurement              *Measurement `name:"controller" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                     *time.Time   `json:"time"                        required:"false" column:"tag"`
	Name                     string       `json:"name"                        required:"true"  column:"tag"`
	Namespace                string       `json:"namespace"                   required:"true"  column:"tag"`
	ClusterName              string       `json:"cluster_name"                required:"true"  column:"tag"`
	Kind                     Kind         `json:"kind"                        required:"true"  column:"tag"`
	AlamedaScalerName        string       `json:"alameda_scaler_name"         required:"true"  column:"tag"`
	AlamedaScalerNamespace   string       `json:"alameda_scaler_namespace"    required:"true"  column:"tag"`
	AlamedaScalerScalingTool ScalingTool  `json:"alameda_scaler_scaling_tool" required:"true"  column:"tag"`
	MinReplicas              int32        `json:"resource_k8s_min_replicas"   required:"false" column:"field"`
	MaxReplicas              int32        `json:"resource_k8s_max_replicas"   required:"false" column:"field"`
	Policy                   string       `json:"policy"                      required:"false" column:"field"`
	EnableExecution          bool         `json:"enable_execution"            required:"false" column:"field"`
	RawSpec                  string       `json:"raw_spec"                    required:"false" column:"field"`
}

type TargetKafkaConsumerGroup struct {
	DatahubEntity            `scope:"target" category:"kafka" type:"consumer_group"`
	Measurement              *Measurement `name:"kafka_consumer_group" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                     *time.Time   `json:"time"                        required:"false" column:"tag"`
	Name                     string       `json:"name"                        required:"true"  column:"tag"`
	ClusterName              string       `json:"cluster_name"                required:"true"  column:"tag"`
	ExporterNamespace        string       `json:"exporter_namespace"          required:"true"  column:"tag"`
	TopicName                string       `json:"topic_name"                  required:"true"  column:"tag"`
	AlamedaScalerName        string       `json:"alameda_scaler_name"         required:"true"  column:"tag"`
	AlamedaScalerNamespace   string       `json:"alameda_scaler_namespace"    required:"true"  column:"tag"`
	AlamedaScalerScalingTool ScalingTool  `json:"alameda_scaler_scaling_tool" required:"true"  column:"tag"`
	ResourceK8sName          string       `json:"resource_k8s_name"           required:"true"  column:"tag"`
	ResourceK8sNamespace     string       `json:"resource_k8s_namespace"      required:"true"  column:"tag"`
	ResourceK8sKind          Kind         `json:"resource_k8s_kind"           required:"true"  column:"tag"`
	ResourceK8sMinReplicas   int32        `json:"resource_k8s_min_replicas"   required:"false" column:"field"`
	ResourceK8sMaxReplicas   int32        `json:"resource_k8s_max_replicas"   required:"false" column:"field"`
	Policy                   Policy       `json:"policy"                      required:"false" column:"field"`
	GroupId                  string       `json:"group_id"                    required:"true"  column:"field"`
	EnableExecution          bool         `json:"enable_execution"            required:"false" column:"field"`
	RawSpec                  string       `json:"raw_spec"                    required:"false" column:"field"`
}
