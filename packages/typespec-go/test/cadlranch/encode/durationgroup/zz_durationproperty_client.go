// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package durationgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DurationPropertyClient contains the methods for the Encode.Duration namespace.
// Don't use this type directly, use [DurationClient.NewDurationPropertyClient] instead.
type DurationPropertyClient struct {
	internal *azcore.Client
}

// Default -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DurationPropertyClientDefaultOptions contains the optional parameters for the DurationPropertyClient.Default
//     method.
func (client *DurationPropertyClient) Default(ctx context.Context, body DefaultDurationProperty, options *DurationPropertyClientDefaultOptions) (DurationPropertyClientDefaultResponse, error) {
	var err error
	const operationName = "DurationPropertyClient.Default"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.defaultCreateRequest(ctx, body, options)
	if err != nil {
		return DurationPropertyClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationPropertyClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DurationPropertyClientDefaultResponse{}, err
	}
	resp, err := client.defaultHandleResponse(httpResp)
	return resp, err
}

// defaultCreateRequest creates the Default request.
func (client *DurationPropertyClient) defaultCreateRequest(ctx context.Context, body DefaultDurationProperty, _ *DurationPropertyClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/property/default"
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
func (client *DurationPropertyClient) defaultHandleResponse(resp *http.Response) (DurationPropertyClientDefaultResponse, error) {
	result := DurationPropertyClientDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefaultDurationProperty); err != nil {
		return DurationPropertyClientDefaultResponse{}, err
	}
	return result, nil
}

// Float64Seconds -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DurationPropertyClientFloat64SecondsOptions contains the optional parameters for the DurationPropertyClient.Float64Seconds
//     method.
func (client *DurationPropertyClient) Float64Seconds(ctx context.Context, body Float64SecondsDurationProperty, options *DurationPropertyClientFloat64SecondsOptions) (DurationPropertyClientFloat64SecondsResponse, error) {
	var err error
	const operationName = "DurationPropertyClient.Float64Seconds"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.float64SecondsCreateRequest(ctx, body, options)
	if err != nil {
		return DurationPropertyClientFloat64SecondsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationPropertyClientFloat64SecondsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DurationPropertyClientFloat64SecondsResponse{}, err
	}
	resp, err := client.float64SecondsHandleResponse(httpResp)
	return resp, err
}

// float64SecondsCreateRequest creates the Float64Seconds request.
func (client *DurationPropertyClient) float64SecondsCreateRequest(ctx context.Context, body Float64SecondsDurationProperty, _ *DurationPropertyClientFloat64SecondsOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/property/float64-seconds"
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

// float64SecondsHandleResponse handles the Float64Seconds response.
func (client *DurationPropertyClient) float64SecondsHandleResponse(resp *http.Response) (DurationPropertyClientFloat64SecondsResponse, error) {
	result := DurationPropertyClientFloat64SecondsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Float64SecondsDurationProperty); err != nil {
		return DurationPropertyClientFloat64SecondsResponse{}, err
	}
	return result, nil
}

// FloatSeconds -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DurationPropertyClientFloatSecondsOptions contains the optional parameters for the DurationPropertyClient.FloatSeconds
//     method.
func (client *DurationPropertyClient) FloatSeconds(ctx context.Context, body FloatSecondsDurationProperty, options *DurationPropertyClientFloatSecondsOptions) (DurationPropertyClientFloatSecondsResponse, error) {
	var err error
	const operationName = "DurationPropertyClient.FloatSeconds"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.floatSecondsCreateRequest(ctx, body, options)
	if err != nil {
		return DurationPropertyClientFloatSecondsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationPropertyClientFloatSecondsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DurationPropertyClientFloatSecondsResponse{}, err
	}
	resp, err := client.floatSecondsHandleResponse(httpResp)
	return resp, err
}

// floatSecondsCreateRequest creates the FloatSeconds request.
func (client *DurationPropertyClient) floatSecondsCreateRequest(ctx context.Context, body FloatSecondsDurationProperty, _ *DurationPropertyClientFloatSecondsOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/property/float-seconds"
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

// floatSecondsHandleResponse handles the FloatSeconds response.
func (client *DurationPropertyClient) floatSecondsHandleResponse(resp *http.Response) (DurationPropertyClientFloatSecondsResponse, error) {
	result := DurationPropertyClientFloatSecondsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FloatSecondsDurationProperty); err != nil {
		return DurationPropertyClientFloatSecondsResponse{}, err
	}
	return result, nil
}

