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

type Service struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec,omitempty"`
	Status            ServiceStatus `json:"status,omitempty"`
}

type ServiceSpecAuthSettingsActiveDirectory struct {
	// +optional
	AllowedAudiences []string `json:"allowedAudiences,omitempty" tf:"allowed_audiences"`
	ClientID         *string  `json:"clientID" tf:"client_id"`
	// +optional
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
}

type ServiceSpecAuthSettingsFacebook struct {
	AppID     *string `json:"appID" tf:"app_id"`
	AppSecret *string `json:"-" sensitive:"true" tf:"app_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type ServiceSpecAuthSettingsGoogle struct {
	ClientID     *string `json:"clientID" tf:"client_id"`
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type ServiceSpecAuthSettingsMicrosoft struct {
	ClientID     *string `json:"clientID" tf:"client_id"`
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type ServiceSpecAuthSettingsTwitter struct {
	ConsumerKey    *string `json:"consumerKey" tf:"consumer_key"`
	ConsumerSecret *string `json:"-" sensitive:"true" tf:"consumer_secret"`
}

type ServiceSpecAuthSettings struct {
	// +optional
	ActiveDirectory *ServiceSpecAuthSettingsActiveDirectory `json:"activeDirectory,omitempty" tf:"active_directory"`
	// +optional
	AdditionalLoginParams *map[string]string `json:"additionalLoginParams,omitempty" tf:"additional_login_params"`
	// +optional
	AllowedExternalRedirectUrls []string `json:"allowedExternalRedirectUrls,omitempty" tf:"allowed_external_redirect_urls"`
	// +optional
	DefaultProvider *string `json:"defaultProvider,omitempty" tf:"default_provider"`
	Enabled         *bool   `json:"enabled" tf:"enabled"`
	// +optional
	Facebook *ServiceSpecAuthSettingsFacebook `json:"facebook,omitempty" tf:"facebook"`
	// +optional
	Google *ServiceSpecAuthSettingsGoogle `json:"google,omitempty" tf:"google"`
	// +optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer"`
	// +optional
	Microsoft *ServiceSpecAuthSettingsMicrosoft `json:"microsoft,omitempty" tf:"microsoft"`
	// +optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version"`
	// +optional
	TokenRefreshExtensionHours *float64 `json:"tokenRefreshExtensionHours,omitempty" tf:"token_refresh_extension_hours"`
	// +optional
	TokenStoreEnabled *bool `json:"tokenStoreEnabled,omitempty" tf:"token_store_enabled"`
	// +optional
	Twitter *ServiceSpecAuthSettingsTwitter `json:"twitter,omitempty" tf:"twitter"`
	// +optional
	UnauthenticatedClientAction *string `json:"unauthenticatedClientAction,omitempty" tf:"unauthenticated_client_action"`
}

type ServiceSpecBackupSchedule struct {
	FrequencyInterval *int64  `json:"frequencyInterval" tf:"frequency_interval"`
	FrequencyUnit     *string `json:"frequencyUnit" tf:"frequency_unit"`
	// +optional
	KeepAtLeastOneBackup *bool `json:"keepAtLeastOneBackup,omitempty" tf:"keep_at_least_one_backup"`
	// +optional
	RetentionPeriodInDays *int64 `json:"retentionPeriodInDays,omitempty" tf:"retention_period_in_days"`
	// +optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time"`
}

type ServiceSpecBackup struct {
	// +optional
	Enabled           *bool                      `json:"enabled,omitempty" tf:"enabled"`
	Name              *string                    `json:"name" tf:"name"`
	Schedule          *ServiceSpecBackupSchedule `json:"schedule" tf:"schedule"`
	StorageAccountURL *string                    `json:"-" sensitive:"true" tf:"storage_account_url"`
}

type ServiceSpecConnectionString struct {
	Name  *string `json:"name" tf:"name"`
	Type  *string `json:"type" tf:"type"`
	Value *string `json:"-" sensitive:"true" tf:"value"`
}

type ServiceSpecIdentity struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type ServiceSpecLogsApplicationLogsAzureBlobStorage struct {
	Level           *string `json:"level" tf:"level"`
	RetentionInDays *int64  `json:"retentionInDays" tf:"retention_in_days"`
	SasURL          *string `json:"-" sensitive:"true" tf:"sas_url"`
}

type ServiceSpecLogsApplicationLogs struct {
	// +optional
	AzureBlobStorage *ServiceSpecLogsApplicationLogsAzureBlobStorage `json:"azureBlobStorage,omitempty" tf:"azure_blob_storage"`
	// +optional
	FileSystemLevel *string `json:"fileSystemLevel,omitempty" tf:"file_system_level"`
}

type ServiceSpecLogsHttpLogsAzureBlobStorage struct {
	RetentionInDays *int64  `json:"retentionInDays" tf:"retention_in_days"`
	SasURL          *string `json:"-" sensitive:"true" tf:"sas_url"`
}

