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

// ExpressRouteCrossConnectionsOperations contains the methods for the ExpressRouteCrossConnections group.
type ExpressRouteCrossConnectionsOperations interface {
	// BeginCreateOrUpdate - Update the specified ExpressRouteCrossConnection.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsCreateOrUpdateOptions) (*ExpressRouteCrossConnectionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPoller, error)
	// Get - Gets details about the specified ExpressRouteCrossConnection.
	Get(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsGetOptions) (*ExpressRouteCrossConnectionResponse, error)
	// List - Retrieves all the ExpressRouteCrossConnections in a subscription.
	List(options *ExpressRouteCrossConnectionsListOptions) ExpressRouteCrossConnectionListResultPager
	// BeginListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
	BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListArpTableOptions) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error)
	// ResumeListArpTable - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error)
	// ListByResourceGroup - Retrieves all the ExpressRouteCrossConnections in a resource group.
	ListByResourceGroup(resourceGroupName string, options *ExpressRouteCrossConnectionsListByResourceGroupOptions) ExpressRouteCrossConnectionListResultPager
	// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource group.
	BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableOptions) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error)
	// ResumeListRoutesTable - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error)
	// BeginListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource group.
	BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableSummaryOptions) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse, error)
	// ResumeListRoutesTableSummary - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListRoutesTableSummary(token string) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller, error)
	// UpdateTags - Updates an express route cross connection tags.
	UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsUpdateTagsOptions) (*ExpressRouteCrossConnectionResponse, error)
}

// ExpressRouteCrossConnectionsClient implements the ExpressRouteCrossConnectionsOperations interface.
// Don't use this type directly, use NewExpressRouteCrossConnectionsClient() instead.
type ExpressRouteCrossConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCrossConnectionsClient creates a new instance of ExpressRouteCrossConnectionsClient with the specified values.
func NewExpressRouteCrossConnectionsClient(con *armcore.Connection, subscriptionID string) ExpressRouteCrossConnectionsOperations {
	return &ExpressRouteCrossConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *ExpressRouteCrossConnectionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *ExpressRouteCrossConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsCreateOrUpdateOptions) (*ExpressRouteCrossConnectionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCrossConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCrossConnectionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCrossConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCrossConnectionsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Update the specified ExpressRouteCrossConnection.
func (client *ExpressRouteCrossConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteCrossConnectionsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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
func (client *ExpressRouteCrossConnectionsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionResponse, error) {
	result := ExpressRouteCrossConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnection)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCrossConnectionsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets details about the specified ExpressRouteCrossConnection.
func (client *ExpressRouteCrossConnectionsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsGetOptions) (*ExpressRouteCrossConnectionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
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
func (client *ExpressRouteCrossConnectionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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
func (client *ExpressRouteCrossConnectionsClient) GetHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionResponse, error) {
	result := ExpressRouteCrossConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnection)
}

// GetHandleError handles the Get error response.
func (client *ExpressRouteCrossConnectionsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Retrieves all the ExpressRouteCrossConnections in a subscription.
func (client *ExpressRouteCrossConnectionsClient) List(options *ExpressRouteCrossConnectionsListOptions) ExpressRouteCrossConnectionListResultPager {
	return &expressRouteCrossConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCrossConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionsClient) ListCreateRequest(ctx context.Context, options *ExpressRouteCrossConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections"
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
func (client *ExpressRouteCrossConnectionsClient) ListHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionListResultResponse, error) {
	result := ExpressRouteCrossConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionListResult)
}

// ListHandleError handles the List error response.
func (client *ExpressRouteCrossConnectionsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCrossConnectionsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListArpTableOptions) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	resp, err := client.ListArpTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsArpTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListArpTable", "location", resp, client.ListArpTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsArpTableListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCrossConnectionsClient) ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListArpTable", token, client.ListArpTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsArpTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
func (client *ExpressRouteCrossConnectionsClient) ListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListArpTableOptions) (*azcore.Response, error) {
	req, err := client.ListArpTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ListArpTableHandleError(resp)
	}
	return resp, nil
}

// ListArpTableCreateRequest creates the ListArpTable request.
func (client *ExpressRouteCrossConnectionsClient) ListArpTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListArpTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/arpTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListArpTableHandleResponse handles the ListArpTable response.
func (client *ExpressRouteCrossConnectionsClient) ListArpTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
	result := ExpressRouteCircuitsArpTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsArpTableListResult)
}

// ListArpTableHandleError handles the ListArpTable error response.
func (client *ExpressRouteCrossConnectionsClient) ListArpTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Retrieves all the ExpressRouteCrossConnections in a resource group.
func (client *ExpressRouteCrossConnectionsClient) ListByResourceGroup(resourceGroupName string, options *ExpressRouteCrossConnectionsListByResourceGroupOptions) ExpressRouteCrossConnectionListResultPager {
	return &expressRouteCrossConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCrossConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRouteCrossConnectionsClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteCrossConnectionsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections"
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
func (client *ExpressRouteCrossConnectionsClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionListResultResponse, error) {
	result := ExpressRouteCrossConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ExpressRouteCrossConnectionsClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableOptions) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	resp, err := client.ListRoutesTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsRoutesTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListRoutesTable", "location", resp, client.ListRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCrossConnectionsClient) ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListRoutesTable", token, client.ListRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource group.
func (client *ExpressRouteCrossConnectionsClient) ListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableOptions) (*azcore.Response, error) {
	req, err := client.ListRoutesTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ListRoutesTableHandleError(resp)
	}
	return resp, nil
}

// ListRoutesTableCreateRequest creates the ListRoutesTable request.
func (client *ExpressRouteCrossConnectionsClient) ListRoutesTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListRoutesTableHandleResponse handles the ListRoutesTable response.
func (client *ExpressRouteCrossConnectionsClient) ListRoutesTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
	result := ExpressRouteCircuitsRoutesTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsRoutesTableListResult)
}

// ListRoutesTableHandleError handles the ListRoutesTable error response.
func (client *ExpressRouteCrossConnectionsClient) ListRoutesTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableSummaryOptions) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse, error) {
	resp, err := client.ListRoutesTableSummary(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListRoutesTableSummary", "location", resp, client.ListRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCrossConnectionsRoutesTableSummaryListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCrossConnectionsClient) ResumeListRoutesTableSummary(token string) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListRoutesTableSummary", token, client.ListRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionsRoutesTableSummaryListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource group.
func (client *ExpressRouteCrossConnectionsClient) ListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableSummaryOptions) (*azcore.Response, error) {
	req, err := client.ListRoutesTableSummaryCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ListRoutesTableSummaryHandleError(resp)
	}
	return resp, nil
}

// ListRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client *ExpressRouteCrossConnectionsClient) ListRoutesTableSummaryCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableSummaryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListRoutesTableSummaryHandleResponse handles the ListRoutesTableSummary response.
func (client *ExpressRouteCrossConnectionsClient) ListRoutesTableSummaryHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse, error) {
	result := ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionsRoutesTableSummaryListResult)
}

// ListRoutesTableSummaryHandleError handles the ListRoutesTableSummary error response.
func (client *ExpressRouteCrossConnectionsClient) ListRoutesTableSummaryHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates an express route cross connection tags.
func (client *ExpressRouteCrossConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsUpdateTagsOptions) (*ExpressRouteCrossConnectionResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, crossConnectionName, crossConnectionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRouteCrossConnectionsClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(crossConnectionParameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *ExpressRouteCrossConnectionsClient) UpdateTagsHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionResponse, error) {
	result := ExpressRouteCrossConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnection)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *ExpressRouteCrossConnectionsClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
