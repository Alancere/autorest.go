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

// RouteFilterRulesClient contains the methods for the RouteFilterRules group.
// Don't use this type directly, use NewRouteFilterRulesClient() instead.
type RouteFilterRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewRouteFilterRulesClient creates a new instance of RouteFilterRulesClient with the specified values.
func NewRouteFilterRulesClient(con *armcore.Connection, subscriptionID string) RouteFilterRulesClient {
	return RouteFilterRulesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client RouteFilterRulesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates a route in the specified route filter.
func (client RouteFilterRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule, options *RouteFilterRulesCreateOrUpdateOptions) (RouteFilterRulePollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, routeFilterName, ruleName, routeFilterRuleParameters, options)
	if err != nil {
		return RouteFilterRulePollerResponse{}, err
	}
	result := RouteFilterRulePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RouteFilterRulesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return RouteFilterRulePollerResponse{}, err
	}
	poller := &routeFilterRulePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (RouteFilterRuleResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new RouteFilterRulePoller from the specified resume token.
// token - The value must come from a previous call to RouteFilterRulePoller.ResumeToken().
func (client RouteFilterRulesClient) ResumeCreateOrUpdate(token string) (RouteFilterRulePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RouteFilterRulesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &routeFilterRulePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a route in the specified route filter.
func (client RouteFilterRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule, options *RouteFilterRulesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeFilterName, ruleName, routeFilterRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client RouteFilterRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule, options *RouteFilterRulesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(routeFilterRuleParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client RouteFilterRulesClient) createOrUpdateHandleResponse(resp *azcore.Response) (RouteFilterRuleResponse, error) {
	result := RouteFilterRuleResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.RouteFilterRule)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client RouteFilterRulesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified rule from a route filter.
func (client RouteFilterRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, routeFilterName, ruleName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RouteFilterRulesClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
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

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client RouteFilterRulesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RouteFilterRulesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified rule from a route filter.
func (client RouteFilterRulesClient) Delete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeFilterName, ruleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client RouteFilterRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client RouteFilterRulesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified rule from a route filter.
func (client RouteFilterRulesClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesGetOptions) (RouteFilterRuleResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeFilterName, ruleName, options)
	if err != nil {
		return RouteFilterRuleResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return RouteFilterRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RouteFilterRuleResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return RouteFilterRuleResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client RouteFilterRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client RouteFilterRulesClient) getHandleResponse(resp *azcore.Response) (RouteFilterRuleResponse, error) {
	result := RouteFilterRuleResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.RouteFilterRule)
	return result, err
}

// getHandleError handles the Get error response.
func (client RouteFilterRulesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByRouteFilter - Gets all RouteFilterRules in a route filter.
func (client RouteFilterRulesClient) ListByRouteFilter(resourceGroupName string, routeFilterName string, options *RouteFilterRulesListByRouteFilterOptions) RouteFilterRuleListResultPager {
	return &routeFilterRuleListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByRouteFilterCreateRequest(ctx, resourceGroupName, routeFilterName, options)
		},
		responder: client.listByRouteFilterHandleResponse,
		errorer:   client.listByRouteFilterHandleError,
		advancer: func(ctx context.Context, resp RouteFilterRuleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RouteFilterRuleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByRouteFilterCreateRequest creates the ListByRouteFilter request.
func (client RouteFilterRulesClient) listByRouteFilterCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFilterRulesListByRouteFilterOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listByRouteFilterHandleResponse handles the ListByRouteFilter response.
func (client RouteFilterRulesClient) listByRouteFilterHandleResponse(resp *azcore.Response) (RouteFilterRuleListResultResponse, error) {
	result := RouteFilterRuleListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.RouteFilterRuleListResult)
	return result, err
}

// listByRouteFilterHandleError handles the ListByRouteFilter error response.
func (client RouteFilterRulesClient) listByRouteFilterHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
