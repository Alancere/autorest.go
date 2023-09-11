//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HeaderClient contains the methods for the Header group.
// Don't use this type directly, use a constructor function instead.
type HeaderClient struct {
	internal *azcore.Client
}

// CustomNamedRequestID - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - fooClientRequestID - The fooRequestId
//   - options - HeaderClientCustomNamedRequestIDOptions contains the optional parameters for the HeaderClient.CustomNamedRequestID
//     method.
func (client *HeaderClient) CustomNamedRequestID(ctx context.Context, fooClientRequestID string, options *HeaderClientCustomNamedRequestIDOptions) (HeaderClientCustomNamedRequestIDResponse, error) {
	var err error
	const operationName = "HeaderClient.CustomNamedRequestID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.customNamedRequestIDCreateRequest(ctx, fooClientRequestID, options)
	if err != nil {
		return HeaderClientCustomNamedRequestIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientCustomNamedRequestIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientCustomNamedRequestIDResponse{}, err
	}
	resp, err := client.customNamedRequestIDHandleResponse(httpResp)
	return resp, err
}

// customNamedRequestIDCreateRequest creates the CustomNamedRequestID request.
func (client *HeaderClient) customNamedRequestIDCreateRequest(ctx context.Context, fooClientRequestID string, options *HeaderClientCustomNamedRequestIDOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/customNamedRequestId"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["foo-client-request-id"] = []string{fooClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// customNamedRequestIDHandleResponse handles the CustomNamedRequestID response.
func (client *HeaderClient) customNamedRequestIDHandleResponse(resp *http.Response) (HeaderClientCustomNamedRequestIDResponse, error) {
	result := HeaderClientCustomNamedRequestIDResponse{}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// CustomNamedRequestIDHead - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
//
// Generated from API version 2015-07-01-preview
//   - fooClientRequestID - The fooRequestId
//   - options - HeaderClientCustomNamedRequestIDHeadOptions contains the optional parameters for the HeaderClient.CustomNamedRequestIDHead
//     method.
func (client *HeaderClient) CustomNamedRequestIDHead(ctx context.Context, fooClientRequestID string, options *HeaderClientCustomNamedRequestIDHeadOptions) (HeaderClientCustomNamedRequestIDHeadResponse, error) {
	var err error
	const operationName = "HeaderClient.CustomNamedRequestIDHead"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.customNamedRequestIDHeadCreateRequest(ctx, fooClientRequestID, options)
	if err != nil {
		return HeaderClientCustomNamedRequestIDHeadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientCustomNamedRequestIDHeadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientCustomNamedRequestIDHeadResponse{}, err
	}
	resp, err := client.customNamedRequestIDHeadHandleResponse(httpResp)
	return resp, err
}

// customNamedRequestIDHeadCreateRequest creates the CustomNamedRequestIDHead request.
func (client *HeaderClient) customNamedRequestIDHeadCreateRequest(ctx context.Context, fooClientRequestID string, options *HeaderClientCustomNamedRequestIDHeadOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdHead"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["foo-client-request-id"] = []string{fooClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// customNamedRequestIDHeadHandleResponse handles the CustomNamedRequestIDHead response.
func (client *HeaderClient) customNamedRequestIDHeadHandleResponse(resp *http.Response) (HeaderClientCustomNamedRequestIDHeadResponse, error) {
	result := HeaderClientCustomNamedRequestIDHeadResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// CustomNamedRequestIDParamGrouping - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of
// the request, via a parameter group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01-preview
//   - HeaderClientCustomNamedRequestIDParamGroupingParameters - HeaderClientCustomNamedRequestIDParamGroupingParameters contains
//     a group of parameters for the HeaderClient.CustomNamedRequestIDParamGrouping method.
//   - options - HeaderClientCustomNamedRequestIDParamGroupingOptions contains the optional parameters for the HeaderClient.CustomNamedRequestIDParamGrouping
//     method.
func (client *HeaderClient) CustomNamedRequestIDParamGrouping(ctx context.Context, headerClientCustomNamedRequestIDParamGroupingParameters HeaderClientCustomNamedRequestIDParamGroupingParameters, options *HeaderClientCustomNamedRequestIDParamGroupingOptions) (HeaderClientCustomNamedRequestIDParamGroupingResponse, error) {
	var err error
	const operationName = "HeaderClient.CustomNamedRequestIDParamGrouping"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.customNamedRequestIDParamGroupingCreateRequest(ctx, headerClientCustomNamedRequestIDParamGroupingParameters, options)
	if err != nil {
		return HeaderClientCustomNamedRequestIDParamGroupingResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientCustomNamedRequestIDParamGroupingResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientCustomNamedRequestIDParamGroupingResponse{}, err
	}
	resp, err := client.customNamedRequestIDParamGroupingHandleResponse(httpResp)
	return resp, err
}

// customNamedRequestIDParamGroupingCreateRequest creates the CustomNamedRequestIDParamGrouping request.
func (client *HeaderClient) customNamedRequestIDParamGroupingCreateRequest(ctx context.Context, headerClientCustomNamedRequestIDParamGroupingParameters HeaderClientCustomNamedRequestIDParamGroupingParameters, options *HeaderClientCustomNamedRequestIDParamGroupingOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdParamGrouping"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["foo-client-request-id"] = []string{headerClientCustomNamedRequestIDParamGroupingParameters.FooClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// customNamedRequestIDParamGroupingHandleResponse handles the CustomNamedRequestIDParamGrouping response.
func (client *HeaderClient) customNamedRequestIDParamGroupingHandleResponse(resp *http.Response) (HeaderClientCustomNamedRequestIDParamGroupingResponse, error) {
	result := HeaderClientCustomNamedRequestIDParamGroupingResponse{}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}
