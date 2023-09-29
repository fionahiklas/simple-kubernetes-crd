/*
Copyright 2023 Fiona Bianchi.

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

// ElectricTreesSpec defines the desired state of ElectricTrees
type ElectricTreesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ElectricTrees. Edit electrictrees_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ElectricTreesStatus defines the observed state of ElectricTrees
type ElectricTreesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ElectricTrees is the Schema for the electrictrees API
type ElectricTrees struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElectricTreesSpec   `json:"spec,omitempty"`
	Status ElectricTreesStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ElectricTreesList contains a list of ElectricTrees
type ElectricTreesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElectricTrees `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ElectricTrees{}, &ElectricTreesList{})
}
