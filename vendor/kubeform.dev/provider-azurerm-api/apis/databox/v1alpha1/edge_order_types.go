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

type EdgeOrder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeOrderSpec   `json:"spec,omitempty"`
	Status            EdgeOrderStatus `json:"status,omitempty"`
}

type EdgeOrderSpecContact struct {
	CompanyName *string  `json:"companyName" tf:"company_name"`
	Emails      []string `json:"emails" tf:"emails"`
	Name        *string  `json:"name" tf:"name"`
	PhoneNumber *string  `json:"phoneNumber" tf:"phone_number"`
}

type EdgeOrderSpecReturnTracking struct {
	// +optional
	CarrierName *string `json:"carrierName,omitempty" tf:"carrier_name"`
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number"`
	// +optional
	TrackingID *string `json:"trackingID,omitempty" tf:"tracking_id"`
	// +optional
	TrackingURL *string `json:"trackingURL,omitempty" tf:"tracking_url"`
}

type EdgeOrderSpecShipmentAddress struct {
	// +kubebuilder:validation:MaxItems=3
	Address    []string `json:"address" tf:"address"`
	City       *string  `json:"city" tf:"city"`
	Country    *string  `json:"country" tf:"country"`
	PostalCode *string  `json:"postalCode" tf:"postal_code"`
	State      *string  `json:"state" tf:"state"`
}

type EdgeOrderSpecShipmentHistory struct {
	// +optional
	AdditionalDetails *map[string]string `json:"additionalDetails,omitempty" tf:"additional_details"`
	// +optional
	Comments *string `json:"comments,omitempty" tf:"comments"`
	// +optional
	LastUpdate *string `json:"lastUpdate,omitempty" tf:"last_update"`
}

type EdgeOrderSpecShipmentTracking struct {
	// +optional
	CarrierName *string `json:"carrierName,omitempty" tf:"carrier_name"`
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number"`
	// +optional
	TrackingID *string `json:"trackingID,omitempty" tf:"tracking_id"`
	// +optional
	TrackingURL *string `json:"trackingURL,omitempty" tf:"tracking_url"`
}

type EdgeOrderSpecStatus struct {
	// +optional
	AdditionalDetails *map[string]string `json:"additionalDetails,omitempty" tf:"additional_details"`
	// +optional
	Comments *string `json:"comments,omitempty" tf:"comments"`
	// +optional
	Info *string `json:"info,omitempty" tf:"info"`
	// +optional
	LastUpdate *string `json:"lastUpdate,omitempty" tf:"last_update"`
}

type EdgeOrderSpec struct {
	State *EdgeOrderSpecResource `json:"state,omitempty" tf:"-"`

	Resource EdgeOrderSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EdgeOrderSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Contact    *EdgeOrderSpecContact `json:"contact" tf:"contact"`
	DeviceName *string               `json:"deviceName" tf:"device_name"`
	// +optional
	Name              *string `json:"name,omitempty" tf:"name"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ReturnTracking []EdgeOrderSpecReturnTracking `json:"returnTracking,omitempty" tf:"return_tracking"`
	// +optional
	SerialNumber    *string                       `json:"serialNumber,omitempty" tf:"serial_number"`
	ShipmentAddress *EdgeOrderSpecShipmentAddress `json:"shipmentAddress" tf:"shipment_address"`
	// +optional
	ShipmentHistory []EdgeOrderSpecShipmentHistory `json:"shipmentHistory,omitempty" tf:"shipment_history"`
	// +optional
	ShipmentTracking []EdgeOrderSpecShipmentTracking `json:"shipmentTracking,omitempty" tf:"shipment_tracking"`
	// +optional
	Status []EdgeOrderSpecStatus `json:"status,omitempty" tf:"status"`
}

type EdgeOrderStatus struct {
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

// EdgeOrderList is a list of EdgeOrders
type EdgeOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EdgeOrder CRD objects
	Items []EdgeOrder `json:"items,omitempty"`
}
