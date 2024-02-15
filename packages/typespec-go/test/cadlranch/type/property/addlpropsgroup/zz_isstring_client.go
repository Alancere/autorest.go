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

// IsStringClient contains the methods for the Type.Property.AdditionalProperties namespace.
// Don't use this type directly, use [AdditionalPropertiesClient.NewIsStringClient] instead.
type IsStringClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - IsStringClientGetOptions contains the optional parameters for the IsStringClient.Get method.
func (client *IsStringClient) Get(ctx context.Context, options *IsStringClientGetOptions) (IsStringClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return IsStringClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IsStringClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IsStringClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IsStringClient) getCreateRequest(ctx context.Context, options *IsStringClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/isRecordstring"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IsStringClient) getHandleResponse(resp *http.Response) (IsStringClientGetResponse, error) {
	result := IsStringClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IsStringAdditionalProperties); err != nil {
		return IsStringClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - IsStringClientPutOptions contains the optional parameters for the IsStringClient.Put method.
func (client *IsStringClient) Put(ctx context.Context, body IsStringAdditionalProperties, options *IsStringClientPutOptions) (IsStringClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return IsStringClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IsStringClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IsStringClientPutResponse{}, err
	}
	return IsStringClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *IsStringClient) putCreateRequest(ctx context.Context, body IsStringAdditionalProperties, options *IsStringClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/isRecordstring"
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
