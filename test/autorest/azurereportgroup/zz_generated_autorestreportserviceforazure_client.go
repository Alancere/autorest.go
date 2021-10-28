//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AutoRestReportServiceForAzureClient contains the methods for the AutoRestReportServiceForAzure group.
// Don't use this type directly, use NewAutoRestReportServiceForAzureClient() instead.
type AutoRestReportServiceForAzureClient struct {
	con *Connection
}

// NewAutoRestReportServiceForAzureClient creates a new instance of AutoRestReportServiceForAzureClient with the specified values.
func NewAutoRestReportServiceForAzureClient(con *Connection) *AutoRestReportServiceForAzureClient {
	return &AutoRestReportServiceForAzureClient{con: con}
}

// GetReport - Get test coverage report
// If the operation fails it returns the *Error error type.
func (client *AutoRestReportServiceForAzureClient) GetReport(ctx context.Context, options *AutoRestReportServiceForAzureGetReportOptions) (AutoRestReportServiceForAzureGetReportResponse, error) {
	req, err := client.getReportCreateRequest(ctx, options)
	if err != nil {
		return AutoRestReportServiceForAzureGetReportResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AutoRestReportServiceForAzureGetReportResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoRestReportServiceForAzureGetReportResponse{}, client.getReportHandleError(resp)
	}
	return client.getReportHandleResponse(resp)
}

// getReportCreateRequest creates the GetReport request.
func (client *AutoRestReportServiceForAzureClient) getReportCreateRequest(ctx context.Context, options *AutoRestReportServiceForAzureGetReportOptions) (*policy.Request, error) {
	urlPath := "/report/azure"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Qualifier != nil {
		reqQP.Set("qualifier", *options.Qualifier)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getReportHandleResponse handles the GetReport response.
func (client *AutoRestReportServiceForAzureClient) getReportHandleResponse(resp *http.Response) (AutoRestReportServiceForAzureGetReportResponse, error) {
	result := AutoRestReportServiceForAzureGetReportResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return AutoRestReportServiceForAzureGetReportResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getReportHandleError handles the GetReport error response.
func (client *AutoRestReportServiceForAzureClient) getReportHandleError(resp *http.Response) error {
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
