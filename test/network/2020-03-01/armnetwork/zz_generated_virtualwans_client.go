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

// VirtualWansClient contains the methods for the VirtualWans group.
// Don't use this type directly, use NewVirtualWansClient() instead.
type VirtualWansClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualWansClient creates a new instance of VirtualWansClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualWansClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualWansClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &VirtualWansClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a VirtualWAN resource if it doesn't exist else updates the existing VirtualWAN.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The resource group name of the VirtualWan.
// virtualWANName - The name of the VirtualWAN being created or updated.
// wanParameters - Parameters supplied to create or update VirtualWAN.
// options - VirtualWansBeginCreateOrUpdateOptions contains the optional parameters for the VirtualWansClient.BeginCreateOrUpdate
// method.
func (client *VirtualWansClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters VirtualWAN, options *VirtualWansBeginCreateOrUpdateOptions) (VirtualWansCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualWANName, wanParameters, options)
	if err != nil {
		return VirtualWansCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualWansCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualWansClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualWansCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualWansCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a VirtualWAN resource if it doesn't exist else updates the existing VirtualWAN.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters VirtualWAN, options *VirtualWansBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualWANName, wanParameters, options)
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
func (client *VirtualWansClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters VirtualWAN, options *VirtualWansBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWANName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, wanParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualWansClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a VirtualWAN.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The resource group name of the VirtualWan.
// virtualWANName - The name of the VirtualWAN being deleted.
// options - VirtualWansBeginDeleteOptions contains the optional parameters for the VirtualWansClient.BeginDelete method.
func (client *VirtualWansClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansBeginDeleteOptions) (VirtualWansDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return VirtualWansDeletePollerResponse{}, err
	}
	result := VirtualWansDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualWansClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualWansDeletePollerResponse{}, err
	}
	result.Poller = &VirtualWansDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a VirtualWAN.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualWansClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualWANName, options)
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
func (client *VirtualWansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWANName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
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
func (client *VirtualWansClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieves the details of a VirtualWAN.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The resource group name of the VirtualWan.
// virtualWANName - The name of the VirtualWAN being retrieved.
// options - VirtualWansGetOptions contains the optional parameters for the VirtualWansClient.Get method.
func (client *VirtualWansClient) Get(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansGetOptions) (VirtualWansGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return VirtualWansGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualWansGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualWansGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualWansClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, options *VirtualWansGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWANName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualWansClient) getHandleResponse(resp *http.Response) (VirtualWansGetResponse, error) {
	result := VirtualWansGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualWAN); err != nil {
		return VirtualWansGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualWansClient) getHandleError(resp *http.Response) error {
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

// List - Lists all the VirtualWANs in a subscription.
// If the operation fails it returns the *CloudError error type.
// options - VirtualWansListOptions contains the optional parameters for the VirtualWansClient.List method.
func (client *VirtualWansClient) List(options *VirtualWansListOptions) *VirtualWansListPager {
	return &VirtualWansListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VirtualWansListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVirtualWANsResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualWansClient) listCreateRequest(ctx context.Context, options *VirtualWansListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualWans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *VirtualWansClient) listHandleResponse(resp *http.Response) (VirtualWansListResponse, error) {
	result := VirtualWansListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualWANsResult); err != nil {
		return VirtualWansListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualWansClient) listHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Lists all the VirtualWANs in a resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The resource group name of the VirtualWan.
// options - VirtualWansListByResourceGroupOptions contains the optional parameters for the VirtualWansClient.ListByResourceGroup
// method.
func (client *VirtualWansClient) ListByResourceGroup(resourceGroupName string, options *VirtualWansListByResourceGroupOptions) *VirtualWansListByResourceGroupPager {
	return &VirtualWansListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp VirtualWansListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVirtualWANsResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualWansClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualWansListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *VirtualWansClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualWansListByResourceGroupResponse, error) {
	result := VirtualWansListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualWANsResult); err != nil {
		return VirtualWansListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VirtualWansClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// UpdateTags - Updates a VirtualWAN tags.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The resource group name of the VirtualWan.
// virtualWANName - The name of the VirtualWAN being updated.
// wanParameters - Parameters supplied to Update VirtualWAN tags.
// options - VirtualWansUpdateTagsOptions contains the optional parameters for the VirtualWansClient.UpdateTags method.
func (client *VirtualWansClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters TagsObject, options *VirtualWansUpdateTagsOptions) (VirtualWansUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualWANName, wanParameters, options)
	if err != nil {
		return VirtualWansUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualWansUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualWansUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualWansClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, wanParameters TagsObject, options *VirtualWansUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWANName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, wanParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualWansClient) updateTagsHandleResponse(resp *http.Response) (VirtualWansUpdateTagsResponse, error) {
	result := VirtualWansUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualWAN); err != nil {
		return VirtualWansUpdateTagsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VirtualWansClient) updateTagsHandleError(resp *http.Response) error {
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
