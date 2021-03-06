//go:build !ignore_autogenerated
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

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesServer) DeepCopyInto(out *ServicesServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesServer.
func (in *ServicesServer) DeepCopy() *ServicesServer {
	if in == nil {
		return nil
	}
	out := new(ServicesServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesServerList) DeepCopyInto(out *ServicesServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicesServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesServerList.
func (in *ServicesServerList) DeepCopy() *ServicesServerList {
	if in == nil {
		return nil
	}
	out := new(ServicesServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesServerSpec) DeepCopyInto(out *ServicesServerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ServicesServerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesServerSpec.
func (in *ServicesServerSpec) DeepCopy() *ServicesServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesServerSpecIpv4FirewallRule) DeepCopyInto(out *ServicesServerSpecIpv4FirewallRule) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RangeEnd != nil {
		in, out := &in.RangeEnd, &out.RangeEnd
		*out = new(string)
		**out = **in
	}
	if in.RangeStart != nil {
		in, out := &in.RangeStart, &out.RangeStart
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesServerSpecIpv4FirewallRule.
func (in *ServicesServerSpecIpv4FirewallRule) DeepCopy() *ServicesServerSpecIpv4FirewallRule {
	if in == nil {
		return nil
	}
	out := new(ServicesServerSpecIpv4FirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesServerSpecResource) DeepCopyInto(out *ServicesServerSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AdminUsers != nil {
		in, out := &in.AdminUsers, &out.AdminUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BackupBlobContainerURI != nil {
		in, out := &in.BackupBlobContainerURI, &out.BackupBlobContainerURI
		*out = new(string)
		**out = **in
	}
	if in.EnablePowerBiService != nil {
		in, out := &in.EnablePowerBiService, &out.EnablePowerBiService
		*out = new(bool)
		**out = **in
	}
	if in.Ipv4FirewallRule != nil {
		in, out := &in.Ipv4FirewallRule, &out.Ipv4FirewallRule
		*out = make([]ServicesServerSpecIpv4FirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.QuerypoolConnectionMode != nil {
		in, out := &in.QuerypoolConnectionMode, &out.QuerypoolConnectionMode
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ServerFullName != nil {
		in, out := &in.ServerFullName, &out.ServerFullName
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesServerSpecResource.
func (in *ServicesServerSpecResource) DeepCopy() *ServicesServerSpecResource {
	if in == nil {
		return nil
	}
	out := new(ServicesServerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesServerStatus) DeepCopyInto(out *ServicesServerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesServerStatus.
func (in *ServicesServerStatus) DeepCopy() *ServicesServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesServerStatus)
	in.DeepCopyInto(out)
	return out
}
