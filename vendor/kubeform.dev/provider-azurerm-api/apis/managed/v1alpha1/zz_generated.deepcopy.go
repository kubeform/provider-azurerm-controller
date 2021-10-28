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
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinition) DeepCopyInto(out *ApplicationDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinition.
func (in *ApplicationDefinition) DeepCopy() *ApplicationDefinition {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinitionList) DeepCopyInto(out *ApplicationDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinitionList.
func (in *ApplicationDefinitionList) DeepCopy() *ApplicationDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinitionSpec) DeepCopyInto(out *ApplicationDefinitionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ApplicationDefinitionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinitionSpec.
func (in *ApplicationDefinitionSpec) DeepCopy() *ApplicationDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinitionSpecAuthorization) DeepCopyInto(out *ApplicationDefinitionSpecAuthorization) {
	*out = *in
	if in.RoleDefinitionID != nil {
		in, out := &in.RoleDefinitionID, &out.RoleDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.ServicePrincipalID != nil {
		in, out := &in.ServicePrincipalID, &out.ServicePrincipalID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinitionSpecAuthorization.
func (in *ApplicationDefinitionSpecAuthorization) DeepCopy() *ApplicationDefinitionSpecAuthorization {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinitionSpecAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinitionSpecResource) DeepCopyInto(out *ApplicationDefinitionSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = make([]ApplicationDefinitionSpecAuthorization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateUiDefinition != nil {
		in, out := &in.CreateUiDefinition, &out.CreateUiDefinition
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LockLevel != nil {
		in, out := &in.LockLevel, &out.LockLevel
		*out = new(string)
		**out = **in
	}
	if in.MainTemplate != nil {
		in, out := &in.MainTemplate, &out.MainTemplate
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PackageEnabled != nil {
		in, out := &in.PackageEnabled, &out.PackageEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PackageFileURI != nil {
		in, out := &in.PackageFileURI, &out.PackageFileURI
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinitionSpecResource.
func (in *ApplicationDefinitionSpecResource) DeepCopy() *ApplicationDefinitionSpecResource {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinitionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDefinitionStatus) DeepCopyInto(out *ApplicationDefinitionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDefinitionStatus.
func (in *ApplicationDefinitionStatus) DeepCopy() *ApplicationDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ApplicationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpecPlan) DeepCopyInto(out *ApplicationSpecPlan) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Product != nil {
		in, out := &in.Product, &out.Product
		*out = new(string)
		**out = **in
	}
	if in.PromotionCode != nil {
		in, out := &in.PromotionCode, &out.PromotionCode
		*out = new(string)
		**out = **in
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpecPlan.
func (in *ApplicationSpecPlan) DeepCopy() *ApplicationSpecPlan {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpecPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpecResource) DeepCopyInto(out *ApplicationSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ApplicationDefinitionID != nil {
		in, out := &in.ApplicationDefinitionID, &out.ApplicationDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.ParameterValues != nil {
		in, out := &in.ParameterValues, &out.ParameterValues
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(ApplicationSpecPlan)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpecResource.
func (in *ApplicationSpecResource) DeepCopy() *ApplicationSpecResource {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Disk) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskList) DeepCopyInto(out *DiskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Disk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskList.
func (in *DiskList) DeepCopy() *DiskList {
	if in == nil {
		return nil
	}
	out := new(DiskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpec) DeepCopyInto(out *DiskSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DiskSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpec.
func (in *DiskSpec) DeepCopy() *DiskSpec {
	if in == nil {
		return nil
	}
	out := new(DiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpecEncryptionSettings) DeepCopyInto(out *DiskSpecEncryptionSettings) {
	*out = *in
	if in.DiskEncryptionKey != nil {
		in, out := &in.DiskEncryptionKey, &out.DiskEncryptionKey
		*out = new(DiskSpecEncryptionSettingsDiskEncryptionKey)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyEncryptionKey != nil {
		in, out := &in.KeyEncryptionKey, &out.KeyEncryptionKey
		*out = new(DiskSpecEncryptionSettingsKeyEncryptionKey)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpecEncryptionSettings.
func (in *DiskSpecEncryptionSettings) DeepCopy() *DiskSpecEncryptionSettings {
	if in == nil {
		return nil
	}
	out := new(DiskSpecEncryptionSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpecEncryptionSettingsDiskEncryptionKey) DeepCopyInto(out *DiskSpecEncryptionSettingsDiskEncryptionKey) {
	*out = *in
	if in.SecretURL != nil {
		in, out := &in.SecretURL, &out.SecretURL
		*out = new(string)
		**out = **in
	}
	if in.SourceVaultID != nil {
		in, out := &in.SourceVaultID, &out.SourceVaultID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpecEncryptionSettingsDiskEncryptionKey.
func (in *DiskSpecEncryptionSettingsDiskEncryptionKey) DeepCopy() *DiskSpecEncryptionSettingsDiskEncryptionKey {
	if in == nil {
		return nil
	}
	out := new(DiskSpecEncryptionSettingsDiskEncryptionKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpecEncryptionSettingsKeyEncryptionKey) DeepCopyInto(out *DiskSpecEncryptionSettingsKeyEncryptionKey) {
	*out = *in
	if in.KeyURL != nil {
		in, out := &in.KeyURL, &out.KeyURL
		*out = new(string)
		**out = **in
	}
	if in.SourceVaultID != nil {
		in, out := &in.SourceVaultID, &out.SourceVaultID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpecEncryptionSettingsKeyEncryptionKey.
func (in *DiskSpecEncryptionSettingsKeyEncryptionKey) DeepCopy() *DiskSpecEncryptionSettingsKeyEncryptionKey {
	if in == nil {
		return nil
	}
	out := new(DiskSpecEncryptionSettingsKeyEncryptionKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpecResource) DeepCopyInto(out *DiskSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateOption != nil {
		in, out := &in.CreateOption, &out.CreateOption
		*out = new(string)
		**out = **in
	}
	if in.DiskAccessID != nil {
		in, out := &in.DiskAccessID, &out.DiskAccessID
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSetID != nil {
		in, out := &in.DiskEncryptionSetID, &out.DiskEncryptionSetID
		*out = new(string)
		**out = **in
	}
	if in.DiskIopsReadWrite != nil {
		in, out := &in.DiskIopsReadWrite, &out.DiskIopsReadWrite
		*out = new(int64)
		**out = **in
	}
	if in.DiskMbpsReadWrite != nil {
		in, out := &in.DiskMbpsReadWrite, &out.DiskMbpsReadWrite
		*out = new(int64)
		**out = **in
	}
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.EncryptionSettings != nil {
		in, out := &in.EncryptionSettings, &out.EncryptionSettings
		*out = new(DiskSpecEncryptionSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageReferenceID != nil {
		in, out := &in.ImageReferenceID, &out.ImageReferenceID
		*out = new(string)
		**out = **in
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
	if in.NetworkAccessPolicy != nil {
		in, out := &in.NetworkAccessPolicy, &out.NetworkAccessPolicy
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SourceResourceID != nil {
		in, out := &in.SourceResourceID, &out.SourceResourceID
		*out = new(string)
		**out = **in
	}
	if in.SourceURI != nil {
		in, out := &in.SourceURI, &out.SourceURI
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
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
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpecResource.
func (in *DiskSpecResource) DeepCopy() *DiskSpecResource {
	if in == nil {
		return nil
	}
	out := new(DiskSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskStatus) DeepCopyInto(out *DiskStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskStatus.
func (in *DiskStatus) DeepCopy() *DiskStatus {
	if in == nil {
		return nil
	}
	out := new(DiskStatus)
	in.DeepCopyInto(out)
	return out
}
