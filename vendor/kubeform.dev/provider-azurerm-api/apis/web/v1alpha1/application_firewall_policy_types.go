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

type ApplicationFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationFirewallPolicySpec   `json:"spec,omitempty"`
	Status            ApplicationFirewallPolicyStatus `json:"status,omitempty"`
}

type ApplicationFirewallPolicySpecCustomRulesMatchConditionsMatchVariables struct {
	// +optional
	Selector     *string `json:"selector,omitempty" tf:"selector"`
	VariableName *string `json:"variableName" tf:"variable_name"`
}

type ApplicationFirewallPolicySpecCustomRulesMatchConditions struct {
	MatchValues    []string                                                                `json:"matchValues" tf:"match_values"`
	MatchVariables []ApplicationFirewallPolicySpecCustomRulesMatchConditionsMatchVariables `json:"matchVariables" tf:"match_variables"`
	// +optional
	NegationCondition *bool   `json:"negationCondition,omitempty" tf:"negation_condition"`
	Operator          *string `json:"operator" tf:"operator"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type ApplicationFirewallPolicySpecCustomRules struct {
	Action          *string                                                   `json:"action" tf:"action"`
	MatchConditions []ApplicationFirewallPolicySpecCustomRulesMatchConditions `json:"matchConditions" tf:"match_conditions"`
	// +optional
	Name     *string `json:"name,omitempty" tf:"name"`
	Priority *int64  `json:"priority" tf:"priority"`
	RuleType *string `json:"ruleType" tf:"rule_type"`
}

type ApplicationFirewallPolicySpecManagedRulesExclusion struct {
	MatchVariable         *string `json:"matchVariable" tf:"match_variable"`
	Selector              *string `json:"selector" tf:"selector"`
	SelectorMatchOperator *string `json:"selectorMatchOperator" tf:"selector_match_operator"`
}

type ApplicationFirewallPolicySpecManagedRulesManagedRuleSetRuleGroupOverride struct {
	// +optional
	DisabledRules []string `json:"disabledRules,omitempty" tf:"disabled_rules"`
	RuleGroupName *string  `json:"ruleGroupName" tf:"rule_group_name"`
}

type ApplicationFirewallPolicySpecManagedRulesManagedRuleSet struct {
	// +optional
	RuleGroupOverride []ApplicationFirewallPolicySpecManagedRulesManagedRuleSetRuleGroupOverride `json:"ruleGroupOverride,omitempty" tf:"rule_group_override"`
	// +optional
	Type    *string `json:"type,omitempty" tf:"type"`
	Version *string `json:"version" tf:"version"`
}

type ApplicationFirewallPolicySpecManagedRules struct {
	// +optional
	Exclusion      []ApplicationFirewallPolicySpecManagedRulesExclusion      `json:"exclusion,omitempty" tf:"exclusion"`
	ManagedRuleSet []ApplicationFirewallPolicySpecManagedRulesManagedRuleSet `json:"managedRuleSet" tf:"managed_rule_set"`
}

type ApplicationFirewallPolicySpecPolicySettings struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	FileUploadLimitInMb *int64 `json:"fileUploadLimitInMb,omitempty" tf:"file_upload_limit_in_mb"`
	// +optional
	MaxRequestBodySizeInKb *int64 `json:"maxRequestBodySizeInKb,omitempty" tf:"max_request_body_size_in_kb"`
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// +optional
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty" tf:"request_body_check"`
}

type ApplicationFirewallPolicySpec struct {
	State *ApplicationFirewallPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource ApplicationFirewallPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ApplicationFirewallPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CustomRules []ApplicationFirewallPolicySpecCustomRules `json:"customRules,omitempty" tf:"custom_rules"`
	// +optional
	HttpListenerIDS []string                                   `json:"httpListenerIDS,omitempty" tf:"http_listener_ids"`
	Location        *string                                    `json:"location" tf:"location"`
	ManagedRules    *ApplicationFirewallPolicySpecManagedRules `json:"managedRules" tf:"managed_rules"`
	Name            *string                                    `json:"name" tf:"name"`
	// +optional
	PathBasedRuleIDS []string `json:"pathBasedRuleIDS,omitempty" tf:"path_based_rule_ids"`
	// +optional
	PolicySettings    *ApplicationFirewallPolicySpecPolicySettings `json:"policySettings,omitempty" tf:"policy_settings"`
	ResourceGroupName *string                                      `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ApplicationFirewallPolicyStatus struct {
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

// ApplicationFirewallPolicyList is a list of ApplicationFirewallPolicys
type ApplicationFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationFirewallPolicy CRD objects
	Items []ApplicationFirewallPolicy `json:"items,omitempty"`
}
