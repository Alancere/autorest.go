//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bytesgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PropertyClient contains the methods for the Encode.Bytes namespace.
// Don't use this type directly, use [BytesClient.NewPropertyClient] instead.
type PropertyClient struct {
	internal *azcore.Client
}

// - options - PropertyClientBase64Options contains the optional parameters for the PropertyClient.Base64 method.
func (client *PropertyClient) Base64(ctx context.Context, body Base64BytesProperty, options *PropertyClientBase64Options) (PropertyClientBase64Response, error) {
	var err error
	req, err := client.base64CreateRequest(ctx, body, options)
	if err != nil {
		return PropertyClientBase64Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PropertyClientBase64Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PropertyClientBase64Response{}, err
	}
	resp, err := client.base64HandleResponse(httpResp)
	return resp, err
}

// base64CreateRequest creates the Base64 request.
func (client *PropertyClient) base64CreateRequest(ctx context.Context, body Base64BytesProperty, options *PropertyClientBase64Options) (*policy.Request, error) {
	urlPath := "/encode/bytes/property/base64"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// base64HandleResponse handles the Base64 response.
func (client *PropertyClient) base64HandleResponse(resp *http.Response) (PropertyClientBase64Response, error) {
	result := PropertyClientBase64Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Base64BytesProperty); err != nil {
		return PropertyClientBase64Response{}, err
	}
	return result, nil
}

// - options - PropertyClientBase64URLOptions contains the optional parameters for the PropertyClient.Base64URL method.
func (client *PropertyClient) Base64URL(ctx context.Context, body Base64URLBytesProperty, options *PropertyClientBase64URLOptions) (PropertyClientBase64URLResponse, error) {
	var err error
	req, err := client.base64URLCreateRequest(ctx, body, options)
	if err != nil {
		return PropertyClientBase64URLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PropertyClientBase64URLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PropertyClientBase64URLResponse{}, err
	}
	resp, err := client.base64URLHandleResponse(httpResp)
	return resp, err
}

// base64URLCreateRequest creates the Base64URL request.
func (client *PropertyClient) base64URLCreateRequest(ctx context.Context, body Base64URLBytesProperty, options *PropertyClientBase64URLOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/property/base64url"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// base64URLHandleResponse handles the Base64URL response.
func (client *PropertyClient) base64URLHandleResponse(resp *http.Response) (PropertyClientBase64URLResponse, error) {
	result := PropertyClientBase64URLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Base64URLBytesProperty); err != nil {
		return PropertyClientBase64URLResponse{}, err
	}
	return result, nil
}

// - options - PropertyClientBase64URLArrayOptions contains the optional parameters for the PropertyClient.Base64URLArray method.
func (client *PropertyClient) Base64URLArray(ctx context.Context, body Base64URLArrayBytesProperty, options *PropertyClientBase64URLArrayOptions) (PropertyClientBase64URLArrayResponse, error) {
	var err error
	req, err := client.base64URLArrayCreateRequest(ctx, body, options)
	if err != nil {
		return PropertyClientBase64URLArrayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PropertyClientBase64URLArrayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PropertyClientBase64URLArrayResponse{}, err
	}
	resp, err := client.base64URLArrayHandleResponse(httpResp)
	return resp, err
}

// base64URLArrayCreateRequest creates the Base64URLArray request.
func (client *PropertyClient) base64URLArrayCreateRequest(ctx context.Context, body Base64URLArrayBytesProperty, options *PropertyClientBase64URLArrayOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/property/base64url-array"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// base64URLArrayHandleResponse handles the Base64URLArray response.
func (client *PropertyClient) base64URLArrayHandleResponse(resp *http.Response) (PropertyClientBase64URLArrayResponse, error) {
	result := PropertyClientBase64URLArrayResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Base64URLArrayBytesProperty); err != nil {
		return PropertyClientBase64URLArrayResponse{}, err
	}
	return result, nil
}

// - options - PropertyClientDefaultOptions contains the optional parameters for the PropertyClient.Default method.
func (client *PropertyClient) Default(ctx context.Context, body DefaultBytesProperty, options *PropertyClientDefaultOptions) (PropertyClientDefaultResponse, error) {
	var err error
	req, err := client.defaultCreateRequest(ctx, body, options)
	if err != nil {
		return PropertyClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PropertyClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PropertyClientDefaultResponse{}, err
	}
	resp, err := client.defaultHandleResponse(httpResp)
	return resp, err
}

// defaultCreateRequest creates the Default request.
func (client *PropertyClient) defaultCreateRequest(ctx context.Context, body DefaultBytesProperty, options *PropertyClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/property/default"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// defaultHandleResponse handles the Default response.
func (client *PropertyClient) defaultHandleResponse(resp *http.Response) (PropertyClientDefaultResponse, error) {
	result := PropertyClientDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefaultBytesProperty); err != nil {
		return PropertyClientDefaultResponse{}, err
	}
	return result, nil
}
