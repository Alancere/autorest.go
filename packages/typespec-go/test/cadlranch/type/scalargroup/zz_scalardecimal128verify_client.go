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
)

// ScalarDecimal128VerifyClient - Decimal128 type verification
// Don't use this type directly, use [ScalarClient.NewScalarDecimal128VerifyClient] instead.
type ScalarDecimal128VerifyClient struct {
	internal *azcore.Client
}

// PrepareVerify -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ScalarDecimal128VerifyClientPrepareVerifyOptions contains the optional parameters for the ScalarDecimal128VerifyClient.PrepareVerify
//     method.
func (client *ScalarDecimal128VerifyClient) PrepareVerify(ctx context.Context, options *ScalarDecimal128VerifyClientPrepareVerifyOptions) (ScalarDecimal128VerifyClientPrepareVerifyResponse, error) {
	var err error
	const operationName = "ScalarDecimal128VerifyClient.PrepareVerify"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.prepareVerifyCreateRequest(ctx, options)
	if err != nil {
		return ScalarDecimal128VerifyClientPrepareVerifyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarDecimal128VerifyClientPrepareVerifyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScalarDecimal128VerifyClientPrepareVerifyResponse{}, err
	}
	resp, err := client.prepareVerifyHandleResponse(httpResp)
	return resp, err
}

// prepareVerifyCreateRequest creates the PrepareVerify request.
func (client *ScalarDecimal128VerifyClient) prepareVerifyCreateRequest(ctx context.Context, _ *ScalarDecimal128VerifyClientPrepareVerifyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal128/prepare_verify"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// prepareVerifyHandleResponse handles the PrepareVerify response.
func (client *ScalarDecimal128VerifyClient) prepareVerifyHandleResponse(resp *http.Response) (ScalarDecimal128VerifyClientPrepareVerifyResponse, error) {
	result := ScalarDecimal128VerifyClientPrepareVerifyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ScalarDecimal128VerifyClientPrepareVerifyResponse{}, err
	}
	return result, nil
}

// Verify -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - ScalarDecimal128VerifyClientVerifyOptions contains the optional parameters for the ScalarDecimal128VerifyClient.Verify
//     method.
func (client *ScalarDecimal128VerifyClient) Verify(ctx context.Context, body float64, options *ScalarDecimal128VerifyClientVerifyOptions) (ScalarDecimal128VerifyClientVerifyResponse, error) {
	var err error
	const operationName = "ScalarDecimal128VerifyClient.Verify"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.verifyCreateRequest(ctx, body, options)
	if err != nil {
		return ScalarDecimal128VerifyClientVerifyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarDecimal128VerifyClientVerifyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScalarDecimal128VerifyClientVerifyResponse{}, err
	}
	return ScalarDecimal128VerifyClientVerifyResponse{}, nil
}

// verifyCreateRequest creates the Verify request.
func (client *ScalarDecimal128VerifyClient) verifyCreateRequest(ctx context.Context, body float64, _ *ScalarDecimal128VerifyClientVerifyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal128/verify"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
