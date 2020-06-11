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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AlamedaRecommendationSpec defines the desired state of AlamedaRecommendation
type AlamedaRecommendationSpec struct {
	Containers []AlamedaContainer `json:"containers" protobuf:"bytes,1,opt,name=containers"`
}

// AlamedaRecommendationStatus defines the observed state of AlamedaRecommendation
type AlamedaRecommendationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

type AlamedaRecommendation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlamedaRecommendationSpec   `json:"spec,omitempty"`
	Status AlamedaRecommendationStatus `json:"status,omitempty"`
}

func (ar *AlamedaRecommendation) GetNamespacedName() string {
	return fmt.Sprintf("%s/%s", ar.Namespace, ar.Name)
}

// +kubebuilder:object:root=true

// AlamedaRecommendationList contains a list of AlamedaRecommendation
type AlamedaRecommendationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlamedaRecommendation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlamedaRecommendation{}, &AlamedaRecommendationList{})
}
