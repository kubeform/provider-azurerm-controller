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

type ManagedInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedInstanceSpec   `json:"spec,omitempty"`
	Status            ManagedInstanceStatus `json:"status,omitempty"`
}

type ManagedInstanceSpecIdentity struct {
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type ManagedInstanceSpec struct {
	State *ManagedInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagedInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagedInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AdministratorLogin         *string `json:"administratorLogin" tf:"administrator_login"`
	AdministratorLoginPassword *string `json:"-" sensitive:"true" tf:"administrator_login_password"`
	// +optional
	Collation *string `json:"collation,omitempty" tf:"collation"`
	// +optional
	DnsZonePartnerID *string `json:"dnsZonePartnerID,omitempty" tf:"dns_zone_partner_id"`
	// +optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn"`
	// +optional
	Identity    *ManagedInstanceSpecIdentity `json:"identity,omitempty" tf:"identity"`
	LicenseType *string                      `json:"licenseType" tf:"license_type"`
	Location    *string                      `json:"location" tf:"location"`
	// +optional
	MinimumTlsVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version"`
	Name              *string `json:"name" tf:"name"`
	// +optional
	ProxyOverride *string `json:"proxyOverride,omitempty" tf:"proxy_override"`
	// +optional
	PublicDataEndpointEnabled *bool   `json:"publicDataEndpointEnabled,omitempty" tf:"public_data_endpoint_enabled"`
	ResourceGroupName         *string `json:"resourceGroupName" tf:"resource_group_name"`
	SkuName                   *string `json:"skuName" tf:"sku_name"`
	// +optional
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type"`
	StorageSizeInGb    *int64  `json:"storageSizeInGb" tf:"storage_size_in_gb"`
	SubnetID           *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TimezoneID *string `json:"timezoneID,omitempty" tf:"timezone_id"`
	Vcores     *int64  `json:"vcores" tf:"vcores"`
}

type ManagedInstanceStatus struct {
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

// ManagedInstanceList is a list of ManagedInstances
type ManagedInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedInstance CRD objects
	Items []ManagedInstance `json:"items,omitempty"`
}