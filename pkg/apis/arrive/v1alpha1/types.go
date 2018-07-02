package v1alpha1

import (
	//"github.com/mattkelly/arrive/pkg/query"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/selection"
)

// TODO remove noStatus

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArriveTest is a specification for a ArriveTest resource
type ArriveTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ArriveTestSpec `json:"spec"`
	//Status ArriveTestStatus `json:"status"`
}

// ArriveTestSpec is the spec for a ArriveTest resource
type ArriveTestSpec struct {
	Timeout       string        `json:timeout`
	Input         string        `json:input`
	ExpectedState ExpectedState `json:expectedState`
}

//type ArriveTestStatus struct {
//}

// TODO names
type ExpectedState struct {
	MatchFilters []Filter     `json:"matchFilters"`
	Result       FilterResult `json:"result"`
}

// TODO worst name ever
type FilterResult struct {
	Operator selection.Operator `json:"operator"`
	Count    Operand            `json:"count"`
}

type Operand struct {
	// TODO more than just strings
	// Pointers to determine unset vs explicit zero value
	Value     *string         `json:"value,omitempty"`
	ValueFrom *ValueReference `json:"valueFrom,omitempty"`
}

type Filter struct {
	OperandLeft  Operand            `json:"operandLeft,omitempty"`
	Operator     selection.Operator `json:"operator"`
	OperandRight Operand            `json:"operandRight,omitempty"`
}

type FieldSelector struct {
	FieldPath string `json"fieldPath"`
}

type ValueReference struct {
	// TODO would be nice to use corev1.ObjectFieldRef but it's limited in
	// which fields it supports right now.
	FieldRef *FieldSelector `json:"fieldRef,omitempty"`
	// Selects a resource of the container: only resources limits and requests
	// (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	// +optional
	ResourceFieldRef *corev1.ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	// Selects a key of a secret in the pod's namespace
	// +optional
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArriveTestList is a list of ArriveTest resources
type ArriveTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ArriveTest `json:"items"`
}
