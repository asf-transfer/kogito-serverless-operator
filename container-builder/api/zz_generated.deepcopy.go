//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package api

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuild) DeepCopyInto(out *ContainerBuild) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuild.
func (in *ContainerBuild) DeepCopy() *ContainerBuild {
	if in == nil {
		return nil
	}
	out := new(ContainerBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuildBaseTask) DeepCopyInto(out *ContainerBuildBaseTask) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuildBaseTask.
func (in *ContainerBuildBaseTask) DeepCopy() *ContainerBuildBaseTask {
	if in == nil {
		return nil
	}
	out := new(ContainerBuildBaseTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuildCondition) DeepCopyInto(out *ContainerBuildCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuildCondition.
func (in *ContainerBuildCondition) DeepCopy() *ContainerBuildCondition {
	if in == nil {
		return nil
	}
	out := new(ContainerBuildCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuildFailure) DeepCopyInto(out *ContainerBuildFailure) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	in.Recovery.DeepCopyInto(&out.Recovery)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuildFailure.
func (in *ContainerBuildFailure) DeepCopy() *ContainerBuildFailure {
	if in == nil {
		return nil
	}
	out := new(ContainerBuildFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuildFailureRecovery) DeepCopyInto(out *ContainerBuildFailureRecovery) {
	*out = *in
	in.AttemptTime.DeepCopyInto(&out.AttemptTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuildFailureRecovery.
func (in *ContainerBuildFailureRecovery) DeepCopy() *ContainerBuildFailureRecovery {
	if in == nil {
		return nil
	}
	out := new(ContainerBuildFailureRecovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuildResourceVolume) DeepCopyInto(out *ContainerBuildResourceVolume) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuildResourceVolume.
func (in *ContainerBuildResourceVolume) DeepCopy() *ContainerBuildResourceVolume {
	if in == nil {
		return nil
	}
	out := new(ContainerBuildResourceVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuildSpec) DeepCopyInto(out *ContainerBuildSpec) {
	*out = *in
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]ContainerBuildTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Timeout = in.Timeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuildSpec.
func (in *ContainerBuildSpec) DeepCopy() *ContainerBuildSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuildStatus) DeepCopyInto(out *ContainerBuildStatus) {
	*out = *in
	if in.Failure != nil {
		in, out := &in.Failure, &out.Failure
		*out = new(ContainerBuildFailure)
		(*in).DeepCopyInto(*out)
	}
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ContainerBuildCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceVolume != nil {
		in, out := &in.ResourceVolume, &out.ResourceVolume
		*out = new(ContainerBuildResourceVolume)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuildStatus.
func (in *ContainerBuildStatus) DeepCopy() *ContainerBuildStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerBuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBuildTask) DeepCopyInto(out *ContainerBuildTask) {
	*out = *in
	if in.Kaniko != nil {
		in, out := &in.Kaniko, &out.Kaniko
		*out = new(KanikoTask)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBuildTask.
func (in *ContainerBuildTask) DeepCopy() *ContainerBuildTask {
	if in == nil {
		return nil
	}
	out := new(ContainerBuildTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRegistrySpec) DeepCopyInto(out *ContainerRegistrySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRegistrySpec.
func (in *ContainerRegistrySpec) DeepCopy() *ContainerRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(ContainerRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KanikoTask) DeepCopyInto(out *KanikoTask) {
	*out = *in
	out.ContainerBuildBaseTask = in.ContainerBuildBaseTask
	out.PublishTask = in.PublishTask
	if in.Verbose != nil {
		in, out := &in.Verbose, &out.Verbose
		*out = new(bool)
		**out = **in
	}
	in.Cache.DeepCopyInto(&out.Cache)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.AdditionalFlags != nil {
		in, out := &in.AdditionalFlags, &out.AdditionalFlags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KanikoTask.
func (in *KanikoTask) DeepCopy() *KanikoTask {
	if in == nil {
		return nil
	}
	out := new(KanikoTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KanikoTaskCache) DeepCopyInto(out *KanikoTaskCache) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KanikoTaskCache.
func (in *KanikoTaskCache) DeepCopy() *KanikoTaskCache {
	if in == nil {
		return nil
	}
	out := new(KanikoTaskCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformContainerBuild) DeepCopyInto(out *PlatformContainerBuild) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformContainerBuild.
func (in *PlatformContainerBuild) DeepCopy() *PlatformContainerBuild {
	if in == nil {
		return nil
	}
	out := new(PlatformContainerBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformContainerBuildSpec) DeepCopyInto(out *PlatformContainerBuildSpec) {
	*out = *in
	out.Registry = in.Registry
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.BuildStrategyOptions != nil {
		in, out := &in.BuildStrategyOptions, &out.BuildStrategyOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformContainerBuildSpec.
func (in *PlatformContainerBuildSpec) DeepCopy() *PlatformContainerBuildSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformContainerBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishTask) DeepCopyInto(out *PublishTask) {
	*out = *in
	out.Registry = in.Registry
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishTask.
func (in *PublishTask) DeepCopy() *PublishTask {
	if in == nil {
		return nil
	}
	out := new(PublishTask)
	in.DeepCopyInto(out)
	return out
}
