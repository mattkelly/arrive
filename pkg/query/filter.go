package query

import (
	// NOTE: struct
	//"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	// NOTE: interface
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"
)

type Filter struct {
	//OperandLeft  Operand            `json:"operandLeft,omitempty"`
	Operator selection.Operator `json:"operator"`
	//OperandRight Operand            `json:"operandRight,omitempty"`
}

func (f *Filter) Match(obj *runtime.Unstructured) bool {
	return false
}
