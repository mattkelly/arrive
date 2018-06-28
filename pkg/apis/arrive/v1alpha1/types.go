package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArriveTest is a specification for a ArriveTest resource
type ArriveTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArriveTestSpec   `json:"spec"`
	Status ArriveTestStatus `json:"status"`
}

// ArriveTestSpec is the spec for a ArriveTest resource
type ArriveTestSpec struct {
	Hallo string
}

// ArriveTestList is a list of ArriveTest resources
type ArriveTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ArriveTest `json:"items"`
}
