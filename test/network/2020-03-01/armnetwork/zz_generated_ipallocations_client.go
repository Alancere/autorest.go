//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IPAllocationsClient contains the methods for the IPAllocations group.
// Don't use this type directly, use NewIPAllocationsClient() instead.
type IPAllocationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIPAllocationsClient creates a new instance of IPAllocationsClient with the specified values.
func NewIPAllocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IPAllocationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &IPAllocationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates an IpAllocation in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (IPAllocationsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return IPAllocationsCreateOrUpdatePollerResponse{}, err
	}
	result := IPAllocationsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IPAllocationsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return IPAllocationsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &IPAllocationsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an IpAllocation in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IPAllocationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *IPAllocationsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified IpAllocation.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) BeginDelete(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (IPAllocationsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return IPAllocationsDeletePollerResponse{}, err
	}
	result := IPAllocationsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IPAllocationsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return IPAllocationsDeletePollerResponse{}, err
	}
	result.Poller = &IPAllocationsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified IpAllocation.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) deleteOperation(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IPAllocationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *IPAllocationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified IpAllocation by resource group.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) Get(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsGetOptions) (IPAllocationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return IPAllocationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPAllocationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPAllocationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPAllocationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPAllocationsClient) getHandleResponse(resp *http.Response) (IPAllocationsGetResponse, error) {
	result := IPAllocationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAllocation); err != nil {
		return IPAllocationsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IPAllocationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all IpAllocations in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) List(options *IPAllocationsListOptions) *IPAllocationsListPager {
	return &IPAllocationsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp IPAllocationsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IPAllocationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IPAllocationsClient) listCreateRequest(ctx context.Context, options *IPAllocationsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/IpAllocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IPAllocationsClient) listHandleResponse(resp *http.Response) (IPAllocationsListResponse, error) {
	result := IPAllocationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAllocationListResult); err != nil {
		return IPAllocationsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IPAllocationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets all IpAllocations in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) ListByResourceGroup(resourceGroupName string, options *IPAllocationsListByResourceGroupOptions) *IPAllocationsListByResourceGroupPager {
	return &IPAllocationsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp IPAllocationsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IPAllocationListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IPAllocationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPAllocationsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IPAllocationsClient) listByResourceGroupHandleResponse(resp *http.Response) (IPAllocationsListByResourceGroupResponse, error) {
	result := IPAllocationsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAllocationListResult); err != nil {
		return IPAllocationsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *IPAllocationsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateTags - Updates a IpAllocation tags.
// If the operation fails it returns the *CloudError error type.
func (client *IPAllocationsClient) UpdateTags(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters TagsObject, options *IPAllocationsUpdateTagsOptions) (IPAllocationsUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return IPAllocationsUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPAllocationsUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPAllocationsUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *IPAllocationsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters TagsObject, options *IPAllocationsUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipAllocationName == "" {
		return nil, errors.New("parameter ipAllocationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *IPAllocationsClient) updateTagsHandleResponse(resp *http.Response) (IPAllocationsUpdateTagsResponse, error) {
	result := IPAllocationsUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPAllocation); err != nil {
		return IPAllocationsUpdateTagsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *IPAllocationsClient) updateTagsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
