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
func (in *Cache) DeepCopyInto(out *Cache) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cache.
func (in *Cache) DeepCopy() *Cache {
	if in == nil {
		return nil
	}
	out := new(Cache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cache) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheList) DeepCopyInto(out *CacheList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cache, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheList.
func (in *CacheList) DeepCopy() *CacheList {
	if in == nil {
		return nil
	}
	out := new(CacheList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CacheList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSpec) DeepCopyInto(out *CacheSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CacheSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSpec.
func (in *CacheSpec) DeepCopy() *CacheSpec {
	if in == nil {
		return nil
	}
	out := new(CacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSpecPatchSchedule) DeepCopyInto(out *CacheSpecPatchSchedule) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSpecPatchSchedule.
func (in *CacheSpecPatchSchedule) DeepCopy() *CacheSpecPatchSchedule {
	if in == nil {
		return nil
	}
	out := new(CacheSpecPatchSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSpecRedisConfiguration) DeepCopyInto(out *CacheSpecRedisConfiguration) {
	*out = *in
	if in.AofBackupEnabled != nil {
		in, out := &in.AofBackupEnabled, &out.AofBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofStorageConnectionString0 != nil {
		in, out := &in.AofStorageConnectionString0, &out.AofStorageConnectionString0
		*out = new(string)
		**out = **in
	}
	if in.AofStorageConnectionString1 != nil {
		in, out := &in.AofStorageConnectionString1, &out.AofStorageConnectionString1
		*out = new(string)
		**out = **in
	}
	if in.EnableAuthentication != nil {
		in, out := &in.EnableAuthentication, &out.EnableAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.Maxclients != nil {
		in, out := &in.Maxclients, &out.Maxclients
		*out = new(int64)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(int64)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(int64)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
		*out = new(int64)
		**out = **in
	}
	if in.NotifyKeyspaceEvents != nil {
		in, out := &in.NotifyKeyspaceEvents, &out.NotifyKeyspaceEvents
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(int64)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(int64)
		**out = **in
	}
	if in.RdbStorageConnectionString != nil {
		in, out := &in.RdbStorageConnectionString, &out.RdbStorageConnectionString
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSpecRedisConfiguration.
func (in *CacheSpecRedisConfiguration) DeepCopy() *CacheSpecRedisConfiguration {
	if in == nil {
		return nil
	}
	out := new(CacheSpecRedisConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSpecResource) DeepCopyInto(out *CacheSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int64)
		**out = **in
	}
	if in.EnableNonSslPort != nil {
		in, out := &in.EnableNonSslPort, &out.EnableNonSslPort
		*out = new(bool)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PatchSchedule != nil {
		in, out := &in.PatchSchedule, &out.PatchSchedule
		*out = make([]CacheSpecPatchSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.PrimaryAccessKey != nil {
		in, out := &in.PrimaryAccessKey, &out.PrimaryAccessKey
		*out = new(string)
		**out = **in
	}
	if in.PrimaryConnectionString != nil {
		in, out := &in.PrimaryConnectionString, &out.PrimaryConnectionString
		*out = new(string)
		**out = **in
	}
	if in.PrivateStaticIPAddress != nil {
		in, out := &in.PrivateStaticIPAddress, &out.PrivateStaticIPAddress
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(CacheSpecRedisConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(int64)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SecondaryAccessKey != nil {
		in, out := &in.SecondaryAccessKey, &out.SecondaryAccessKey
		*out = new(string)
		**out = **in
	}
	if in.SecondaryConnectionString != nil {
		in, out := &in.SecondaryConnectionString, &out.SecondaryConnectionString
		*out = new(string)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int64)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SslPort != nil {
		in, out := &in.SslPort, &out.SslPort
		*out = new(int64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
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
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSpecResource.
func (in *CacheSpecResource) DeepCopy() *CacheSpecResource {
	if in == nil {
		return nil
	}
	out := new(CacheSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheStatus) DeepCopyInto(out *CacheStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheStatus.
func (in *CacheStatus) DeepCopy() *CacheStatus {
	if in == nil {
		return nil
	}
	out := new(CacheStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseCluster) DeepCopyInto(out *EnterpriseCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseCluster.
func (in *EnterpriseCluster) DeepCopy() *EnterpriseCluster {
	if in == nil {
		return nil
	}
	out := new(EnterpriseCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseClusterList) DeepCopyInto(out *EnterpriseClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnterpriseCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseClusterList.
func (in *EnterpriseClusterList) DeepCopy() *EnterpriseClusterList {
	if in == nil {
		return nil
	}
	out := new(EnterpriseClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseClusterSpec) DeepCopyInto(out *EnterpriseClusterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(EnterpriseClusterSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseClusterSpec.
func (in *EnterpriseClusterSpec) DeepCopy() *EnterpriseClusterSpec {
	if in == nil {
		return nil
	}
	out := new(EnterpriseClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseClusterSpecResource) DeepCopyInto(out *EnterpriseClusterSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(string)
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
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
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
	if in.Version != nil {
		in, out := &in.Version, &out.Version
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseClusterSpecResource.
func (in *EnterpriseClusterSpecResource) DeepCopy() *EnterpriseClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(EnterpriseClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseClusterStatus) DeepCopyInto(out *EnterpriseClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseClusterStatus.
func (in *EnterpriseClusterStatus) DeepCopy() *EnterpriseClusterStatus {
	if in == nil {
		return nil
	}
	out := new(EnterpriseClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseDatabase) DeepCopyInto(out *EnterpriseDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseDatabase.
func (in *EnterpriseDatabase) DeepCopy() *EnterpriseDatabase {
	if in == nil {
		return nil
	}
	out := new(EnterpriseDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseDatabaseList) DeepCopyInto(out *EnterpriseDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnterpriseDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseDatabaseList.
func (in *EnterpriseDatabaseList) DeepCopy() *EnterpriseDatabaseList {
	if in == nil {
		return nil
	}
	out := new(EnterpriseDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseDatabaseSpec) DeepCopyInto(out *EnterpriseDatabaseSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(EnterpriseDatabaseSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseDatabaseSpec.
func (in *EnterpriseDatabaseSpec) DeepCopy() *EnterpriseDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(EnterpriseDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseDatabaseSpecModule) DeepCopyInto(out *EnterpriseDatabaseSpecModule) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseDatabaseSpecModule.
func (in *EnterpriseDatabaseSpecModule) DeepCopy() *EnterpriseDatabaseSpecModule {
	if in == nil {
		return nil
	}
	out := new(EnterpriseDatabaseSpecModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseDatabaseSpecResource) DeepCopyInto(out *EnterpriseDatabaseSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(string)
		**out = **in
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(string)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(string)
		**out = **in
	}
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = make([]EnterpriseDatabaseSpecModule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.PrimaryAccessKey != nil {
		in, out := &in.PrimaryAccessKey, &out.PrimaryAccessKey
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SecondaryAccessKey != nil {
		in, out := &in.SecondaryAccessKey, &out.SecondaryAccessKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseDatabaseSpecResource.
func (in *EnterpriseDatabaseSpecResource) DeepCopy() *EnterpriseDatabaseSpecResource {
	if in == nil {
		return nil
	}
	out := new(EnterpriseDatabaseSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseDatabaseStatus) DeepCopyInto(out *EnterpriseDatabaseStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseDatabaseStatus.
func (in *EnterpriseDatabaseStatus) DeepCopy() *EnterpriseDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(EnterpriseDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

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
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
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
	if in.EndIP != nil {
		in, out := &in.EndIP, &out.EndIP
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RedisCacheName != nil {
		in, out := &in.RedisCacheName, &out.RedisCacheName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.StartIP != nil {
		in, out := &in.StartIP, &out.StartIP
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
		*out = make([]apiv1.Condition, len(*in))
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
func (in *LinkedServer) DeepCopyInto(out *LinkedServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedServer.
func (in *LinkedServer) DeepCopy() *LinkedServer {
	if in == nil {
		return nil
	}
	out := new(LinkedServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LinkedServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedServerList) DeepCopyInto(out *LinkedServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LinkedServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedServerList.
func (in *LinkedServerList) DeepCopy() *LinkedServerList {
	if in == nil {
		return nil
	}
	out := new(LinkedServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LinkedServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedServerSpec) DeepCopyInto(out *LinkedServerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LinkedServerSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedServerSpec.
func (in *LinkedServerSpec) DeepCopy() *LinkedServerSpec {
	if in == nil {
		return nil
	}
	out := new(LinkedServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedServerSpecResource) DeepCopyInto(out *LinkedServerSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.LinkedRedisCacheID != nil {
		in, out := &in.LinkedRedisCacheID, &out.LinkedRedisCacheID
		*out = new(string)
		**out = **in
	}
	if in.LinkedRedisCacheLocation != nil {
		in, out := &in.LinkedRedisCacheLocation, &out.LinkedRedisCacheLocation
		*out = new(string)
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
	if in.ServerRole != nil {
		in, out := &in.ServerRole, &out.ServerRole
		*out = new(string)
		**out = **in
	}
	if in.TargetRedisCacheName != nil {
		in, out := &in.TargetRedisCacheName, &out.TargetRedisCacheName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedServerSpecResource.
func (in *LinkedServerSpecResource) DeepCopy() *LinkedServerSpecResource {
	if in == nil {
		return nil
	}
	out := new(LinkedServerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedServerStatus) DeepCopyInto(out *LinkedServerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedServerStatus.
func (in *LinkedServerStatus) DeepCopy() *LinkedServerStatus {
	if in == nil {
		return nil
	}
	out := new(LinkedServerStatus)
	in.DeepCopyInto(out)
	return out
}
