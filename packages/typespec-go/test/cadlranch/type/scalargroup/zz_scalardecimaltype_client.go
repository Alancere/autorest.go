// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package scalargroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// ScalarDecimalTypeClient - Decimal type
// Don't use this type directly, use [ScalarClient.NewScalarDecimalTypeClient] instead.
type ScalarDecimalTypeClient struct {
	internal *azcore.Client
}

// RequestBody -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ScalarDecimalTypeClientRequestBodyOptions contains the optional parameters for the ScalarDecimalTypeClient.RequestBody
//     method.
func (client *ScalarDecimalTypeClient) RequestBody(ctx context.Context, body float64, options *ScalarDecimalTypeClientRequestBodyOptions) (ScalarDecimalTypeClientRequestBodyResponse, error) {
	var err error
	const operationName = "ScalarDecimalTypeClient.RequestBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.requestBodyCreateRequest(ctx, body, options)
	if err != nil {
		return ScalarDecimalTypeClientRequestBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarDecimalTypeClientRequestBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScalarDecimalTypeClientRequestBodyResponse{}, err
	}
	return ScalarDecimalTypeClientRequestBodyResponse{}, nil
}

// requestBodyCreateRequest creates the RequestBody request.
func (client *ScalarDecimalTypeClient) requestBodyCreateRequest(ctx context.Context, body float64, _ *ScalarDecimalTypeClientRequestBodyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal/resquest_body"
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

// RequestParameter -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ScalarDecimalTypeClientRequestParameterOptions contains the optional parameters for the ScalarDecimalTypeClient.RequestParameter
//     method.
func (client *ScalarDecimalTypeClient) RequestParameter(ctx context.Context, value float64, options *ScalarDecimalTypeClientRequestParameterOptions) (ScalarDecimalTypeClientRequestParameterResponse, error) {
	var err error
	const operationName = "ScalarDecimalTypeClient.RequestParameter"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.requestParameterCreateRequest(ctx, value, options)
	if err != nil {
		return ScalarDecimalTypeClientRequestParameterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarDecimalTypeClientRequestParameterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScalarDecimalTypeClientRequestParameterResponse{}, err
	}
	return ScalarDecimalTypeClientRequestParameterResponse{}, nil
}

// requestParameterCreateRequest creates the RequestParameter request.
func (client *ScalarDecimalTypeClient) requestParameterCreateRequest(ctx context.Context, value float64, _ *ScalarDecimalTypeClientRequestParameterOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal/request_parameter"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", strconv.FormatFloat(value, 'f', -1, 64))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// ResponseBody -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ScalarDecimalTypeClientResponseBodyOptions contains the optional parameters for the ScalarDecimalTypeClient.ResponseBody
//     method.
func (client *ScalarDecimalTypeClient) ResponseBody(ctx context.Context, options *ScalarDecimalTypeClientResponseBodyOptions) (ScalarDecimalTypeClientResponseBodyResponse, error) {
	var err error
	const operationName = "ScalarDecimalTypeClient.ResponseBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.responseBodyCreateRequest(ctx, options)
	if err != nil {
		return ScalarDecimalTypeClientResponseBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarDecimalTypeClientResponseBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScalarDecimalTypeClientResponseBodyResponse{}, err
	}
	resp, err := client.responseBodyHandleResponse(httpResp)
	return resp, err
}

// responseBodyCreateRequest creates the ResponseBody request.
func (client *ScalarDecimalTypeClient) responseBodyCreateRequest(ctx context.Context, _ *ScalarDecimalTypeClientResponseBodyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal/response_body"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// responseBodyHandleResponse handles the ResponseBody response.
func (client *ScalarDecimalTypeClient) responseBodyHandleResponse(resp *http.Response) (ScalarDecimalTypeClientResponseBodyResponse, error) {
	result := ScalarDecimalTypeClientResponseBodyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ScalarDecimalTypeClientResponseBodyResponse{}, err
	}
	return result, nil
}
