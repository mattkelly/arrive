package query

import (
	// NOTE: struct
	//"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	// NOTE: interface
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"
)

type Filter struct {
	OperandLeft  Operand            `json:"operandLeft"`
	Operator     selection.Operator `json:"operator"`
	OperandRight Operand            `json:"operandRight"`
}

func (f *Filter) Match(obj *runtime.Unstructured) bool {
	return false
}

// TODO move to types.go when it exists
type ExpectedState struct {
	Operator selection.Operator `json:"operator"`
	Count    Operand            `json:"count"`
}