type ServiceSpecLogsHttpLogsFileSystem struct {
	RetentionInDays *int64 `json:"retentionInDays" tf:"retention_in_days"`
	RetentionInMb   *int64 `json:"retentionInMb" tf:"retention_in_mb"`
}

type ServiceSpecLogsHttpLogs struct {
	// +optional
	AzureBlobStorage *ServiceSpecLogsHttpLogsAzureBlobStorage `json:"azureBlobStorage,omitempty" tf:"azure_blob_storage"`
	// +optional
	FileSystem *ServiceSpecLogsHttpLogsFileSystem `json:"fileSystem,omitempty" tf:"file_system"`
}

type ServiceSpecLogs struct {
	// +optional
	ApplicationLogs *ServiceSpecLogsApplicationLogs `json:"applicationLogs,omitempty" tf:"application_logs"`
	// +optional
	DetailedErrorMessagesEnabled *bool `json:"detailedErrorMessagesEnabled,omitempty" tf:"detailed_error_messages_enabled"`
	// +optional
	FailedRequestTracingEnabled *bool `json:"failedRequestTracingEnabled,omitempty" tf:"failed_request_tracing_enabled"`
	// +optional
	HttpLogs *ServiceSpecLogsHttpLogs `json:"httpLogs,omitempty" tf:"http_logs"`
}

type ServiceSpecSiteConfigCors struct {
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	SupportCredentials *bool `json:"supportCredentials,omitempty" tf:"support_credentials"`
}

type ServiceSpecSiteConfigIpRestrictionHeaders struct {
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XAzureFdid []string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid"`
	// +optional
	XFdHealthProbe []string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe"`
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XForwardedFor []string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for"`
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XForwardedHost []string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host"`
}

type ServiceSpecSiteConfigIpRestriction struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	Headers *ServiceSpecSiteConfigIpRestrictionHeaders `json:"headers,omitempty" tf:"headers"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag"`
	// +optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetID,omitempty" tf:"virtual_network_subnet_id"`
}

type ServiceSpecSiteConfigScmIPRestrictionHeaders struct {
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XAzureFdid []string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid"`
	// +optional
	XFdHealthProbe []string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe"`
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XForwardedFor []string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for"`
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XForwardedHost []string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host"`
}

type ServiceSpecSiteConfigScmIPRestriction struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	Headers *ServiceSpecSiteConfigScmIPRestrictionHeaders `json:"headers,omitempty" tf:"headers"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag"`
	// +optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetID,omitempty" tf:"virtual_network_subnet_id"`
}

type ServiceSpecSiteConfig struct {
	// +optional
	AcrUseManagedIdentityCredentials *bool `json:"acrUseManagedIdentityCredentials,omitempty" tf:"acr_use_managed_identity_credentials"`
	// +optional
	AcrUserManagedIdentityClientID *string `json:"acrUserManagedIdentityClientID,omitempty" tf:"acr_user_managed_identity_client_id"`
	// +optional
	AlwaysOn *bool `json:"alwaysOn,omitempty" tf:"always_on"`
	// +optional
	AppCommandLine *string `json:"appCommandLine,omitempty" tf:"app_command_line"`
	// +optional
	AutoSwapSlotName *string `json:"autoSwapSlotName,omitempty" tf:"auto_swap_slot_name"`
	// +optional
	Cors *ServiceSpecSiteConfigCors `json:"cors,omitempty" tf:"cors"`
	// +optional
	DefaultDocuments []string `json:"defaultDocuments,omitempty" tf:"default_documents"`
	// +optional
	DotnetFrameworkVersion *string `json:"dotnetFrameworkVersion,omitempty" tf:"dotnet_framework_version"`
	// +optional
	FtpsState *string `json:"ftpsState,omitempty" tf:"ftps_state"`
	// +optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path"`
	// +optional
	Http2Enabled *bool `json:"http2Enabled,omitempty" tf:"http2_enabled"`
	// +optional
	IpRestriction []ServiceSpecSiteConfigIpRestriction `json:"ipRestriction,omitempty" tf:"ip_restriction"`
	// +optional
	JavaContainer *string `json:"javaContainer,omitempty" tf:"java_container"`
	// +optional
	JavaContainerVersion *string `json:"javaContainerVersion,omitempty" tf:"java_container_version"`
	// +optional
	JavaVersion *string `json:"javaVersion,omitempty" tf:"java_version"`
	// +optional
	LinuxFxVersion *string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version"`
	// +optional
	LocalMysqlEnabled *bool `json:"localMysqlEnabled,omitempty" tf:"local_mysql_enabled"`
	// +optional
	ManagedPipelineMode *string `json:"managedPipelineMode,omitempty" tf:"managed_pipeline_mode"`
	// +optional
	MinTlsVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version"`
	// +optional
	NumberOfWorkers *int64 `json:"numberOfWorkers,omitempty" tf:"number_of_workers"`
	// +optional
	PhpVersion *string `json:"phpVersion,omitempty" tf:"php_version"`
	// +optional
	PythonVersion *string `json:"pythonVersion,omitempty" tf:"python_version"`
	// +optional
	RemoteDebuggingEnabled *bool `json:"remoteDebuggingEnabled,omitempty" tf:"remote_debugging_enabled"`
	// +optional
	RemoteDebuggingVersion *string `json:"remoteDebuggingVersion,omitempty" tf:"remote_debugging_version"`
	// +optional
	ScmIPRestriction []ServiceSpecSiteConfigScmIPRestriction `json:"scmIPRestriction,omitempty" tf:"scm_ip_restriction"`
	// +optional
	ScmType *string `json:"scmType,omitempty" tf:"scm_type"`
	// +optional
	ScmUseMainIPRestriction *bool `json:"scmUseMainIPRestriction,omitempty" tf:"scm_use_main_ip_restriction"`
	// +optional
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process"`
	// +optional
	VnetRouteAllEnabled *bool `json:"vnetRouteAllEnabled,omitempty" tf:"vnet_route_all_enabled"`
	// +optional
	WebsocketsEnabled *bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled"`
	// +optional
	WindowsFxVersion *string `json:"windowsFxVersion,omitempty" tf:"windows_fx_version"`
}

