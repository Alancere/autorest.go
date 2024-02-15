//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// IsModelClient contains the methods for the Type.Property.AdditionalProperties namespace.
// Don't use this type directly, use [AdditionalPropertiesClient.NewIsModelClient] instead.
type IsModelClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - IsModelClientGetOptions contains the optional parameters for the IsModelClient.Get method.
func (client *IsModelClient) Get(ctx context.Context, options *IsModelClientGetOptions) (IsModelClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return IsModelClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IsModelClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IsModelClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IsModelClient) getCreateRequest(ctx context.Context, options *IsModelClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/isRecordModel"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IsModelClient) getHandleResponse(resp *http.Response) (IsModelClientGetResponse, error) {
	result := IsModelClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IsModelAdditionalProperties); err != nil {
		return IsModelClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - IsModelClientPutOptions contains the optional parameters for the IsModelClient.Put method.
func (client *IsModelClient) Put(ctx context.Context, body IsModelAdditionalProperties, options *IsModelClientPutOptions) (IsModelClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return IsModelClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IsModelClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IsModelClientPutResponse{}, err
	}
	return IsModelClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *IsModelClient) putCreateRequest(ctx context.Context, body IsModelAdditionalProperties, options *IsModelClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/isRecordModel"
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
