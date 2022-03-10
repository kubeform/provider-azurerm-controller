package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CatalogCollationType enumerates the values for catalog collation type.
type CatalogCollationType string

const (
	// DATABASEDEFAULT ...
	DATABASEDEFAULT CatalogCollationType = "DATABASE_DEFAULT"
	// SQLLatin1GeneralCP1CIAS ...
	SQLLatin1GeneralCP1CIAS CatalogCollationType = "SQL_Latin1_General_CP1_CI_AS"
)

// PossibleCatalogCollationTypeValues returns an array of possible values for the CatalogCollationType const type.
func PossibleCatalogCollationTypeValues() []CatalogCollationType {
	return []CatalogCollationType{DATABASEDEFAULT, SQLLatin1GeneralCP1CIAS}
}

// DatabaseState enumerates the values for database state.
type DatabaseState string

const (
	// All ...
	All DatabaseState = "All"
	// Deleted ...
	Deleted DatabaseState = "Deleted"
	// Live ...
	Live DatabaseState = "Live"
)

// PossibleDatabaseStateValues returns an array of possible values for the DatabaseState const type.
func PossibleDatabaseStateValues() []DatabaseState {
	return []DatabaseState{All, Deleted, Live}
}

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// SystemAssigned ...
	SystemAssigned IdentityType = "SystemAssigned"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{SystemAssigned}
}

// InstancePoolLicenseType enumerates the values for instance pool license type.
type InstancePoolLicenseType string

const (
	// BasePrice ...
	BasePrice InstancePoolLicenseType = "BasePrice"
	// LicenseIncluded ...
	LicenseIncluded InstancePoolLicenseType = "LicenseIncluded"
)

// PossibleInstancePoolLicenseTypeValues returns an array of possible values for the InstancePoolLicenseType const type.
func PossibleInstancePoolLicenseTypeValues() []InstancePoolLicenseType {
	return []InstancePoolLicenseType{BasePrice, LicenseIncluded}
}

// ManagedDatabaseCreateMode enumerates the values for managed database create mode.
type ManagedDatabaseCreateMode string

const (
	// Default ...
	Default ManagedDatabaseCreateMode = "Default"
	// PointInTimeRestore ...
	PointInTimeRestore ManagedDatabaseCreateMode = "PointInTimeRestore"
	// Recovery ...
	Recovery ManagedDatabaseCreateMode = "Recovery"
	// RestoreExternalBackup ...
	RestoreExternalBackup ManagedDatabaseCreateMode = "RestoreExternalBackup"
	// RestoreLongTermRetentionBackup ...
	RestoreLongTermRetentionBackup ManagedDatabaseCreateMode = "RestoreLongTermRetentionBackup"
)

// PossibleManagedDatabaseCreateModeValues returns an array of possible values for the ManagedDatabaseCreateMode const type.
func PossibleManagedDatabaseCreateModeValues() []ManagedDatabaseCreateMode {
	return []ManagedDatabaseCreateMode{Default, PointInTimeRestore, Recovery, RestoreExternalBackup, RestoreLongTermRetentionBackup}
}

// ManagedDatabaseStatus enumerates the values for managed database status.
type ManagedDatabaseStatus string

const (
	// Creating ...
	Creating ManagedDatabaseStatus = "Creating"
	// Inaccessible ...
	Inaccessible ManagedDatabaseStatus = "Inaccessible"
	// Offline ...
	Offline ManagedDatabaseStatus = "Offline"
	// Online ...
	Online ManagedDatabaseStatus = "Online"
	// Restoring ...
	Restoring ManagedDatabaseStatus = "Restoring"
	// Shutdown ...
	Shutdown ManagedDatabaseStatus = "Shutdown"
	// Updating ...
	Updating ManagedDatabaseStatus = "Updating"
)

// PossibleManagedDatabaseStatusValues returns an array of possible values for the ManagedDatabaseStatus const type.
func PossibleManagedDatabaseStatusValues() []ManagedDatabaseStatus {
	return []ManagedDatabaseStatus{Creating, Inaccessible, Offline, Online, Restoring, Shutdown, Updating}
}

// ManagedInstanceLicenseType enumerates the values for managed instance license type.
type ManagedInstanceLicenseType string

const (
	// ManagedInstanceLicenseTypeBasePrice ...
	ManagedInstanceLicenseTypeBasePrice ManagedInstanceLicenseType = "BasePrice"
	// ManagedInstanceLicenseTypeLicenseIncluded ...
	ManagedInstanceLicenseTypeLicenseIncluded ManagedInstanceLicenseType = "LicenseIncluded"
)

// PossibleManagedInstanceLicenseTypeValues returns an array of possible values for the ManagedInstanceLicenseType const type.
func PossibleManagedInstanceLicenseTypeValues() []ManagedInstanceLicenseType {
	return []ManagedInstanceLicenseType{ManagedInstanceLicenseTypeBasePrice, ManagedInstanceLicenseTypeLicenseIncluded}
}

// ManagedInstanceProxyOverride enumerates the values for managed instance proxy override.
type ManagedInstanceProxyOverride string

const (
	// ManagedInstanceProxyOverrideDefault ...
	ManagedInstanceProxyOverrideDefault ManagedInstanceProxyOverride = "Default"
	// ManagedInstanceProxyOverrideProxy ...
	ManagedInstanceProxyOverrideProxy ManagedInstanceProxyOverride = "Proxy"
	// ManagedInstanceProxyOverrideRedirect ...
	ManagedInstanceProxyOverrideRedirect ManagedInstanceProxyOverride = "Redirect"
)

