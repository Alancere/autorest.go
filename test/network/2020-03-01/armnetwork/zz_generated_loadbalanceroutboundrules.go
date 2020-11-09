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

// LoadBalancerOutboundRulesOperations contains the methods for the LoadBalancerOutboundRules group.
type LoadBalancerOutboundRulesOperations interface {
	// Get - Gets the specified load balancer outbound rule.
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesGetOptions) (*OutboundRuleResponse, error)
	// List - Gets all the outbound rules in a load balancer.
	List(resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesListOptions) LoadBalancerOutboundRuleListResultPager
}

// LoadBalancerOutboundRulesClient implements the LoadBalancerOutboundRulesOperations interface.
// Don't use this type directly, use NewLoadBalancerOutboundRulesClient() instead.
type LoadBalancerOutboundRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerOutboundRulesClient creates a new instance of LoadBalancerOutboundRulesClient with the specified values.
func NewLoadBalancerOutboundRulesClient(con *armcore.Connection, subscriptionID string) LoadBalancerOutboundRulesOperations {
	return &LoadBalancerOutboundRulesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *LoadBalancerOutboundRulesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Gets the specified load balancer outbound rule.
func (client *LoadBalancerOutboundRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesGetOptions) (*OutboundRuleResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, loadBalancerName, outboundRuleName, options)
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
func (client *LoadBalancerOutboundRulesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleName}", url.PathEscape(outboundRuleName))
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
func (client *LoadBalancerOutboundRulesClient) GetHandleResponse(resp *azcore.Response) (*OutboundRuleResponse, error) {
	result := OutboundRuleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.OutboundRule)
}

// GetHandleError handles the Get error response.
func (client *LoadBalancerOutboundRulesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the outbound rules in a load balancer.
func (client *LoadBalancerOutboundRulesClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesListOptions) LoadBalancerOutboundRuleListResultPager {
	return &loadBalancerOutboundRuleListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *LoadBalancerOutboundRuleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerOutboundRuleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *LoadBalancerOutboundRulesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *LoadBalancerOutboundRulesClient) ListHandleResponse(resp *azcore.Response) (*LoadBalancerOutboundRuleListResultResponse, error) {
	result := LoadBalancerOutboundRuleListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancerOutboundRuleListResult)
}

// ListHandleError handles the List error response.
func (client *LoadBalancerOutboundRulesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
