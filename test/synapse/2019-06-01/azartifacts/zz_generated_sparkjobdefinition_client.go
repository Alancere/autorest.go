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

type sparkJobDefinitionClient struct {
	con *connection
}

// BeginCreateOrUpdateSparkJobDefinition - Creates or updates a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginCreateOrUpdateSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionBeginCreateOrUpdateSparkJobDefinitionOptions) (SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse, error) {
	resp, err := client.createOrUpdateSparkJobDefinition(ctx, sparkJobDefinitionName, sparkJobDefinition, options)
	if err != nil {
		return SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("sparkJobDefinitionClient.CreateOrUpdateSparkJobDefinition", resp, client.con.Pipeline(), client.createOrUpdateSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPollerResponse{}, err
	}
	result.Poller = &SparkJobDefinitionCreateOrUpdateSparkJobDefinitionPoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdateSparkJobDefinition - Creates or updates a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionBeginCreateOrUpdateSparkJobDefinitionOptions) (*http.Response, error) {
	req, err := client.createOrUpdateSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, sparkJobDefinition, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateSparkJobDefinitionCreateRequest creates the CreateOrUpdateSparkJobDefinition request.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionBeginCreateOrUpdateSparkJobDefinitionOptions) (*policy.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sparkJobDefinition)
}

// createOrUpdateSparkJobDefinitionHandleError handles the CreateOrUpdateSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinitionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDebugSparkJobDefinition - Debug the spark job definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginDebugSparkJobDefinition(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionBeginDebugSparkJobDefinitionOptions) (SparkJobDefinitionDebugSparkJobDefinitionPollerResponse, error) {
	resp, err := client.debugSparkJobDefinition(ctx, sparkJobDefinitionAzureResource, options)
	if err != nil {
		return SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("sparkJobDefinitionClient.DebugSparkJobDefinition", resp, client.con.Pipeline(), client.debugSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionDebugSparkJobDefinitionPollerResponse{}, err
	}
	result.Poller = &SparkJobDefinitionDebugSparkJobDefinitionPoller{
		pt: pt,
	}
	return result, nil
}

// DebugSparkJobDefinition - Debug the spark job definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) debugSparkJobDefinition(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionBeginDebugSparkJobDefinitionOptions) (*http.Response, error) {
	req, err := client.debugSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionAzureResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.debugSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// debugSparkJobDefinitionCreateRequest creates the DebugSparkJobDefinition request.
func (client *sparkJobDefinitionClient) debugSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionBeginDebugSparkJobDefinitionOptions) (*policy.Request, error) {
	urlPath := "/debugSparkJobDefinition"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sparkJobDefinitionAzureResource)
}

// debugSparkJobDefinitionHandleError handles the DebugSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) debugSparkJobDefinitionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDeleteSparkJobDefinition - Deletes a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginDeleteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginDeleteSparkJobDefinitionOptions) (SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse, error) {
	resp, err := client.deleteSparkJobDefinition(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("sparkJobDefinitionClient.DeleteSparkJobDefinition", resp, client.con.Pipeline(), client.deleteSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionDeleteSparkJobDefinitionPollerResponse{}, err
	}
	result.Poller = &SparkJobDefinitionDeleteSparkJobDefinitionPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteSparkJobDefinition - Deletes a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginDeleteSparkJobDefinitionOptions) (*http.Response, error) {
	req, err := client.deleteSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// deleteSparkJobDefinitionCreateRequest creates the DeleteSparkJobDefinition request.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginDeleteSparkJobDefinitionOptions) (*policy.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteSparkJobDefinitionHandleError handles the DeleteSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinitionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginExecuteSparkJobDefinition - Executes the spark job definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginExecuteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginExecuteSparkJobDefinitionOptions) (SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse, error) {
	resp, err := client.executeSparkJobDefinition(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("sparkJobDefinitionClient.ExecuteSparkJobDefinition", resp, client.con.Pipeline(), client.executeSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionExecuteSparkJobDefinitionPollerResponse{}, err
	}
	result.Poller = &SparkJobDefinitionExecuteSparkJobDefinitionPoller{
		pt: pt,
	}
	return result, nil
}

// ExecuteSparkJobDefinition - Executes the spark job definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) executeSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginExecuteSparkJobDefinitionOptions) (*http.Response, error) {
	req, err := client.executeSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.executeSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// executeSparkJobDefinitionCreateRequest creates the ExecuteSparkJobDefinition request.
func (client *sparkJobDefinitionClient) executeSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginExecuteSparkJobDefinitionOptions) (*policy.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}/execute"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// executeSparkJobDefinitionHandleError handles the ExecuteSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) executeSparkJobDefinitionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetSparkJobDefinition - Gets a Spark Job Definition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) GetSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionGetSparkJobDefinitionOptions) (SparkJobDefinitionGetSparkJobDefinitionResponse, error) {
	req, err := client.getSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return SparkJobDefinitionGetSparkJobDefinitionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkJobDefinitionGetSparkJobDefinitionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return SparkJobDefinitionGetSparkJobDefinitionResponse{}, client.getSparkJobDefinitionHandleError(resp)
	}
	return client.getSparkJobDefinitionHandleResponse(resp)
}

