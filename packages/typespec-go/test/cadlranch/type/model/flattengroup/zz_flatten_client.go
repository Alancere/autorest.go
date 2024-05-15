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

// FlattenClient - Illustrates the model flatten cases.
// Don't use this type directly, use a constructor function instead.
type FlattenClient struct {
	internal *azcore.Client
}

// PutFlattenModel -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - FlattenClientPutFlattenModelOptions contains the optional parameters for the FlattenClient.PutFlattenModel method.
func (client *FlattenClient) PutFlattenModel(ctx context.Context, input FlattenModel, options *FlattenClientPutFlattenModelOptions) (FlattenClientPutFlattenModelResponse, error) {
	var err error
	const operationName = "FlattenClient.PutFlattenModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putFlattenModelCreateRequest(ctx, input, options)
	if err != nil {
		return FlattenClientPutFlattenModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FlattenClientPutFlattenModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FlattenClientPutFlattenModelResponse{}, err
	}
	resp, err := client.putFlattenModelHandleResponse(httpResp)
	return resp, err
}

// putFlattenModelCreateRequest creates the PutFlattenModel request.
func (client *FlattenClient) putFlattenModelCreateRequest(ctx context.Context, input FlattenModel, _ *FlattenClientPutFlattenModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/flatten/flattenModel"
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
func (client *FlattenClient) putFlattenModelHandleResponse(resp *http.Response) (FlattenClientPutFlattenModelResponse, error) {
	result := FlattenClientPutFlattenModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FlattenModel); err != nil {
		return FlattenClientPutFlattenModelResponse{}, err
	}
	return result, nil
}

// PutNestedFlattenModel -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - FlattenClientPutNestedFlattenModelOptions contains the optional parameters for the FlattenClient.PutNestedFlattenModel
//     method.
func (client *FlattenClient) PutNestedFlattenModel(ctx context.Context, input NestedFlattenModel, options *FlattenClientPutNestedFlattenModelOptions) (FlattenClientPutNestedFlattenModelResponse, error) {
	var err error
	const operationName = "FlattenClient.PutNestedFlattenModel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putNestedFlattenModelCreateRequest(ctx, input, options)
	if err != nil {
		return FlattenClientPutNestedFlattenModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FlattenClientPutNestedFlattenModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FlattenClientPutNestedFlattenModelResponse{}, err
	}
	resp, err := client.putNestedFlattenModelHandleResponse(httpResp)
	return resp, err
}

// putNestedFlattenModelCreateRequest creates the PutNestedFlattenModel request.
func (client *FlattenClient) putNestedFlattenModelCreateRequest(ctx context.Context, input NestedFlattenModel, _ *FlattenClientPutNestedFlattenModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/flatten/nestedFlattenModel"
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
func (client *FlattenClient) putNestedFlattenModelHandleResponse(resp *http.Response) (FlattenClientPutNestedFlattenModelResponse, error) {
	result := FlattenClientPutNestedFlattenModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NestedFlattenModel); err != nil {
		return FlattenClientPutNestedFlattenModelResponse{}, err
	}
	return result, nil
}
