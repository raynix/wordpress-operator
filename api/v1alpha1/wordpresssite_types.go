/*


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

// WordPressSiteSpec defines the desired state of WordPressSite
type WordPressSiteSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of WordPressSite. Edit WordPressSite_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// WordPressSiteStatus defines the observed state of WordPressSite
type WordPressSiteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// WordPressSite is the Schema for the wordpresssites API
type WordPressSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WordPressSiteSpec   `json:"spec,omitempty"`
	Status WordPressSiteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WordPressSiteList contains a list of WordPressSite
type WordPressSiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WordPressSite `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WordPressSite{}, &WordPressSiteList{})
}
