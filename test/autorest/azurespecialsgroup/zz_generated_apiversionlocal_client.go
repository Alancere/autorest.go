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

// APIVersionLocalClient contains the methods for the APIVersionLocal group.
// Don't use this type directly, use NewAPIVersionLocalClient() instead.
type APIVersionLocalClient struct {
	pl runtime.Pipeline
}

// NewAPIVersionLocalClient creates a new instance of APIVersionLocalClient with the specified values.
// options - pass nil to accept the default values.
func NewAPIVersionLocalClient(options *azcore.ClientOptions) *APIVersionLocalClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &APIVersionLocalClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// GetMethodLocalNull - Get method with api-version modeled in the method. pass in api-version = null to succeed
// If the operation fails it returns the *Error error type.
// options - APIVersionLocalGetMethodLocalNullOptions contains the optional parameters for the APIVersionLocalClient.GetMethodLocalNull
// method.
func (client *APIVersionLocalClient) GetMethodLocalNull(ctx context.Context, options *APIVersionLocalGetMethodLocalNullOptions) (APIVersionLocalGetMethodLocalNullResponse, error) {
	req, err := client.getMethodLocalNullCreateRequest(ctx, options)
	if err != nil {
		return APIVersionLocalGetMethodLocalNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionLocalGetMethodLocalNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionLocalGetMethodLocalNullResponse{}, client.getMethodLocalNullHandleError(resp)
	}
	return APIVersionLocalGetMethodLocalNullResponse{RawResponse: resp}, nil
}

// getMethodLocalNullCreateRequest creates the GetMethodLocalNull request.
func (client *APIVersionLocalClient) getMethodLocalNullCreateRequest(ctx context.Context, options *APIVersionLocalGetMethodLocalNullOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/local/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.APIVersion != nil {
		reqQP.Set("api-version", *options.APIVersion)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodLocalNullHandleError handles the GetMethodLocalNull error response.
func (client *APIVersionLocalClient) getMethodLocalNullHandleError(resp *http.Response) error {
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

// GetMethodLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
// If the operation fails it returns the *Error error type.
// options - APIVersionLocalGetMethodLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetMethodLocalValid
// method.
func (client *APIVersionLocalClient) GetMethodLocalValid(ctx context.Context, options *APIVersionLocalGetMethodLocalValidOptions) (APIVersionLocalGetMethodLocalValidResponse, error) {
	req, err := client.getMethodLocalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionLocalGetMethodLocalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionLocalGetMethodLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionLocalGetMethodLocalValidResponse{}, client.getMethodLocalValidHandleError(resp)
	}
	return APIVersionLocalGetMethodLocalValidResponse{RawResponse: resp}, nil
}

// getMethodLocalValidCreateRequest creates the GetMethodLocalValid request.
func (client *APIVersionLocalClient) getMethodLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalGetMethodLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/local/2.0"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2.0")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodLocalValidHandleError handles the GetMethodLocalValid error response.
func (client *APIVersionLocalClient) getMethodLocalValidHandleError(resp *http.Response) error {
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

// GetPathLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
// If the operation fails it returns the *Error error type.
// options - APIVersionLocalGetPathLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetPathLocalValid
// method.
func (client *APIVersionLocalClient) GetPathLocalValid(ctx context.Context, options *APIVersionLocalGetPathLocalValidOptions) (APIVersionLocalGetPathLocalValidResponse, error) {
	req, err := client.getPathLocalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionLocalGetPathLocalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionLocalGetPathLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionLocalGetPathLocalValidResponse{}, client.getPathLocalValidHandleError(resp)
	}
	return APIVersionLocalGetPathLocalValidResponse{RawResponse: resp}, nil
}

// getPathLocalValidCreateRequest creates the GetPathLocalValid request.
func (client *APIVersionLocalClient) getPathLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalGetPathLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/path/string/none/query/local/2.0"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2.0")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPathLocalValidHandleError handles the GetPathLocalValid error response.
func (client *APIVersionLocalClient) getPathLocalValidHandleError(resp *http.Response) error {
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

// GetSwaggerLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
// If the operation fails it returns the *Error error type.
// options - APIVersionLocalGetSwaggerLocalValidOptions contains the optional parameters for the APIVersionLocalClient.GetSwaggerLocalValid
// method.
func (client *APIVersionLocalClient) GetSwaggerLocalValid(ctx context.Context, options *APIVersionLocalGetSwaggerLocalValidOptions) (APIVersionLocalGetSwaggerLocalValidResponse, error) {
	req, err := client.getSwaggerLocalValidCreateRequest(ctx, options)
	if err != nil {
		return APIVersionLocalGetSwaggerLocalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionLocalGetSwaggerLocalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionLocalGetSwaggerLocalValidResponse{}, client.getSwaggerLocalValidHandleError(resp)
	}
	return APIVersionLocalGetSwaggerLocalValidResponse{RawResponse: resp}, nil
}

// getSwaggerLocalValidCreateRequest creates the GetSwaggerLocalValid request.
func (client *APIVersionLocalClient) getSwaggerLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalGetSwaggerLocalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/apiVersion/swagger/string/none/query/local/2.0"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2.0")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSwaggerLocalValidHandleError handles the GetSwaggerLocalValid error response.
func (client *APIVersionLocalClient) getSwaggerLocalValidHandleError(resp *http.Response) error {
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
