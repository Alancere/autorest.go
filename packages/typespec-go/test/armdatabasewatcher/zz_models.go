// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasewatcher

import "time"

// Datastore - The properties of a data store.
type Datastore struct {
	// REQUIRED; The Kusto cluster URI.
	KustoClusterURI *string

	// REQUIRED; The Kusto data ingestion URI.
	KustoDataIngestionURI *string

	// REQUIRED; The name of a Kusto database.
	KustoDatabaseName *string

	// REQUIRED; The Kusto management URL.
	KustoManagementURL *string

	// REQUIRED; The type of a Kusto offering.
	KustoOfferingType *KustoOfferingType

	// The Azure ResourceId of an Azure Data Explorer cluster.
	AdxClusterResourceID *string

	// The Kusto cluster display name.
	KustoClusterDisplayName *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info *ErrorAdditionalInfoInfo

	// READ-ONLY; The additional info type.
	Type *string
}

type ErrorAdditionalInfoInfo struct {
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; The type of managed identity assigned to this resource.
	Type *ManagedServiceIdentityType

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure
	// Resource Manager/control-plane operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SQLDbElasticPoolTargetProperties - The properties specific to elastic pool in Azure SQL Database.
type SQLDbElasticPoolTargetProperties struct {
	// REQUIRED; The Azure ResourceId of the anchor database used to connect to an elastic pool.
	AnchorDatabaseResourceID *string

	// REQUIRED; The server name to use in the connection string when connecting to a target. Port number and instance name must
	// be specified separately.
	ConnectionServerName *string

	// REQUIRED; The Azure ResourceId of an Azure SQL DB elastic pool target.
	SQLEpResourceID *string

	// REQUIRED; The type of authentication to use when connecting to a target.
	TargetAuthenticationType *TargetAuthenticationType

	// CONSTANT; The Azure SQL DB elastic pool target.
	// Field has constant value "SqlEp", any specified value is ignored.
	TargetType *string

	// Set to true to monitor a high availability replica of specified target, if any.
	ReadIntent *bool

	// To use SQL authentication when connecting to targets, specify the vault where the login name and password secrets are stored.
	TargetVault *VaultSecret

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ResourceProvisioningState
}

// GetTargetProperties implements the TargetPropertiesClassification interface for type SQLDbElasticPoolTargetProperties.
func (s *SQLDbElasticPoolTargetProperties) GetTargetProperties() *TargetProperties {
	return &TargetProperties{
		ConnectionServerName:     s.ConnectionServerName,
		ProvisioningState:        s.ProvisioningState,
		TargetAuthenticationType: s.TargetAuthenticationType,
		TargetType:               s.TargetType,
		TargetVault:              s.TargetVault,
	}
}

// SQLDbSingleDatabaseTargetProperties - The properties specific to single database in Azure SQL Database.
type SQLDbSingleDatabaseTargetProperties struct {
	// REQUIRED; The server name to use in the connection string when connecting to a target. Port number and instance name must
	// be specified separately.
	ConnectionServerName *string

	// REQUIRED; The Azure ResourceId of an Azure SQL DB single database target.
	SQLDbResourceID *string

	// REQUIRED; The type of authentication to use when connecting to a target.
	TargetAuthenticationType *TargetAuthenticationType

	// CONSTANT; The Azure SQL DB single database target.
	// Field has constant value "SqlDb", any specified value is ignored.
	TargetType *string

	// Set to true to monitor a high availability replica of specified target, if any.
	ReadIntent *bool

	// To use SQL authentication when connecting to targets, specify the vault where the login name and password secrets are stored.
	TargetVault *VaultSecret

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ResourceProvisioningState
}

// GetTargetProperties implements the TargetPropertiesClassification interface for type SQLDbSingleDatabaseTargetProperties.
func (s *SQLDbSingleDatabaseTargetProperties) GetTargetProperties() *TargetProperties {
	return &TargetProperties{
		ConnectionServerName:     s.ConnectionServerName,
		ProvisioningState:        s.ProvisioningState,
		TargetAuthenticationType: s.TargetAuthenticationType,
		TargetType:               s.TargetType,
		TargetVault:              s.TargetVault,
	}
}

// SQLMiTargetProperties - The properties specific to Azure SQL Managed Instance targets.
type SQLMiTargetProperties struct {
	// REQUIRED; The server name to use in the connection string when connecting to a target. Port number and instance name must
	// be specified separately.
	ConnectionServerName *string

	// REQUIRED; The Azure ResourceId of an Azure SQL Managed Instance target.
	SQLMiResourceID *string

	// REQUIRED; The type of authentication to use when connecting to a target.
	TargetAuthenticationType *TargetAuthenticationType

	// CONSTANT; The Azure SQL Managed Instance target.
	// Field has constant value "SqlMi", any specified value is ignored.
	TargetType *string

	// The TCP port number to optionally use in the connection string when connecting to an Azure SQL Managed Instance target.
	ConnectionTCPPort *int32

	// Set to true to monitor a high availability replica of specified target, if any.
	ReadIntent *bool

	// To use SQL authentication when connecting to targets, specify the vault where the login name and password secrets are stored.
	TargetVault *VaultSecret

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ResourceProvisioningState
}

// GetTargetProperties implements the TargetPropertiesClassification interface for type SQLMiTargetProperties.
func (s *SQLMiTargetProperties) GetTargetProperties() *TargetProperties {
	return &TargetProperties{
		ConnectionServerName:     s.ConnectionServerName,
		ProvisioningState:        s.ProvisioningState,
		TargetAuthenticationType: s.TargetAuthenticationType,
		TargetType:               s.TargetType,
		TargetVault:              s.TargetVault,
	}
}

// SQLVMTargetProperties - The properties specific to Azure SQL VM targets.
type SQLVMTargetProperties struct {
	// REQUIRED; The server name to use in the connection string when connecting to a target. Port number and instance name must
	// be specified separately.
	ConnectionServerName *string

	// REQUIRED; The Azure ResourceId of an Azure SQL VM target.
	SQLVMResourceID *string

	// REQUIRED; The type of authentication to use when connecting to a target.
	TargetAuthenticationType *TargetAuthenticationType

	// CONSTANT; The Azure SQL VM target.
	// Field has constant value "SqlVm", any specified value is ignored.
	TargetType *string

	// The TCP port number to optionally use in the connection string when connecting to an Azure SQL VM target.
	ConnectionTCPPort *int32

	// The SQL instance name to optionally use in the connection string when connecting to an Azure SQL VM target.
	SQLNamedInstanceName *string

	// To use SQL authentication when connecting to targets, specify the vault where the login name and password secrets are stored.
	TargetVault *VaultSecret

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ResourceProvisioningState
}

// GetTargetProperties implements the TargetPropertiesClassification interface for type SQLVMTargetProperties.
func (s *SQLVMTargetProperties) GetTargetProperties() *TargetProperties {
	return &TargetProperties{
		ConnectionServerName:     s.ConnectionServerName,
		ProvisioningState:        s.ProvisioningState,
		TargetAuthenticationType: s.TargetAuthenticationType,
		TargetType:               s.TargetType,
		TargetVault:              s.TargetVault,
	}
}

// SharedPrivateLinkResource - Concrete proxy resource types can be created by aliasing this type using a specific property
// type.
type SharedPrivateLinkResource struct {
	// The resource-specific properties for this resource.
	Properties *SharedPrivateLinkResourceProperties

	// READ-ONLY; The Shared Private Link resource name.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SharedPrivateLinkResourceListResult - The response of a SharedPrivateLinkResource list operation.
type SharedPrivateLinkResourceListResult struct {
	// REQUIRED; The SharedPrivateLinkResource items on this page
	Value []*SharedPrivateLinkResource

	// The link to the next page of items
	NextLink *string
}

// SharedPrivateLinkResourceProperties - The generic properties of a Shared Private Link resource.
type SharedPrivateLinkResourceProperties struct {
	// REQUIRED; The group id from the provider of resource the shared private link resource is for.
	GroupID *string

	// REQUIRED; The resource id of the resource the shared private link resource is for.
	PrivateLinkResourceID *string

	// REQUIRED; The request message for requesting approval of the shared private link resource.
	RequestMessage *string

	// The DNS zone to be included in the DNS name of the shared private link. Value is service-specific.
	DNSZone *string

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ResourceProvisioningState

	// READ-ONLY; Status of the shared private link resource. Can be Pending, Approved, Rejected or Disconnected.
	Status *SharedPrivateLinkResourceStatus
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// Target - Concrete proxy resource types can be created by aliasing this type using a specific property type.
type Target struct {
	// The resource-specific properties for this resource.
	Properties TargetPropertiesClassification

	// READ-ONLY; The target resource name.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TargetListResult - The response of a Target list operation.
type TargetListResult struct {
	// REQUIRED; The Target items on this page
	Value []*Target

	// The link to the next page of items
	NextLink *string
}

// TargetProperties - The generic properties of a target.
type TargetProperties struct {
	// REQUIRED; The server name to use in the connection string when connecting to a target. Port number and instance name must
	// be specified separately.
	ConnectionServerName *string

	// REQUIRED; The type of authentication to use when connecting to a target.
	TargetAuthenticationType *TargetAuthenticationType

	// REQUIRED; Discriminator property for TargetProperties.
	TargetType *string

	// To use SQL authentication when connecting to targets, specify the vault where the login name and password secrets are stored.
	TargetVault *VaultSecret

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ResourceProvisioningState
}

// GetTargetProperties implements the TargetPropertiesClassification interface for type TargetProperties.
func (t *TargetProperties) GetTargetProperties() *TargetProperties { return t }

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// VaultSecret - The vault specific details required if using SQL authentication to connect to a target.
type VaultSecret struct {
	// The Azure ResourceId of the Key Vault instance storing database authentication secrets.
	AkvResourceID *string

	// The path to the Key Vault secret storing the password for authentication to a target.
	AkvTargetPassword *string

	// The path to the Key Vault secret storing the login name (aka user name, aka account name) for authentication to a target.
	AkvTargetUser *string
}

// Watcher - The DatabaseWatcherProviderHub resource.
type Watcher struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The database watcher name.
	Name *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *WatcherProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// WatcherListResult - The response of a Watcher list operation.
type WatcherListResult struct {
	// REQUIRED; The Watcher items on this page
	Value []*Watcher

	// The link to the next page of items
	NextLink *string
}

// WatcherProperties - The RP specific properties of the resource.
type WatcherProperties struct {
	// The data store for collected monitoring data.
	Datastore *Datastore

	// READ-ONLY; The provisioning state of the resource watcher.
	ProvisioningState *DatabaseWatcherProvisioningState

	// READ-ONLY; The monitoring collection status of the watcher.
	Status *WatcherStatus
}
