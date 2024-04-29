// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasewatcher

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	PagedOperation
}

// SharedPrivateLinkResourcesClientCreateResponse contains the response from method SharedPrivateLinkResourcesClient.BeginCreate.
type SharedPrivateLinkResourcesClientCreateResponse struct {
	// Concrete proxy resource types can be created by aliasing this type using a specific property type.
	SharedPrivateLinkResource
}

// SharedPrivateLinkResourcesClientDeleteResponse contains the response from method SharedPrivateLinkResourcesClient.BeginDelete.
type SharedPrivateLinkResourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// SharedPrivateLinkResourcesClientGetResponse contains the response from method SharedPrivateLinkResourcesClient.Get.
type SharedPrivateLinkResourcesClientGetResponse struct {
	// Concrete proxy resource types can be created by aliasing this type using a specific property type.
	SharedPrivateLinkResource
}

// SharedPrivateLinkResourcesClientListByWatcherResponse contains the response from method SharedPrivateLinkResourcesClient.NewListByWatcherPager.
type SharedPrivateLinkResourcesClientListByWatcherResponse struct {
	// The response of a SharedPrivateLinkResource list operation.
	SharedPrivateLinkResourceListResult
}

// TargetsClientCreateOrUpdateResponse contains the response from method TargetsClient.CreateOrUpdate.
type TargetsClientCreateOrUpdateResponse struct {
	// Concrete proxy resource types can be created by aliasing this type using a specific property type.
	Target
}

// TargetsClientDeleteResponse contains the response from method TargetsClient.Delete.
type TargetsClientDeleteResponse struct {
	// placeholder for future response values
}

// TargetsClientGetResponse contains the response from method TargetsClient.Get.
type TargetsClientGetResponse struct {
	// Concrete proxy resource types can be created by aliasing this type using a specific property type.
	Target
}

// TargetsClientListByWatcherResponse contains the response from method TargetsClient.NewListByWatcherPager.
type TargetsClientListByWatcherResponse struct {
	// The response of a Target list operation.
	TargetListResult
}

// WatchersClientCreateOrUpdateResponse contains the response from method WatchersClient.BeginCreateOrUpdate.
type WatchersClientCreateOrUpdateResponse struct {
	// The DatabaseWatcherProviderHub resource.
	Watcher
}

// WatchersClientDeleteResponse contains the response from method WatchersClient.BeginDelete.
type WatchersClientDeleteResponse struct {
	// placeholder for future response values
}

// WatchersClientGetResponse contains the response from method WatchersClient.Get.
type WatchersClientGetResponse struct {
	// The DatabaseWatcherProviderHub resource.
	Watcher
}

// WatchersClientListByResourceGroupResponse contains the response from method WatchersClient.NewListByResourceGroupPager.
type WatchersClientListByResourceGroupResponse struct {
	// The response of a Watcher list operation.
	WatcherListResult
}

// WatchersClientListBySubscriptionResponse contains the response from method WatchersClient.NewListBySubscriptionPager.
type WatchersClientListBySubscriptionResponse struct {
	// The response of a Watcher list operation.
	WatcherListResult
}

// WatchersClientStartResponse contains the response from method WatchersClient.BeginStart.
type WatchersClientStartResponse struct {
	// placeholder for future response values
}

// WatchersClientStopResponse contains the response from method WatchersClient.BeginStop.
type WatchersClientStopResponse struct {
	// placeholder for future response values
}

// WatchersClientUpdateResponse contains the response from method WatchersClient.BeginUpdate.
type WatchersClientUpdateResponse struct {
	// The DatabaseWatcherProviderHub resource.
	Watcher
}
