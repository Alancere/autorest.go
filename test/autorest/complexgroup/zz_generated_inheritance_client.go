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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// InheritanceClient contains the methods for the Inheritance group.
// Don't use this type directly, use NewInheritanceClient() instead.
type InheritanceClient struct {
	con *Connection
}

// NewInheritanceClient creates a new instance of InheritanceClient with the specified values.
func NewInheritanceClient(con *Connection) *InheritanceClient {
	return &InheritanceClient{con: con}
}

// GetValid - Get complex types that extend others
// If the operation fails it returns the *Error error type.
func (client *InheritanceClient) GetValid(ctx context.Context, options *InheritanceGetValidOptions) (InheritanceGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return InheritanceGetValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return InheritanceGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InheritanceGetValidResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *InheritanceClient) getValidCreateRequest(ctx context.Context, options *InheritanceGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/inheritance/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *InheritanceClient) getValidHandleResponse(resp *http.Response) (InheritanceGetValidResponse, error) {
	result := InheritanceGetValidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Siamese); err != nil {
		return InheritanceGetValidResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getValidHandleError handles the GetValid error response.
func (client *InheritanceClient) getValidHandleError(resp *http.Response) error {
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

// PutValid - Put complex types that extend others
// If the operation fails it returns the *Error error type.
func (client *InheritanceClient) PutValid(ctx context.Context, complexBody Siamese, options *InheritancePutValidOptions) (InheritancePutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return InheritancePutValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return InheritancePutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InheritancePutValidResponse{}, client.putValidHandleError(resp)
	}
	return InheritancePutValidResponse{RawResponse: resp}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *InheritanceClient) putValidCreateRequest(ctx context.Context, complexBody Siamese, options *InheritancePutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/inheritance/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *InheritanceClient) putValidHandleError(resp *http.Response) error {
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
