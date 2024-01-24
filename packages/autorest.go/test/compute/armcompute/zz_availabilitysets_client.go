//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AvailabilitySetsClient contains the methods for the AvailabilitySets group.
// Don't use this type directly, use NewAvailabilitySetsClient() instead.
type AvailabilitySetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAvailabilitySetsClient creates a new instance of AvailabilitySetsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvailabilitySetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailabilitySetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvailabilitySetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an availability set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group.
//   - availabilitySetName - The name of the availability set.
//   - parameters - Parameters supplied to the Create Availability Set operation.
//   - options - AvailabilitySetsClientCreateOrUpdateOptions contains the optional parameters for the AvailabilitySetsClient.CreateOrUpdate
//     method.
func (client *AvailabilitySetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet, options *AvailabilitySetsClientCreateOrUpdateOptions) (AvailabilitySetsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "AvailabilitySetsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, availabilitySetName, parameters, options)
	if err != nil {
		return AvailabilitySetsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AvailabilitySetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AvailabilitySetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AvailabilitySetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet, options *AvailabilitySetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AvailabilitySetsClient) createOrUpdateHandleResponse(resp *http.Response) (AvailabilitySetsClientCreateOrUpdateResponse, error) {
	result := AvailabilitySetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySet); err != nil {
		return AvailabilitySetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an availability set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group.
//   - availabilitySetName - The name of the availability set.
//   - options - AvailabilitySetsClientDeleteOptions contains the optional parameters for the AvailabilitySetsClient.Delete method.
func (client *AvailabilitySetsClient) Delete(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientDeleteOptions) (AvailabilitySetsClientDeleteResponse, error) {
	var err error
	const operationName = "AvailabilitySetsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
	if err != nil {
		return AvailabilitySetsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AvailabilitySetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AvailabilitySetsClientDeleteResponse{}, err
	}
	return AvailabilitySetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AvailabilitySetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about an availability set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group.
//   - availabilitySetName - The name of the availability set.
//   - options - AvailabilitySetsClientGetOptions contains the optional parameters for the AvailabilitySetsClient.Get method.
func (client *AvailabilitySetsClient) Get(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientGetOptions) (AvailabilitySetsClientGetResponse, error) {
	var err error
	const operationName = "AvailabilitySetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
	if err != nil {
		return AvailabilitySetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AvailabilitySetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AvailabilitySetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AvailabilitySetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AvailabilitySetsClient) getHandleResponse(resp *http.Response) (AvailabilitySetsClientGetResponse, error) {
	result := AvailabilitySetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySet); err != nil {
		return AvailabilitySetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all availability sets in a resource group.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group.
//   - options - AvailabilitySetsClientListOptions contains the optional parameters for the AvailabilitySetsClient.NewListPager
//     method.
func (client *AvailabilitySetsClient) NewListPager(resourceGroupName string, options *AvailabilitySetsClientListOptions) *runtime.Pager[AvailabilitySetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailabilitySetsClientListResponse]{
		More: func(page AvailabilitySetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailabilitySetsClientListResponse) (AvailabilitySetsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AvailabilitySetsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AvailabilitySetsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AvailabilitySetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *AvailabilitySetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailabilitySetsClient) listHandleResponse(resp *http.Response) (AvailabilitySetsClientListResponse, error) {
	result := AvailabilitySetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySetListResult); err != nil {
		return AvailabilitySetsClientListResponse{}, err
	}
	return result, nil
}

// NewListAvailableSizesPager - Lists all available virtual machine sizes that can be used to create a new virtual machine
// in an existing availability set.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group.
//   - availabilitySetName - The name of the availability set.
//   - options - AvailabilitySetsClientListAvailableSizesOptions contains the optional parameters for the AvailabilitySetsClient.NewListAvailableSizesPager
//     method.
func (client *AvailabilitySetsClient) NewListAvailableSizesPager(resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientListAvailableSizesOptions) *runtime.Pager[AvailabilitySetsClientListAvailableSizesResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailabilitySetsClientListAvailableSizesResponse]{
		More: func(page AvailabilitySetsClientListAvailableSizesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AvailabilitySetsClientListAvailableSizesResponse) (AvailabilitySetsClientListAvailableSizesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AvailabilitySetsClient.NewListAvailableSizesPager")
			req, err := client.listAvailableSizesCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
			if err != nil {
				return AvailabilitySetsClientListAvailableSizesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AvailabilitySetsClientListAvailableSizesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailabilitySetsClientListAvailableSizesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAvailableSizesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAvailableSizesCreateRequest creates the ListAvailableSizes request.
func (client *AvailabilitySetsClient) listAvailableSizesCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsClientListAvailableSizesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAvailableSizesHandleResponse handles the ListAvailableSizes response.
func (client *AvailabilitySetsClient) listAvailableSizesHandleResponse(resp *http.Response) (AvailabilitySetsClientListAvailableSizesResponse, error) {
	result := AvailabilitySetsClientListAvailableSizesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineSizeListResult); err != nil {
		return AvailabilitySetsClientListAvailableSizesResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all availability sets in a subscription.
//
// Generated from API version 2021-11-01
//   - options - AvailabilitySetsClientListBySubscriptionOptions contains the optional parameters for the AvailabilitySetsClient.NewListBySubscriptionPager
//     method.
func (client *AvailabilitySetsClient) NewListBySubscriptionPager(options *AvailabilitySetsClientListBySubscriptionOptions) *runtime.Pager[AvailabilitySetsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailabilitySetsClientListBySubscriptionResponse]{
		More: func(page AvailabilitySetsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailabilitySetsClientListBySubscriptionResponse) (AvailabilitySetsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AvailabilitySetsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AvailabilitySetsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AvailabilitySetsClient) listBySubscriptionCreateRequest(ctx context.Context, options *AvailabilitySetsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/availabilitySets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AvailabilitySetsClient) listBySubscriptionHandleResponse(resp *http.Response) (AvailabilitySetsClientListBySubscriptionResponse, error) {
	result := AvailabilitySetsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySetListResult); err != nil {
		return AvailabilitySetsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an availability set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group.
//   - availabilitySetName - The name of the availability set.
//   - parameters - Parameters supplied to the Update Availability Set operation.
//   - options - AvailabilitySetsClientUpdateOptions contains the optional parameters for the AvailabilitySetsClient.Update method.
func (client *AvailabilitySetsClient) Update(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate, options *AvailabilitySetsClientUpdateOptions) (AvailabilitySetsClientUpdateResponse, error) {
	var err error
	const operationName = "AvailabilitySetsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, availabilitySetName, parameters, options)
	if err != nil {
		return AvailabilitySetsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AvailabilitySetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AvailabilitySetsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AvailabilitySetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate, options *AvailabilitySetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if availabilitySetName == "" {
		return nil, errors.New("parameter availabilitySetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AvailabilitySetsClient) updateHandleResponse(resp *http.Response) (AvailabilitySetsClientUpdateResponse, error) {
	result := AvailabilitySetsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilitySet); err != nil {
		return AvailabilitySetsClientUpdateResponse{}, err
	}
	return result, nil
}
