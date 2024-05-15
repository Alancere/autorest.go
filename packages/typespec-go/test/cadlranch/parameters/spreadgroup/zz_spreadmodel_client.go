// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package spreadgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SpreadModelClient contains the methods for the Parameters.Spread namespace.
// Don't use this type directly, use [SpreadClient.NewSpreadModelClient] instead.
type SpreadModelClient struct {
	internal *azcore.Client
}

// SpreadAsRequestBody -
// If the operation fails it returns an *azcore.ResponseError type.
//   - bodyParameter - This is a simple model.
//   - options - SpreadModelClientSpreadAsRequestBodyOptions contains the optional parameters for the SpreadModelClient.SpreadAsRequestBody
//     method.
func (client *SpreadModelClient) SpreadAsRequestBody(ctx context.Context, bodyParameter BodyParameter, options *SpreadModelClientSpreadAsRequestBodyOptions) (SpreadModelClientSpreadAsRequestBodyResponse, error) {
	var err error
	const operationName = "SpreadModelClient.SpreadAsRequestBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadAsRequestBodyCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return SpreadModelClientSpreadAsRequestBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadModelClientSpreadAsRequestBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadModelClientSpreadAsRequestBodyResponse{}, err
	}
	return SpreadModelClientSpreadAsRequestBodyResponse{}, nil
}

// spreadAsRequestBodyCreateRequest creates the SpreadAsRequestBody request.
func (client *SpreadModelClient) spreadAsRequestBodyCreateRequest(ctx context.Context, bodyParameter BodyParameter, _ *SpreadModelClientSpreadAsRequestBodyOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/model/request-body"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, bodyParameter); err != nil {
		return nil, err
	}
	return req, nil
}

// SpreadCompositeRequest -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - SpreadModelClientSpreadCompositeRequestOptions contains the optional parameters for the SpreadModelClient.SpreadCompositeRequest
//     method.
func (client *SpreadModelClient) SpreadCompositeRequest(ctx context.Context, name string, testHeader string, body BodyParameter, options *SpreadModelClientSpreadCompositeRequestOptions) (SpreadModelClientSpreadCompositeRequestResponse, error) {
	var err error
	const operationName = "SpreadModelClient.SpreadCompositeRequest"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadCompositeRequestCreateRequest(ctx, name, testHeader, body, options)
	if err != nil {
		return SpreadModelClientSpreadCompositeRequestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadModelClientSpreadCompositeRequestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadModelClientSpreadCompositeRequestResponse{}, err
	}
	return SpreadModelClientSpreadCompositeRequestResponse{}, nil
}

// spreadCompositeRequestCreateRequest creates the SpreadCompositeRequest request.
func (client *SpreadModelClient) spreadCompositeRequestCreateRequest(ctx context.Context, name string, testHeader string, body BodyParameter, _ *SpreadModelClientSpreadCompositeRequestOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/model/composite-request/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	req.Raw().Header["test-header"] = []string{testHeader}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// SpreadCompositeRequestMix -
// If the operation fails it returns an *azcore.ResponseError type.
//   - compositeRequestMix - This is a model with non-body http request decorator.
//   - options - SpreadModelClientSpreadCompositeRequestMixOptions contains the optional parameters for the SpreadModelClient.SpreadCompositeRequestMix
//     method.
func (client *SpreadModelClient) SpreadCompositeRequestMix(ctx context.Context, name string, testHeader string, compositeRequestMix CompositeRequestMix, options *SpreadModelClientSpreadCompositeRequestMixOptions) (SpreadModelClientSpreadCompositeRequestMixResponse, error) {
	var err error
	const operationName = "SpreadModelClient.SpreadCompositeRequestMix"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadCompositeRequestMixCreateRequest(ctx, name, testHeader, compositeRequestMix, options)
	if err != nil {
		return SpreadModelClientSpreadCompositeRequestMixResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadModelClientSpreadCompositeRequestMixResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadModelClientSpreadCompositeRequestMixResponse{}, err
	}
	return SpreadModelClientSpreadCompositeRequestMixResponse{}, nil
}

// spreadCompositeRequestMixCreateRequest creates the SpreadCompositeRequestMix request.
func (client *SpreadModelClient) spreadCompositeRequestMixCreateRequest(ctx context.Context, name string, testHeader string, compositeRequestMix CompositeRequestMix, _ *SpreadModelClientSpreadCompositeRequestMixOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/model/composite-request-mix/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	req.Raw().Header["test-header"] = []string{testHeader}
	if err := runtime.MarshalAsJSON(req, compositeRequestMix); err != nil {
		return nil, err
	}
	return req, nil
}

// SpreadCompositeRequestOnlyWithBody -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - SpreadModelClientSpreadCompositeRequestOnlyWithBodyOptions contains the optional parameters for the SpreadModelClient.SpreadCompositeRequestOnlyWithBody
//     method.
func (client *SpreadModelClient) SpreadCompositeRequestOnlyWithBody(ctx context.Context, body BodyParameter, options *SpreadModelClientSpreadCompositeRequestOnlyWithBodyOptions) (SpreadModelClientSpreadCompositeRequestOnlyWithBodyResponse, error) {
	var err error
	const operationName = "SpreadModelClient.SpreadCompositeRequestOnlyWithBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadCompositeRequestOnlyWithBodyCreateRequest(ctx, body, options)
	if err != nil {
		return SpreadModelClientSpreadCompositeRequestOnlyWithBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadModelClientSpreadCompositeRequestOnlyWithBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadModelClientSpreadCompositeRequestOnlyWithBodyResponse{}, err
	}
	return SpreadModelClientSpreadCompositeRequestOnlyWithBodyResponse{}, nil
}

// spreadCompositeRequestOnlyWithBodyCreateRequest creates the SpreadCompositeRequestOnlyWithBody request.
func (client *SpreadModelClient) spreadCompositeRequestOnlyWithBodyCreateRequest(ctx context.Context, body BodyParameter, _ *SpreadModelClientSpreadCompositeRequestOnlyWithBodyOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/model/composite-request-only-with-body"
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

// SpreadCompositeRequestWithoutBody -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - SpreadModelClientSpreadCompositeRequestWithoutBodyOptions contains the optional parameters for the SpreadModelClient.SpreadCompositeRequestWithoutBody
//     method.
func (client *SpreadModelClient) SpreadCompositeRequestWithoutBody(ctx context.Context, name string, testHeader string, options *SpreadModelClientSpreadCompositeRequestWithoutBodyOptions) (SpreadModelClientSpreadCompositeRequestWithoutBodyResponse, error) {
	var err error
	const operationName = "SpreadModelClient.SpreadCompositeRequestWithoutBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.spreadCompositeRequestWithoutBodyCreateRequest(ctx, name, testHeader, options)
	if err != nil {
		return SpreadModelClientSpreadCompositeRequestWithoutBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpreadModelClientSpreadCompositeRequestWithoutBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpreadModelClientSpreadCompositeRequestWithoutBodyResponse{}, err
	}
	return SpreadModelClientSpreadCompositeRequestWithoutBodyResponse{}, nil
}

// spreadCompositeRequestWithoutBodyCreateRequest creates the SpreadCompositeRequestWithoutBody request.
func (client *SpreadModelClient) spreadCompositeRequestWithoutBodyCreateRequest(ctx context.Context, name string, testHeader string, _ *SpreadModelClientSpreadCompositeRequestWithoutBodyOptions) (*policy.Request, error) {
	urlPath := "/parameters/spread/model/composite-request-without-body/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["test-header"] = []string{testHeader}
	return req, nil
}