// PossibleManagedInstanceProxyOverrideValues returns an array of possible values for the ManagedInstanceProxyOverride const type.
func PossibleManagedInstanceProxyOverrideValues() []ManagedInstanceProxyOverride {
	return []ManagedInstanceProxyOverride{ManagedInstanceProxyOverrideDefault, ManagedInstanceProxyOverrideProxy, ManagedInstanceProxyOverrideRedirect}
}

// ManagedServerCreateMode enumerates the values for managed server create mode.
type ManagedServerCreateMode string

const (
	// ManagedServerCreateModeDefault ...
	ManagedServerCreateModeDefault ManagedServerCreateMode = "Default"
	// ManagedServerCreateModePointInTimeRestore ...
	ManagedServerCreateModePointInTimeRestore ManagedServerCreateMode = "PointInTimeRestore"
)

// PossibleManagedServerCreateModeValues returns an array of possible values for the ManagedServerCreateMode const type.
func PossibleManagedServerCreateModeValues() []ManagedServerCreateMode {
	return []ManagedServerCreateMode{ManagedServerCreateModeDefault, ManagedServerCreateModePointInTimeRestore}
}

// ManagementOperationState enumerates the values for management operation state.
type ManagementOperationState string

const (
	// CancelInProgress ...
	CancelInProgress ManagementOperationState = "CancelInProgress"
	// Cancelled ...
	Cancelled ManagementOperationState = "Cancelled"
	// Failed ...
	Failed ManagementOperationState = "Failed"
	// InProgress ...
	InProgress ManagementOperationState = "InProgress"
	// Pending ...
	Pending ManagementOperationState = "Pending"
	// Succeeded ...
	Succeeded ManagementOperationState = "Succeeded"
)

// PossibleManagementOperationStateValues returns an array of possible values for the ManagementOperationState const type.
func PossibleManagementOperationStateValues() []ManagementOperationState {
	return []ManagementOperationState{CancelInProgress, Cancelled, Failed, InProgress, Pending, Succeeded}
}

// ReplicaType enumerates the values for replica type.
type ReplicaType string

const (
	// Primary ...
	Primary ReplicaType = "Primary"
	// ReadableSecondary ...
	ReadableSecondary ReplicaType = "ReadableSecondary"
)

// PossibleReplicaTypeValues returns an array of possible values for the ReplicaType const type.
func PossibleReplicaTypeValues() []ReplicaType {
	return []ReplicaType{Primary, ReadableSecondary}
}

// SecurityAlertPolicyState enumerates the values for security alert policy state.
type SecurityAlertPolicyState string

const (
	// SecurityAlertPolicyStateDisabled ...
	SecurityAlertPolicyStateDisabled SecurityAlertPolicyState = "Disabled"
	// SecurityAlertPolicyStateEnabled ...
	SecurityAlertPolicyStateEnabled SecurityAlertPolicyState = "Enabled"
	// SecurityAlertPolicyStateNew ...
	SecurityAlertPolicyStateNew SecurityAlertPolicyState = "New"
)

// PossibleSecurityAlertPolicyStateValues returns an array of possible values for the SecurityAlertPolicyState const type.
func PossibleSecurityAlertPolicyStateValues() []SecurityAlertPolicyState {
	return []SecurityAlertPolicyState{SecurityAlertPolicyStateDisabled, SecurityAlertPolicyStateEnabled, SecurityAlertPolicyStateNew}
}

// SensitivityLabelRank enumerates the values for sensitivity label rank.
type SensitivityLabelRank string

const (
	// Critical ...
	Critical SensitivityLabelRank = "Critical"
	// High ...
	High SensitivityLabelRank = "High"
	// Low ...
	Low SensitivityLabelRank = "Low"
	// Medium ...
	Medium SensitivityLabelRank = "Medium"
	// None ...
	None SensitivityLabelRank = "None"
)

// PossibleSensitivityLabelRankValues returns an array of possible values for the SensitivityLabelRank const type.
func PossibleSensitivityLabelRankValues() []SensitivityLabelRank {
	return []SensitivityLabelRank{Critical, High, Low, Medium, None}
}

// SensitivityLabelSource enumerates the values for sensitivity label source.
type SensitivityLabelSource string

const (
	// Current ...
	Current SensitivityLabelSource = "current"
	// Recommended ...
	Recommended SensitivityLabelSource = "recommended"
)

// PossibleSensitivityLabelSourceValues returns an array of possible values for the SensitivityLabelSource const type.
func PossibleSensitivityLabelSourceValues() []SensitivityLabelSource {
	return []SensitivityLabelSource{Current, Recommended}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusCanceled ...
	StatusCanceled Status = "Canceled"
	// StatusCompleted ...
	StatusCompleted Status = "Completed"
	// StatusFailed ...
	StatusFailed Status = "Failed"
	// StatusInProgress ...
	StatusInProgress Status = "InProgress"
	// StatusNotStarted ...
	StatusNotStarted Status = "NotStarted"
	// StatusSlowedDown ...
	StatusSlowedDown Status = "SlowedDown"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusCanceled, StatusCompleted, StatusFailed, StatusInProgress, StatusNotStarted, StatusSlowedDown}
}
