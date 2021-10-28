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

type DiagnosticSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiagnosticSettingSpec   `json:"spec,omitempty"`
	Status            DiagnosticSettingStatus `json:"status,omitempty"`
}

type DiagnosticSettingSpecLogRetentionPolicy struct {
	// +optional
	Days    *int64 `json:"days,omitempty" tf:"days"`
	Enabled *bool  `json:"enabled" tf:"enabled"`
}

type DiagnosticSettingSpecLog struct {
	Category *string `json:"category" tf:"category"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	RetentionPolicy *DiagnosticSettingSpecLogRetentionPolicy `json:"retentionPolicy,omitempty" tf:"retention_policy"`
}

type DiagnosticSettingSpecMetricRetentionPolicy struct {
	// +optional
	Days    *int64 `json:"days,omitempty" tf:"days"`
	Enabled *bool  `json:"enabled" tf:"enabled"`
}

type DiagnosticSettingSpecMetric struct {
	Category *string `json:"category" tf:"category"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	RetentionPolicy *DiagnosticSettingSpecMetricRetentionPolicy `json:"retentionPolicy,omitempty" tf:"retention_policy"`
}

type DiagnosticSettingSpec struct {
	State *DiagnosticSettingSpecResource `json:"state,omitempty" tf:"-"`

	Resource DiagnosticSettingSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DiagnosticSettingSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EventhubAuthorizationRuleID *string `json:"eventhubAuthorizationRuleID,omitempty" tf:"eventhub_authorization_rule_id"`
	// +optional
	EventhubName *string `json:"eventhubName,omitempty" tf:"eventhub_name"`
	// +optional
	Log []DiagnosticSettingSpecLog `json:"log,omitempty" tf:"log"`
	// +optional
	LogAnalyticsDestinationType *string `json:"logAnalyticsDestinationType,omitempty" tf:"log_analytics_destination_type"`
	// +optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceID,omitempty" tf:"log_analytics_workspace_id"`
	// +optional
	Metric []DiagnosticSettingSpecMetric `json:"metric,omitempty" tf:"metric"`
	Name   *string                       `json:"name" tf:"name"`
	// +optional
	StorageAccountID *string `json:"storageAccountID,omitempty" tf:"storage_account_id"`
	TargetResourceID *string `json:"targetResourceID" tf:"target_resource_id"`
}

type DiagnosticSettingStatus struct {
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

// DiagnosticSettingList is a list of DiagnosticSettings
type DiagnosticSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DiagnosticSetting CRD objects
	Items []DiagnosticSetting `json:"items,omitempty"`
}
