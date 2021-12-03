//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FlattencomplexClient contains the methods for the Flattencomplex group.
// Don't use this type directly, use NewFlattencomplexClient() instead.
type FlattencomplexClient struct {
	pl runtime.Pipeline
}

// NewFlattencomplexClient creates a new instance of FlattencomplexClient with the specified values.
// options - pass nil to accept the default values.
func NewFlattencomplexClient(options *azcore.ClientOptions) *FlattencomplexClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &FlattencomplexClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// GetValid -
// If the operation fails it returns a generic error.
// options - FlattencomplexGetValidOptions contains the optional parameters for the FlattencomplexClient.GetValid method.
func (client *FlattencomplexClient) GetValid(ctx context.Context, options *FlattencomplexGetValidOptions) (FlattencomplexGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return FlattencomplexGetValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FlattencomplexGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FlattencomplexGetValidResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *FlattencomplexClient) getValidCreateRequest(ctx context.Context, options *FlattencomplexGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/flatten/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *FlattencomplexClient) getValidHandleResponse(resp *http.Response) (FlattencomplexGetValidResponse, error) {
	result := FlattencomplexGetValidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return FlattencomplexGetValidResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getValidHandleError handles the GetValid error response.
func (client *FlattencomplexClient) getValidHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
