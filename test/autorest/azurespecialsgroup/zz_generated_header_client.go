//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HeaderClient contains the methods for the Header group.
// Don't use this type directly, use NewHeaderClient() instead.
type HeaderClient struct {
	pl runtime.Pipeline
}

// NewHeaderClient creates a new instance of HeaderClient with the specified values.
// options - pass nil to accept the default values.
func NewHeaderClient(options *azcore.ClientOptions) *HeaderClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &HeaderClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// CustomNamedRequestID - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
// If the operation fails it returns the *Error error type.
// fooClientRequestID - The fooRequestId
// options - HeaderCustomNamedRequestIDOptions contains the optional parameters for the HeaderClient.CustomNamedRequestID
// method.
func (client *HeaderClient) CustomNamedRequestID(ctx context.Context, fooClientRequestID string, options *HeaderCustomNamedRequestIDOptions) (HeaderCustomNamedRequestIDResponse, error) {
	req, err := client.customNamedRequestIDCreateRequest(ctx, fooClientRequestID, options)
	if err != nil {
		return HeaderCustomNamedRequestIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HeaderCustomNamedRequestIDResponse{}, client.customNamedRequestIDHandleError(resp)
	}
	return client.customNamedRequestIDHandleResponse(resp)
}

// customNamedRequestIDCreateRequest creates the CustomNamedRequestID request.
func (client *HeaderClient) customNamedRequestIDCreateRequest(ctx context.Context, fooClientRequestID string, options *HeaderCustomNamedRequestIDOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/customNamedRequestId"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("foo-client-request-id", fooClientRequestID)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIDHandleResponse handles the CustomNamedRequestID response.
func (client *HeaderClient) customNamedRequestIDHandleResponse(resp *http.Response) (HeaderCustomNamedRequestIDResponse, error) {
	result := HeaderCustomNamedRequestIDResponse{RawResponse: resp}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// customNamedRequestIDHandleError handles the CustomNamedRequestID error response.
func (client *HeaderClient) customNamedRequestIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CustomNamedRequestIDHead - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
// If the operation fails it returns the *Error error type.
// fooClientRequestID - The fooRequestId
// options - HeaderCustomNamedRequestIDHeadOptions contains the optional parameters for the HeaderClient.CustomNamedRequestIDHead
// method.
func (client *HeaderClient) CustomNamedRequestIDHead(ctx context.Context, fooClientRequestID string, options *HeaderCustomNamedRequestIDHeadOptions) (HeaderCustomNamedRequestIDHeadResponse, error) {
	req, err := client.customNamedRequestIDHeadCreateRequest(ctx, fooClientRequestID, options)
	if err != nil {
		return HeaderCustomNamedRequestIDHeadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDHeadResponse{}, err
	}
	return client.customNamedRequestIDHeadHandleResponse(resp)
}

// customNamedRequestIDHeadCreateRequest creates the CustomNamedRequestIDHead request.
func (client *HeaderClient) customNamedRequestIDHeadCreateRequest(ctx context.Context, fooClientRequestID string, options *HeaderCustomNamedRequestIDHeadOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdHead"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("foo-client-request-id", fooClientRequestID)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIDHeadHandleResponse handles the CustomNamedRequestIDHead response.
func (client *HeaderClient) customNamedRequestIDHeadHandleResponse(resp *http.Response) (HeaderCustomNamedRequestIDHeadResponse, error) {
	result := HeaderCustomNamedRequestIDHeadResponse{RawResponse: resp}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// CustomNamedRequestIDParamGrouping - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of
// the request, via a parameter group
// If the operation fails it returns the *Error error type.
// HeaderCustomNamedRequestIDParamGroupingParameters - HeaderCustomNamedRequestIDParamGroupingParameters contains a group
// of parameters for the Header.CustomNamedRequestIDParamGrouping method.
func (client *HeaderClient) CustomNamedRequestIDParamGrouping(ctx context.Context, headerCustomNamedRequestIDParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (HeaderCustomNamedRequestIDParamGroupingResponse, error) {
	req, err := client.customNamedRequestIDParamGroupingCreateRequest(ctx, headerCustomNamedRequestIDParamGroupingParameters)
	if err != nil {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, client.customNamedRequestIDParamGroupingHandleError(resp)
	}
	return client.customNamedRequestIDParamGroupingHandleResponse(resp)
}

// customNamedRequestIDParamGroupingCreateRequest creates the CustomNamedRequestIDParamGrouping request.
func (client *HeaderClient) customNamedRequestIDParamGroupingCreateRequest(ctx context.Context, headerCustomNamedRequestIDParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (*policy.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdParamGrouping"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("foo-client-request-id", headerCustomNamedRequestIDParamGroupingParameters.FooClientRequestID)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIDParamGroupingHandleResponse handles the CustomNamedRequestIDParamGrouping response.
func (client *HeaderClient) customNamedRequestIDParamGroupingHandleResponse(resp *http.Response) (HeaderCustomNamedRequestIDParamGroupingResponse, error) {
	result := HeaderCustomNamedRequestIDParamGroupingResponse{RawResponse: resp}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// customNamedRequestIDParamGroupingHandleError handles the CustomNamedRequestIDParamGrouping error response.
func (client *HeaderClient) customNamedRequestIDParamGroupingHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
