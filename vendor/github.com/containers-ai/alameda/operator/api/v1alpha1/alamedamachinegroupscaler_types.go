/*
Copyright 2020 The Alameda Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AlamedaMachineGroupScalerSpec defines the desired state of AlamedaMachineGroupScaler
type AlamedaMachineGroupScalerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	MachineSets []MachineSet          `json:"machineSets"`
	Metrics     map[string]MetricRule `json:"metrics"`
}

// AlamedaMachineGroupScalerStatus defines the observed state of AlamedaMachineGroupScaler
type AlamedaMachineGroupScalerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// AlamedaMachineGroupScaler is the Schema for the alamedamachinegroupscalers API
type AlamedaMachineGroupScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlamedaMachineGroupScalerSpec   `json:"spec,omitempty"`
	Status AlamedaMachineGroupScalerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlamedaMachineGroupScalerList contains a list of AlamedaMachineGroupScaler
type AlamedaMachineGroupScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlamedaMachineGroupScaler `json:"items"`
}

type MachineSet struct {
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	MinReplicas    *int32 `json:"minReplicas,omitempty"`
	MaxReplicas    *int32 `json:"maxReplicas,omitempty"`
	ScaleUpDelay   *int64 `json:"scaleUpDelay,omitempty"`
	ScaleDownDelay *int64 `json:"scaleDownDelay,omitempty"`
}

type MetricRule struct {
	UtilizationTarget               *int32 `json:"utilizationTarget,omitempty"`
	ScaleUpGap                      *int32 `json:"scaleUpGap,omitempty"`
	ScaleDownGap                    *int32 `json:"scaleDownGap,omitempty"`
	DurationUpThresholdPercentage   *int32 `json:"durationUpThresholdPercentage,omitempty"`
	DurationDownThresholdPercentage *int32 `json:"durationDownThresholdPercentage,omitempty"`
}

func init() {
	SchemeBuilder.Register(&AlamedaMachineGroupScaler{}, &AlamedaMachineGroupScalerList{})
}
