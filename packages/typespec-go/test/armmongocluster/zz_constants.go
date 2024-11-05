// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongocluster

const (
	moduleName    = "armmongocluster"
	moduleVersion = "v0.1.0"
)

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CheckNameAvailabilityReason - Possible reasons for a name not being available.
type CheckNameAvailabilityReason string

const (
	// CheckNameAvailabilityReasonAlreadyExists - Name already exists.
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	// CheckNameAvailabilityReasonInvalid - Name is invalid.
	CheckNameAvailabilityReasonInvalid CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// CreateMode - The mode that the Mongo Cluster is created with.
type CreateMode string

const (
	// CreateModeDefault - Create a new mongo cluster.
	CreateModeDefault CreateMode = "Default"
	// CreateModeGeoReplica - Create a replica cluster in distinct geographic region from the source cluster.
	CreateModeGeoReplica CreateMode = "GeoReplica"
	// CreateModePointInTimeRestore - Create a mongo cluster from a restore point-in-time.
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	// CreateModeReplica - Create a replica cluster in the same geographic region as the source cluster.
	CreateModeReplica CreateMode = "Replica"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModeGeoReplica,
		CreateModePointInTimeRestore,
		CreateModeReplica,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// HighAvailabilityMode - The high availability modes for a cluster.
type HighAvailabilityMode string

const (
	// HighAvailabilityModeDisabled - High availability mode is disabled. This mode is can see availability impact during faults
	// or maintenance and is not recommended for production.
	HighAvailabilityModeDisabled HighAvailabilityMode = "Disabled"
	// HighAvailabilityModeSameZone - High availability mode is enabled, where each server in a shard is placed in the same availability
	// zone.
	HighAvailabilityModeSameZone HighAvailabilityMode = "SameZone"
	// HighAvailabilityModeZoneRedundantPreferred - High availability mode is enabled and preferences ZoneRedundant if availability
	// zones capacity is available in the region, otherwise falls-back to provisioning with SameZone.
	HighAvailabilityModeZoneRedundantPreferred HighAvailabilityMode = "ZoneRedundantPreferred"
)

// PossibleHighAvailabilityModeValues returns the possible values for the HighAvailabilityMode const type.
func PossibleHighAvailabilityModeValues() []HighAvailabilityMode {
	return []HighAvailabilityMode{
		HighAvailabilityModeDisabled,
		HighAvailabilityModeSameZone,
		HighAvailabilityModeZoneRedundantPreferred,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PreviewFeature - Preview features that can be enabled on a mongo cluster.
type PreviewFeature string

const (
	// PreviewFeatureGeoReplicas - Enables geo replicas preview feature. The feature must be set at create-time on new cluster
	// to enable linking a geo-replica cluster to it.
	PreviewFeatureGeoReplicas PreviewFeature = "GeoReplicas"
)

// PossiblePreviewFeatureValues returns the possible values for the PreviewFeature const type.
func PossiblePreviewFeatureValues() []PreviewFeature {
	return []PreviewFeature{
		PreviewFeatureGeoReplicas,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateCreating - Connection is being created
	PrivateEndpointConnectionProvisioningStateCreating PrivateEndpointConnectionProvisioningState = "Creating"
	// PrivateEndpointConnectionProvisioningStateDeleting - Connection is being deleted
	PrivateEndpointConnectionProvisioningStateDeleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// PrivateEndpointConnectionProvisioningStateFailed - Connection provisioning has failed
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded - Connection has been provisioned
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// PrivateEndpointServiceConnectionStatusApproved - Connection approved
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	// PrivateEndpointServiceConnectionStatusPending - Connectionaiting for approval or rejection
	PrivateEndpointServiceConnectionStatusPending PrivateEndpointServiceConnectionStatus = "Pending"
	// PrivateEndpointServiceConnectionStatusRejected - Connection Rejected
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// PromoteMode - The mode to apply to a promote operation.
type PromoteMode string

const (
	// PromoteModeSwitchover - Promotion will switch the current replica cluster to the primary role and the original primary
	// will be switched to a replica role, maintaining the replication link.
	PromoteModeSwitchover PromoteMode = "Switchover"
)

// PossiblePromoteModeValues returns the possible values for the PromoteMode const type.
func PossiblePromoteModeValues() []PromoteMode {
	return []PromoteMode{
		PromoteModeSwitchover,
	}
}

// PromoteOption - The option to apply to a promote operation.
type PromoteOption string

const (
	// PromoteOptionForced - Promote option forces the promotion without waiting for the replica to be caught up to the primary.
	// This can result in data-loss so should only be used during disaster recovery scenarios.
	PromoteOptionForced PromoteOption = "Forced"
)

// PossiblePromoteOptionValues returns the possible values for the PromoteOption const type.
func PossiblePromoteOptionValues() []PromoteOption {
	return []PromoteOption{
		PromoteOptionForced,
	}
}

// ProvisioningState - The provisioning state of the last accepted operation.
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDropping - A drop operation is in-progress on the resource.
	ProvisioningStateDropping ProvisioningState = "Dropping"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateInProgress - An operation is in-progress on the resource.
	ProvisioningStateInProgress ProvisioningState = "InProgress"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - An update operation is in-progress on the resource.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateDropping,
		ProvisioningStateFailed,
		ProvisioningStateInProgress,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Whether or not public endpoint access is allowed for this Mongo cluster. Value is optional and default
// value is 'Enabled'
type PublicNetworkAccess string

const (
	// PublicNetworkAccessDisabled - If set, the private endpoints are the exclusive access method.
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	// PublicNetworkAccessEnabled - If set, mongo cluster can be accessed through private and public methods.
	PublicNetworkAccessEnabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ReplicationRole - Replication role of the mongo cluster.
type ReplicationRole string

const (
	// ReplicationRoleAsyncReplica - The cluster is a local asynchronous replica.
	ReplicationRoleAsyncReplica ReplicationRole = "AsyncReplica"
	// ReplicationRoleGeoAsyncReplica - The cluster is a geo-asynchronous replica.
	ReplicationRoleGeoAsyncReplica ReplicationRole = "GeoAsyncReplica"
	// ReplicationRolePrimary - The cluster is a primary replica.
	ReplicationRolePrimary ReplicationRole = "Primary"
)

// PossibleReplicationRoleValues returns the possible values for the ReplicationRole const type.
func PossibleReplicationRoleValues() []ReplicationRole {
	return []ReplicationRole{
		ReplicationRoleAsyncReplica,
		ReplicationRoleGeoAsyncReplica,
		ReplicationRolePrimary,
	}
}

// ReplicationState - The state of the replication link between the replica and source cluster.
type ReplicationState string

const (
	// ReplicationStateActive - Replication link is active.
	ReplicationStateActive ReplicationState = "Active"
	// ReplicationStateBroken - Replication link is broken and the replica may need to be recreated.
	ReplicationStateBroken ReplicationState = "Broken"
	// ReplicationStateCatchup - Replica is catching-up with the primary. This can occur after the replica is created or after
	// a promotion is triggered.
	ReplicationStateCatchup ReplicationState = "Catchup"
	// ReplicationStateProvisioning - Replica and replication link to the primary is being created.
	ReplicationStateProvisioning ReplicationState = "Provisioning"
	// ReplicationStateReconfiguring - Replication link is re-configuring due to a promotion event.
	ReplicationStateReconfiguring ReplicationState = "Reconfiguring"
	// ReplicationStateUpdating - Replication link is being updated due to a change on the replica or an upgrade.
	ReplicationStateUpdating ReplicationState = "Updating"
)

// PossibleReplicationStateValues returns the possible values for the ReplicationState const type.
func PossibleReplicationStateValues() []ReplicationState {
	return []ReplicationState{
		ReplicationStateActive,
		ReplicationStateBroken,
		ReplicationStateCatchup,
		ReplicationStateProvisioning,
		ReplicationStateReconfiguring,
		ReplicationStateUpdating,
	}
}

// Status - The status of the Mongo cluster resource.
type Status string

const (
	// StatusDropping - The mongo cluster resource is being dropped.
	StatusDropping Status = "Dropping"
	// StatusProvisioning - The mongo cluster resource is being provisioned.
	StatusProvisioning Status = "Provisioning"
	// StatusReady - The mongo cluster resource is ready for use.
	StatusReady Status = "Ready"
	// StatusStarting - The mongo cluster resource is being started.
	StatusStarting Status = "Starting"
	// StatusStopped - The mongo cluster resource is stopped.
	StatusStopped Status = "Stopped"
	// StatusStopping - The mongo cluster resource is being stopped.
	StatusStopping Status = "Stopping"
	// StatusUpdating - The mongo cluster resource is being updated.
	StatusUpdating Status = "Updating"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusDropping,
		StatusProvisioning,
		StatusReady,
		StatusStarting,
		StatusStopped,
		StatusStopping,
		StatusUpdating,
	}
}
