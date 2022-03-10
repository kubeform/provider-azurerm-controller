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

type TagRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagRuleSpec   `json:"spec,omitempty"`
	Status            TagRuleStatus `json:"status,omitempty"`
}

type TagRuleSpecTagFilter struct {
	Action *string `json:"action" tf:"action"`
	Name   *string `json:"name" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type TagRuleSpec struct {
	State *TagRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource TagRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TagRuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	LogzMonitorID *string `json:"logzMonitorID" tf:"logz_monitor_id"`
	// +optional
	SendAadLogs *bool `json:"sendAadLogs,omitempty" tf:"send_aad_logs"`
	// +optional
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs"`
	// +optional
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	TagFilter []TagRuleSpecTagFilter `json:"tagFilter,omitempty" tf:"tag_filter"`
}

type TagRuleStatus struct {
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

// TagRuleList is a list of TagRules
type TagRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TagRule CRD objects
	Items []TagRule `json:"items,omitempty"`
}
