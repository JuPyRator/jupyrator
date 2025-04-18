//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelConnectionConfig) DeepCopyInto(out *KernelConnectionConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelConnectionConfig.
func (in *KernelConnectionConfig) DeepCopy() *KernelConnectionConfig {
	if in == nil {
		return nil
	}
	out := new(KernelConnectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelManager) DeepCopyInto(out *KernelManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelManager.
func (in *KernelManager) DeepCopy() *KernelManager {
	if in == nil {
		return nil
	}
	out := new(KernelManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KernelManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelManagerCondition) DeepCopyInto(out *KernelManagerCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelManagerCondition.
func (in *KernelManagerCondition) DeepCopy() *KernelManagerCondition {
	if in == nil {
		return nil
	}
	out := new(KernelManagerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelManagerList) DeepCopyInto(out *KernelManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KernelManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelManagerList.
func (in *KernelManagerList) DeepCopy() *KernelManagerList {
	if in == nil {
		return nil
	}
	out := new(KernelManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KernelManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelManagerSpec) DeepCopyInto(out *KernelManagerSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	out.ConnectionConfig = in.ConnectionConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelManagerSpec.
func (in *KernelManagerSpec) DeepCopy() *KernelManagerSpec {
	if in == nil {
		return nil
	}
	out := new(KernelManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelManagerStatus) DeepCopyInto(out *KernelManagerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KernelManagerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ContainerState.DeepCopyInto(&out.ContainerState)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelManagerStatus.
func (in *KernelManagerStatus) DeepCopy() *KernelManagerStatus {
	if in == nil {
		return nil
	}
	out := new(KernelManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelManagerTemplateSpec) DeepCopyInto(out *KernelManagerTemplateSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelManagerTemplateSpec.
func (in *KernelManagerTemplateSpec) DeepCopy() *KernelManagerTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(KernelManagerTemplateSpec)
	in.DeepCopyInto(out)
	return out
}
