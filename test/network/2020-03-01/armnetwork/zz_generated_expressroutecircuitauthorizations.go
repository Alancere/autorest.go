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

// ExpressRouteCircuitAuthorizationsOperations contains the methods for the ExpressRouteCircuitAuthorizations group.
type ExpressRouteCircuitAuthorizationsOperations interface {
	// BeginCreateOrUpdate - Creates or updates an authorization in the specified express route circuit.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsCreateOrUpdateOptions) (*ExpressRouteCircuitAuthorizationPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRouteCircuitAuthorizationPoller, error)
	// BeginDelete - Deletes the specified authorization from the specified express route circuit.
	BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified authorization from the specified express route circuit.
	Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsGetOptions) (*ExpressRouteCircuitAuthorizationResponse, error)
	// List - Gets all authorizations in an express route circuit.
	List(resourceGroupName string, circuitName string, options *ExpressRouteCircuitAuthorizationsListOptions) AuthorizationListResultPager
}

// ExpressRouteCircuitAuthorizationsClient implements the ExpressRouteCircuitAuthorizationsOperations interface.
// Don't use this type directly, use NewExpressRouteCircuitAuthorizationsClient() instead.
type ExpressRouteCircuitAuthorizationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCircuitAuthorizationsClient creates a new instance of ExpressRouteCircuitAuthorizationsClient with the specified values.
func NewExpressRouteCircuitAuthorizationsClient(con *armcore.Connection, subscriptionID string) ExpressRouteCircuitAuthorizationsOperations {
	return &ExpressRouteCircuitAuthorizationsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *ExpressRouteCircuitAuthorizationsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *ExpressRouteCircuitAuthorizationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsCreateOrUpdateOptions) (*ExpressRouteCircuitAuthorizationPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, circuitName, authorizationName, authorizationParameters, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitAuthorizationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitAuthorizationsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitAuthorizationPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitAuthorizationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteCircuitAuthorizationsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCircuitAuthorizationPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitAuthorizationsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitAuthorizationPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates an authorization in the specified express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, authorizationParameters, options)
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
func (client *ExpressRouteCircuitAuthorizationsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(authorizationParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExpressRouteCircuitAuthorizationsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitAuthorizationResponse, error) {
	result := ExpressRouteCircuitAuthorizationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitAuthorization)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCircuitAuthorizationsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ExpressRouteCircuitAuthorizationsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, circuitName, authorizationName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitAuthorizationsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *ExpressRouteCircuitAuthorizationsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitAuthorizationsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified authorization from the specified express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, options)
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
func (client *ExpressRouteCircuitAuthorizationsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
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
func (client *ExpressRouteCircuitAuthorizationsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified authorization from the specified express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsGetOptions) (*ExpressRouteCircuitAuthorizationResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, options)
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
func (client *ExpressRouteCircuitAuthorizationsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
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
func (client *ExpressRouteCircuitAuthorizationsClient) GetHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitAuthorizationResponse, error) {
	result := ExpressRouteCircuitAuthorizationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitAuthorization)
}

// GetHandleError handles the Get error response.
func (client *ExpressRouteCircuitAuthorizationsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all authorizations in an express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) List(resourceGroupName string, circuitName string, options *ExpressRouteCircuitAuthorizationsListOptions) AuthorizationListResultPager {
	return &authorizationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, circuitName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *AuthorizationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AuthorizationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *ExpressRouteCircuitAuthorizationsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitAuthorizationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *ExpressRouteCircuitAuthorizationsClient) ListHandleResponse(resp *azcore.Response) (*AuthorizationListResultResponse, error) {
	result := AuthorizationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AuthorizationListResult)
}

// ListHandleError handles the List error response.
func (client *ExpressRouteCircuitAuthorizationsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
