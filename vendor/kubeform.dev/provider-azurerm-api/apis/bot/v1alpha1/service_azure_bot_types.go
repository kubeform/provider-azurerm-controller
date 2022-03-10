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

type ServiceAzureBot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAzureBotSpec   `json:"spec,omitempty"`
	Status            ServiceAzureBotStatus `json:"status,omitempty"`
}

type ServiceAzureBotSpec struct {
	State *ServiceAzureBotSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceAzureBotSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceAzureBotSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DeveloperAppInsightsAPIKey *string `json:"-" sensitive:"true" tf:"developer_app_insights_api_key"`
	// +optional
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationID,omitempty" tf:"developer_app_insights_application_id"`
	// +optional
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	Location *string `json:"location" tf:"location"`
	// +optional
	LuisAppIDS []string `json:"luisAppIDS,omitempty" tf:"luis_app_ids"`
	// +optional
	LuisKey           *string `json:"-" sensitive:"true" tf:"luis_key"`
	MicrosoftAppID    *string `json:"microsoftAppID" tf:"microsoft_app_id"`
	Name              *string `json:"name" tf:"name"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	Sku               *string `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ServiceAzureBotStatus struct {
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

// ServiceAzureBotList is a list of ServiceAzureBots
type ServiceAzureBotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceAzureBot CRD objects
	Items []ServiceAzureBot `json:"items,omitempty"`
}
