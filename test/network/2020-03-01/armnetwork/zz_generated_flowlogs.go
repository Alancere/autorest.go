// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FlowLogsClient contains the methods for the FlowLogs group.
// Don't use this type directly, use NewFlowLogsClient() instead.
type FlowLogsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFlowLogsClient creates a new instance of FlowLogsClient with the specified values.
func NewFlowLogsClient(con *armcore.Connection, subscriptionID string) *FlowLogsClient {
	return &FlowLogsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a flow log for the specified network security group.
func (client *FlowLogsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsBeginCreateOrUpdateOptions) (FlowLogPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkWatcherName, flowLogName, parameters, options)
	if err != nil {
		return FlowLogPollerResponse{}, err
	}
	result := FlowLogPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FlowLogsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return FlowLogPollerResponse{}, err
	}
	poller := &flowLogPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (FlowLogResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new FlowLogPoller from the specified resume token.
// token - The value must come from a previous call to FlowLogPoller.ResumeToken().
func (client *FlowLogsClient) ResumeCreateOrUpdate(token string) (FlowLogPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FlowLogsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &flowLogPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a flow log for the specified network security group.
func (client *FlowLogsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FlowLogsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if flowLogName == "" {
		return nil, errors.New("parameter flowLogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FlowLogsClient) createOrUpdateHandleResponse(resp *azcore.Response) (FlowLogResponse, error) {
	var val *FlowLog
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FlowLogResponse{}, err
	}
	return FlowLogResponse{RawResponse: resp.Response, FlowLog: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FlowLogsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified flow log resource.
func (client *FlowLogsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FlowLogsClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *FlowLogsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FlowLogsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified flow log resource.
func (client *FlowLogsClient) delete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FlowLogsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if flowLogName == "" {
		return nil, errors.New("parameter flowLogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
func (client *FlowLogsClient) deleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets a flow log resource by name.
func (client *FlowLogsClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsGetOptions) (FlowLogResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
	if err != nil {
		return FlowLogResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FlowLogResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FlowLogResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FlowLogsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if flowLogName == "" {
		return nil, errors.New("parameter flowLogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
func (client *FlowLogsClient) getHandleResponse(resp *azcore.Response) (FlowLogResponse, error) {
	var val *FlowLog
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FlowLogResponse{}, err
	}
	return FlowLogResponse{RawResponse: resp.Response, FlowLog: val}, nil
}

// getHandleError handles the Get error response.
func (client *FlowLogsClient) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all flow log resources for the specified Network Watcher.
func (client *FlowLogsClient) List(resourceGroupName string, networkWatcherName string, options *FlowLogsListOptions) FlowLogListResultPager {
	return &flowLogListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp FlowLogListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FlowLogListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *FlowLogsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *FlowLogsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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

// listHandleResponse handles the List response.
func (client *FlowLogsClient) listHandleResponse(resp *azcore.Response) (FlowLogListResultResponse, error) {
	var val *FlowLogListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FlowLogListResultResponse{}, err
	}
	return FlowLogListResultResponse{RawResponse: resp.Response, FlowLogListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *FlowLogsClient) listHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
