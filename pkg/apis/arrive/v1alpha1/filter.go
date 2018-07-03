package v1alpha1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (f *Filter) Match(obj *unstructured.Unstructured) (bool, error) {
	// TODO handle case where only one Operand specified
	left, err := f.OperandLeft.Resolve(obj)
	fmt.Printf("left = %v (%T)\n", left, left)
	if err != nil {
		// TODO wrap
		return false, err
	}
	leftStr, ok := left.(string)
	if !ok {
		panic("")
	}
	_ = leftStr

	right, err := f.OperandRight.Resolve(obj)
	fmt.Printf("right = %v (%T)\n", right, right)
	if err != nil {
		// TODO wrap
		return false, err
	}

	// TODO actually use Operator
	return left == right, nil
}
