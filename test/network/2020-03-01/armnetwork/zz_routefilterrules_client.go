//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RouteFilterRulesClient contains the methods for the RouteFilterRules group.
// Don't use this type directly, use NewRouteFilterRulesClient() instead.
type RouteFilterRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRouteFilterRulesClient creates a new instance of RouteFilterRulesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRouteFilterRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RouteFilterRulesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RouteFilterRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a route in the specified route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - routeFilterName - The name of the route filter.
//   - ruleName - The name of the route filter rule.
//   - routeFilterRuleParameters - Parameters supplied to the create or update route filter rule operation.
//   - options - RouteFilterRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the RouteFilterRulesClient.BeginCreateOrUpdate
//     method.
func (client *RouteFilterRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule, options *RouteFilterRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[RouteFilterRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, routeFilterName, ruleName, routeFilterRuleParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[RouteFilterRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RouteFilterRulesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a route in the specified route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *RouteFilterRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule, options *RouteFilterRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeFilterName, ruleName, routeFilterRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteFilterRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule, options *RouteFilterRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, routeFilterRuleParameters)
}

// BeginDelete - Deletes the specified rule from a route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - routeFilterName - The name of the route filter.
//   - ruleName - The name of the rule.
//   - options - RouteFilterRulesClientBeginDeleteOptions contains the optional parameters for the RouteFilterRulesClient.BeginDelete
//     method.
func (client *RouteFilterRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesClientBeginDeleteOptions) (*runtime.Poller[RouteFilterRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, routeFilterName, ruleName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[RouteFilterRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RouteFilterRulesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified rule from a route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *RouteFilterRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeFilterName, ruleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RouteFilterRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified rule from a route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - routeFilterName - The name of the route filter.
//   - ruleName - The name of the rule.
//   - options - RouteFilterRulesClientGetOptions contains the optional parameters for the RouteFilterRulesClient.Get method.
func (client *RouteFilterRulesClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesClientGetOptions) (RouteFilterRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeFilterName, ruleName, options)
	if err != nil {
		return RouteFilterRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RouteFilterRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RouteFilterRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RouteFilterRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, options *RouteFilterRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RouteFilterRulesClient) getHandleResponse(resp *http.Response) (RouteFilterRulesClientGetResponse, error) {
	result := RouteFilterRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilterRule); err != nil {
		return RouteFilterRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRouteFilterPager - Gets all RouteFilterRules in a route filter.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - routeFilterName - The name of the route filter.
//   - options - RouteFilterRulesClientListByRouteFilterOptions contains the optional parameters for the RouteFilterRulesClient.NewListByRouteFilterPager
//     method.
func (client *RouteFilterRulesClient) NewListByRouteFilterPager(resourceGroupName string, routeFilterName string, options *RouteFilterRulesClientListByRouteFilterOptions) *runtime.Pager[RouteFilterRulesClientListByRouteFilterResponse] {
	return runtime.NewPager(runtime.PagingHandler[RouteFilterRulesClientListByRouteFilterResponse]{
		More: func(page RouteFilterRulesClientListByRouteFilterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RouteFilterRulesClientListByRouteFilterResponse) (RouteFilterRulesClientListByRouteFilterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRouteFilterCreateRequest(ctx, resourceGroupName, routeFilterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RouteFilterRulesClientListByRouteFilterResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RouteFilterRulesClientListByRouteFilterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RouteFilterRulesClientListByRouteFilterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRouteFilterHandleResponse(resp)
		},
	})
}

// listByRouteFilterCreateRequest creates the ListByRouteFilter request.
func (client *RouteFilterRulesClient) listByRouteFilterCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFilterRulesClientListByRouteFilterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRouteFilterHandleResponse handles the ListByRouteFilter response.
func (client *RouteFilterRulesClient) listByRouteFilterHandleResponse(resp *http.Response) (RouteFilterRulesClientListByRouteFilterResponse, error) {
	result := RouteFilterRulesClientListByRouteFilterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilterRuleListResult); err != nil {
		return RouteFilterRulesClientListByRouteFilterResponse{}, err
	}
	return result, nil
}
