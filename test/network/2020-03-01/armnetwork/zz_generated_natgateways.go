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

// NatGatewaysClient contains the methods for the NatGateways group.
// Don't use this type directly, use NewNatGatewaysClient() instead.
type NatGatewaysClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNatGatewaysClient creates a new instance of NatGatewaysClient with the specified values.
func NewNatGatewaysClient(con *armcore.Connection, subscriptionID string) *NatGatewaysClient {
	return &NatGatewaysClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a nat gateway.
func (client *NatGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysBeginCreateOrUpdateOptions) (NatGatewayPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, natGatewayName, parameters, options)
	if err != nil {
		return NatGatewayPollerResponse{}, err
	}
	result := NatGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NatGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return NatGatewayPollerResponse{}, err
	}
	poller := &natGatewayPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NatGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new NatGatewayPoller from the specified resume token.
// token - The value must come from a previous call to NatGatewayPoller.ResumeToken().
func (client *NatGatewaysClient) ResumeCreateOrUpdate(token string) (NatGatewayPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NatGatewaysClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &natGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a nat gateway.
func (client *NatGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, natGatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NatGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
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
func (client *NatGatewaysClient) createOrUpdateHandleResponse(resp *azcore.Response) (NatGatewayResponse, error) {
	var val *NatGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NatGatewayResponse{}, err
	}
	return NatGatewayResponse{RawResponse: resp.Response, NatGateway: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NatGatewaysClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified nat gateway.
func (client *NatGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, natGatewayName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NatGatewaysClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *NatGatewaysClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NatGatewaysClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified nat gateway.
func (client *NatGatewaysClient) delete(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, natGatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NatGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
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
func (client *NatGatewaysClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified nat gateway in a specified resource group.
func (client *NatGatewaysClient) Get(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysGetOptions) (NatGatewayResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, natGatewayName, options)
	if err != nil {
		return NatGatewayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NatGatewayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NatGatewayResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NatGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
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
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NatGatewaysClient) getHandleResponse(resp *azcore.Response) (NatGatewayResponse, error) {
	var val *NatGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NatGatewayResponse{}, err
	}
	return NatGatewayResponse{RawResponse: resp.Response, NatGateway: val}, nil
}

// getHandleError handles the Get error response.
func (client *NatGatewaysClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all nat gateways in a resource group.
func (client *NatGatewaysClient) List(resourceGroupName string, options *NatGatewaysListOptions) NatGatewayListResultPager {
	return &natGatewayListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp NatGatewayListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NatGatewayListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *NatGatewaysClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *NatGatewaysListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *NatGatewaysClient) listHandleResponse(resp *azcore.Response) (NatGatewayListResultResponse, error) {
	var val *NatGatewayListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NatGatewayListResultResponse{}, err
	}
	return NatGatewayListResultResponse{RawResponse: resp.Response, NatGatewayListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *NatGatewaysClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the Nat Gateways in a subscription.
func (client *NatGatewaysClient) ListAll(options *NatGatewaysListAllOptions) NatGatewayListResultPager {
	return &natGatewayListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp NatGatewayListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NatGatewayListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *NatGatewaysClient) listAllCreateRequest(ctx context.Context, options *NatGatewaysListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways"
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

// listAllHandleResponse handles the ListAll response.
func (client *NatGatewaysClient) listAllHandleResponse(resp *azcore.Response) (NatGatewayListResultResponse, error) {
	var val *NatGatewayListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NatGatewayListResultResponse{}, err
	}
	return NatGatewayListResultResponse{RawResponse: resp.Response, NatGatewayListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *NatGatewaysClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates nat gateway tags.
func (client *NatGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysUpdateTagsOptions) (NatGatewayResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, natGatewayName, parameters, options)
	if err != nil {
		return NatGatewayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NatGatewayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NatGatewayResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *NatGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// updateTagsHandleResponse handles the UpdateTags response.
func (client *NatGatewaysClient) updateTagsHandleResponse(resp *azcore.Response) (NatGatewayResponse, error) {
	var val *NatGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NatGatewayResponse{}, err
	}
	return NatGatewayResponse{RawResponse: resp.Response, NatGateway: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *NatGatewaysClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
