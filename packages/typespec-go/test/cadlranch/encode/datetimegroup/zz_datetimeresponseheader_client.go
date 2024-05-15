// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package datetimegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"strconv"
	"time"
)

// DatetimeResponseHeaderClient contains the methods for the Encode.Datetime namespace.
// Don't use this type directly, use [DatetimeClient.NewDatetimeResponseHeaderClient] instead.
type DatetimeResponseHeaderClient struct {
	internal *azcore.Client
}

// Default -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DatetimeResponseHeaderClientDefaultOptions contains the optional parameters for the DatetimeResponseHeaderClient.Default
//     method.
func (client *DatetimeResponseHeaderClient) Default(ctx context.Context, options *DatetimeResponseHeaderClientDefaultOptions) (DatetimeResponseHeaderClientDefaultResponse, error) {
	var err error
	const operationName = "DatetimeResponseHeaderClient.Default"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.defaultCreateRequest(ctx, options)
	if err != nil {
		return DatetimeResponseHeaderClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeResponseHeaderClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeResponseHeaderClientDefaultResponse{}, err
	}
	resp, err := client.defaultHandleResponse(httpResp)
	return resp, err
}

// defaultCreateRequest creates the Default request.
func (client *DatetimeResponseHeaderClient) defaultCreateRequest(ctx context.Context, _ *DatetimeResponseHeaderClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/responseheader/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// defaultHandleResponse handles the Default response.
func (client *DatetimeResponseHeaderClient) defaultHandleResponse(resp *http.Response) (DatetimeResponseHeaderClientDefaultResponse, error) {
	result := DatetimeResponseHeaderClientDefaultResponse{}
	if val := resp.Header.Get("value"); val != "" {
		value, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return DatetimeResponseHeaderClientDefaultResponse{}, err
		}
		result.Value = &value
	}
	return result, nil
}

// RFC3339 -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DatetimeResponseHeaderClientRFC3339Options contains the optional parameters for the DatetimeResponseHeaderClient.RFC3339
//     method.
func (client *DatetimeResponseHeaderClient) RFC3339(ctx context.Context, options *DatetimeResponseHeaderClientRFC3339Options) (DatetimeResponseHeaderClientRFC3339Response, error) {
	var err error
	const operationName = "DatetimeResponseHeaderClient.RFC3339"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rfc3339CreateRequest(ctx, options)
	if err != nil {
		return DatetimeResponseHeaderClientRFC3339Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeResponseHeaderClientRFC3339Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeResponseHeaderClientRFC3339Response{}, err
	}
	resp, err := client.rfc3339HandleResponse(httpResp)
	return resp, err
}

// rfc3339CreateRequest creates the RFC3339 request.
func (client *DatetimeResponseHeaderClient) rfc3339CreateRequest(ctx context.Context, _ *DatetimeResponseHeaderClientRFC3339Options) (*policy.Request, error) {
	urlPath := "/encode/datetime/responseheader/rfc3339"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// rfc3339HandleResponse handles the RFC3339 response.
func (client *DatetimeResponseHeaderClient) rfc3339HandleResponse(resp *http.Response) (DatetimeResponseHeaderClientRFC3339Response, error) {
	result := DatetimeResponseHeaderClientRFC3339Response{}
	if val := resp.Header.Get("value"); val != "" {
		value, err := time.Parse(time.RFC3339Nano, val)
		if err != nil {
			return DatetimeResponseHeaderClientRFC3339Response{}, err
		}
		result.Value = &value
	}
	return result, nil
}

// RFC7231 -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DatetimeResponseHeaderClientRFC7231Options contains the optional parameters for the DatetimeResponseHeaderClient.RFC7231
//     method.
func (client *DatetimeResponseHeaderClient) RFC7231(ctx context.Context, options *DatetimeResponseHeaderClientRFC7231Options) (DatetimeResponseHeaderClientRFC7231Response, error) {
	var err error
	const operationName = "DatetimeResponseHeaderClient.RFC7231"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rfc7231CreateRequest(ctx, options)
	if err != nil {
		return DatetimeResponseHeaderClientRFC7231Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeResponseHeaderClientRFC7231Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeResponseHeaderClientRFC7231Response{}, err
	}
	resp, err := client.rfc7231HandleResponse(httpResp)
	return resp, err
}

// rfc7231CreateRequest creates the RFC7231 request.
func (client *DatetimeResponseHeaderClient) rfc7231CreateRequest(ctx context.Context, _ *DatetimeResponseHeaderClientRFC7231Options) (*policy.Request, error) {
	urlPath := "/encode/datetime/responseheader/rfc7231"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// rfc7231HandleResponse handles the RFC7231 response.
func (client *DatetimeResponseHeaderClient) rfc7231HandleResponse(resp *http.Response) (DatetimeResponseHeaderClientRFC7231Response, error) {
	result := DatetimeResponseHeaderClientRFC7231Response{}
	if val := resp.Header.Get("value"); val != "" {
		value, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return DatetimeResponseHeaderClientRFC7231Response{}, err
		}
		result.Value = &value
	}
	return result, nil
}

// UnixTimestamp -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DatetimeResponseHeaderClientUnixTimestampOptions contains the optional parameters for the DatetimeResponseHeaderClient.UnixTimestamp
//     method.
func (client *DatetimeResponseHeaderClient) UnixTimestamp(ctx context.Context, options *DatetimeResponseHeaderClientUnixTimestampOptions) (DatetimeResponseHeaderClientUnixTimestampResponse, error) {
	var err error
	const operationName = "DatetimeResponseHeaderClient.UnixTimestamp"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unixTimestampCreateRequest(ctx, options)
	if err != nil {
		return DatetimeResponseHeaderClientUnixTimestampResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeResponseHeaderClientUnixTimestampResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeResponseHeaderClientUnixTimestampResponse{}, err
	}
	resp, err := client.unixTimestampHandleResponse(httpResp)
	return resp, err
}

// unixTimestampCreateRequest creates the UnixTimestamp request.
func (client *DatetimeResponseHeaderClient) unixTimestampCreateRequest(ctx context.Context, _ *DatetimeResponseHeaderClientUnixTimestampOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/responseheader/unix-timestamp"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// unixTimestampHandleResponse handles the UnixTimestamp response.
func (client *DatetimeResponseHeaderClient) unixTimestampHandleResponse(resp *http.Response) (DatetimeResponseHeaderClientUnixTimestampResponse, error) {
	result := DatetimeResponseHeaderClientUnixTimestampResponse{}
	if val := resp.Header.Get("value"); val != "" {
		sec, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return DatetimeResponseHeaderClientUnixTimestampResponse{}, err
		}
		result.Value = to.Ptr(time.Unix(sec, 0))
	}
	return result, nil
}
