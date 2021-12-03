//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type bigDataPoolsClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newBigDataPoolsClient creates a new instance of bigDataPoolsClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newBigDataPoolsClient(endpoint string, pl runtime.Pipeline) *bigDataPoolsClient {
	client := &bigDataPoolsClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Get - Get Big Data Pool
// If the operation fails it returns the *ErrorContract error type.
// bigDataPoolName - The Big Data Pool name
// options - BigDataPoolsGetOptions contains the optional parameters for the bigDataPoolsClient.Get method.
func (client *bigDataPoolsClient) Get(ctx context.Context, bigDataPoolName string, options *BigDataPoolsGetOptions) (BigDataPoolsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, bigDataPoolName, options)
	if err != nil {
		return BigDataPoolsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BigDataPoolsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BigDataPoolsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *bigDataPoolsClient) getCreateRequest(ctx context.Context, bigDataPoolName string, options *BigDataPoolsGetOptions) (*policy.Request, error) {
	urlPath := "/bigDataPools/{bigDataPoolName}"
	if bigDataPoolName == "" {
		return nil, errors.New("parameter bigDataPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bigDataPoolName}", url.PathEscape(bigDataPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *bigDataPoolsClient) getHandleResponse(resp *http.Response) (BigDataPoolsGetResponse, error) {
	result := BigDataPoolsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BigDataPoolResourceInfo); err != nil {
		return BigDataPoolsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *bigDataPoolsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List Big Data Pools
// If the operation fails it returns the *ErrorContract error type.
// options - BigDataPoolsListOptions contains the optional parameters for the bigDataPoolsClient.List method.
func (client *bigDataPoolsClient) List(ctx context.Context, options *BigDataPoolsListOptions) (BigDataPoolsListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return BigDataPoolsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BigDataPoolsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BigDataPoolsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *bigDataPoolsClient) listCreateRequest(ctx context.Context, options *BigDataPoolsListOptions) (*policy.Request, error) {
	urlPath := "/bigDataPools"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *bigDataPoolsClient) listHandleResponse(resp *http.Response) (BigDataPoolsListResponse, error) {
	result := BigDataPoolsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BigDataPoolResourceInfoListResult); err != nil {
		return BigDataPoolsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *bigDataPoolsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