type ServiceSpecSiteCredential struct {
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type ServiceSpecSourceControl struct {
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// +optional
	ManualIntegration *bool `json:"manualIntegration,omitempty" tf:"manual_integration"`
	// +optional
	RepoURL *string `json:"repoURL,omitempty" tf:"repo_url"`
	// +optional
	RollbackEnabled *bool `json:"rollbackEnabled,omitempty" tf:"rollback_enabled"`
	// +optional
	UseMercurial *bool `json:"useMercurial,omitempty" tf:"use_mercurial"`
}

type ServiceSpecStorageAccount struct {
	AccessKey   *string `json:"-" sensitive:"true" tf:"access_key"`
	AccountName *string `json:"accountName" tf:"account_name"`
	// +optional
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path"`
	Name      *string `json:"name" tf:"name"`
	ShareName *string `json:"shareName" tf:"share_name"`
	Type      *string `json:"type" tf:"type"`
}

type ServiceSpec struct {
	State *ServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AppServicePlanID *string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// +optional
	AppSettings *map[string]string `json:"appSettings,omitempty" tf:"app_settings"`
	// +optional
	AuthSettings *ServiceSpecAuthSettings `json:"authSettings,omitempty" tf:"auth_settings"`
	// +optional
	Backup *ServiceSpecBackup `json:"backup,omitempty" tf:"backup"`
	// +optional
	ClientAffinityEnabled *bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled"`
	// +optional
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty" tf:"client_cert_enabled"`
	// +optional
	ClientCertMode *string `json:"clientCertMode,omitempty" tf:"client_cert_mode"`
	// +optional
	ConnectionString []ServiceSpecConnectionString `json:"connectionString,omitempty" tf:"connection_string"`
	// +optional
	CustomDomainVerificationID *string `json:"customDomainVerificationID,omitempty" tf:"custom_domain_verification_id"`
	// +optional
	DefaultSiteHostname *string `json:"defaultSiteHostname,omitempty" tf:"default_site_hostname"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	HttpsOnly *bool `json:"httpsOnly,omitempty" tf:"https_only"`
	// +optional
	Identity *ServiceSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	KeyVaultReferenceIdentityID *string `json:"keyVaultReferenceIdentityID,omitempty" tf:"key_vault_reference_identity_id"`
	Location                    *string `json:"location" tf:"location"`
	// +optional
	Logs *ServiceSpecLogs `json:"logs,omitempty" tf:"logs"`
	Name *string          `json:"name" tf:"name"`
	// +optional
	OutboundIPAddressList []string `json:"outboundIPAddressList,omitempty" tf:"outbound_ip_address_list"`
	// +optional
	OutboundIPAddresses *string `json:"outboundIPAddresses,omitempty" tf:"outbound_ip_addresses"`
	// +optional
	PossibleOutboundIPAddressList []string `json:"possibleOutboundIPAddressList,omitempty" tf:"possible_outbound_ip_address_list"`
	// +optional
	PossibleOutboundIPAddresses *string `json:"possibleOutboundIPAddresses,omitempty" tf:"possible_outbound_ip_addresses"`
	ResourceGroupName           *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SiteConfig *ServiceSpecSiteConfig `json:"siteConfig,omitempty" tf:"site_config"`
	// +optional
	SiteCredential []ServiceSpecSiteCredential `json:"siteCredential,omitempty" tf:"site_credential"`
	// +optional
	SourceControl *ServiceSpecSourceControl `json:"sourceControl,omitempty" tf:"source_control"`
	// +optional
	StorageAccount []ServiceSpecStorageAccount `json:"storageAccount,omitempty" tf:"storage_account"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ServiceStatus struct {
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

// ServiceList is a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Service CRD objects
	Items []Service `json:"items,omitempty"`
}
