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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Volume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec,omitempty"`
	Status            VolumeStatus `json:"status,omitempty"`
}

type VolumeSpecDataProtectionReplication struct {
	// +optional
	EndpointType           *string `json:"endpointType,omitempty" tf:"endpoint_type"`
	RemoteVolumeLocation   *string `json:"remoteVolumeLocation" tf:"remote_volume_location"`
	RemoteVolumeResourceID *string `json:"remoteVolumeResourceID" tf:"remote_volume_resource_id"`
	ReplicationFrequency   *string `json:"replicationFrequency" tf:"replication_frequency"`
}

type VolumeSpecDataProtectionSnapshotPolicy struct {
	SnapshotPolicyID *string `json:"snapshotPolicyID" tf:"snapshot_policy_id"`
}

type VolumeSpecExportPolicyRule struct {
	AllowedClients []string `json:"allowedClients" tf:"allowed_clients"`
	// +optional
	// Deprecated
	CifsEnabled *bool `json:"cifsEnabled,omitempty" tf:"cifs_enabled"`
	// +optional
	// Deprecated
	Nfsv3Enabled *bool `json:"nfsv3Enabled,omitempty" tf:"nfsv3_enabled"`
	// +optional
	// Deprecated
	Nfsv4Enabled *bool `json:"nfsv4Enabled,omitempty" tf:"nfsv4_enabled"`
	// +optional
	ProtocolsEnabled []string `json:"protocolsEnabled,omitempty" tf:"protocols_enabled"`
	// +optional
	RootAccessEnabled *bool  `json:"rootAccessEnabled,omitempty" tf:"root_access_enabled"`
	RuleIndex         *int64 `json:"ruleIndex" tf:"rule_index"`
	// +optional
	UnixReadOnly *bool `json:"unixReadOnly,omitempty" tf:"unix_read_only"`
	// +optional
	UnixReadWrite *bool `json:"unixReadWrite,omitempty" tf:"unix_read_write"`
}

type VolumeSpec struct {
	State *VolumeSpecResource `json:"state,omitempty" tf:"-"`

	Resource VolumeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VolumeSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountName *string `json:"accountName" tf:"account_name"`
	// +optional
	CreateFromSnapshotResourceID *string `json:"createFromSnapshotResourceID,omitempty" tf:"create_from_snapshot_resource_id"`
	// +optional
	DataProtectionReplication *VolumeSpecDataProtectionReplication `json:"dataProtectionReplication,omitempty" tf:"data_protection_replication"`
	// +optional
	DataProtectionSnapshotPolicy *VolumeSpecDataProtectionSnapshotPolicy `json:"dataProtectionSnapshotPolicy,omitempty" tf:"data_protection_snapshot_policy"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	ExportPolicyRule []VolumeSpecExportPolicyRule `json:"exportPolicyRule,omitempty" tf:"export_policy_rule"`
	Location         *string                      `json:"location" tf:"location"`
	// +optional
	MountIPAddresses []string `json:"mountIPAddresses,omitempty" tf:"mount_ip_addresses"`
	Name             *string  `json:"name" tf:"name"`
	PoolName         *string  `json:"poolName" tf:"pool_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=2
	Protocols         []string `json:"protocols,omitempty" tf:"protocols"`
	ResourceGroupName *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecurityStyle *string `json:"securityStyle,omitempty" tf:"security_style"`
	ServiceLevel  *string `json:"serviceLevel" tf:"service_level"`
	// +optional
	SnapshotDirectoryVisible *bool   `json:"snapshotDirectoryVisible,omitempty" tf:"snapshot_directory_visible"`
	StorageQuotaInGb         *int64  `json:"storageQuotaInGb" tf:"storage_quota_in_gb"`
	SubnetID                 *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	ThroughputInMibps *float64 `json:"throughputInMibps,omitempty" tf:"throughput_in_mibps"`
	VolumePath        *string  `json:"volumePath" tf:"volume_path"`
}

type VolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VolumeList is a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Volume CRD objects
	Items []Volume `json:"items,omitempty"`
}
