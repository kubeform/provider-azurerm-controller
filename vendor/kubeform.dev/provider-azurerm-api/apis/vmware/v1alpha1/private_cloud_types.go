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

type PrivateCloud struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateCloudSpec   `json:"spec,omitempty"`
	Status            PrivateCloudStatus `json:"status,omitempty"`
}

type PrivateCloudSpecCircuit struct {
	// +optional
	ExpressRouteID *string `json:"expressRouteID,omitempty" tf:"express_route_id"`
	// +optional
	ExpressRoutePrivatePeeringID *string `json:"expressRoutePrivatePeeringID,omitempty" tf:"express_route_private_peering_id"`
	// +optional
	PrimarySubnetCIDR *string `json:"primarySubnetCIDR,omitempty" tf:"primary_subnet_cidr"`
	// +optional
	SecondarySubnetCIDR *string `json:"secondarySubnetCIDR,omitempty" tf:"secondary_subnet_cidr"`
}

type PrivateCloudSpecManagementCluster struct {
	// +optional
	Hosts []string `json:"hosts,omitempty" tf:"hosts"`
	// +optional
	ID   *int64 `json:"ID,omitempty" tf:"id"`
	Size *int64 `json:"size" tf:"size"`
}

type PrivateCloudSpec struct {
	State *PrivateCloudSpecResource `json:"state,omitempty" tf:"-"`

	Resource PrivateCloudSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PrivateCloudSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Circuit []PrivateCloudSpecCircuit `json:"circuit,omitempty" tf:"circuit"`
	// +optional
	HcxCloudManagerEndpoint *string `json:"hcxCloudManagerEndpoint,omitempty" tf:"hcx_cloud_manager_endpoint"`
	// +optional
	InternetConnectionEnabled *bool                              `json:"internetConnectionEnabled,omitempty" tf:"internet_connection_enabled"`
	Location                  *string                            `json:"location" tf:"location"`
	ManagementCluster         *PrivateCloudSpecManagementCluster `json:"managementCluster" tf:"management_cluster"`
	// +optional
	ManagementSubnetCIDR *string `json:"managementSubnetCIDR,omitempty" tf:"management_subnet_cidr"`
	Name                 *string `json:"name" tf:"name"`
	NetworkSubnetCIDR    *string `json:"networkSubnetCIDR" tf:"network_subnet_cidr"`
	// +optional
	NsxtCertificateThumbprint *string `json:"nsxtCertificateThumbprint,omitempty" tf:"nsxt_certificate_thumbprint"`
	// +optional
	NsxtManagerEndpoint *string `json:"nsxtManagerEndpoint,omitempty" tf:"nsxt_manager_endpoint"`
	// +optional
	NsxtPassword *string `json:"-" sensitive:"true" tf:"nsxt_password"`
	// +optional
	ProvisioningSubnetCIDR *string `json:"provisioningSubnetCIDR,omitempty" tf:"provisioning_subnet_cidr"`
	ResourceGroupName      *string `json:"resourceGroupName" tf:"resource_group_name"`
	SkuName                *string `json:"skuName" tf:"sku_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VcenterCertificateThumbprint *string `json:"vcenterCertificateThumbprint,omitempty" tf:"vcenter_certificate_thumbprint"`
	// +optional
	VcenterPassword *string `json:"-" sensitive:"true" tf:"vcenter_password"`
	// +optional
	VcsaEndpoint *string `json:"vcsaEndpoint,omitempty" tf:"vcsa_endpoint"`
	// +optional
	VmotionSubnetCIDR *string `json:"vmotionSubnetCIDR,omitempty" tf:"vmotion_subnet_cidr"`
}

type PrivateCloudStatus struct {
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

// PrivateCloudList is a list of PrivateClouds
type PrivateCloudList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PrivateCloud CRD objects
	Items []PrivateCloud `json:"items,omitempty"`
}
