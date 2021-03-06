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

type FactoryIntegrationRuntimeManaged struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryIntegrationRuntimeManagedSpec   `json:"spec,omitempty"`
	Status            FactoryIntegrationRuntimeManagedStatus `json:"status,omitempty"`
}

type FactoryIntegrationRuntimeManagedSpecCatalogInfo struct {
	// +optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login"`
	// +optional
	AdministratorPassword *string `json:"-" sensitive:"true" tf:"administrator_password"`
	// +optional
	PricingTier    *string `json:"pricingTier,omitempty" tf:"pricing_tier"`
	ServerEndpoint *string `json:"serverEndpoint" tf:"server_endpoint"`
}

type FactoryIntegrationRuntimeManagedSpecCustomSetupScript struct {
	BlobContainerURI *string `json:"blobContainerURI" tf:"blob_container_uri"`
	SasToken         *string `json:"-" sensitive:"true" tf:"sas_token"`
}

type FactoryIntegrationRuntimeManagedSpecVnetIntegration struct {
	SubnetName *string `json:"subnetName" tf:"subnet_name"`
	VnetID     *string `json:"vnetID" tf:"vnet_id"`
}

type FactoryIntegrationRuntimeManagedSpec struct {
	State *FactoryIntegrationRuntimeManagedSpecResource `json:"state,omitempty" tf:"-"`

	Resource FactoryIntegrationRuntimeManagedSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FactoryIntegrationRuntimeManagedSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CatalogInfo *FactoryIntegrationRuntimeManagedSpecCatalogInfo `json:"catalogInfo,omitempty" tf:"catalog_info"`
	// +optional
	CustomSetupScript *FactoryIntegrationRuntimeManagedSpecCustomSetupScript `json:"customSetupScript,omitempty" tf:"custom_setup_script"`
	DataFactoryName   *string                                                `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Edition *string `json:"edition,omitempty" tf:"edition"`
	// +optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type"`
	Location    *string `json:"location" tf:"location"`
	// +optional
	MaxParallelExecutionsPerNode *int64  `json:"maxParallelExecutionsPerNode,omitempty" tf:"max_parallel_executions_per_node"`
	Name                         *string `json:"name" tf:"name"`
	NodeSize                     *string `json:"nodeSize" tf:"node_size"`
	// +optional
	NumberOfNodes     *int64  `json:"numberOfNodes,omitempty" tf:"number_of_nodes"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	VnetIntegration *FactoryIntegrationRuntimeManagedSpecVnetIntegration `json:"vnetIntegration,omitempty" tf:"vnet_integration"`
}

type FactoryIntegrationRuntimeManagedStatus struct {
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

// FactoryIntegrationRuntimeManagedList is a list of FactoryIntegrationRuntimeManageds
type FactoryIntegrationRuntimeManagedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FactoryIntegrationRuntimeManaged CRD objects
	Items []FactoryIntegrationRuntimeManaged `json:"items,omitempty"`
}
