package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebAppSpec defines the desired state of WebApp
type WebAppSpec struct {
	Count    int32  `json:"count"`
	Image    string `json:"image"`
	Port     int32  `json:"port"`
	Webgroup string `json:"webgroup"`
	Message  string `json:"message"`
}

// WebAppStatus defines the observed state of WebApp
type WebAppStatus struct {
	// Nodes are the names of the Instance Pods
	Nodes []string `json:"nodes"`
	// Env Message Passed to the Pods
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebApp is the Schema for the webapps API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=webapps,scope=Namespaced
type WebApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebAppSpec   `json:"spec,omitempty"`
	Status WebAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebAppList contains a list of WebApp
type WebAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebApp{}, &WebAppList{})
}
