// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRule) DeepCopyInto(out *FirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRule.
func (in *FirewallRule) DeepCopy() *FirewallRule {
	if in == nil {
		return nil
	}
	out := new(FirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleList) DeepCopyInto(out *FirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleList.
func (in *FirewallRuleList) DeepCopy() *FirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleSpec) DeepCopyInto(out *FirewallRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(FirewallRuleSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleSpec.
func (in *FirewallRuleSpec) DeepCopy() *FirewallRuleSpec {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleSpecResource) DeepCopyInto(out *FirewallRuleSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.EndIPAddress != nil {
		in, out := &in.EndIPAddress, &out.EndIPAddress
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.StartIPAddress != nil {
		in, out := &in.StartIPAddress, &out.StartIPAddress
		*out = new(string)
		**out = **in
	}
	if in.SynapseWorkspaceID != nil {
		in, out := &in.SynapseWorkspaceID, &out.SynapseWorkspaceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleSpecResource.
func (in *FirewallRuleSpecResource) DeepCopy() *FirewallRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleStatus) DeepCopyInto(out *FirewallRuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleStatus.
func (in *FirewallRuleStatus) DeepCopy() *FirewallRuleStatus {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPrivateEndpoint) DeepCopyInto(out *ManagedPrivateEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPrivateEndpoint.
func (in *ManagedPrivateEndpoint) DeepCopy() *ManagedPrivateEndpoint {
	if in == nil {
		return nil
	}
	out := new(ManagedPrivateEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedPrivateEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPrivateEndpointList) DeepCopyInto(out *ManagedPrivateEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedPrivateEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPrivateEndpointList.
func (in *ManagedPrivateEndpointList) DeepCopy() *ManagedPrivateEndpointList {
	if in == nil {
		return nil
	}
	out := new(ManagedPrivateEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedPrivateEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPrivateEndpointSpec) DeepCopyInto(out *ManagedPrivateEndpointSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ManagedPrivateEndpointSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPrivateEndpointSpec.
func (in *ManagedPrivateEndpointSpec) DeepCopy() *ManagedPrivateEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedPrivateEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPrivateEndpointSpecResource) DeepCopyInto(out *ManagedPrivateEndpointSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SubresourceName != nil {
		in, out := &in.SubresourceName, &out.SubresourceName
		*out = new(string)
		**out = **in
	}
	if in.SynapseWorkspaceID != nil {
		in, out := &in.SynapseWorkspaceID, &out.SynapseWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.TargetResourceID != nil {
		in, out := &in.TargetResourceID, &out.TargetResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPrivateEndpointSpecResource.
func (in *ManagedPrivateEndpointSpecResource) DeepCopy() *ManagedPrivateEndpointSpecResource {
	if in == nil {
		return nil
	}
	out := new(ManagedPrivateEndpointSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPrivateEndpointStatus) DeepCopyInto(out *ManagedPrivateEndpointStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPrivateEndpointStatus.
func (in *ManagedPrivateEndpointStatus) DeepCopy() *ManagedPrivateEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedPrivateEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignment) DeepCopyInto(out *RoleAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignment.
func (in *RoleAssignment) DeepCopy() *RoleAssignment {
	if in == nil {
		return nil
	}
	out := new(RoleAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentList) DeepCopyInto(out *RoleAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentList.
func (in *RoleAssignmentList) DeepCopy() *RoleAssignmentList {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentSpec) DeepCopyInto(out *RoleAssignmentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RoleAssignmentSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentSpec.
func (in *RoleAssignmentSpec) DeepCopy() *RoleAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentSpecResource) DeepCopyInto(out *RoleAssignmentSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.SynapseSparkPoolID != nil {
		in, out := &in.SynapseSparkPoolID, &out.SynapseSparkPoolID
		*out = new(string)
		**out = **in
	}
	if in.SynapseWorkspaceID != nil {
		in, out := &in.SynapseWorkspaceID, &out.SynapseWorkspaceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentSpecResource.
func (in *RoleAssignmentSpecResource) DeepCopy() *RoleAssignmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentStatus) DeepCopyInto(out *RoleAssignmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentStatus.
func (in *RoleAssignmentStatus) DeepCopy() *RoleAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPool) DeepCopyInto(out *SparkPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPool.
func (in *SparkPool) DeepCopy() *SparkPool {
	if in == nil {
		return nil
	}
	out := new(SparkPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPoolList) DeepCopyInto(out *SparkPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SparkPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPoolList.
func (in *SparkPoolList) DeepCopy() *SparkPoolList {
	if in == nil {
		return nil
	}
	out := new(SparkPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPoolSpec) DeepCopyInto(out *SparkPoolSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SparkPoolSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPoolSpec.
func (in *SparkPoolSpec) DeepCopy() *SparkPoolSpec {
	if in == nil {
		return nil
	}
	out := new(SparkPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPoolSpecAutoPause) DeepCopyInto(out *SparkPoolSpecAutoPause) {
	*out = *in
	if in.DelayInMinutes != nil {
		in, out := &in.DelayInMinutes, &out.DelayInMinutes
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPoolSpecAutoPause.
func (in *SparkPoolSpecAutoPause) DeepCopy() *SparkPoolSpecAutoPause {
	if in == nil {
		return nil
	}
	out := new(SparkPoolSpecAutoPause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPoolSpecAutoScale) DeepCopyInto(out *SparkPoolSpecAutoScale) {
	*out = *in
	if in.MaxNodeCount != nil {
		in, out := &in.MaxNodeCount, &out.MaxNodeCount
		*out = new(int64)
		**out = **in
	}
	if in.MinNodeCount != nil {
		in, out := &in.MinNodeCount, &out.MinNodeCount
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPoolSpecAutoScale.
func (in *SparkPoolSpecAutoScale) DeepCopy() *SparkPoolSpecAutoScale {
	if in == nil {
		return nil
	}
	out := new(SparkPoolSpecAutoScale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPoolSpecLibraryRequirement) DeepCopyInto(out *SparkPoolSpecLibraryRequirement) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Filename != nil {
		in, out := &in.Filename, &out.Filename
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPoolSpecLibraryRequirement.
func (in *SparkPoolSpecLibraryRequirement) DeepCopy() *SparkPoolSpecLibraryRequirement {
	if in == nil {
		return nil
	}
	out := new(SparkPoolSpecLibraryRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPoolSpecResource) DeepCopyInto(out *SparkPoolSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoPause != nil {
		in, out := &in.AutoPause, &out.AutoPause
		*out = new(SparkPoolSpecAutoPause)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoScale != nil {
		in, out := &in.AutoScale, &out.AutoScale
		*out = new(SparkPoolSpecAutoScale)
		(*in).DeepCopyInto(*out)
	}
	if in.LibraryRequirement != nil {
		in, out := &in.LibraryRequirement, &out.LibraryRequirement
		*out = new(SparkPoolSpecLibraryRequirement)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int64)
		**out = **in
	}
	if in.NodeSize != nil {
		in, out := &in.NodeSize, &out.NodeSize
		*out = new(string)
		**out = **in
	}
	if in.NodeSizeFamily != nil {
		in, out := &in.NodeSizeFamily, &out.NodeSizeFamily
		*out = new(string)
		**out = **in
	}
	if in.SparkEventsFolder != nil {
		in, out := &in.SparkEventsFolder, &out.SparkEventsFolder
		*out = new(string)
		**out = **in
	}
	if in.SparkLogFolder != nil {
		in, out := &in.SparkLogFolder, &out.SparkLogFolder
		*out = new(string)
		**out = **in
	}
	if in.SparkVersion != nil {
		in, out := &in.SparkVersion, &out.SparkVersion
		*out = new(string)
		**out = **in
	}
	if in.SynapseWorkspaceID != nil {
		in, out := &in.SynapseWorkspaceID, &out.SynapseWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPoolSpecResource.
func (in *SparkPoolSpecResource) DeepCopy() *SparkPoolSpecResource {
	if in == nil {
		return nil
	}
	out := new(SparkPoolSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPoolStatus) DeepCopyInto(out *SparkPoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPoolStatus.
func (in *SparkPoolStatus) DeepCopy() *SparkPoolStatus {
	if in == nil {
		return nil
	}
	out := new(SparkPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqlPool) DeepCopyInto(out *SqlPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqlPool.
func (in *SqlPool) DeepCopy() *SqlPool {
	if in == nil {
		return nil
	}
	out := new(SqlPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SqlPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqlPoolList) DeepCopyInto(out *SqlPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SqlPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqlPoolList.
func (in *SqlPoolList) DeepCopy() *SqlPoolList {
	if in == nil {
		return nil
	}
	out := new(SqlPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SqlPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqlPoolSpec) DeepCopyInto(out *SqlPoolSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SqlPoolSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqlPoolSpec.
func (in *SqlPoolSpec) DeepCopy() *SqlPoolSpec {
	if in == nil {
		return nil
	}
	out := new(SqlPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqlPoolSpecResource) DeepCopyInto(out *SqlPoolSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
		*out = new(string)
		**out = **in
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.DataEncrypted != nil {
		in, out := &in.DataEncrypted, &out.DataEncrypted
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryDatabaseID != nil {
		in, out := &in.RecoveryDatabaseID, &out.RecoveryDatabaseID
		*out = new(string)
		**out = **in
	}
	if in.Restore != nil {
		in, out := &in.Restore, &out.Restore
		*out = new(SqlPoolSpecRestore)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SynapseWorkspaceID != nil {
		in, out := &in.SynapseWorkspaceID, &out.SynapseWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqlPoolSpecResource.
func (in *SqlPoolSpecResource) DeepCopy() *SqlPoolSpecResource {
	if in == nil {
		return nil
	}
	out := new(SqlPoolSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqlPoolSpecRestore) DeepCopyInto(out *SqlPoolSpecRestore) {
	*out = *in
	if in.PointInTime != nil {
		in, out := &in.PointInTime, &out.PointInTime
		*out = new(string)
		**out = **in
	}
	if in.SourceDatabaseID != nil {
		in, out := &in.SourceDatabaseID, &out.SourceDatabaseID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqlPoolSpecRestore.
func (in *SqlPoolSpecRestore) DeepCopy() *SqlPoolSpecRestore {
	if in == nil {
		return nil
	}
	out := new(SqlPoolSpecRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqlPoolStatus) DeepCopyInto(out *SqlPoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqlPoolStatus.
func (in *SqlPoolStatus) DeepCopy() *SqlPoolStatus {
	if in == nil {
		return nil
	}
	out := new(SqlPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpec) DeepCopyInto(out *WorkspaceSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(WorkspaceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpec.
func (in *WorkspaceSpec) DeepCopy() *WorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpecAadAdmin) DeepCopyInto(out *WorkspaceSpecAadAdmin) {
	*out = *in
	if in.Login != nil {
		in, out := &in.Login, &out.Login
		*out = new(string)
		**out = **in
	}
	if in.ObjectID != nil {
		in, out := &in.ObjectID, &out.ObjectID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpecAadAdmin.
func (in *WorkspaceSpecAadAdmin) DeepCopy() *WorkspaceSpecAadAdmin {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpecAadAdmin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpecAzureDevopsRepo) DeepCopyInto(out *WorkspaceSpecAzureDevopsRepo) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.ProjectName != nil {
		in, out := &in.ProjectName, &out.ProjectName
		*out = new(string)
		**out = **in
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RootFolder != nil {
		in, out := &in.RootFolder, &out.RootFolder
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpecAzureDevopsRepo.
func (in *WorkspaceSpecAzureDevopsRepo) DeepCopy() *WorkspaceSpecAzureDevopsRepo {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpecAzureDevopsRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpecGithubRepo) DeepCopyInto(out *WorkspaceSpecGithubRepo) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.GitURL != nil {
		in, out := &in.GitURL, &out.GitURL
		*out = new(string)
		**out = **in
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RootFolder != nil {
		in, out := &in.RootFolder, &out.RootFolder
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpecGithubRepo.
func (in *WorkspaceSpecGithubRepo) DeepCopy() *WorkspaceSpecGithubRepo {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpecGithubRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpecIdentity) DeepCopyInto(out *WorkspaceSpecIdentity) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpecIdentity.
func (in *WorkspaceSpecIdentity) DeepCopy() *WorkspaceSpecIdentity {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpecIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpecResource) DeepCopyInto(out *WorkspaceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AadAdmin != nil {
		in, out := &in.AadAdmin, &out.AadAdmin
		*out = new(WorkspaceSpecAadAdmin)
		(*in).DeepCopyInto(*out)
	}
	if in.AzureDevopsRepo != nil {
		in, out := &in.AzureDevopsRepo, &out.AzureDevopsRepo
		*out = new(WorkspaceSpecAzureDevopsRepo)
		(*in).DeepCopyInto(*out)
	}
	if in.ConnectivityEndpoints != nil {
		in, out := &in.ConnectivityEndpoints, &out.ConnectivityEndpoints
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.CustomerManagedKeyVersionlessID != nil {
		in, out := &in.CustomerManagedKeyVersionlessID, &out.CustomerManagedKeyVersionlessID
		*out = new(string)
		**out = **in
	}
	if in.DataExfiltrationProtectionEnabled != nil {
		in, out := &in.DataExfiltrationProtectionEnabled, &out.DataExfiltrationProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.GithubRepo != nil {
		in, out := &in.GithubRepo, &out.GithubRepo
		*out = new(WorkspaceSpecGithubRepo)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]WorkspaceSpecIdentity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ManagedResourceGroupName != nil {
		in, out := &in.ManagedResourceGroupName, &out.ManagedResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ManagedVirtualNetworkEnabled != nil {
		in, out := &in.ManagedVirtualNetworkEnabled, &out.ManagedVirtualNetworkEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SqlAdministratorLogin != nil {
		in, out := &in.SqlAdministratorLogin, &out.SqlAdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.SqlAdministratorLoginPassword != nil {
		in, out := &in.SqlAdministratorLoginPassword, &out.SqlAdministratorLoginPassword
		*out = new(string)
		**out = **in
	}
	if in.SqlIdentityControlEnabled != nil {
		in, out := &in.SqlIdentityControlEnabled, &out.SqlIdentityControlEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StorageDataLakeGen2FilesystemID != nil {
		in, out := &in.StorageDataLakeGen2FilesystemID, &out.StorageDataLakeGen2FilesystemID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpecResource.
func (in *WorkspaceSpecResource) DeepCopy() *WorkspaceSpecResource {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceStatus) DeepCopyInto(out *WorkspaceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceStatus.
func (in *WorkspaceStatus) DeepCopy() *WorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}
