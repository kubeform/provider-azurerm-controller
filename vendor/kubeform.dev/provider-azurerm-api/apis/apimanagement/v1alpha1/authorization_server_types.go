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

type AuthorizationServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthorizationServerSpec   `json:"spec,omitempty"`
	Status            AuthorizationServerStatus `json:"status,omitempty"`
}

type AuthorizationServerSpecTokenBodyParameter struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type AuthorizationServerSpec struct {
	State *AuthorizationServerSpecResource `json:"state,omitempty" tf:"-"`

	Resource AuthorizationServerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AuthorizationServerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiManagementName     *string  `json:"apiManagementName" tf:"api_management_name"`
	AuthorizationEndpoint *string  `json:"authorizationEndpoint" tf:"authorization_endpoint"`
	AuthorizationMethods  []string `json:"authorizationMethods" tf:"authorization_methods"`
	// +optional
	BearerTokenSendingMethods []string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods"`
	// +optional
	ClientAuthenticationMethod []string `json:"clientAuthenticationMethod,omitempty" tf:"client_authentication_method"`
	ClientID                   *string  `json:"clientID" tf:"client_id"`
	ClientRegistrationEndpoint *string  `json:"clientRegistrationEndpoint" tf:"client_registration_endpoint"`
	// +optional
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	DefaultScope *string `json:"defaultScope,omitempty" tf:"default_scope"`
	// +optional
	Description       *string  `json:"description,omitempty" tf:"description"`
	DisplayName       *string  `json:"displayName" tf:"display_name"`
	GrantTypes        []string `json:"grantTypes" tf:"grant_types"`
	Name              *string  `json:"name" tf:"name"`
	ResourceGroupName *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ResourceOwnerPassword *string `json:"-" sensitive:"true" tf:"resource_owner_password"`
	// +optional
	ResourceOwnerUsername *string `json:"resourceOwnerUsername,omitempty" tf:"resource_owner_username"`
	// +optional
	SupportState *bool `json:"supportState,omitempty" tf:"support_state"`
	// +optional
	TokenBodyParameter []AuthorizationServerSpecTokenBodyParameter `json:"tokenBodyParameter,omitempty" tf:"token_body_parameter"`
	// +optional
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint"`
}

type AuthorizationServerStatus struct {
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

// AuthorizationServerList is a list of AuthorizationServers
type AuthorizationServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AuthorizationServer CRD objects
	Items []AuthorizationServer `json:"items,omitempty"`
}
