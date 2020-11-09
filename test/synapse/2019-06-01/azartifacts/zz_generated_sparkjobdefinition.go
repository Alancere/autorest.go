// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type sparkJobDefinitionClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *sparkJobDefinitionClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateOrUpdateSparkJobDefinition - Creates or updates a Spark Job Definition.
func (client *sparkJobDefinitionClient) CreateOrUpdateSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionCreateOrUpdateSparkJobDefinitionOptions) (*SparkJobDefinitionResourceResponse, error) {
	req, err := client.CreateOrUpdateSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, sparkJobDefinition, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateOrUpdateSparkJobDefinitionHandleError(resp)
	}
	result, err := client.CreateOrUpdateSparkJobDefinitionHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateSparkJobDefinitionCreateRequest creates the CreateOrUpdateSparkJobDefinition request.
func (client *sparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionCreateOrUpdateSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkJobDefinition)
}

// CreateOrUpdateSparkJobDefinitionHandleResponse handles the CreateOrUpdateSparkJobDefinition response.
func (client *sparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionHandleResponse(resp *azcore.Response) (*SparkJobDefinitionResourceResponse, error) {
	result := SparkJobDefinitionResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SparkJobDefinitionResource)
}

// CreateOrUpdateSparkJobDefinitionHandleError handles the CreateOrUpdateSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) CreateOrUpdateSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DebugSparkJobDefinition - Debug the spark job definition.
func (client *sparkJobDefinitionClient) DebugSparkJobDefinition(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionDebugSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.DebugSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionAzureResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.DebugSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// DebugSparkJobDefinitionCreateRequest creates the DebugSparkJobDefinition request.
func (client *sparkJobDefinitionClient) DebugSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionDebugSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/debugSparkJobDefinition"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkJobDefinitionAzureResource)
}

// DebugSparkJobDefinitionHandleResponse handles the DebugSparkJobDefinition response.
func (client *sparkJobDefinitionClient) DebugSparkJobDefinitionHandleResponse(resp *azcore.Response) (*SparkBatchJobResponse, error) {
	result := SparkBatchJobResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SparkBatchJob)
}

// DebugSparkJobDefinitionHandleError handles the DebugSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) DebugSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteSparkJobDefinition - Deletes a Spark Job Definition.
func (client *sparkJobDefinitionClient) DeleteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionDeleteSparkJobDefinitionOptions) (*http.Response, error) {
	req, err := client.DeleteSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.DeleteSparkJobDefinitionHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteSparkJobDefinitionCreateRequest creates the DeleteSparkJobDefinition request.
func (client *sparkJobDefinitionClient) DeleteSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionDeleteSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteSparkJobDefinitionHandleError handles the DeleteSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) DeleteSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ExecuteSparkJobDefinition - Executes the spark job definition.
func (client *sparkJobDefinitionClient) ExecuteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionExecuteSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.ExecuteSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.ExecuteSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// ExecuteSparkJobDefinitionCreateRequest creates the ExecuteSparkJobDefinition request.
func (client *sparkJobDefinitionClient) ExecuteSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionExecuteSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}/execute"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ExecuteSparkJobDefinitionHandleResponse handles the ExecuteSparkJobDefinition response.
func (client *sparkJobDefinitionClient) ExecuteSparkJobDefinitionHandleResponse(resp *azcore.Response) (*SparkBatchJobResponse, error) {
	result := SparkBatchJobResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SparkBatchJob)
}

// ExecuteSparkJobDefinitionHandleError handles the ExecuteSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) ExecuteSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSparkJobDefinition - Gets a Spark Job Definition.
func (client *sparkJobDefinitionClient) GetSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionGetSparkJobDefinitionOptions) (*SparkJobDefinitionResourceResponse, error) {
	req, err := client.GetSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return nil, client.GetSparkJobDefinitionHandleError(resp)
	}
	result, err := client.GetSparkJobDefinitionHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSparkJobDefinitionCreateRequest creates the GetSparkJobDefinition request.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionGetSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetSparkJobDefinitionHandleResponse handles the GetSparkJobDefinition response.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionHandleResponse(resp *azcore.Response) (*SparkJobDefinitionResourceResponse, error) {
	result := SparkJobDefinitionResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SparkJobDefinitionResource)
}

// GetSparkJobDefinitionHandleError handles the GetSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSparkJobDefinitionsByWorkspace - Lists spark job definitions.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspace(options *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOptions) SparkJobDefinitionsListResponsePager {
	return &sparkJobDefinitionsListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetSparkJobDefinitionsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.GetSparkJobDefinitionsByWorkspaceHandleResponse,
		errorer:   client.GetSparkJobDefinitionsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *SparkJobDefinitionsListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SparkJobDefinitionsListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// GetSparkJobDefinitionsByWorkspaceCreateRequest creates the GetSparkJobDefinitionsByWorkspace request.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceCreateRequest(ctx context.Context, options *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetSparkJobDefinitionsByWorkspaceHandleResponse handles the GetSparkJobDefinitionsByWorkspace response.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceHandleResponse(resp *azcore.Response) (*SparkJobDefinitionsListResponseResponse, error) {
	result := SparkJobDefinitionsListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SparkJobDefinitionsListResponse)
}

// GetSparkJobDefinitionsByWorkspaceHandleError handles the GetSparkJobDefinitionsByWorkspace error response.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
