// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ArrayOperations contains the methods for the Array group.
type ArrayOperations interface {
	// GetEmpty - Get complex types with array property which is empty
	GetEmpty(ctx context.Context, options *ArrayGetEmptyOptions) (*ArrayWrapperResponse, error)
	// GetNotProvided - Get complex types with array property while server doesn't provide a response payload
	GetNotProvided(ctx context.Context, options *ArrayGetNotProvidedOptions) (*ArrayWrapperResponse, error)
	// GetValid - Get complex types with array property
	GetValid(ctx context.Context, options *ArrayGetValidOptions) (*ArrayWrapperResponse, error)
	// PutEmpty - Put complex types with array property which is empty
	PutEmpty(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutEmptyOptions) (*http.Response, error)
	// PutValid - Put complex types with array property
	PutValid(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutValidOptions) (*http.Response, error)
}

// ArrayClient implements the ArrayOperations interface.
// Don't use this type directly, use NewArrayClient() instead.
type ArrayClient struct {
	con *Connection
}

// NewArrayClient creates a new instance of ArrayClient with the specified values.
func NewArrayClient(con *Connection) ArrayOperations {
	return &ArrayClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *ArrayClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetEmpty - Get complex types with array property which is empty
func (client *ArrayClient) GetEmpty(ctx context.Context, options *ArrayGetEmptyOptions) (*ArrayWrapperResponse, error) {
	req, err := client.GetEmptyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetEmptyHandleError(resp)
	}
	result, err := client.GetEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEmptyCreateRequest creates the GetEmpty request.
func (client *ArrayClient) GetEmptyCreateRequest(ctx context.Context, options *ArrayGetEmptyOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (client *ArrayClient) GetEmptyHandleResponse(resp *azcore.Response) (*ArrayWrapperResponse, error) {
	result := ArrayWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// GetEmptyHandleError handles the GetEmpty error response.
func (client *ArrayClient) GetEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotProvided - Get complex types with array property while server doesn't provide a response payload
func (client *ArrayClient) GetNotProvided(ctx context.Context, options *ArrayGetNotProvidedOptions) (*ArrayWrapperResponse, error) {
	req, err := client.GetNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNotProvidedHandleError(resp)
	}
	result, err := client.GetNotProvidedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNotProvidedCreateRequest creates the GetNotProvided request.
func (client *ArrayClient) GetNotProvidedCreateRequest(ctx context.Context, options *ArrayGetNotProvidedOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/notprovided"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNotProvidedHandleResponse handles the GetNotProvided response.
func (client *ArrayClient) GetNotProvidedHandleResponse(resp *azcore.Response) (*ArrayWrapperResponse, error) {
	result := ArrayWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// GetNotProvidedHandleError handles the GetNotProvided error response.
func (client *ArrayClient) GetNotProvidedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetValid - Get complex types with array property
func (client *ArrayClient) GetValid(ctx context.Context, options *ArrayGetValidOptions) (*ArrayWrapperResponse, error) {
	req, err := client.GetValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetValidHandleError(resp)
	}
	result, err := client.GetValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetValidCreateRequest creates the GetValid request.
func (client *ArrayClient) GetValidCreateRequest(ctx context.Context, options *ArrayGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetValidHandleResponse handles the GetValid response.
func (client *ArrayClient) GetValidHandleResponse(resp *azcore.Response) (*ArrayWrapperResponse, error) {
	result := ArrayWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// GetValidHandleError handles the GetValid error response.
func (client *ArrayClient) GetValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutEmpty - Put complex types with array property which is empty
func (client *ArrayClient) PutEmpty(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutEmptyOptions) (*http.Response, error) {
	req, err := client.PutEmptyCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// PutEmptyCreateRequest creates the PutEmpty request.
func (client *ArrayClient) PutEmptyCreateRequest(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutEmptyOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/empty"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutEmptyHandleError handles the PutEmpty error response.
func (client *ArrayClient) PutEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Put complex types with array property
func (client *ArrayClient) PutValid(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutValidOptions) (*http.Response, error) {
	req, err := client.PutValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutValidHandleError(resp)
	}
	return resp.Response, nil
}

// PutValidCreateRequest creates the PutValid request.
func (client *ArrayClient) PutValidCreateRequest(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutValidHandleError handles the PutValid error response.
func (client *ArrayClient) PutValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
