// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package flattengroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FlattenPropertyClient - Illustrates the model flatten cases.
// Don't use this type directly, use a constructor function instead.
type FlattenPropertyClient struct {
	internal *azcore.Client
}

// PutFlattenModel -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - FlattenPropertyClientPutFlattenModelOptions contains the optional parameters for the FlattenPropertyClient.PutFlattenModel
//     method.
func (client *FlattenPropertyClient) PutFlattenModel(ctx context.Context, input FlattenModel, options *FlattenPropertyClientPutFlattenModelOptions) (FlattenPropertyClientPutFlattenModelResponse, error) {
	var err error
	const operationName = "FlattenPropertyClient.PutFlattenModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putFlattenModelCreateRequest(ctx, input, options)
	if err != nil {
		return FlattenPropertyClientPutFlattenModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FlattenPropertyClientPutFlattenModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FlattenPropertyClientPutFlattenModelResponse{}, err
	}
	resp, err := client.putFlattenModelHandleResponse(httpResp)
	return resp, err
}

// putFlattenModelCreateRequest creates the PutFlattenModel request.
func (client *FlattenPropertyClient) putFlattenModelCreateRequest(ctx context.Context, input FlattenModel, _ *FlattenPropertyClientPutFlattenModelOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/flatten-property/flattenModel"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// putFlattenModelHandleResponse handles the PutFlattenModel response.
func (client *FlattenPropertyClient) putFlattenModelHandleResponse(resp *http.Response) (FlattenPropertyClientPutFlattenModelResponse, error) {
	result := FlattenPropertyClientPutFlattenModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FlattenModel); err != nil {
		return FlattenPropertyClientPutFlattenModelResponse{}, err
	}
	return result, nil
}

// PutNestedFlattenModel -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - FlattenPropertyClientPutNestedFlattenModelOptions contains the optional parameters for the FlattenPropertyClient.PutNestedFlattenModel
//     method.
func (client *FlattenPropertyClient) PutNestedFlattenModel(ctx context.Context, input NestedFlattenModel, options *FlattenPropertyClientPutNestedFlattenModelOptions) (FlattenPropertyClientPutNestedFlattenModelResponse, error) {
	var err error
	const operationName = "FlattenPropertyClient.PutNestedFlattenModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putNestedFlattenModelCreateRequest(ctx, input, options)
	if err != nil {
		return FlattenPropertyClientPutNestedFlattenModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FlattenPropertyClientPutNestedFlattenModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FlattenPropertyClientPutNestedFlattenModelResponse{}, err
	}
	resp, err := client.putNestedFlattenModelHandleResponse(httpResp)
	return resp, err
}

// putNestedFlattenModelCreateRequest creates the PutNestedFlattenModel request.
func (client *FlattenPropertyClient) putNestedFlattenModelCreateRequest(ctx context.Context, input NestedFlattenModel, _ *FlattenPropertyClientPutNestedFlattenModelOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/flatten-property/nestedFlattenModel"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// putNestedFlattenModelHandleResponse handles the PutNestedFlattenModel response.
func (client *FlattenPropertyClient) putNestedFlattenModelHandleResponse(resp *http.Response) (FlattenPropertyClientPutNestedFlattenModelResponse, error) {
	result := FlattenPropertyClientPutNestedFlattenModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NestedFlattenModel); err != nil {
		return FlattenPropertyClientPutNestedFlattenModelResponse{}, err
	}
	return result, nil
}