// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineRunCommandsClient contains the methods for the VirtualMachineRunCommands group.
// Don't use this type directly, use NewVirtualMachineRunCommandsClient() instead.
type VirtualMachineRunCommandsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineRunCommandsClient creates a new instance of VirtualMachineRunCommandsClient with the specified values.
func NewVirtualMachineRunCommandsClient(con *armcore.Connection, subscriptionID string) *VirtualMachineRunCommandsClient {
	return &VirtualMachineRunCommandsClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets specific run command for a subscription in a location.
func (client *VirtualMachineRunCommandsClient) Get(ctx context.Context, location string, commandID string, options *VirtualMachineRunCommandsGetOptions) (RunCommandDocumentResponse, error) {
	req, err := client.getCreateRequest(ctx, location, commandID, options)
	if err != nil {
		return RunCommandDocumentResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RunCommandDocumentResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RunCommandDocumentResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineRunCommandsClient) getCreateRequest(ctx context.Context, location string, commandID string, options *VirtualMachineRunCommandsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if commandID == "" {
		return nil, errors.New("parameter commandID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commandId}", url.PathEscape(commandID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineRunCommandsClient) getHandleResponse(resp *azcore.Response) (RunCommandDocumentResponse, error) {
	var val *RunCommandDocument
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RunCommandDocumentResponse{}, err
	}
	return RunCommandDocumentResponse{RawResponse: resp.Response, RunCommandDocument: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineRunCommandsClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Lists all available run commands for a subscription in a location.
func (client *VirtualMachineRunCommandsClient) List(location string, options *VirtualMachineRunCommandsListOptions) RunCommandListResultPager {
	return &runCommandListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp RunCommandListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RunCommandListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachineRunCommandsClient) listCreateRequest(ctx context.Context, location string, options *VirtualMachineRunCommandsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineRunCommandsClient) listHandleResponse(resp *azcore.Response) (RunCommandListResultResponse, error) {
	var val *RunCommandListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RunCommandListResultResponse{}, err
	}
	return RunCommandListResultResponse{RawResponse: resp.Response, RunCommandListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineRunCommandsClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
