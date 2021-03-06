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

type ApiManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSpec   `json:"spec,omitempty"`
	Status            ApiManagementStatus `json:"status,omitempty"`
}

type ApiManagementSpecAdditionalLocationVirtualNetworkConfiguration struct {
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
}

type ApiManagementSpecAdditionalLocation struct {
	// +optional
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`
	// +optional
	GatewayRegionalURL *string `json:"gatewayRegionalURL,omitempty" tf:"gateway_regional_url"`
	Location           *string `json:"location" tf:"location"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses"`
	// +optional
	PublicIPAddressID *string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses"`
	// +optional
	VirtualNetworkConfiguration *ApiManagementSpecAdditionalLocationVirtualNetworkConfiguration `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type ApiManagementSpecCertificate struct {
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	EncodedCertificate  *string `json:"-" sensitive:"true" tf:"encoded_certificate"`
	// +optional
	Expiry    *string `json:"expiry,omitempty" tf:"expiry"`
	StoreName *string `json:"storeName" tf:"store_name"`
	// +optional
	Subject *string `json:"subject,omitempty" tf:"subject"`
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint"`
}

type ApiManagementSpecHostnameConfigurationDeveloperPortal struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	// +optional
	Expiry   *string `json:"expiry,omitempty" tf:"expiry"`
	HostName *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
	// +optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientID,omitempty" tf:"ssl_keyvault_identity_client_id"`
	// +optional
	Subject *string `json:"subject,omitempty" tf:"subject"`
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint"`
}

type ApiManagementSpecHostnameConfigurationManagement struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	// +optional
	Expiry   *string `json:"expiry,omitempty" tf:"expiry"`
	HostName *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
	// +optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientID,omitempty" tf:"ssl_keyvault_identity_client_id"`
	// +optional
	Subject *string `json:"subject,omitempty" tf:"subject"`
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint"`
}

type ApiManagementSpecHostnameConfigurationPortal struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	// +optional
	Expiry   *string `json:"expiry,omitempty" tf:"expiry"`
	HostName *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
	// +optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientID,omitempty" tf:"ssl_keyvault_identity_client_id"`
	// +optional
	Subject *string `json:"subject,omitempty" tf:"subject"`
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint"`
}

type ApiManagementSpecHostnameConfigurationProxy struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	// +optional
	DefaultSslBinding *bool `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding"`
	// +optional
	Expiry   *string `json:"expiry,omitempty" tf:"expiry"`
	HostName *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
	// +optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientID,omitempty" tf:"ssl_keyvault_identity_client_id"`
	// +optional
	Subject *string `json:"subject,omitempty" tf:"subject"`
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint"`
}

type ApiManagementSpecHostnameConfigurationScm struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	// +optional
	Expiry   *string `json:"expiry,omitempty" tf:"expiry"`
	HostName *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
	// +optional
	SslKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientID,omitempty" tf:"ssl_keyvault_identity_client_id"`
	// +optional
	Subject *string `json:"subject,omitempty" tf:"subject"`
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint"`
}

type ApiManagementSpecHostnameConfiguration struct {
	// +optional
	DeveloperPortal []ApiManagementSpecHostnameConfigurationDeveloperPortal `json:"developerPortal,omitempty" tf:"developer_portal"`
	// +optional
	Management []ApiManagementSpecHostnameConfigurationManagement `json:"management,omitempty" tf:"management"`
	// +optional
	Portal []ApiManagementSpecHostnameConfigurationPortal `json:"portal,omitempty" tf:"portal"`
	// +optional
	Proxy []ApiManagementSpecHostnameConfigurationProxy `json:"proxy,omitempty" tf:"proxy"`
	// +optional
	Scm []ApiManagementSpecHostnameConfigurationScm `json:"scm,omitempty" tf:"scm"`
}

type ApiManagementSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type ApiManagementSpecPolicy struct {
	// +optional
	XmlContent *string `json:"xmlContent,omitempty" tf:"xml_content"`
	// +optional
	XmlLink *string `json:"xmlLink,omitempty" tf:"xml_link"`
}

type ApiManagementSpecProtocols struct {
	// +optional
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2"`
}

type ApiManagementSpecSecurity struct {
	// +optional
	EnableBackendSsl30 *bool `json:"enableBackendSsl30,omitempty" tf:"enable_backend_ssl30"`
	// +optional
	EnableBackendTls10 *bool `json:"enableBackendTls10,omitempty" tf:"enable_backend_tls10"`
	// +optional
	EnableBackendTls11 *bool `json:"enableBackendTls11,omitempty" tf:"enable_backend_tls11"`
	// +optional
	EnableFrontendSsl30 *bool `json:"enableFrontendSsl30,omitempty" tf:"enable_frontend_ssl30"`
	// +optional
	EnableFrontendTls10 *bool `json:"enableFrontendTls10,omitempty" tf:"enable_frontend_tls10"`
	// +optional
	EnableFrontendTls11 *bool `json:"enableFrontendTls11,omitempty" tf:"enable_frontend_tls11"`
	// +optional
	// Deprecated
	EnableTripleDESCiphers *bool `json:"enableTripleDESCiphers,omitempty" tf:"enable_triple_des_ciphers"`
	// +optional
	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled"`
	// +optional
	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled"`
	// +optional
	TlsEcdheRsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled"`
	// +optional
	TlsEcdheRsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled"`
	// +optional
	TlsRsaWithAes128CbcSha256CiphersEnabled *bool `json:"tlsRsaWithAes128CbcSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_cbc_sha256_ciphers_enabled"`
	// +optional
	TlsRsaWithAes128CbcShaCiphersEnabled *bool `json:"tlsRsaWithAes128CbcShaCiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_cbc_sha_ciphers_enabled"`
	// +optional
	TlsRsaWithAes128GcmSha256CiphersEnabled *bool `json:"tlsRsaWithAes128GcmSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes128_gcm_sha256_ciphers_enabled"`
	// +optional
	TlsRsaWithAes256CbcSha256CiphersEnabled *bool `json:"tlsRsaWithAes256CbcSha256CiphersEnabled,omitempty" tf:"tls_rsa_with_aes256_cbc_sha256_ciphers_enabled"`
	// +optional
	TlsRsaWithAes256CbcShaCiphersEnabled *bool `json:"tlsRsaWithAes256CbcShaCiphersEnabled,omitempty" tf:"tls_rsa_with_aes256_cbc_sha_ciphers_enabled"`
	// +optional
	TripleDESCiphersEnabled *bool `json:"tripleDESCiphersEnabled,omitempty" tf:"triple_des_ciphers_enabled"`
}

