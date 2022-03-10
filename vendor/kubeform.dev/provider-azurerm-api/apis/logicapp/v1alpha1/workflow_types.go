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

type Workflow struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkflowSpec   `json:"spec,omitempty"`
	Status            WorkflowStatus `json:"status,omitempty"`
}

type WorkflowSpecAccessControlAction struct {
	AllowedCallerIPAddressRange []string `json:"allowedCallerIPAddressRange" tf:"allowed_caller_ip_address_range"`
}

type WorkflowSpecAccessControlContent struct {
	AllowedCallerIPAddressRange []string `json:"allowedCallerIPAddressRange" tf:"allowed_caller_ip_address_range"`
}

type WorkflowSpecAccessControlTriggerOpenAuthenticationPolicyClaim struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type WorkflowSpecAccessControlTriggerOpenAuthenticationPolicy struct {
	Claim []WorkflowSpecAccessControlTriggerOpenAuthenticationPolicyClaim `json:"claim" tf:"claim"`
	Name  *string                                                         `json:"name" tf:"name"`
}

type WorkflowSpecAccessControlTrigger struct {
	AllowedCallerIPAddressRange []string `json:"allowedCallerIPAddressRange" tf:"allowed_caller_ip_address_range"`
	// +optional
	OpenAuthenticationPolicy []WorkflowSpecAccessControlTriggerOpenAuthenticationPolicy `json:"openAuthenticationPolicy,omitempty" tf:"open_authentication_policy"`
}

type WorkflowSpecAccessControlWorkflowManagement struct {
	AllowedCallerIPAddressRange []string `json:"allowedCallerIPAddressRange" tf:"allowed_caller_ip_address_range"`
}

type WorkflowSpecAccessControl struct {
	// +optional
	Action *WorkflowSpecAccessControlAction `json:"action,omitempty" tf:"action"`
	// +optional
	Content *WorkflowSpecAccessControlContent `json:"content,omitempty" tf:"content"`
	// +optional
	Trigger *WorkflowSpecAccessControlTrigger `json:"trigger,omitempty" tf:"trigger"`
	// +optional
	WorkflowManagement *WorkflowSpecAccessControlWorkflowManagement `json:"workflowManagement,omitempty" tf:"workflow_management"`
}

type WorkflowSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type WorkflowSpec struct {
	State *WorkflowSpecResource `json:"state,omitempty" tf:"-"`

	Resource WorkflowSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type WorkflowSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccessControl *WorkflowSpecAccessControl `json:"accessControl,omitempty" tf:"access_control"`
	// +optional
	AccessEndpoint *string `json:"accessEndpoint,omitempty" tf:"access_endpoint"`
	// +optional
	ConnectorEndpointIPAddresses []string `json:"connectorEndpointIPAddresses,omitempty" tf:"connector_endpoint_ip_addresses"`
	// +optional
	ConnectorOutboundIPAddresses []string `json:"connectorOutboundIPAddresses,omitempty" tf:"connector_outbound_ip_addresses"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	Identity *WorkflowSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	IntegrationServiceEnvironmentID *string `json:"integrationServiceEnvironmentID,omitempty" tf:"integration_service_environment_id"`
	Location                        *string `json:"location" tf:"location"`
	// +optional
	LogicAppIntegrationAccountID *string `json:"logicAppIntegrationAccountID,omitempty" tf:"logic_app_integration_account_id"`
	Name                         *string `json:"name" tf:"name"`
	// +optional
	Parameters        *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	ResourceGroupName *string            `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	WorkflowEndpointIPAddresses []string `json:"workflowEndpointIPAddresses,omitempty" tf:"workflow_endpoint_ip_addresses"`
	// +optional
	WorkflowOutboundIPAddresses []string `json:"workflowOutboundIPAddresses,omitempty" tf:"workflow_outbound_ip_addresses"`
	// +optional
	WorkflowParameters *map[string]string `json:"workflowParameters,omitempty" tf:"workflow_parameters"`
	// +optional
	WorkflowSchema *string `json:"workflowSchema,omitempty" tf:"workflow_schema"`
	// +optional
	WorkflowVersion *string `json:"workflowVersion,omitempty" tf:"workflow_version"`
}

type WorkflowStatus struct {
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

// WorkflowList is a list of Workflows
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Workflow CRD objects
	Items []Workflow `json:"items,omitempty"`
}
