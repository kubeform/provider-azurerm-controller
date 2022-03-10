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

type Factory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactorySpec   `json:"spec,omitempty"`
	Status            FactoryStatus `json:"status,omitempty"`
}

type FactorySpecGithubConfiguration struct {
	AccountName    *string `json:"accountName" tf:"account_name"`
	BranchName     *string `json:"branchName" tf:"branch_name"`
	GitURL         *string `json:"gitURL" tf:"git_url"`
	RepositoryName *string `json:"repositoryName" tf:"repository_name"`
	RootFolder     *string `json:"rootFolder" tf:"root_folder"`
}

type FactorySpecGlobalParameter struct {
	Name  *string `json:"name" tf:"name"`
	Type  *string `json:"type" tf:"type"`
	Value *string `json:"value" tf:"value"`
}

type FactorySpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type FactorySpecVstsConfiguration struct {
	AccountName    *string `json:"accountName" tf:"account_name"`
	BranchName     *string `json:"branchName" tf:"branch_name"`
	ProjectName    *string `json:"projectName" tf:"project_name"`
	RepositoryName *string `json:"repositoryName" tf:"repository_name"`
	RootFolder     *string `json:"rootFolder" tf:"root_folder"`
	TenantID       *string `json:"tenantID" tf:"tenant_id"`
}

type FactorySpec struct {
	State *FactorySpecResource `json:"state,omitempty" tf:"-"`

	Resource FactorySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FactorySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CustomerManagedKeyID *string `json:"customerManagedKeyID,omitempty" tf:"customer_managed_key_id"`
	// +optional
	GithubConfiguration *FactorySpecGithubConfiguration `json:"githubConfiguration,omitempty" tf:"github_configuration"`
	// +optional
	GlobalParameter []FactorySpecGlobalParameter `json:"globalParameter,omitempty" tf:"global_parameter"`
	// +optional
	Identity *FactorySpecIdentity `json:"identity,omitempty" tf:"identity"`
	Location *string              `json:"location" tf:"location"`
	// +optional
	ManagedVirtualNetworkEnabled *bool   `json:"managedVirtualNetworkEnabled,omitempty" tf:"managed_virtual_network_enabled"`
	Name                         *string `json:"name" tf:"name"`
	// +optional
	PublicNetworkEnabled *bool   `json:"publicNetworkEnabled,omitempty" tf:"public_network_enabled"`
	ResourceGroupName    *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VstsConfiguration *FactorySpecVstsConfiguration `json:"vstsConfiguration,omitempty" tf:"vsts_configuration"`
}

type FactoryStatus struct {
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

// FactoryList is a list of Factorys
type FactoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Factory CRD objects
	Items []Factory `json:"items,omitempty"`
}
