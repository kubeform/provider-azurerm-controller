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

type Pubsub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubSpec   `json:"spec,omitempty"`
	Status            PubsubStatus `json:"status,omitempty"`
}

type PubsubSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type PubsubSpecLiveTrace struct {
	// +optional
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	HttpRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled"`
	// +optional
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled"`
}

type PubsubSpec struct {
	State *PubsubSpecResource `json:"state,omitempty" tf:"-"`

	Resource PubsubSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PubsubSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AadAuthEnabled *bool `json:"aadAuthEnabled,omitempty" tf:"aad_auth_enabled"`
	// +optional
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`
	// +optional
	ExternalIP *string `json:"externalIP,omitempty" tf:"external_ip"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// +optional
	Identity *PubsubSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	LiveTrace *PubsubSpecLiveTrace `json:"liveTrace,omitempty" tf:"live_trace"`
	// +optional
	LocalAuthEnabled *bool   `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled"`
	Location         *string `json:"location" tf:"location"`
	Name             *string `json:"name" tf:"name"`
	// +optional
	PrimaryAccessKey *string `json:"-" sensitive:"true" tf:"primary_access_key"`
	// +optional
	PrimaryConnectionString *string `json:"-" sensitive:"true" tf:"primary_connection_string"`
	// +optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	// +optional
	PublicPort        *int64  `json:"publicPort,omitempty" tf:"public_port"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey *string `json:"-" sensitive:"true" tf:"secondary_access_key"`
	// +optional
	SecondaryConnectionString *string `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	// +optional
	ServerPort *int64  `json:"serverPort,omitempty" tf:"server_port"`
	Sku        *string `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TlsClientCertEnabled *bool `json:"tlsClientCertEnabled,omitempty" tf:"tls_client_cert_enabled"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type PubsubStatus struct {
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

// PubsubList is a list of Pubsubs
type PubsubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Pubsub CRD objects
	Items []Pubsub `json:"items,omitempty"`
}