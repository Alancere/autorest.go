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
)

// LoadBalancerProbesOperations contains the methods for the LoadBalancerProbes group.
type LoadBalancerProbesOperations interface {
	// Get - Gets load balancer probe.
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string, options *LoadBalancerProbesGetOptions) (*ProbeResponse, error)
	// List - Gets all the load balancer probes.
	List(resourceGroupName string, loadBalancerName string, options *LoadBalancerProbesListOptions) LoadBalancerProbeListResultPager
}

// LoadBalancerProbesClient implements the LoadBalancerProbesOperations interface.
// Don't use this type directly, use NewLoadBalancerProbesClient() instead.
type LoadBalancerProbesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerProbesClient creates a new instance of LoadBalancerProbesClient with the specified values.
func NewLoadBalancerProbesClient(con *armcore.Connection, subscriptionID string) LoadBalancerProbesOperations {
	return &LoadBalancerProbesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *LoadBalancerProbesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Gets load balancer probe.
func (client *LoadBalancerProbesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string, options *LoadBalancerProbesGetOptions) (*ProbeResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, loadBalancerName, probeName, options)
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
func (client *LoadBalancerProbesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string, options *LoadBalancerProbesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{probeName}", url.PathEscape(probeName))
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

// GetHandleResponse handles the Get response.
func (client *LoadBalancerProbesClient) GetHandleResponse(resp *azcore.Response) (*ProbeResponse, error) {
	result := ProbeResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Probe)
}

// GetHandleError handles the Get error response.
func (client *LoadBalancerProbesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the load balancer probes.
func (client *LoadBalancerProbesClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerProbesListOptions) LoadBalancerProbeListResultPager {
	return &loadBalancerProbeListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *LoadBalancerProbeListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerProbeListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *LoadBalancerProbesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerProbesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *LoadBalancerProbesClient) ListHandleResponse(resp *azcore.Response) (*LoadBalancerProbeListResultResponse, error) {
	result := LoadBalancerProbeListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancerProbeListResult)
}

// ListHandleError handles the List error response.
func (client *LoadBalancerProbesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
