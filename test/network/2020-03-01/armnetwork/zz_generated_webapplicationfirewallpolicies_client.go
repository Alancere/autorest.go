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

// WebApplicationFirewallPoliciesClient contains the methods for the WebApplicationFirewallPolicies group.
// Don't use this type directly, use NewWebApplicationFirewallPoliciesClient() instead.
type WebApplicationFirewallPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWebApplicationFirewallPoliciesClient creates a new instance of WebApplicationFirewallPoliciesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWebApplicationFirewallPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WebApplicationFirewallPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &WebApplicationFirewallPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or update policy with specified rule set name within a resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// policyName - The name of the policy.
// parameters - Policy to be created.
// options - WebApplicationFirewallPoliciesCreateOrUpdateOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.CreateOrUpdate
// method.
func (client *WebApplicationFirewallPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *WebApplicationFirewallPoliciesCreateOrUpdateOptions) (WebApplicationFirewallPoliciesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, policyName, parameters, options)
	if err != nil {
		return WebApplicationFirewallPoliciesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebApplicationFirewallPoliciesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return WebApplicationFirewallPoliciesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WebApplicationFirewallPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *WebApplicationFirewallPoliciesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WebApplicationFirewallPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (WebApplicationFirewallPoliciesCreateOrUpdateResponse, error) {
	result := WebApplicationFirewallPoliciesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicy); err != nil {
		return WebApplicationFirewallPoliciesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *WebApplicationFirewallPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes Policy.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// policyName - The name of the policy.
// options - WebApplicationFirewallPoliciesBeginDeleteOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.BeginDelete
// method.
func (client *WebApplicationFirewallPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesBeginDeleteOptions) (WebApplicationFirewallPoliciesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return WebApplicationFirewallPoliciesDeletePollerResponse{}, err
	}
	result := WebApplicationFirewallPoliciesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WebApplicationFirewallPoliciesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return WebApplicationFirewallPoliciesDeletePollerResponse{}, err
	}
	result.Poller = &WebApplicationFirewallPoliciesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes Policy.
// If the operation fails it returns the *CloudError error type.
func (client *WebApplicationFirewallPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, policyName, options)
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
func (client *WebApplicationFirewallPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
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
func (client *WebApplicationFirewallPoliciesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieve protection policy with specified name within a resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// policyName - The name of the policy.
// options - WebApplicationFirewallPoliciesGetOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.Get
// method.
func (client *WebApplicationFirewallPoliciesClient) Get(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesGetOptions) (WebApplicationFirewallPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return WebApplicationFirewallPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebApplicationFirewallPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebApplicationFirewallPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WebApplicationFirewallPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
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
func (client *WebApplicationFirewallPoliciesClient) getHandleResponse(resp *http.Response) (WebApplicationFirewallPoliciesGetResponse, error) {
	result := WebApplicationFirewallPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicy); err != nil {
		return WebApplicationFirewallPoliciesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WebApplicationFirewallPoliciesClient) getHandleError(resp *http.Response) error {
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

// List - Lists all of the protection policies within a resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// options - WebApplicationFirewallPoliciesListOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.List
// method.
func (client *WebApplicationFirewallPoliciesClient) List(resourceGroupName string, options *WebApplicationFirewallPoliciesListOptions) *WebApplicationFirewallPoliciesListPager {
	return &WebApplicationFirewallPoliciesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp WebApplicationFirewallPoliciesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebApplicationFirewallPolicyListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WebApplicationFirewallPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *WebApplicationFirewallPoliciesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *WebApplicationFirewallPoliciesClient) listHandleResponse(resp *http.Response) (WebApplicationFirewallPoliciesListResponse, error) {
	result := WebApplicationFirewallPoliciesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicyListResult); err != nil {
		return WebApplicationFirewallPoliciesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WebApplicationFirewallPoliciesClient) listHandleError(resp *http.Response) error {
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

// ListAll - Gets all the WAF policies in a subscription.
// If the operation fails it returns the *CloudError error type.
// options - WebApplicationFirewallPoliciesListAllOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.ListAll
// method.
func (client *WebApplicationFirewallPoliciesClient) ListAll(options *WebApplicationFirewallPoliciesListAllOptions) *WebApplicationFirewallPoliciesListAllPager {
	return &WebApplicationFirewallPoliciesListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp WebApplicationFirewallPoliciesListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebApplicationFirewallPolicyListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *WebApplicationFirewallPoliciesClient) listAllCreateRequest(ctx context.Context, options *WebApplicationFirewallPoliciesListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
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

// listAllHandleResponse handles the ListAll response.
func (client *WebApplicationFirewallPoliciesClient) listAllHandleResponse(resp *http.Response) (WebApplicationFirewallPoliciesListAllResponse, error) {
	result := WebApplicationFirewallPoliciesListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicyListResult); err != nil {
		return WebApplicationFirewallPoliciesListAllResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *WebApplicationFirewallPoliciesClient) listAllHandleError(resp *http.Response) error {
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
