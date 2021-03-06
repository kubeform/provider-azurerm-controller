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

type CloudAppCosmosdbAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudAppCosmosdbAssociationSpec   `json:"spec,omitempty"`
	Status            CloudAppCosmosdbAssociationStatus `json:"status,omitempty"`
}

type CloudAppCosmosdbAssociationSpec struct {
	State *CloudAppCosmosdbAssociationSpecResource `json:"state,omitempty" tf:"-"`

	Resource CloudAppCosmosdbAssociationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CloudAppCosmosdbAssociationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiType           *string `json:"apiType" tf:"api_type"`
	CosmosdbAccessKey *string `json:"cosmosdbAccessKey" tf:"cosmosdb_access_key"`
	CosmosdbAccountID *string `json:"cosmosdbAccountID" tf:"cosmosdb_account_id"`
	// +optional
	CosmosdbCassandraKeyspaceName *string `json:"cosmosdbCassandraKeyspaceName,omitempty" tf:"cosmosdb_cassandra_keyspace_name"`
	// +optional
	CosmosdbGremlinDatabaseName *string `json:"cosmosdbGremlinDatabaseName,omitempty" tf:"cosmosdb_gremlin_database_name"`
	// +optional
	CosmosdbGremlinGraphName *string `json:"cosmosdbGremlinGraphName,omitempty" tf:"cosmosdb_gremlin_graph_name"`
	// +optional
	CosmosdbMongoDatabaseName *string `json:"cosmosdbMongoDatabaseName,omitempty" tf:"cosmosdb_mongo_database_name"`
	// +optional
	CosmosdbSQLDatabaseName *string `json:"cosmosdbSQLDatabaseName,omitempty" tf:"cosmosdb_sql_database_name"`
	Name                    *string `json:"name" tf:"name"`
	SpringCloudAppID        *string `json:"springCloudAppID" tf:"spring_cloud_app_id"`
}

type CloudAppCosmosdbAssociationStatus struct {
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

// CloudAppCosmosdbAssociationList is a list of CloudAppCosmosdbAssociations
type CloudAppCosmosdbAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudAppCosmosdbAssociation CRD objects
	Items []CloudAppCosmosdbAssociation `json:"items,omitempty"`
}
