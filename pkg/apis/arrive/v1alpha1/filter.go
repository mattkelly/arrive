package v1alpha1

import (
	"fmt"

	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/selection"
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

	return operatorResult(f.Operator, left, right)
}

func operatorResult(op selection.Operator, left, right interface{}) (bool, error) {
	switch op {
	case selection.DoubleEquals:
		fallthrough
	case selection.Equals:
		return left == right, nil
	case selection.In:
		fallthrough
	case selection.NotEquals:
		fallthrough
	case selection.NotIn:
		fallthrough
	case selection.Exists:
		fallthrough
	case selection.GreaterThan:
		fallthrough
	case selection.LessThan:
		fallthrough
	default:
		return false, errors.Errorf("operator %q not supported", op)
	}
}
