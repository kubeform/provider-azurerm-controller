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

type TriggerRecurrence struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerRecurrenceSpec   `json:"spec,omitempty"`
	Status            TriggerRecurrenceStatus `json:"status,omitempty"`
}

type TriggerRecurrenceSpecSchedule struct {
	// +optional
	AtTheseHours []int64 `json:"atTheseHours,omitempty" tf:"at_these_hours"`
	// +optional
	AtTheseMinutes []int64 `json:"atTheseMinutes,omitempty" tf:"at_these_minutes"`
	// +optional
	OnTheseDays []string `json:"onTheseDays,omitempty" tf:"on_these_days"`
}

type TriggerRecurrenceSpec struct {
	State *TriggerRecurrenceSpecResource `json:"state,omitempty" tf:"-"`

	Resource TriggerRecurrenceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TriggerRecurrenceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Frequency  *string `json:"frequency" tf:"frequency"`
	Interval   *int64  `json:"interval" tf:"interval"`
	LogicAppID *string `json:"logicAppID" tf:"logic_app_id"`
	Name       *string `json:"name" tf:"name"`
	// +optional
	Schedule *TriggerRecurrenceSpecSchedule `json:"schedule,omitempty" tf:"schedule"`
	// +optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time"`
	// +optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
}

type TriggerRecurrenceStatus struct {
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

// TriggerRecurrenceList is a list of TriggerRecurrences
type TriggerRecurrenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TriggerRecurrence CRD objects
	Items []TriggerRecurrence `json:"items,omitempty"`
}
