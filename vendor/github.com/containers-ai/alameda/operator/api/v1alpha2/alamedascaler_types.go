/*
Copyright 2019 The Alameda Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha2

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:validation:Enum=generic;kafka
type TargetType string

const (
	GenericTarget TargetType = "generic"
	KafkaTarget   TargetType = "kafka"
)

// +kubebuilder:validation:Enum=hpa;predictOnly
type ScalingType string

const (
	NonScaling ScalingType = "predictOnly"
	HPAScaling ScalingType = "hpa"
)

// Raw value stored in datahub
var ScalingTypeMap = map[ScalingType]entities.ScalingTool{
	NonScaling: entities.None,
	HPAScaling: entities.HPA,
}

// +kubebuilder:validation:Enum=Deployment;StatefulSet;DeploymentConfig
type ControllerKind string

const (
	DeploymentController       ControllerKind = "Deployment"
	StatefulSetController      ControllerKind = "StatefulSet"
	DeploymentConfigController ControllerKind = "DeploymentConfig"
)

// Raw value stored in datahub
var ControllerKindMap = map[ControllerKind]entities.Kind{
	DeploymentController:       entities.Deployment,
	StatefulSetController:      entities.StatefulSet,
	DeploymentConfigController: entities.DeploymentConfig,
}

type Target struct {
	// controller namespace
	Namespace string `json:"namespace"`
	// controller name
	Name string `json:"name"`
	// controller kind (deployment/deploymentConfig/statefulSet)
	Kind ControllerKind `json:"kind"`
}

type GenericHPAParameters struct {
	// minimum limit of number of replicas
	// +optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// maximum limit of number of replicas
	// +optional
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
}

type Generic struct {
	// reference to generic application to be managed
	Target Target `json:"target"`
	// HPA autoscaling parameters for generic application
	// +optional
	HpaParameters *GenericHPAParameters `json:"hpaParameters,omitempty"`
}

type ConsumerGroup struct {
	// namespace of the consumer group
	Namespace string `json:"namespace"`
	// name of the consumer group
	Name string `json:"name"`
	// controller kind of the consumer group
	Kind ControllerKind `json:"kind"`
	// topic name that the consumer group subscribed
	Topic string `json:"topic"`
	// the name of the consumer group a Kafka consumer belongs to.
	// It’s the group ID given by Kafka, not the consumer group deployment name.
	// +optional
	GroupId *string `json:"groupId"`
}

type KafkaHPAParameters struct {
	// minimum limit of number of replicas
	// +optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// maximum limit of number of replicas
	// +optional
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
}

type Kafka struct {
	// reference to kafka consumer group
	ConsumerGroup ConsumerGroup `json:"consumerGroup"`
	// namespace of kafka broker; namespace of the metrics service where kafka metrics are read from
	ExporterNamespce string `json:"exporterNamespace"`
	// HPA autoscaling parameters for kafka
	// +optional
	HpaParameters *KafkaHPAParameters `json:"hpaParameters,omitempty"`
}

type Controller struct {
	// enable Federator.ai autoscaling execution.
	// This flag is to control the execution by Federator.ai executor.
	// It is usable only if the application and Federator.ai are running in the same cluster.
	// In the cases of using Datadog WPA to do execution or the application is running in a different target cluster,
	// ‘enableExecution’ is noneffective.
	// +optional
	EnableExecution *bool `json:"enableExecution,omitempty"`
	// controller type (generic, kafka, nginx)
	Type TargetType `json:"type"`
	// scaling methods (hpa, predictionOnly)
	Scaling ScalingType `json:"scaling"`
	// generic application metadata
	// +optional
	Generic *Generic `json:"generic,omitempty"`
	// kafka metadata
	// +optional
	Kafka *Kafka `json:"kafka,omitempty"`
}

// AlamedaScalerSpec defines the desired state of AlamedaScaler
type AlamedaScalerSpec struct {
	// target cluster name; the cluster where applications to be managed are running
	ClusterName string `json:"clusterName"`
	// list of controllers (deployment/deploymentConfig/statefulSet) to be managed
	// +optional
	Controllers []Controller `json:"controllers,omitempty"`
}

// AlamedaScalerStatus defines the observed state of AlamedaScaler
type AlamedaScalerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// AlamedaScaler is the Schema for the alamedascalers API
type AlamedaScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlamedaScalerSpec   `json:"spec,omitempty"`
	Status AlamedaScalerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlamedaScalerList contains a list of AlamedaScaler
type AlamedaScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlamedaScaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlamedaScaler{}, &AlamedaScalerList{})
}
