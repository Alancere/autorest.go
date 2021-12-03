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

// VirtualRouterPeeringsClient contains the methods for the VirtualRouterPeerings group.
// Don't use this type directly, use NewVirtualRouterPeeringsClient() instead.
type VirtualRouterPeeringsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualRouterPeeringsClient creates a new instance of VirtualRouterPeeringsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualRouterPeeringsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualRouterPeeringsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &VirtualRouterPeeringsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the specified Virtual Router Peering.
// If the operation fails it returns the *Error error type.
// resourceGroupName - The name of the resource group.
// virtualRouterName - The name of the Virtual Router.
// peeringName - The name of the Virtual Router Peering.
// parameters - Parameters supplied to the create or update Virtual Router Peering operation.
// options - VirtualRouterPeeringsBeginCreateOrUpdateOptions contains the optional parameters for the VirtualRouterPeeringsClient.BeginCreateOrUpdate
// method.
func (client *VirtualRouterPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsBeginCreateOrUpdateOptions) (VirtualRouterPeeringsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualRouterName, peeringName, parameters, options)
	if err != nil {
		return VirtualRouterPeeringsCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualRouterPeeringsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualRouterPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualRouterPeeringsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualRouterPeeringsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified Virtual Router Peering.
// If the operation fails it returns the *Error error type.
func (client *VirtualRouterPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, parameters, options)
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
func (client *VirtualRouterPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualRouterName == "" {
		return nil, errors.New("parameter virtualRouterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
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
func (client *VirtualRouterPeeringsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified peering from a Virtual Router.
// If the operation fails it returns the *Error error type.
// resourceGroupName - The name of the resource group.
// virtualRouterName - The name of the Virtual Router.
// peeringName - The name of the peering.
// options - VirtualRouterPeeringsBeginDeleteOptions contains the optional parameters for the VirtualRouterPeeringsClient.BeginDelete
// method.
func (client *VirtualRouterPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsBeginDeleteOptions) (VirtualRouterPeeringsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualRouterName, peeringName, options)
	if err != nil {
		return VirtualRouterPeeringsDeletePollerResponse{}, err
	}
	result := VirtualRouterPeeringsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualRouterPeeringsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualRouterPeeringsDeletePollerResponse{}, err
	}
	result.Poller = &VirtualRouterPeeringsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified peering from a Virtual Router.
// If the operation fails it returns the *Error error type.
func (client *VirtualRouterPeeringsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, options)
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
func (client *VirtualRouterPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualRouterName == "" {
		return nil, errors.New("parameter virtualRouterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *VirtualRouterPeeringsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified Virtual Router Peering.
// If the operation fails it returns the *Error error type.
// resourceGroupName - The name of the resource group.
// virtualRouterName - The name of the Virtual Router.
// peeringName - The name of the Virtual Router Peering.
// options - VirtualRouterPeeringsGetOptions contains the optional parameters for the VirtualRouterPeeringsClient.Get method.
func (client *VirtualRouterPeeringsClient) Get(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsGetOptions) (VirtualRouterPeeringsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, options)
	if err != nil {
		return VirtualRouterPeeringsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualRouterPeeringsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualRouterPeeringsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualRouterPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualRouterName == "" {
		return nil, errors.New("parameter virtualRouterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
func (client *VirtualRouterPeeringsClient) getHandleResponse(resp *http.Response) (VirtualRouterPeeringsGetResponse, error) {
	result := VirtualRouterPeeringsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualRouterPeering); err != nil {
		return VirtualRouterPeeringsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualRouterPeeringsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all Virtual Router Peerings in a Virtual Router resource.
// If the operation fails it returns the *Error error type.
// resourceGroupName - The name of the resource group.
// virtualRouterName - The name of the Virtual Router.
// options - VirtualRouterPeeringsListOptions contains the optional parameters for the VirtualRouterPeeringsClient.List method.
func (client *VirtualRouterPeeringsClient) List(resourceGroupName string, virtualRouterName string, options *VirtualRouterPeeringsListOptions) *VirtualRouterPeeringsListPager {
	return &VirtualRouterPeeringsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualRouterName, options)
		},
		advancer: func(ctx context.Context, resp VirtualRouterPeeringsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualRouterPeeringListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualRouterPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRouterPeeringsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualRouterName == "" {
		return nil, errors.New("parameter virtualRouterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
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
func (client *VirtualRouterPeeringsClient) listHandleResponse(resp *http.Response) (VirtualRouterPeeringsListResponse, error) {
	result := VirtualRouterPeeringsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualRouterPeeringListResult); err != nil {
		return VirtualRouterPeeringsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualRouterPeeringsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
