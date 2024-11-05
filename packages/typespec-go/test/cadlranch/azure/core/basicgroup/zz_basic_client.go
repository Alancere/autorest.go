// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package basicgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BasicClient - Illustrates bodies templated with Azure Core
// Don't use this type directly, use a constructor function instead.
type BasicClient struct {
	internal *azcore.Client
}

// CreateOrReplace - Adds a user or replaces a user's fields.
//
// Creates or replaces a User
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - id - The user's id.
//   - resource - The resource instance.
//   - options - BasicClientCreateOrReplaceOptions contains the optional parameters for the BasicClient.CreateOrReplace method.
func (client *BasicClient) CreateOrReplace(ctx context.Context, id int32, resource User, options *BasicClientCreateOrReplaceOptions) (BasicClientCreateOrReplaceResponse, error) {
	var err error
	const operationName = "BasicClient.CreateOrReplace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrReplaceCreateRequest(ctx, id, resource, options)
	if err != nil {
		return BasicClientCreateOrReplaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientCreateOrReplaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientCreateOrReplaceResponse{}, err
	}
	resp, err := client.createOrReplaceHandleResponse(httpResp)
	return resp, err
}

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *BasicClient) createOrReplaceCreateRequest(ctx context.Context, id int32, resource User, _ *BasicClientCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrReplaceHandleResponse handles the CreateOrReplace response.
func (client *BasicClient) createOrReplaceHandleResponse(resp *http.Response) (BasicClientCreateOrReplaceResponse, error) {
	result := BasicClientCreateOrReplaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return BasicClientCreateOrReplaceResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Adds a user or updates a user's fields.
//
// Creates or updates a User
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - id - The user's id.
//   - resource - The resource instance.
//   - options - BasicClientCreateOrUpdateOptions contains the optional parameters for the BasicClient.CreateOrUpdate method.
func (client *BasicClient) CreateOrUpdate(ctx context.Context, id int32, resource User, options *BasicClientCreateOrUpdateOptions) (BasicClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "BasicClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, id, resource, options)
	if err != nil {
		return BasicClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BasicClient) createOrUpdateCreateRequest(ctx context.Context, id int32, resource User, _ *BasicClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/merge-patch+json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BasicClient) createOrUpdateHandleResponse(resp *http.Response) (BasicClientCreateOrUpdateResponse, error) {
	result := BasicClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return BasicClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a user.
//
// Deletes a User
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - id - The user's id.
//   - options - BasicClientDeleteOptions contains the optional parameters for the BasicClient.Delete method.
func (client *BasicClient) Delete(ctx context.Context, id int32, options *BasicClientDeleteOptions) (BasicClientDeleteResponse, error) {
	var err error
	const operationName = "BasicClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, id, options)
	if err != nil {
		return BasicClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientDeleteResponse{}, err
	}
	return BasicClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BasicClient) deleteCreateRequest(ctx context.Context, id int32, _ *BasicClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Export - Exports a user.
//
// Exports a User
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - id - The user's id.
//   - formatParam - The format of the data.
//   - options - BasicClientExportOptions contains the optional parameters for the BasicClient.Export method.
func (client *BasicClient) Export(ctx context.Context, id int32, formatParam string, options *BasicClientExportOptions) (BasicClientExportResponse, error) {
	var err error
	const operationName = "BasicClient.Export"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportCreateRequest(ctx, id, formatParam, options)
	if err != nil {
		return BasicClientExportResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientExportResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientExportResponse{}, err
	}
	resp, err := client.exportHandleResponse(httpResp)
	return resp, err
}

// exportCreateRequest creates the Export request.
func (client *BasicClient) exportCreateRequest(ctx context.Context, id int32, formatParam string, _ *BasicClientExportOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}:export"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	reqQP.Set("format", formatParam)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// exportHandleResponse handles the Export response.
func (client *BasicClient) exportHandleResponse(resp *http.Response) (BasicClientExportResponse, error) {
	result := BasicClientExportResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return BasicClientExportResponse{}, err
	}
	return result, nil
}

// ExportAllUsers - Exports all users.
//
// Exports all users
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - formatParam - The format of the data.
//   - options - BasicClientExportAllUsersOptions contains the optional parameters for the BasicClient.ExportAllUsers method.
func (client *BasicClient) ExportAllUsers(ctx context.Context, formatParam string, options *BasicClientExportAllUsersOptions) (BasicClientExportAllUsersResponse, error) {
	var err error
	const operationName = "BasicClient.ExportAllUsers"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportAllUsersCreateRequest(ctx, formatParam, options)
	if err != nil {
		return BasicClientExportAllUsersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientExportAllUsersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientExportAllUsersResponse{}, err
	}
	resp, err := client.exportAllUsersHandleResponse(httpResp)
	return resp, err
}

// exportAllUsersCreateRequest creates the ExportAllUsers request.
func (client *BasicClient) exportAllUsersCreateRequest(ctx context.Context, formatParam string, _ *BasicClientExportAllUsersOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users:exportallusers"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	reqQP.Set("format", formatParam)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// exportAllUsersHandleResponse handles the ExportAllUsers response.
func (client *BasicClient) exportAllUsersHandleResponse(resp *http.Response) (BasicClientExportAllUsersResponse, error) {
	result := BasicClientExportAllUsersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserList); err != nil {
		return BasicClientExportAllUsersResponse{}, err
	}
	return result, nil
}

// Get - Gets a user.
//
// Gets a User
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - id - The user's id.
//   - options - BasicClientGetOptions contains the optional parameters for the BasicClient.Get method.
func (client *BasicClient) Get(ctx context.Context, id int32, options *BasicClientGetOptions) (BasicClientGetResponse, error) {
	var err error
	const operationName = "BasicClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, id, options)
	if err != nil {
		return BasicClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BasicClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BasicClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BasicClient) getCreateRequest(ctx context.Context, id int32, _ *BasicClientGetOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users/{id}"
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BasicClient) getHandleResponse(resp *http.Response) (BasicClientGetResponse, error) {
	result := BasicClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return BasicClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all users.
//
// # Lists all Users
//
// Generated from API version 2022-12-01-preview
//   - options - BasicClientListOptions contains the optional parameters for the BasicClient.NewListPager method.
func (client *BasicClient) NewListPager(options *BasicClientListOptions) *runtime.Pager[BasicClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BasicClientListResponse]{
		More: func(page BasicClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BasicClientListResponse) (BasicClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BasicClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return BasicClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BasicClient) listCreateRequest(ctx context.Context, options *BasicClientListOptions) (*policy.Request, error) {
	urlPath := "/azure/core/basic/users"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	if options != nil && options.Expand != nil {
		for _, qv := range options.Expand {
			reqQP.Add("expand", qv)
		}
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	if options != nil && options.Orderby != nil {
		for _, qv := range options.Orderby {
			reqQP.Add("orderby", qv)
		}
	}
	if options != nil && options.SelectParam != nil {
		for _, qv := range options.SelectParam {
			reqQP.Add("select", qv)
		}
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BasicClient) listHandleResponse(resp *http.Response) (BasicClientListResponse, error) {
	result := BasicClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedUser); err != nil {
		return BasicClientListResponse{}, err
	}
	return result, nil
}
