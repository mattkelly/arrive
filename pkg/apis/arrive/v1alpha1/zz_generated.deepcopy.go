// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArriveTest) DeepCopyInto(out *ArriveTest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArriveTest.
func (in *ArriveTest) DeepCopy() *ArriveTest {
	if in == nil {
		return nil
	}
	out := new(ArriveTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArriveTest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArriveTestList) DeepCopyInto(out *ArriveTestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArriveTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArriveTestList.
func (in *ArriveTestList) DeepCopy() *ArriveTestList {
	if in == nil {
		return nil
	}
	out := new(ArriveTestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArriveTestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArriveTestSpec) DeepCopyInto(out *ArriveTestSpec) {
	*out = *in
	in.ExpectedState.DeepCopyInto(&out.ExpectedState)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArriveTestSpec.
func (in *ArriveTestSpec) DeepCopy() *ArriveTestSpec {
	if in == nil {
		return nil
	}
	out := new(ArriveTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpectedState) DeepCopyInto(out *ExpectedState) {
	*out = *in
	in.Count.DeepCopyInto(&out.Count)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpectedState.
func (in *ExpectedState) DeepCopy() *ExpectedState {
	if in == nil {
		return nil
	}
	out := new(ExpectedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
	in.OperandLeft.DeepCopyInto(&out.OperandLeft)
	in.OperandRight.DeepCopyInto(&out.OperandRight)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operand) DeepCopyInto(out *Operand) {
	*out = *in
	in.ValueFrom.DeepCopyInto(&out.ValueFrom)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operand.
func (in *Operand) DeepCopy() *Operand {
	if in == nil {
		return nil
	}
	out := new(Operand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueReference) DeepCopyInto(out *ValueReference) {
	*out = *in
	if in.FieldRef != nil {
		in, out := &in.FieldRef, &out.FieldRef
		*out = new(v1.ObjectFieldSelector)
		**out = **in
	}
	if in.ResourceFieldRef != nil {
		in, out := &in.ResourceFieldRef, &out.ResourceFieldRef
		*out = new(v1.ResourceFieldSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueReference.
func (in *ValueReference) DeepCopy() *ValueReference {
	if in == nil {
		return nil
	}
	out := new(ValueReference)
	in.DeepCopyInto(out)
	return out
}
