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

type Interface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InterfaceSpec   `json:"spec,omitempty"`
	Status            InterfaceStatus `json:"status,omitempty"`
}

type InterfaceSpecIpConfiguration struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	Primary *bool `json:"primary,omitempty" tf:"primary"`
	// +optional
	PrivateIPAddress           *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	PrivateIPAddressAllocation *string `json:"privateIPAddressAllocation" tf:"private_ip_address_allocation"`
	// +optional
	PrivateIPAddressVersion *string `json:"privateIPAddressVersion,omitempty" tf:"private_ip_address_version"`
	// +optional
	PublicIPAddressID *string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
}

type InterfaceSpec struct {
	State *InterfaceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InterfaceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InterfaceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppliedDNSServers []string `json:"appliedDNSServers,omitempty" tf:"applied_dns_servers"`
	// +optional
	DnsServers []string `json:"dnsServers,omitempty" tf:"dns_servers"`
	// +optional
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking"`
	// +optional
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty" tf:"enable_ip_forwarding"`
	// +optional
	InternalDNSNameLabel *string `json:"internalDNSNameLabel,omitempty" tf:"internal_dns_name_label"`
	// +optional
	InternalDomainNameSuffix *string                        `json:"internalDomainNameSuffix,omitempty" tf:"internal_domain_name_suffix"`
	IpConfiguration          []InterfaceSpecIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	Location                 *string                        `json:"location" tf:"location"`
	// +optional
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address"`
	Name       *string `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses"`
	ResourceGroupName  *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VirtualMachineID *string `json:"virtualMachineID,omitempty" tf:"virtual_machine_id"`
}

type InterfaceStatus struct {
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

// InterfaceList is a list of Interfaces
type InterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Interface CRD objects
	Items []Interface `json:"items,omitempty"`
}
