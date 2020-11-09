// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualRoutersOperations contains the methods for the VirtualRouters group.
type VirtualRoutersOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified Virtual Router.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, parameters VirtualRouter, options *VirtualRoutersCreateOrUpdateOptions) (*VirtualRouterPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualRouterPoller, error)
	// BeginDelete - Deletes the specified Virtual Router.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRoutersDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified Virtual Router.
	Get(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRoutersGetOptions) (*VirtualRouterResponse, error)
	// List - Gets all the Virtual Routers in a subscription.
	List(options *VirtualRoutersListOptions) VirtualRouterListResultPager
	// ListByResourceGroup - Lists all Virtual Routers in a resource group.
	ListByResourceGroup(resourceGroupName string, options *VirtualRoutersListByResourceGroupOptions) VirtualRouterListResultPager
}

// VirtualRoutersClient implements the VirtualRoutersOperations interface.
// Don't use this type directly, use NewVirtualRoutersClient() instead.
type VirtualRoutersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualRoutersClient creates a new instance of VirtualRoutersClient with the specified values.
func NewVirtualRoutersClient(con *armcore.Connection, subscriptionID string) VirtualRoutersOperations {
	return &VirtualRoutersClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VirtualRoutersClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *VirtualRoutersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, parameters VirtualRouter, options *VirtualRoutersCreateOrUpdateOptions) (*VirtualRouterPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, virtualRouterName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualRouterPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualRoutersClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualRouterPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualRouterResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualRoutersClient) ResumeCreateOrUpdate(token string) (VirtualRouterPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualRoutersClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualRouterPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified Virtual Router.
func (client *VirtualRoutersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, parameters VirtualRouter, options *VirtualRoutersCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, virtualRouterName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualRoutersClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, parameters VirtualRouter, options *VirtualRoutersCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualRoutersClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualRouterResponse, error) {
	result := VirtualRouterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualRouter)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualRoutersClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VirtualRoutersClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRoutersDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, virtualRouterName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualRoutersClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualRoutersClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualRoutersClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified Virtual Router.
func (client *VirtualRoutersClient) Delete(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRoutersDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, virtualRouterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *VirtualRoutersClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRoutersDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *VirtualRoutersClient) DeleteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified Virtual Router.
func (client *VirtualRoutersClient) Get(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRoutersGetOptions) (*VirtualRouterResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, virtualRouterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *VirtualRoutersClient) GetCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRoutersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualRoutersClient) GetHandleResponse(resp *azcore.Response) (*VirtualRouterResponse, error) {
	result := VirtualRouterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualRouter)
}

// GetHandleError handles the Get error response.
func (client *VirtualRoutersClient) GetHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the Virtual Routers in a subscription.
func (client *VirtualRoutersClient) List(options *VirtualRoutersListOptions) VirtualRouterListResultPager {
	return &virtualRouterListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *VirtualRouterListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualRouterListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *VirtualRoutersClient) ListCreateRequest(ctx context.Context, options *VirtualRoutersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualRouters"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *VirtualRoutersClient) ListHandleResponse(resp *azcore.Response) (*VirtualRouterListResultResponse, error) {
	result := VirtualRouterListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualRouterListResult)
}

// ListHandleError handles the List error response.
func (client *VirtualRoutersClient) ListHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all Virtual Routers in a resource group.
func (client *VirtualRoutersClient) ListByResourceGroup(resourceGroupName string, options *VirtualRoutersListByResourceGroupOptions) VirtualRouterListResultPager {
	return &virtualRouterListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *VirtualRouterListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualRouterListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualRoutersClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualRoutersListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualRoutersClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*VirtualRouterListResultResponse, error) {
	result := VirtualRouterListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualRouterListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VirtualRoutersClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
