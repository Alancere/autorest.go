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

// ExpressRouteLinksClient contains the methods for the ExpressRouteLinks group.
// Don't use this type directly, use NewExpressRouteLinksClient() instead.
type ExpressRouteLinksClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteLinksClient creates a new instance of ExpressRouteLinksClient with the specified values.
func NewExpressRouteLinksClient(con *armcore.Connection, subscriptionID string) ExpressRouteLinksClient {
	return ExpressRouteLinksClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client ExpressRouteLinksClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Retrieves the specified ExpressRouteLink resource.
func (client ExpressRouteLinksClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string, options *ExpressRouteLinksGetOptions) (ExpressRouteLinkResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRoutePortName, linkName, options)
	if err != nil {
		return ExpressRouteLinkResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteLinkResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteLinkResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return ExpressRouteLinkResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client ExpressRouteLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string, options *ExpressRouteLinksGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	urlPath = strings.ReplaceAll(urlPath, "{linkName}", url.PathEscape(linkName))
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
func (client ExpressRouteLinksClient) getHandleResponse(resp *azcore.Response) (ExpressRouteLinkResponse, error) {
	result := ExpressRouteLinkResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ExpressRouteLink)
	return result, err
}

// getHandleError handles the Get error response.
func (client ExpressRouteLinksClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
func (client ExpressRouteLinksClient) List(resourceGroupName string, expressRoutePortName string, options *ExpressRouteLinksListOptions) ExpressRouteLinkListResultPager {
	return &expressRouteLinkListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ExpressRouteLinkListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteLinkListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client ExpressRouteLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRouteLinksListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
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
func (client ExpressRouteLinksClient) listHandleResponse(resp *azcore.Response) (ExpressRouteLinkListResultResponse, error) {
	result := ExpressRouteLinkListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ExpressRouteLinkListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client ExpressRouteLinksClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
