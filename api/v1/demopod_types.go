/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DemoPodSpec defines the desired state of DemoPod
type DemoPodSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Pod name
	Name string `json:"name,omitempty"`
	// Pod image
	Image string `json:"image,omitempty"`
}

// DemoPodStatus defines the observed state of DemoPod
type DemoPodStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Name string `json:"name,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DemoPod is the Schema for the demopods API
type DemoPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoPodSpec   `json:"spec,omitempty"`
	Status DemoPodStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DemoPodList contains a list of DemoPod
type DemoPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoPod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoPod{}, &DemoPodList{})
}
