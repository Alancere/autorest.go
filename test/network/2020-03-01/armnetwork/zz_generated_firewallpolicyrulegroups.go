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

// FirewallPolicyRuleGroupsOperations contains the methods for the FirewallPolicyRuleGroups group.
type FirewallPolicyRuleGroupsOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified FirewallPolicyRuleGroup.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsCreateOrUpdateOptions) (*FirewallPolicyRuleGroupPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (FirewallPolicyRuleGroupPoller, error)
	// BeginDelete - Deletes the specified FirewallPolicyRuleGroup.
	BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified FirewallPolicyRuleGroup.
	Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsGetOptions) (*FirewallPolicyRuleGroupResponse, error)
	// List - Lists all FirewallPolicyRuleGroups in a FirewallPolicy resource.
	List(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleGroupsListOptions) FirewallPolicyRuleGroupListResultPager
}

// FirewallPolicyRuleGroupsClient implements the FirewallPolicyRuleGroupsOperations interface.
// Don't use this type directly, use NewFirewallPolicyRuleGroupsClient() instead.
type FirewallPolicyRuleGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFirewallPolicyRuleGroupsClient creates a new instance of FirewallPolicyRuleGroupsClient with the specified values.
func NewFirewallPolicyRuleGroupsClient(con *armcore.Connection, subscriptionID string) FirewallPolicyRuleGroupsOperations {
	return &FirewallPolicyRuleGroupsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *FirewallPolicyRuleGroupsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *FirewallPolicyRuleGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsCreateOrUpdateOptions) (*FirewallPolicyRuleGroupPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &FirewallPolicyRuleGroupPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FirewallPolicyRuleGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &firewallPolicyRuleGroupPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*FirewallPolicyRuleGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *FirewallPolicyRuleGroupsClient) ResumeCreateOrUpdate(token string) (FirewallPolicyRuleGroupPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPolicyRuleGroupsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &firewallPolicyRuleGroupPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified FirewallPolicyRuleGroup.
func (client *FirewallPolicyRuleGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, parameters, options)
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
func (client *FirewallPolicyRuleGroupsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
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

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FirewallPolicyRuleGroupsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleGroupResponse, error) {
	result := FirewallPolicyRuleGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleGroup)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FirewallPolicyRuleGroupsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *FirewallPolicyRuleGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FirewallPolicyRuleGroupsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *FirewallPolicyRuleGroupsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPolicyRuleGroupsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified FirewallPolicyRuleGroup.
func (client *FirewallPolicyRuleGroupsClient) Delete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
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
func (client *FirewallPolicyRuleGroupsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
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

// DeleteHandleError handles the Delete error response.
func (client *FirewallPolicyRuleGroupsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified FirewallPolicyRuleGroup.
func (client *FirewallPolicyRuleGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsGetOptions) (*FirewallPolicyRuleGroupResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
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
func (client *FirewallPolicyRuleGroupsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
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

// GetHandleResponse handles the Get response.
func (client *FirewallPolicyRuleGroupsClient) GetHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleGroupResponse, error) {
	result := FirewallPolicyRuleGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleGroup)
}

// GetHandleError handles the Get error response.
func (client *FirewallPolicyRuleGroupsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all FirewallPolicyRuleGroups in a FirewallPolicy resource.
func (client *FirewallPolicyRuleGroupsClient) List(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleGroupsListOptions) FirewallPolicyRuleGroupListResultPager {
	return &firewallPolicyRuleGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *FirewallPolicyRuleGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyRuleGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *FirewallPolicyRuleGroupsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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

// ListHandleResponse handles the List response.
func (client *FirewallPolicyRuleGroupsClient) ListHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleGroupListResultResponse, error) {
	result := FirewallPolicyRuleGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleGroupListResult)
}

// ListHandleError handles the List error response.
func (client *FirewallPolicyRuleGroupsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
