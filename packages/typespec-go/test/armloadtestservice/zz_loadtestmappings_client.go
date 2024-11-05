// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armloadtestservice

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

// LoadTestMappingsClient contains the methods for the LoadTestMappings group.
// Don't use this type directly, use NewLoadTestMappingsClient() instead.
type LoadTestMappingsClient struct {
	internal *arm.Client
}

// NewLoadTestMappingsClient creates a new instance of LoadTestMappingsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLoadTestMappingsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadTestMappingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LoadTestMappingsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a LoadTestMappingResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - loadTestMappingName - Load Test Mapping name
//   - resource - Resource create parameters.
//   - options - LoadTestMappingsClientCreateOrUpdateOptions contains the optional parameters for the LoadTestMappingsClient.CreateOrUpdate
//     method.
func (client *LoadTestMappingsClient) CreateOrUpdate(ctx context.Context, resourceURI string, loadTestMappingName string, resource LoadTestMappingResource, options *LoadTestMappingsClientCreateOrUpdateOptions) (LoadTestMappingsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "LoadTestMappingsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, loadTestMappingName, resource, options)
	if err != nil {
		return LoadTestMappingsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadTestMappingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return LoadTestMappingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LoadTestMappingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, loadTestMappingName string, resource LoadTestMappingResource, _ *LoadTestMappingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.LoadTestService/loadTestMappings/{loadTestMappingName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if loadTestMappingName == "" {
		return nil, errors.New("parameter loadTestMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestMappingName}", url.PathEscape(loadTestMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LoadTestMappingsClient) createOrUpdateHandleResponse(resp *http.Response) (LoadTestMappingsClientCreateOrUpdateResponse, error) {
	result := LoadTestMappingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestMappingResource); err != nil {
		return LoadTestMappingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a LoadTestMappingResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - loadTestMappingName - Load Test Mapping name
//   - options - LoadTestMappingsClientDeleteOptions contains the optional parameters for the LoadTestMappingsClient.Delete method.
func (client *LoadTestMappingsClient) Delete(ctx context.Context, resourceURI string, loadTestMappingName string, options *LoadTestMappingsClientDeleteOptions) (LoadTestMappingsClientDeleteResponse, error) {
	var err error
	const operationName = "LoadTestMappingsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, loadTestMappingName, options)
	if err != nil {
		return LoadTestMappingsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadTestMappingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return LoadTestMappingsClientDeleteResponse{}, err
	}
	return LoadTestMappingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LoadTestMappingsClient) deleteCreateRequest(ctx context.Context, resourceURI string, loadTestMappingName string, _ *LoadTestMappingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.LoadTestService/loadTestMappings/{loadTestMappingName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if loadTestMappingName == "" {
		return nil, errors.New("parameter loadTestMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestMappingName}", url.PathEscape(loadTestMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a LoadTestMappingResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - loadTestMappingName - Load Test Mapping name
//   - options - LoadTestMappingsClientGetOptions contains the optional parameters for the LoadTestMappingsClient.Get method.
func (client *LoadTestMappingsClient) Get(ctx context.Context, resourceURI string, loadTestMappingName string, options *LoadTestMappingsClientGetOptions) (LoadTestMappingsClientGetResponse, error) {
	var err error
	const operationName = "LoadTestMappingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, loadTestMappingName, options)
	if err != nil {
		return LoadTestMappingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadTestMappingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LoadTestMappingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LoadTestMappingsClient) getCreateRequest(ctx context.Context, resourceURI string, loadTestMappingName string, _ *LoadTestMappingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.LoadTestService/loadTestMappings/{loadTestMappingName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if loadTestMappingName == "" {
		return nil, errors.New("parameter loadTestMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestMappingName}", url.PathEscape(loadTestMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadTestMappingsClient) getHandleResponse(resp *http.Response) (LoadTestMappingsClientGetResponse, error) {
	result := LoadTestMappingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestMappingResource); err != nil {
		return LoadTestMappingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List LoadTestMappingResource resources by parent
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - LoadTestMappingsClientListOptions contains the optional parameters for the LoadTestMappingsClient.NewListPager
//     method.
func (client *LoadTestMappingsClient) NewListPager(resourceURI string, options *LoadTestMappingsClientListOptions) *runtime.Pager[LoadTestMappingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadTestMappingsClientListResponse]{
		More: func(page LoadTestMappingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadTestMappingsClientListResponse) (LoadTestMappingsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LoadTestMappingsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return LoadTestMappingsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LoadTestMappingsClient) listCreateRequest(ctx context.Context, resourceURI string, _ *LoadTestMappingsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.LoadTestService/loadTestMappings"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadTestMappingsClient) listHandleResponse(resp *http.Response) (LoadTestMappingsClientListResponse, error) {
	result := LoadTestMappingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestMappingResourceListResult); err != nil {
		return LoadTestMappingsClientListResponse{}, err
	}
	return result, nil
}

// Update - Update a LoadTestMappingResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - loadTestMappingName - Load Test Mapping name
//   - properties - The resource properties to be updated.
//   - options - LoadTestMappingsClientUpdateOptions contains the optional parameters for the LoadTestMappingsClient.Update method.
func (client *LoadTestMappingsClient) Update(ctx context.Context, resourceURI string, loadTestMappingName string, properties LoadTestMappingResourceUpdate, options *LoadTestMappingsClientUpdateOptions) (LoadTestMappingsClientUpdateResponse, error) {
	var err error
	const operationName = "LoadTestMappingsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceURI, loadTestMappingName, properties, options)
	if err != nil {
		return LoadTestMappingsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadTestMappingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LoadTestMappingsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *LoadTestMappingsClient) updateCreateRequest(ctx context.Context, resourceURI string, loadTestMappingName string, properties LoadTestMappingResourceUpdate, _ *LoadTestMappingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.LoadTestService/loadTestMappings/{loadTestMappingName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if loadTestMappingName == "" {
		return nil, errors.New("parameter loadTestMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestMappingName}", url.PathEscape(loadTestMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *LoadTestMappingsClient) updateHandleResponse(resp *http.Response) (LoadTestMappingsClientUpdateResponse, error) {
	result := LoadTestMappingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestMappingResource); err != nil {
		return LoadTestMappingsClientUpdateResponse{}, err
	}
	return result, nil
}
