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

type Probe struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProbeSpec   `json:"spec,omitempty"`
	Status            ProbeStatus `json:"status,omitempty"`
}

type ProbeSpec struct {
	State *ProbeSpecResource `json:"state,omitempty" tf:"-"`

	Resource ProbeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ProbeSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds"`
	// +optional
	LoadBalancerRules []string `json:"loadBalancerRules,omitempty" tf:"load_balancer_rules"`
	LoadbalancerID    *string  `json:"loadbalancerID" tf:"loadbalancer_id"`
	Name              *string  `json:"name" tf:"name"`
	// +optional
	NumberOfProbes *int64 `json:"numberOfProbes,omitempty" tf:"number_of_probes"`
	Port           *int64 `json:"port" tf:"port"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	RequestPath       *string `json:"requestPath,omitempty" tf:"request_path"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
}

type ProbeStatus struct {
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

// ProbeList is a list of Probes
type ProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Probe CRD objects
	Items []Probe `json:"items,omitempty"`
}
