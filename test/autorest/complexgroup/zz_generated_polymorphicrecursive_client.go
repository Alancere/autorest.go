//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PolymorphicrecursiveClient contains the methods for the Polymorphicrecursive group.
// Don't use this type directly, use NewPolymorphicrecursiveClient() instead.
type PolymorphicrecursiveClient struct {
	pl runtime.Pipeline
}

// NewPolymorphicrecursiveClient creates a new instance of PolymorphicrecursiveClient with the specified values.
// options - pass nil to accept the default values.
func NewPolymorphicrecursiveClient(options *azcore.ClientOptions) *PolymorphicrecursiveClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &PolymorphicrecursiveClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// GetValid - Get complex types that are polymorphic and have recursive references
// If the operation fails it returns the *Error error type.
// options - PolymorphicrecursiveGetValidOptions contains the optional parameters for the PolymorphicrecursiveClient.GetValid
// method.
func (client *PolymorphicrecursiveClient) GetValid(ctx context.Context, options *PolymorphicrecursiveGetValidOptions) (PolymorphicrecursiveGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return PolymorphicrecursiveGetValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphicrecursiveGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphicrecursiveGetValidResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *PolymorphicrecursiveClient) getValidCreateRequest(ctx context.Context, options *PolymorphicrecursiveGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *PolymorphicrecursiveClient) getValidHandleResponse(resp *http.Response) (PolymorphicrecursiveGetValidResponse, error) {
	result := PolymorphicrecursiveGetValidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return PolymorphicrecursiveGetValidResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getValidHandleError handles the GetValid error response.
func (client *PolymorphicrecursiveClient) getValidHandleError(resp *http.Response) error {
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

// PutValid - Put complex types that are polymorphic and have recursive references
// If the operation fails it returns the *Error error type.
// complexBody - Please put a salmon that looks like this: { "fishtype": "salmon", "species": "king", "length": 1, "age":
// 1, "location": "alaska", "iswild": true, "siblings": [ { "fishtype": "shark", "species":
// "predator", "length": 20, "age": 6, "siblings": [ { "fishtype": "salmon", "species": "coho", "length": 2, "age": 2, "location":
// "atlantic", "iswild": true, "siblings": [ { "fishtype": "shark",
// "species": "predator", "length": 20, "age": 6 }, { "fishtype": "sawshark", "species": "dangerous", "length": 10, "age":
// 105 } ] }, { "fishtype": "sawshark", "species": "dangerous", "length": 10,
// "age": 105 } ] }, { "fishtype": "sawshark", "species": "dangerous", "length": 10, "age": 105 } ] }
// options - PolymorphicrecursivePutValidOptions contains the optional parameters for the PolymorphicrecursiveClient.PutValid
// method.
func (client *PolymorphicrecursiveClient) PutValid(ctx context.Context, complexBody FishClassification, options *PolymorphicrecursivePutValidOptions) (PolymorphicrecursivePutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PolymorphicrecursivePutValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolymorphicrecursivePutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphicrecursivePutValidResponse{}, client.putValidHandleError(resp)
	}
	return PolymorphicrecursivePutValidResponse{RawResponse: resp}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *PolymorphicrecursiveClient) putValidCreateRequest(ctx context.Context, complexBody FishClassification, options *PolymorphicrecursivePutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *PolymorphicrecursiveClient) putValidHandleError(resp *http.Response) error {
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
