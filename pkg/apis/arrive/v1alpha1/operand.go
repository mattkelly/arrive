package v1alpha1

import (
	//"fmt"
	//"reflect"

	"github.com/pkg/errors"

	"k8s.io/client-go/util/jsonpath"

	// NOTE: struct
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	// NOTE: interface
	//"k8s.io/apimachinery/pkg/runtime"
)

func (o *Operand) Resolve(obj *unstructured.Unstructured) (interface{}, error) {
	// TODO handle o == nil case?
	if o.Value != nil {
		return *o.Value, nil
	}

	if o.ValueFrom != nil {
		// TODO all of this is obviously extremely unsafe...
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
			return nil, errors.New("TODO only scalar values are supported")
		}

		res := results[0][0].Interface()

		if _, ok := res.(string); !ok {
			return nil, errors.New("TODO only string values are supported")
		}

		return res, nil
	}

	return nil, errors.New("must specify either value or valueFrom")
}
