package v1alpha1

import (
	// NOTE: struct
	//"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	// NOTE: interface
	"k8s.io/apimachinery/pkg/runtime"
)

func (f *Filter) Match(obj *runtime.Unstructured) bool {
	return false
}
