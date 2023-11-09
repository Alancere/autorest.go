//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DscpConfigurationClient contains the methods for the DscpConfiguration group.
// Don't use this type directly, use NewDscpConfigurationClient() instead.
type DscpConfigurationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDscpConfigurationClient creates a new instance of DscpConfigurationClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDscpConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscpConfigurationClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DscpConfigurationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - dscpConfigurationName - The name of the resource.
//   - parameters - Parameters supplied to the create or update dscp configuration operation.
//   - options - DscpConfigurationClientBeginCreateOrUpdateOptions contains the optional parameters for the DscpConfigurationClient.BeginCreateOrUpdate
//     method.
func (client *DscpConfigurationClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, dscpConfigurationName string, parameters DscpConfiguration, options *DscpConfigurationClientBeginCreateOrUpdateOptions) (*runtime.Poller[DscpConfigurationClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, dscpConfigurationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DscpConfigurationClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DscpConfigurationClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *DscpConfigurationClient) createOrUpdate(ctx context.Context, resourceGroupName string, dscpConfigurationName string, parameters DscpConfiguration, options *DscpConfigurationClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DscpConfigurationClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dscpConfigurationName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DscpConfigurationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dscpConfigurationName string, parameters DscpConfiguration, options *DscpConfigurationClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dscpConfigurationName == "" {
		return nil, errors.New("parameter dscpConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dscpConfigurationName}", url.PathEscape(dscpConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - dscpConfigurationName - The name of the resource.
//   - options - DscpConfigurationClientBeginDeleteOptions contains the optional parameters for the DscpConfigurationClient.BeginDelete
//     method.
func (client *DscpConfigurationClient) BeginDelete(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientBeginDeleteOptions) (*runtime.Poller[DscpConfigurationClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, dscpConfigurationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DscpConfigurationClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DscpConfigurationClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *DscpConfigurationClient) deleteOperation(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DscpConfigurationClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dscpConfigurationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DscpConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dscpConfigurationName == "" {
		return nil, errors.New("parameter dscpConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dscpConfigurationName}", url.PathEscape(dscpConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - dscpConfigurationName - The name of the resource.
//   - options - DscpConfigurationClientGetOptions contains the optional parameters for the DscpConfigurationClient.Get method.
func (client *DscpConfigurationClient) Get(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientGetOptions) (DscpConfigurationClientGetResponse, error) {
	var err error
	const operationName = "DscpConfigurationClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, dscpConfigurationName, options)
	if err != nil {
		return DscpConfigurationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscpConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DscpConfigurationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DscpConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dscpConfigurationName == "" {
		return nil, errors.New("parameter dscpConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dscpConfigurationName}", url.PathEscape(dscpConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscpConfigurationClient) getHandleResponse(resp *http.Response) (DscpConfigurationClientGetResponse, error) {
	result := DscpConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscpConfiguration); err != nil {
		return DscpConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a DSCP Configuration.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - options - DscpConfigurationClientListOptions contains the optional parameters for the DscpConfigurationClient.NewListPager
//     method.
func (client *DscpConfigurationClient) NewListPager(resourceGroupName string, options *DscpConfigurationClientListOptions) *runtime.Pager[DscpConfigurationClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscpConfigurationClientListResponse]{
		More: func(page DscpConfigurationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscpConfigurationClientListResponse) (DscpConfigurationClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DscpConfigurationClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DscpConfigurationClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DscpConfigurationClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *DscpConfigurationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DscpConfigurationClient) listHandleResponse(resp *http.Response) (DscpConfigurationClientListResponse, error) {
	result := DscpConfigurationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscpConfigurationListResult); err != nil {
		return DscpConfigurationClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all dscp configurations in a subscription.
//
// Generated from API version 2022-09-01
//   - options - DscpConfigurationClientListAllOptions contains the optional parameters for the DscpConfigurationClient.NewListAllPager
//     method.
func (client *DscpConfigurationClient) NewListAllPager(options *DscpConfigurationClientListAllOptions) *runtime.Pager[DscpConfigurationClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscpConfigurationClientListAllResponse]{
		More: func(page DscpConfigurationClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscpConfigurationClientListAllResponse) (DscpConfigurationClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DscpConfigurationClient.NewListAllPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DscpConfigurationClientListAllResponse{}, err
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *DscpConfigurationClient) listAllCreateRequest(ctx context.Context, options *DscpConfigurationClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dscpConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *DscpConfigurationClient) listAllHandleResponse(resp *http.Response) (DscpConfigurationClientListAllResponse, error) {
	result := DscpConfigurationClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscpConfigurationListResult); err != nil {
		return DscpConfigurationClientListAllResponse{}, err
	}
	return result, nil
}
