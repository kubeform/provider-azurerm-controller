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

type FactoryDatasetPostgresql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryDatasetPostgresqlSpec   `json:"spec,omitempty"`
	Status            FactoryDatasetPostgresqlStatus `json:"status,omitempty"`
}

type FactoryDatasetPostgresqlSpecSchemaColumn struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type FactoryDatasetPostgresqlSpec struct {
	State *FactoryDatasetPostgresqlSpecResource `json:"state,omitempty" tf:"-"`

	Resource FactoryDatasetPostgresqlSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FactoryDatasetPostgresqlSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`
	// +optional
	Annotations []string `json:"annotations,omitempty" tf:"annotations"`
	// +optional
	DataFactoryID *string `json:"dataFactoryID,omitempty" tf:"data_factory_id"`
	// +optional
	// Deprecated
	DataFactoryName *string `json:"dataFactoryName,omitempty" tf:"data_factory_name"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Folder            *string `json:"folder,omitempty" tf:"folder"`
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name"`
	Name              *string `json:"name" tf:"name"`
	// +optional
	Parameters        *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	ResourceGroupName *string            `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SchemaColumn []FactoryDatasetPostgresqlSpecSchemaColumn `json:"schemaColumn,omitempty" tf:"schema_column"`
	// +optional
	TableName *string `json:"tableName,omitempty" tf:"table_name"`
}

type FactoryDatasetPostgresqlStatus struct {
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

// FactoryDatasetPostgresqlList is a list of FactoryDatasetPostgresqls
type FactoryDatasetPostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FactoryDatasetPostgresql CRD objects
	Items []FactoryDatasetPostgresql `json:"items,omitempty"`
}
