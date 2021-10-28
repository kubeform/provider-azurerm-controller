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

type SetDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SetDefinitionSpec   `json:"spec,omitempty"`
	Status            SetDefinitionStatus `json:"status,omitempty"`
}

type SetDefinitionSpecPolicyDefinitionGroup struct {
	// +optional
	AdditionalMetadataResourceID *string `json:"additionalMetadataResourceID,omitempty" tf:"additional_metadata_resource_id"`
	// +optional
	Category *string `json:"category,omitempty" tf:"category"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	Name        *string `json:"name" tf:"name"`
}

type SetDefinitionSpecPolicyDefinitionReference struct {
	// +optional
	ParameterValues *string `json:"parameterValues,omitempty" tf:"parameter_values"`
	// +optional
	// Deprecated
	Parameters         *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	PolicyDefinitionID *string            `json:"policyDefinitionID" tf:"policy_definition_id"`
	// +optional
	PolicyGroupNames []string `json:"policyGroupNames,omitempty" tf:"policy_group_names"`
	// +optional
	ReferenceID *string `json:"referenceID,omitempty" tf:"reference_id"`
}

type SetDefinitionSpec struct {
	State *SetDefinitionSpecResource `json:"state,omitempty" tf:"-"`

	Resource SetDefinitionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SetDefinitionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	// +optional
	// Deprecated
	ManagementGroupID *string `json:"managementGroupID,omitempty" tf:"management_group_id"`
	// +optional
	ManagementGroupName *string `json:"managementGroupName,omitempty" tf:"management_group_name"`
	// +optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	PolicyDefinitionGroup []SetDefinitionSpecPolicyDefinitionGroup `json:"policyDefinitionGroup,omitempty" tf:"policy_definition_group"`
	// +optional
	PolicyDefinitionReference []SetDefinitionSpecPolicyDefinitionReference `json:"policyDefinitionReference,omitempty" tf:"policy_definition_reference"`
	// +optional
	// Deprecated
	PolicyDefinitions *string `json:"policyDefinitions,omitempty" tf:"policy_definitions"`
	PolicyType        *string `json:"policyType" tf:"policy_type"`
}

type SetDefinitionStatus struct {
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

// SetDefinitionList is a list of SetDefinitions
type SetDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SetDefinition CRD objects
	Items []SetDefinition `json:"items,omitempty"`
}
