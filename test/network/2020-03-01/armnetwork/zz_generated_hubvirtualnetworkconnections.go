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

// HubVirtualNetworkConnectionsClient contains the methods for the HubVirtualNetworkConnections group.
// Don't use this type directly, use NewHubVirtualNetworkConnectionsClient() instead.
type HubVirtualNetworkConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewHubVirtualNetworkConnectionsClient creates a new instance of HubVirtualNetworkConnectionsClient with the specified values.
func NewHubVirtualNetworkConnectionsClient(con *armcore.Connection, subscriptionID string) HubVirtualNetworkConnectionsClient {
	return HubVirtualNetworkConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client HubVirtualNetworkConnectionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Retrieves the details of a HubVirtualNetworkConnection.
func (client HubVirtualNetworkConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsGetOptions) (HubVirtualNetworkConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return HubVirtualNetworkConnectionResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return HubVirtualNetworkConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HubVirtualNetworkConnectionResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return HubVirtualNetworkConnectionResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client HubVirtualNetworkConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client HubVirtualNetworkConnectionsClient) getHandleResponse(resp *azcore.Response) (HubVirtualNetworkConnectionResponse, error) {
	result := HubVirtualNetworkConnectionResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.HubVirtualNetworkConnection)
	return result, err
}

// getHandleError handles the Get error response.
func (client HubVirtualNetworkConnectionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Retrieves the details of all HubVirtualNetworkConnections.
func (client HubVirtualNetworkConnectionsClient) List(resourceGroupName string, virtualHubName string, options *HubVirtualNetworkConnectionsListOptions) ListHubVirtualNetworkConnectionsResultPager {
	return &listHubVirtualNetworkConnectionsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ListHubVirtualNetworkConnectionsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListHubVirtualNetworkConnectionsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client HubVirtualNetworkConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *HubVirtualNetworkConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client HubVirtualNetworkConnectionsClient) listHandleResponse(resp *azcore.Response) (ListHubVirtualNetworkConnectionsResultResponse, error) {
	result := ListHubVirtualNetworkConnectionsResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ListHubVirtualNetworkConnectionsResult)
	return result, err
}

// listHandleError handles the List error response.
func (client HubVirtualNetworkConnectionsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
