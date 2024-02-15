//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package dictionarygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// Float32ValueClient - Dictionary of float values
// Don't use this type directly, use [DictionaryClient.NewFloat32ValueClient] instead.
type Float32ValueClient struct {
	internal *azcore.Client
}

// - options - Float32ValueClientGetOptions contains the optional parameters for the Float32ValueClient.Get method.
func (client *Float32ValueClient) Get(ctx context.Context, options *Float32ValueClientGetOptions) (Float32ValueClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return Float32ValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Float32ValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Float32ValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *Float32ValueClient) getCreateRequest(ctx context.Context, options *Float32ValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/float32"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Float32ValueClient) getHandleResponse(resp *http.Response) (Float32ValueClientGetResponse, error) {
	result := Float32ValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return Float32ValueClientGetResponse{}, err
	}
	return result, nil
}

// - options - Float32ValueClientPutOptions contains the optional parameters for the Float32ValueClient.Put method.
func (client *Float32ValueClient) Put(ctx context.Context, body map[string]*float32, options *Float32ValueClientPutOptions) (Float32ValueClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return Float32ValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Float32ValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Float32ValueClientPutResponse{}, err
	}
	return Float32ValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *Float32ValueClient) putCreateRequest(ctx context.Context, body map[string]*float32, options *Float32ValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/float32"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
