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

// +kubebuilder:validation:Enum=hpa;predicitOnly
type ScalingType string

const (
	NonScaling ScalingType = "predicitOnly"
	HPAScaling ScalingType = "hpa"
)

// Raw value stored in datahub
var ScalingTypeMap = map[ScalingType]string{
	NonScaling: "NONE",
	HPAScaling: "HPA",
}

// +kubebuilder:validation:Enum=Deployment;StatefulSet;DeploymentConfig
type ControllerKind string

const (
	DeploymentController       ControllerKind = "Deployment"
	StatefulSetController      ControllerKind = "StatefulSet"
	DeploymentConfigController ControllerKind = "DeploymentConfig"
)

// Raw value stored in datahub
var ControllerKindMap = map[ControllerKind]string{
	DeploymentController:       "DEPLOYMENT",
	StatefulSetController:      "STATEFULSET",
	DeploymentConfigController: "DEPLOYMENTCONFIG",
}

type Target struct {
	Namespace string         `json:"namespace"`
	Name      string         `json:"name"`
	Kind      ControllerKind `json:"kind"`
}

type GenericHPAParameters struct {
	// +optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// +optional
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
}

type Generic struct {
	Target Target `json:"target"`
	// +optional
	HpaParameters *GenericHPAParameters `json:"hpaParameters,omitempty"`
}

type ConsumerGroup struct {
	Namespace string         `json:"namespace"`
	Name      string         `json:"name"`
	Kind      ControllerKind `json:"kind"`
	Topic     string         `json:"topic"`
}

type KafkaHPAParameters struct {
	// +optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// +optional
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
}

type Kafka struct {
	ConsumerGroup    ConsumerGroup `json:"consumerGroup"`
	ExporterNamespce string        `json:"exporterNamespace"`
	// +optional
	HpaParameters *KafkaHPAParameters `json:"hpaParameters,omitempty"`
}

type Controller struct {
	// +optional
	EnableExecution *bool       `json:"enableExecution,omitempty"`
	Type            TargetType  `json:"type"`
	Scaling         ScalingType `json:"scaling"`
	// +optional
	Generic *Generic `json:"generic,omitempty"`
	// +optional
	Kafka *Kafka `json:"kafka,omitempty"`
}

// AlamedaScalerSpec defines the desired state of AlamedaScaler
type AlamedaScalerSpec struct {
	ClusterName string       `json:"clusterName"`
	Controllers []Controller `json:"controllers"`
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