type ApiManagementSpecSignIn struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
}

type ApiManagementSpecSignUpTermsOfService struct {
	ConsentRequired *bool `json:"consentRequired" tf:"consent_required"`
	Enabled         *bool `json:"enabled" tf:"enabled"`
	// +optional
	Text *string `json:"text,omitempty" tf:"text"`
}

type ApiManagementSpecSignUp struct {
	Enabled        *bool                                  `json:"enabled" tf:"enabled"`
	TermsOfService *ApiManagementSpecSignUpTermsOfService `json:"termsOfService" tf:"terms_of_service"`
}

type ApiManagementSpecTenantAccess struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	PrimaryKey *string `json:"-" sensitive:"true" tf:"primary_key"`
	// +optional
	SecondaryKey *string `json:"-" sensitive:"true" tf:"secondary_key"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
}

type ApiManagementSpecVirtualNetworkConfiguration struct {
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
}

type ApiManagementSpec struct {
	State *ApiManagementSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApiManagementSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ApiManagementSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalLocation []ApiManagementSpecAdditionalLocation `json:"additionalLocation,omitempty" tf:"additional_location"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Certificate []ApiManagementSpecCertificate `json:"certificate,omitempty" tf:"certificate"`
	// +optional
	ClientCertificateEnabled *bool `json:"clientCertificateEnabled,omitempty" tf:"client_certificate_enabled"`
	// +optional
	DeveloperPortalURL *string `json:"developerPortalURL,omitempty" tf:"developer_portal_url"`
	// +optional
	GatewayDisabled *bool `json:"gatewayDisabled,omitempty" tf:"gateway_disabled"`
	// +optional
	GatewayRegionalURL *string `json:"gatewayRegionalURL,omitempty" tf:"gateway_regional_url"`
	// +optional
	GatewayURL *string `json:"gatewayURL,omitempty" tf:"gateway_url"`
	// +optional
	HostnameConfiguration *ApiManagementSpecHostnameConfiguration `json:"hostnameConfiguration,omitempty" tf:"hostname_configuration"`
	// +optional
	Identity *ApiManagementSpecIdentity `json:"identity,omitempty" tf:"identity"`
	Location *string                    `json:"location" tf:"location"`
	// +optional
	ManagementAPIURL *string `json:"managementAPIURL,omitempty" tf:"management_api_url"`
	// +optional
	MinAPIVersion *string `json:"minAPIVersion,omitempty" tf:"min_api_version"`
	Name          *string `json:"name" tf:"name"`
	// +optional
	NotificationSenderEmail *string `json:"notificationSenderEmail,omitempty" tf:"notification_sender_email"`
	// +optional
	Policy *ApiManagementSpecPolicy `json:"policy,omitempty" tf:"policy"`
	// +optional
	PortalURL *string `json:"portalURL,omitempty" tf:"portal_url"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses"`
	// +optional
	Protocols *ApiManagementSpecProtocols `json:"protocols,omitempty" tf:"protocols"`
	// +optional
	PublicIPAddressID *string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses"`
	// +optional
	PublicNetworkAccessEnabled *bool   `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	PublisherEmail             *string `json:"publisherEmail" tf:"publisher_email"`
	PublisherName              *string `json:"publisherName" tf:"publisher_name"`
	ResourceGroupName          *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ScmURL *string `json:"scmURL,omitempty" tf:"scm_url"`
	// +optional
	Security *ApiManagementSpecSecurity `json:"security,omitempty" tf:"security"`
	// +optional
	SignIn *ApiManagementSpecSignIn `json:"signIn,omitempty" tf:"sign_in"`
	// +optional
	SignUp  *ApiManagementSpecSignUp `json:"signUp,omitempty" tf:"sign_up"`
	SkuName *string                  `json:"skuName" tf:"sku_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TenantAccess *ApiManagementSpecTenantAccess `json:"tenantAccess,omitempty" tf:"tenant_access"`
	// +optional
	VirtualNetworkConfiguration *ApiManagementSpecVirtualNetworkConfiguration `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration"`
	// +optional
	VirtualNetworkType *string `json:"virtualNetworkType,omitempty" tf:"virtual_network_type"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type ApiManagementStatus struct {
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

// ApiManagementList is a list of ApiManagements
type ApiManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagement CRD objects
	Items []ApiManagement `json:"items,omitempty"`
}
