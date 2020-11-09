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

// AzureFirewallsOperations contains the methods for the AzureFirewalls group.
type AzureFirewallsOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified Azure Firewall.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsCreateOrUpdateOptions) (*AzureFirewallPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (AzureFirewallPoller, error)
	// BeginDelete - Deletes the specified Azure Firewall.
	BeginDelete(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified Azure Firewall.
	Get(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsGetOptions) (*AzureFirewallResponse, error)
	// List - Lists all Azure Firewalls in a resource group.
	List(resourceGroupName string, options *AzureFirewallsListOptions) AzureFirewallListResultPager
	// ListAll - Gets all the Azure Firewalls in a subscription.
	ListAll(options *AzureFirewallsListAllOptions) AzureFirewallListResultPager
	// BeginUpdateTags - Updates tags of an Azure Firewall resource.
	BeginUpdateTags(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsUpdateTagsOptions) (*AzureFirewallPollerResponse, error)
	// ResumeUpdateTags - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdateTags(token string) (AzureFirewallPoller, error)
}

// AzureFirewallsClient implements the AzureFirewallsOperations interface.
// Don't use this type directly, use NewAzureFirewallsClient() instead.
type AzureFirewallsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAzureFirewallsClient creates a new instance of AzureFirewallsClient with the specified values.
func NewAzureFirewallsClient(con *armcore.Connection, subscriptionID string) AzureFirewallsOperations {
	return &AzureFirewallsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *AzureFirewallsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *AzureFirewallsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsCreateOrUpdateOptions) (*AzureFirewallPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, azureFirewallName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &AzureFirewallPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("AzureFirewallsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &azureFirewallPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*AzureFirewallResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *AzureFirewallsClient) ResumeCreateOrUpdate(token string) (AzureFirewallPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("AzureFirewallsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &azureFirewallPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified Azure Firewall.
func (client *AzureFirewallsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, azureFirewallName, parameters, options)
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
func (client *AzureFirewallsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
func (client *AzureFirewallsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*AzureFirewallResponse, error) {
	result := AzureFirewallResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AzureFirewall)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AzureFirewallsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *AzureFirewallsClient) BeginDelete(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, azureFirewallName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("AzureFirewallsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *AzureFirewallsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("AzureFirewallsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified Azure Firewall.
func (client *AzureFirewallsClient) Delete(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, azureFirewallName, options)
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
func (client *AzureFirewallsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
func (client *AzureFirewallsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified Azure Firewall.
func (client *AzureFirewallsClient) Get(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsGetOptions) (*AzureFirewallResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, azureFirewallName, options)
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
func (client *AzureFirewallsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
func (client *AzureFirewallsClient) GetHandleResponse(resp *azcore.Response) (*AzureFirewallResponse, error) {
	result := AzureFirewallResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AzureFirewall)
}

// GetHandleError handles the Get error response.
func (client *AzureFirewallsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all Azure Firewalls in a resource group.
func (client *AzureFirewallsClient) List(resourceGroupName string, options *AzureFirewallsListOptions) AzureFirewallListResultPager {
	return &azureFirewallListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *AzureFirewallListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureFirewallListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *AzureFirewallsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, options *AzureFirewallsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls"
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

// ListHandleResponse handles the List response.
func (client *AzureFirewallsClient) ListHandleResponse(resp *azcore.Response) (*AzureFirewallListResultResponse, error) {
	result := AzureFirewallListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AzureFirewallListResult)
}

// ListHandleError handles the List error response.
func (client *AzureFirewallsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the Azure Firewalls in a subscription.
func (client *AzureFirewallsClient) ListAll(options *AzureFirewallsListAllOptions) AzureFirewallListResultPager {
	return &azureFirewallListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx, options)
		},
		responder: client.ListAllHandleResponse,
		errorer:   client.ListAllHandleError,
		advancer: func(ctx context.Context, resp *AzureFirewallListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureFirewallListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *AzureFirewallsClient) ListAllCreateRequest(ctx context.Context, options *AzureFirewallsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewalls"
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

// ListAllHandleResponse handles the ListAll response.
func (client *AzureFirewallsClient) ListAllHandleResponse(resp *azcore.Response) (*AzureFirewallListResultResponse, error) {
	result := AzureFirewallListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AzureFirewallListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *AzureFirewallsClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *AzureFirewallsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsUpdateTagsOptions) (*AzureFirewallPollerResponse, error) {
	resp, err := client.UpdateTags(ctx, resourceGroupName, azureFirewallName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &AzureFirewallPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("AzureFirewallsClient.UpdateTags", "azure-async-operation", resp, client.UpdateTagsHandleError)
	if err != nil {
		return nil, err
	}
	poller := &azureFirewallPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*AzureFirewallResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *AzureFirewallsClient) ResumeUpdateTags(token string) (AzureFirewallPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("AzureFirewallsClient.UpdateTags", token, client.UpdateTagsHandleError)
	if err != nil {
		return nil, err
	}
	return &azureFirewallPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// UpdateTags - Updates tags of an Azure Firewall resource.
func (client *AzureFirewallsClient) UpdateTags(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsUpdateTagsOptions) (*azcore.Response, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, azureFirewallName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	return resp, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *AzureFirewallsClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *AzureFirewallsClient) UpdateTagsHandleResponse(resp *azcore.Response) (*AzureFirewallResponse, error) {
	result := AzureFirewallResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AzureFirewall)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *AzureFirewallsClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
