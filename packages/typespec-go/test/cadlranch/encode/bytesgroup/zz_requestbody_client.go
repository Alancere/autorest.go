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

// RequestBodyClient contains the methods for the Encode.Bytes group.
// Don't use this type directly, use a constructor function instead.
type RequestBodyClient struct {
	internal *azcore.Client
}

func (client *RequestBodyClient) Base64(ctx context.Context, value []byte, options *RequestBodyClientBase64Options) (RequestBodyClientBase64Response, error) {
	var err error
	req, err := client.base64CreateRequest(ctx, value, options)
	if err != nil {
		return RequestBodyClientBase64Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RequestBodyClientBase64Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RequestBodyClientBase64Response{}, err
	}
	return RequestBodyClientBase64Response{}, nil
}

// base64CreateRequest creates the Base64 request.
func (client *RequestBodyClient) base64CreateRequest(ctx context.Context, value []byte, options *RequestBodyClientBase64Options) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/request/base64"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsByteArray(req, value, runtime.Base64StdFormat); err != nil {
		return nil, err
	}
	return req, nil
}

func (client *RequestBodyClient) Base64url(ctx context.Context, value []byte, options *RequestBodyClientBase64urlOptions) (RequestBodyClientBase64urlResponse, error) {
	var err error
	req, err := client.base64URLCreateRequest(ctx, value, options)
	if err != nil {
		return RequestBodyClientBase64urlResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RequestBodyClientBase64urlResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RequestBodyClientBase64urlResponse{}, err
	}
	return RequestBodyClientBase64urlResponse{}, nil
}

// base64URLCreateRequest creates the Base64url request.
func (client *RequestBodyClient) base64URLCreateRequest(ctx context.Context, value []byte, options *RequestBodyClientBase64urlOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/request/base64url"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsByteArray(req, value, runtime.Base64URLFormat); err != nil {
		return nil, err
	}
	return req, nil
}

func (client *RequestBodyClient) CustomContentType(ctx context.Context, value []byte, options *RequestBodyClientCustomContentTypeOptions) (RequestBodyClientCustomContentTypeResponse, error) {
	var err error
	req, err := client.customContentTypeCreateRequest(ctx, value, options)
	if err != nil {
		return RequestBodyClientCustomContentTypeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RequestBodyClientCustomContentTypeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RequestBodyClientCustomContentTypeResponse{}, err
	}
	return RequestBodyClientCustomContentTypeResponse{}, nil
}

// customContentTypeCreateRequest creates the CustomContentType request.
func (client *RequestBodyClient) customContentTypeCreateRequest(ctx context.Context, value []byte, options *RequestBodyClientCustomContentTypeOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/request/custom-content-type"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["content-type"] = []string{"image/png"}
	if err := runtime.MarshalAsByteArray(req, value, runtime.Base64StdFormat); err != nil {
		return nil, err
	}
	return req, nil
}

func (client *RequestBodyClient) Default(ctx context.Context, value []byte, options *RequestBodyClientDefaultOptions) (RequestBodyClientDefaultResponse, error) {
	var err error
	req, err := client.defaultCreateRequest(ctx, value, options)
	if err != nil {
		return RequestBodyClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RequestBodyClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RequestBodyClientDefaultResponse{}, err
	}
	return RequestBodyClientDefaultResponse{}, nil
}

// defaultCreateRequest creates the Default request.
func (client *RequestBodyClient) defaultCreateRequest(ctx context.Context, value []byte, options *RequestBodyClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/request/default"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsByteArray(req, value, runtime.Base64StdFormat); err != nil {
		return nil, err
	}
	return req, nil
}

func (client *RequestBodyClient) OctetStream(ctx context.Context, value []byte, options *RequestBodyClientOctetStreamOptions) (RequestBodyClientOctetStreamResponse, error) {
	var err error
	req, err := client.octetStreamCreateRequest(ctx, value, options)
	if err != nil {
		return RequestBodyClientOctetStreamResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RequestBodyClientOctetStreamResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RequestBodyClientOctetStreamResponse{}, err
	}
	return RequestBodyClientOctetStreamResponse{}, nil
}

// octetStreamCreateRequest creates the OctetStream request.
func (client *RequestBodyClient) octetStreamCreateRequest(ctx context.Context, value []byte, options *RequestBodyClientOctetStreamOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/request/octet-stream"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["content-type"] = []string{"application/octet-stream"}
	if err := runtime.MarshalAsByteArray(req, value, runtime.Base64StdFormat); err != nil {
		return nil, err
	}
	return req, nil
}