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

type ServiceEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceEnvironmentSpec   `json:"spec,omitempty"`
	Status            ServiceEnvironmentStatus `json:"status,omitempty"`
}

type ServiceEnvironmentSpecClusterSetting struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type ServiceEnvironmentSpec struct {
	State *ServiceEnvironmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceEnvironmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceEnvironmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowedUserIPCidrs []string `json:"allowedUserIPCidrs,omitempty" tf:"allowed_user_ip_cidrs"`
	// +optional
	ClusterSetting []ServiceEnvironmentSpecClusterSetting `json:"clusterSetting,omitempty" tf:"cluster_setting"`
	// +optional
	FrontEndScaleFactor *int64 `json:"frontEndScaleFactor,omitempty" tf:"front_end_scale_factor"`
	// +optional
	InternalIPAddress *string `json:"internalIPAddress,omitempty" tf:"internal_ip_address"`
	// +optional
	InternalLoadBalancingMode *string `json:"internalLoadBalancingMode,omitempty" tf:"internal_load_balancing_mode"`
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	OutboundIPAddresses []string `json:"outboundIPAddresses,omitempty" tf:"outbound_ip_addresses"`
	// +optional
	PricingTier *string `json:"pricingTier,omitempty" tf:"pricing_tier"`
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// +optional
	ServiceIPAddress *string `json:"serviceIPAddress,omitempty" tf:"service_ip_address"`
	SubnetID         *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	// Deprecated
	UserWhitelistedIPRanges []string `json:"userWhitelistedIPRanges,omitempty" tf:"user_whitelisted_ip_ranges"`
}

type ServiceEnvironmentStatus struct {
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

// ServiceEnvironmentList is a list of ServiceEnvironments
type ServiceEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceEnvironment CRD objects
	Items []ServiceEnvironment `json:"items,omitempty"`
}
