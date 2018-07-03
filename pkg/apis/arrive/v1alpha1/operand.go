package v1alpha1

import (
	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/util/jsonpath"
)

func (o *Operand) Resolve(obj *unstructured.Unstructured) (interface{}, error) {
	// TODO handle o == nil case?
	if o.Value != nil {
		return *o.Value, nil
	}

	if o.ValueFrom != nil {
		// TODO this line makes a ton of assumptions
		template := o.ValueFrom.FieldRef.FieldPath

		jp := jsonpath.New("test")
		if err := jp.Parse(template); err != nil {
			return nil, errors.Wrap(err, "parse JSONPath template")
		}

		results, err := jp.FindResults(obj.Object)
		if err != nil {
			return nil, errors.Wrap(err, "find JSONPath results")
		}

		if len(results) != 1 || len(results[0]) != 1 {
			return nil, errors.New("only a single JSONPath result is supported")
		}

		return results[0][0].Interface(), nil
	}

	return nil, errors.New("must specify either value or valueFrom")
}