// getSparkJobDefinitionCreateRequest creates the GetSparkJobDefinition request.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionGetSparkJobDefinitionOptions) (*policy.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkJobDefinitionHandleResponse handles the GetSparkJobDefinition response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionHandleResponse(resp *http.Response) (SparkJobDefinitionGetSparkJobDefinitionResponse, error) {
	result := SparkJobDefinitionGetSparkJobDefinitionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkJobDefinitionResource); err != nil {
		return SparkJobDefinitionGetSparkJobDefinitionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getSparkJobDefinitionHandleError handles the GetSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetSparkJobDefinitionsByWorkspace - Lists spark job definitions.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspace(options *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOptions) *SparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager {
	return &SparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getSparkJobDefinitionsByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SparkJobDefinitionsListResponse.NextLink)
		},
	}
}

// getSparkJobDefinitionsByWorkspaceCreateRequest creates the GetSparkJobDefinitionsByWorkspace request.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceCreateRequest(ctx context.Context, options *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/sparkJobDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkJobDefinitionsByWorkspaceHandleResponse handles the GetSparkJobDefinitionsByWorkspace response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceHandleResponse(resp *http.Response) (SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse, error) {
	result := SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkJobDefinitionsListResponse); err != nil {
		return SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getSparkJobDefinitionsByWorkspaceHandleError handles the GetSparkJobDefinitionsByWorkspace error response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRenameSparkJobDefinition - Renames a sparkJobDefinition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) BeginRenameSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, request ArtifactRenameRequest, options *SparkJobDefinitionBeginRenameSparkJobDefinitionOptions) (SparkJobDefinitionRenameSparkJobDefinitionPollerResponse, error) {
	resp, err := client.renameSparkJobDefinition(ctx, sparkJobDefinitionName, request, options)
	if err != nil {
		return SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{}, err
	}
	result := SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("sparkJobDefinitionClient.RenameSparkJobDefinition", resp, client.con.Pipeline(), client.renameSparkJobDefinitionHandleError)
	if err != nil {
		return SparkJobDefinitionRenameSparkJobDefinitionPollerResponse{}, err
	}
	result.Poller = &SparkJobDefinitionRenameSparkJobDefinitionPoller{
		pt: pt,
	}
	return result, nil
}

// RenameSparkJobDefinition - Renames a sparkJobDefinition.
// If the operation fails it returns the *CloudError error type.
func (client *sparkJobDefinitionClient) renameSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, request ArtifactRenameRequest, options *SparkJobDefinitionBeginRenameSparkJobDefinitionOptions) (*http.Response, error) {
	req, err := client.renameSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.renameSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// renameSparkJobDefinitionCreateRequest creates the RenameSparkJobDefinition request.
func (client *sparkJobDefinitionClient) renameSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, request ArtifactRenameRequest, options *SparkJobDefinitionBeginRenameSparkJobDefinitionOptions) (*policy.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}/rename"
	if sparkJobDefinitionName == "" {
		return nil, errors.New("parameter sparkJobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// renameSparkJobDefinitionHandleError handles the RenameSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) renameSparkJobDefinitionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
