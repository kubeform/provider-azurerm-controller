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

type ActionRuleSuppression struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActionRuleSuppressionSpec   `json:"spec,omitempty"`
	Status            ActionRuleSuppressionStatus `json:"status,omitempty"`
}

type ActionRuleSuppressionSpecConditionAlertContext struct {
	Operator *string  `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type ActionRuleSuppressionSpecConditionAlertRuleID struct {
	Operator *string  `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type ActionRuleSuppressionSpecConditionDescription struct {
	Operator *string  `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type ActionRuleSuppressionSpecConditionMonitor struct {
	Operator *string  `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type ActionRuleSuppressionSpecConditionMonitorService struct {
	Operator *string  `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type ActionRuleSuppressionSpecConditionSeverity struct {
	Operator *string  `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type ActionRuleSuppressionSpecConditionTargetResourceType struct {
	Operator *string  `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type ActionRuleSuppressionSpecCondition struct {
	// +optional
	AlertContext *ActionRuleSuppressionSpecConditionAlertContext `json:"alertContext,omitempty" tf:"alert_context"`
	// +optional
	AlertRuleID *ActionRuleSuppressionSpecConditionAlertRuleID `json:"alertRuleID,omitempty" tf:"alert_rule_id"`
	// +optional
	Description *ActionRuleSuppressionSpecConditionDescription `json:"description,omitempty" tf:"description"`
	// +optional
	Monitor *ActionRuleSuppressionSpecConditionMonitor `json:"monitor,omitempty" tf:"monitor"`
	// +optional
	MonitorService *ActionRuleSuppressionSpecConditionMonitorService `json:"monitorService,omitempty" tf:"monitor_service"`
	// +optional
	Severity *ActionRuleSuppressionSpecConditionSeverity `json:"severity,omitempty" tf:"severity"`
	// +optional
	TargetResourceType *ActionRuleSuppressionSpecConditionTargetResourceType `json:"targetResourceType,omitempty" tf:"target_resource_type"`
}

type ActionRuleSuppressionSpecScope struct {
	ResourceIDS []string `json:"resourceIDS" tf:"resource_ids"`
	Type        *string  `json:"type" tf:"type"`
}

type ActionRuleSuppressionSpecSuppressionSchedule struct {
	EndDateUtc *string `json:"endDateUtc" tf:"end_date_utc"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	RecurrenceMonthly []int64 `json:"recurrenceMonthly,omitempty" tf:"recurrence_monthly"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	RecurrenceWeekly []string `json:"recurrenceWeekly,omitempty" tf:"recurrence_weekly"`
	StartDateUtc     *string  `json:"startDateUtc" tf:"start_date_utc"`
}

type ActionRuleSuppressionSpecSuppression struct {
	RecurrenceType *string `json:"recurrenceType" tf:"recurrence_type"`
	// +optional
	Schedule *ActionRuleSuppressionSpecSuppressionSchedule `json:"schedule,omitempty" tf:"schedule"`
}

type ActionRuleSuppressionSpec struct {
	State *ActionRuleSuppressionSpecResource `json:"state,omitempty" tf:"-"`

	Resource ActionRuleSuppressionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ActionRuleSuppressionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Condition *ActionRuleSuppressionSpecCondition `json:"condition,omitempty" tf:"condition"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Enabled           *bool   `json:"enabled,omitempty" tf:"enabled"`
	Name              *string `json:"name" tf:"name"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Scope       *ActionRuleSuppressionSpecScope       `json:"scope,omitempty" tf:"scope"`
	Suppression *ActionRuleSuppressionSpecSuppression `json:"suppression" tf:"suppression"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ActionRuleSuppressionStatus struct {
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

// ActionRuleSuppressionList is a list of ActionRuleSuppressions
type ActionRuleSuppressionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ActionRuleSuppression CRD objects
	Items []ActionRuleSuppression `json:"items,omitempty"`
}
