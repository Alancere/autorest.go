//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FirewallPolicyRuleGroupsClient contains the methods for the FirewallPolicyRuleGroups group.
// Don't use this type directly, use NewFirewallPolicyRuleGroupsClient() instead.
type FirewallPolicyRuleGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFirewallPolicyRuleGroupsClient creates a new instance of FirewallPolicyRuleGroupsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFirewallPolicyRuleGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FirewallPolicyRuleGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &FirewallPolicyRuleGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// ruleGroupName - The name of the FirewallPolicyRuleGroup.
// parameters - Parameters supplied to the create or update FirewallPolicyRuleGroup operation.
// options - FirewallPolicyRuleGroupsBeginCreateOrUpdateOptions contains the optional parameters for the FirewallPolicyRuleGroupsClient.BeginCreateOrUpdate
// method.
func (client *FirewallPolicyRuleGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsBeginCreateOrUpdateOptions) (FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, parameters, options)
	if err != nil {
		return FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("FirewallPolicyRuleGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return FirewallPolicyRuleGroupsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &FirewallPolicyRuleGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FirewallPolicyRuleGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup, options *FirewallPolicyRuleGroupsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FirewallPolicyRuleGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// ruleGroupName - The name of the FirewallPolicyRuleGroup.
// options - FirewallPolicyRuleGroupsBeginDeleteOptions contains the optional parameters for the FirewallPolicyRuleGroupsClient.BeginDelete
// method.
func (client *FirewallPolicyRuleGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsBeginDeleteOptions) (FirewallPolicyRuleGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
	if err != nil {
		return FirewallPolicyRuleGroupsDeletePollerResponse{}, err
	}
	result := FirewallPolicyRuleGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("FirewallPolicyRuleGroupsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return FirewallPolicyRuleGroupsDeletePollerResponse{}, err
	}
	result.Poller = &FirewallPolicyRuleGroupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FirewallPolicyRuleGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *FirewallPolicyRuleGroupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified FirewallPolicyRuleGroup.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// ruleGroupName - The name of the FirewallPolicyRuleGroup.
// options - FirewallPolicyRuleGroupsGetOptions contains the optional parameters for the FirewallPolicyRuleGroupsClient.Get
// method.
func (client *FirewallPolicyRuleGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsGetOptions) (FirewallPolicyRuleGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, options)
	if err != nil {
		return FirewallPolicyRuleGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallPolicyRuleGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FirewallPolicyRuleGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FirewallPolicyRuleGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, options *FirewallPolicyRuleGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FirewallPolicyRuleGroupsClient) getHandleResponse(resp *http.Response) (FirewallPolicyRuleGroupsGetResponse, error) {
	result := FirewallPolicyRuleGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicyRuleGroup); err != nil {
		return FirewallPolicyRuleGroupsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *FirewallPolicyRuleGroupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all FirewallPolicyRuleGroups in a FirewallPolicy resource.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// options - FirewallPolicyRuleGroupsListOptions contains the optional parameters for the FirewallPolicyRuleGroupsClient.List
// method.
func (client *FirewallPolicyRuleGroupsClient) List(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleGroupsListOptions) *FirewallPolicyRuleGroupsListPager {
	return &FirewallPolicyRuleGroupsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
		},
		advancer: func(ctx context.Context, resp FirewallPolicyRuleGroupsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyRuleGroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *FirewallPolicyRuleGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleGroupsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FirewallPolicyRuleGroupsClient) listHandleResponse(resp *http.Response) (FirewallPolicyRuleGroupsListResponse, error) {
	result := FirewallPolicyRuleGroupsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicyRuleGroupListResult); err != nil {
		return FirewallPolicyRuleGroupsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *FirewallPolicyRuleGroupsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
