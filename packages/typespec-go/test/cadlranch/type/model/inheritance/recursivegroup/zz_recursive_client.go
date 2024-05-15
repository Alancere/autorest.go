// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package recursivegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// RecursiveClient - Illustrates inheritance recursion
// Don't use this type directly, use a constructor function instead.
type RecursiveClient struct {
	internal *azcore.Client
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - RecursiveClientGetOptions contains the optional parameters for the RecursiveClient.Get method.
func (client *RecursiveClient) Get(ctx context.Context, options *RecursiveClientGetOptions) (RecursiveClientGetResponse, error) {
	var err error
	const operationName = "RecursiveClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return RecursiveClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecursiveClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecursiveClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RecursiveClient) getCreateRequest(ctx context.Context, _ *RecursiveClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/recursive"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecursiveClient) getHandleResponse(resp *http.Response) (RecursiveClientGetResponse, error) {
	result := RecursiveClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Extension); err != nil {
		return RecursiveClientGetResponse{}, err
	}
	return result, nil
}

// Put -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - RecursiveClientPutOptions contains the optional parameters for the RecursiveClient.Put method.
func (client *RecursiveClient) Put(ctx context.Context, input Extension, options *RecursiveClientPutOptions) (RecursiveClientPutResponse, error) {
	var err error
	const operationName = "RecursiveClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, input, options)
	if err != nil {
		return RecursiveClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecursiveClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RecursiveClientPutResponse{}, err
	}
	return RecursiveClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *RecursiveClient) putCreateRequest(ctx context.Context, input Extension, _ *RecursiveClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/recursive"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}
