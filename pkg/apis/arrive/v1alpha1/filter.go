package v1alpha1

import (
	"fmt"

	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (f *Filter) Match(obj *unstructured.Unstructured) (bool, error) {
	// TODO handle case where only one Operand specified
	left, err := f.OperandLeft.Resolve(obj)
	fmt.Printf("left = %v (%T)\n", left, left)
	if err != nil {
		return false, errors.Wrap(err, "resolve operandLeft")
	}

	right, err := f.OperandRight.Resolve(obj)
	fmt.Printf("right = %v (%T)\n", right, right)
	if err != nil {
		// TODO wrap
		return false, errors.Wrap(err, "resolve operandRight")
	}

	// TODO actually use Operator
	return left == right, nil
}
