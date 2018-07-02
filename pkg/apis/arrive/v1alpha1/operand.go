package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	// NOTE: struct
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	// NOTE: interface
	//"k8s.io/apimachinery/pkg/runtime"
)

// TODO borrowed from ark, just pull in ark utils maybe.
// GetValue returns the object at root[path], where path is a dot separated string.
func GetValue(root map[string]interface{}, path string) (interface{}, error) {
	if root == nil {
		return "", errors.New("root is nil")
	}

	pathParts := strings.Split(path, ".")
	key := pathParts[0]

	obj, found := root[pathParts[0]]
	if !found {
		return "", errors.Errorf("key %v not found", pathParts[0])
	}

	if len(pathParts) == 1 {
		return obj, nil
	}

	subMap, ok := obj.(map[string]interface{})
	if !ok {
		return "", errors.Errorf("value at key %v is not a map[string]interface{}", key)
	}

	return GetValue(subMap, strings.Join(pathParts[1:], "."))
}

func (o *Operand) Resolve(obj *unstructured.Unstructured) (interface{}, error) {
	// TODO handle o == nil case?
	if o.Value != nil {
		return *o.Value, nil
	}

	if o.ValueFrom != nil {
		return GetValue(obj.Object, o.ValueFrom.FieldRef.FieldPath)
	}

	return nil, fmt.Errorf("must specify either value or valueFrom")
}
