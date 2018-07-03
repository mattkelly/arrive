package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
// ExpectedState is the desired state to reach within the timeout.
type ExpectedState struct {
	// TODO it would be nice to target GVK based on matchFilters alone (just
	// have (first filter(s) specify GVK) but this makes it easier and more
	// clear at least for now
	FilterTarget schema.GroupVersionKind `json:"filterTarget"`
	MatchFilters []Filter                `json:"matchFilters"`
	Result       FilterResult            `json:"result"`
}

// TODO worst name ever
type FilterResult struct {
	Operator selection.Operator `json:"operator"`
	Count    Operand            `json:"count"`
}

// Operand specifies a value to operate on. The value can either be a literal
// or a reference to another value.
type Operand struct {
	// TODO more than just strings

	// Pointers to determine unset vs explicit zero value
	Value     *string         `json:"value,omitempty"`
	ValueFrom *ValueReference `json:"valueFrom,omitempty"`
}

type Filter struct {
	// OperandLeft is the left-hand operand
	OperandLeft Operand `json:"operandLeft,omitempty"`
	// Operator is the operator to perform on the operands
	Operator selection.Operator `json:"operator"`
	// OperandLeft is the right-hand operand
	OperandRight Operand `json:"operandRight,omitempty"`
}

type FieldSelector struct {
	// FieldPath is a JSONPath that selects a field in a resource
	FieldPath string `json"fieldPath"`
}

// ValueReference is a reference to a value in a resource.
type ValueReference struct {
	// FieldRef is a reference to a field in a resource
	FieldRef *FieldSelector `json:"fieldRef,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArriveTestList is a list of ArriveTest resources
type ArriveTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ArriveTest `json:"items"`
}
