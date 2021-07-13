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

type ApiOperation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiOperationSpec   `json:"spec,omitempty"`
	Status            ApiOperationStatus `json:"status,omitempty"`
}

type ApiOperationSpecRequestHeader struct {
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	Required    *bool   `json:"required" tf:"required"`
	Type        *string `json:"type" tf:"type"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type ApiOperationSpecRequestQueryParameter struct {
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	Required    *bool   `json:"required" tf:"required"`
	Type        *string `json:"type" tf:"type"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type ApiOperationSpecRequestRepresentationFormParameter struct {
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	Required    *bool   `json:"required" tf:"required"`
	Type        *string `json:"type" tf:"type"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type ApiOperationSpecRequestRepresentation struct {
	ContentType *string `json:"contentType" tf:"content_type"`
	// +optional
	FormParameter []ApiOperationSpecRequestRepresentationFormParameter `json:"formParameter,omitempty" tf:"form_parameter"`
	// +optional
	Sample *string `json:"sample,omitempty" tf:"sample"`
	// +optional
	SchemaID *string `json:"schemaID,omitempty" tf:"schema_id"`
	// +optional
	TypeName *string `json:"typeName,omitempty" tf:"type_name"`
}

type ApiOperationSpecRequest struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Header []ApiOperationSpecRequestHeader `json:"header,omitempty" tf:"header"`
	// +optional
	QueryParameter []ApiOperationSpecRequestQueryParameter `json:"queryParameter,omitempty" tf:"query_parameter"`
	// +optional
	Representation []ApiOperationSpecRequestRepresentation `json:"representation,omitempty" tf:"representation"`
}

type ApiOperationSpecResponseHeader struct {
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	Required    *bool   `json:"required" tf:"required"`
	Type        *string `json:"type" tf:"type"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type ApiOperationSpecResponseRepresentationFormParameter struct {
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	Required    *bool   `json:"required" tf:"required"`
	Type        *string `json:"type" tf:"type"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type ApiOperationSpecResponseRepresentation struct {
	ContentType *string `json:"contentType" tf:"content_type"`
	// +optional
	FormParameter []ApiOperationSpecResponseRepresentationFormParameter `json:"formParameter,omitempty" tf:"form_parameter"`
	// +optional
	Sample *string `json:"sample,omitempty" tf:"sample"`
	// +optional
	SchemaID *string `json:"schemaID,omitempty" tf:"schema_id"`
	// +optional
	TypeName *string `json:"typeName,omitempty" tf:"type_name"`
}

type ApiOperationSpecResponse struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Header []ApiOperationSpecResponseHeader `json:"header,omitempty" tf:"header"`
	// +optional
	Representation []ApiOperationSpecResponseRepresentation `json:"representation,omitempty" tf:"representation"`
	StatusCode     *int64                                   `json:"statusCode" tf:"status_code"`
}

type ApiOperationSpecTemplateParameter struct {
	// +optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	Required    *bool   `json:"required" tf:"required"`
	Type        *string `json:"type" tf:"type"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type ApiOperationSpec struct {
	State *ApiOperationSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApiOperationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApiOperationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiManagementName *string `json:"apiManagementName" tf:"api_management_name"`
	ApiName           *string `json:"apiName" tf:"api_name"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	Method      *string `json:"method" tf:"method"`
	OperationID *string `json:"operationID" tf:"operation_id"`
	// +optional
	Request           *ApiOperationSpecRequest `json:"request,omitempty" tf:"request"`
	ResourceGroupName *string                  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Response []ApiOperationSpecResponse `json:"response,omitempty" tf:"response"`
	// +optional
	TemplateParameter []ApiOperationSpecTemplateParameter `json:"templateParameter,omitempty" tf:"template_parameter"`
	UrlTemplate       *string                             `json:"urlTemplate" tf:"url_template"`
}

type ApiOperationStatus struct {
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

// ApiOperationList is a list of ApiOperations
type ApiOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiOperation CRD objects
	Items []ApiOperation `json:"items,omitempty"`
}
