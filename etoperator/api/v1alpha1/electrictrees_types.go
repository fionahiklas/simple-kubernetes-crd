/*
   Copyright 2023 Fiona Bianchi.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ElectricTreesSpec defines the desired state of ElectricTrees
type ElectricTreesSpec struct {
	TreeName   string `json:"tree_name,omitempty"`
	Try        bool   `json:"try,omitempty"`
	HowFarAway int    `json:"how_far_away,omitempty"`
	EyesClosed bool   `json:"eyes_closed,omitempty"`
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
