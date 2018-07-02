package v1alpha1

import (
	"fmt"

	// NOTE: struct
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	// NOTE: interface
	//"k8s.io/apimachinery/pkg/runtime"
)

func (f *Filter) Match(obj *unstructured.Unstructured) (bool, error) {
	// TODO handle case where only one Operand specified
	left, err := f.OperandLeft.Resolve()
	fmt.Printf("left = %v\n", left)
	if err != nil {
		// TODO wrap
		return false, err
	}

	right, err := f.OperandRight.Resolve()
	fmt.Printf("right = %v\n", right)
	if err != nil {
		// TODO wrap
		return false, err
	}

	// TODO actually use Operator
	return left == right, nil
}