// FloatSecondsArray -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DurationPropertyClientFloatSecondsArrayOptions contains the optional parameters for the DurationPropertyClient.FloatSecondsArray
//     method.
func (client *DurationPropertyClient) FloatSecondsArray(ctx context.Context, body FloatSecondsDurationArrayProperty, options *DurationPropertyClientFloatSecondsArrayOptions) (DurationPropertyClientFloatSecondsArrayResponse, error) {
	var err error
	const operationName = "DurationPropertyClient.FloatSecondsArray"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.floatSecondsArrayCreateRequest(ctx, body, options)
	if err != nil {
		return DurationPropertyClientFloatSecondsArrayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationPropertyClientFloatSecondsArrayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DurationPropertyClientFloatSecondsArrayResponse{}, err
	}
	resp, err := client.floatSecondsArrayHandleResponse(httpResp)
	return resp, err
}

// floatSecondsArrayCreateRequest creates the FloatSecondsArray request.
func (client *DurationPropertyClient) floatSecondsArrayCreateRequest(ctx context.Context, body FloatSecondsDurationArrayProperty, _ *DurationPropertyClientFloatSecondsArrayOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/property/float-seconds-array"
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

// floatSecondsArrayHandleResponse handles the FloatSecondsArray response.
func (client *DurationPropertyClient) floatSecondsArrayHandleResponse(resp *http.Response) (DurationPropertyClientFloatSecondsArrayResponse, error) {
	result := DurationPropertyClientFloatSecondsArrayResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FloatSecondsDurationArrayProperty); err != nil {
		return DurationPropertyClientFloatSecondsArrayResponse{}, err
	}
	return result, nil
}

// ISO8601 -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DurationPropertyClientISO8601Options contains the optional parameters for the DurationPropertyClient.ISO8601
//     method.
func (client *DurationPropertyClient) ISO8601(ctx context.Context, body ISO8601DurationProperty, options *DurationPropertyClientISO8601Options) (DurationPropertyClientISO8601Response, error) {
	var err error
	const operationName = "DurationPropertyClient.ISO8601"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.iso8601CreateRequest(ctx, body, options)
	if err != nil {
		return DurationPropertyClientISO8601Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationPropertyClientISO8601Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DurationPropertyClientISO8601Response{}, err
	}
	resp, err := client.iso8601HandleResponse(httpResp)
	return resp, err
}

// iso8601CreateRequest creates the ISO8601 request.
func (client *DurationPropertyClient) iso8601CreateRequest(ctx context.Context, body ISO8601DurationProperty, _ *DurationPropertyClientISO8601Options) (*policy.Request, error) {
	urlPath := "/encode/duration/property/iso8601"
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

// iso8601HandleResponse handles the ISO8601 response.
func (client *DurationPropertyClient) iso8601HandleResponse(resp *http.Response) (DurationPropertyClientISO8601Response, error) {
	result := DurationPropertyClientISO8601Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ISO8601DurationProperty); err != nil {
		return DurationPropertyClientISO8601Response{}, err
	}
	return result, nil
}

// Int32Seconds -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - DurationPropertyClientInt32SecondsOptions contains the optional parameters for the DurationPropertyClient.Int32Seconds
//     method.
func (client *DurationPropertyClient) Int32Seconds(ctx context.Context, body Int32SecondsDurationProperty, options *DurationPropertyClientInt32SecondsOptions) (DurationPropertyClientInt32SecondsResponse, error) {
	var err error
	const operationName = "DurationPropertyClient.Int32Seconds"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.int32SecondsCreateRequest(ctx, body, options)
	if err != nil {
		return DurationPropertyClientInt32SecondsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationPropertyClientInt32SecondsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DurationPropertyClientInt32SecondsResponse{}, err
	}
	resp, err := client.int32SecondsHandleResponse(httpResp)
	return resp, err
}

// int32SecondsCreateRequest creates the Int32Seconds request.
func (client *DurationPropertyClient) int32SecondsCreateRequest(ctx context.Context, body Int32SecondsDurationProperty, _ *DurationPropertyClientInt32SecondsOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/property/int32-seconds"
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

// int32SecondsHandleResponse handles the Int32Seconds response.
func (client *DurationPropertyClient) int32SecondsHandleResponse(resp *http.Response) (DurationPropertyClientInt32SecondsResponse, error) {
	result := DurationPropertyClientInt32SecondsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Int32SecondsDurationProperty); err != nil {
		return DurationPropertyClientInt32SecondsResponse{}, err
	}
	return result, nil
}
