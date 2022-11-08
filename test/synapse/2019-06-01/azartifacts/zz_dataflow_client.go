//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type dataFlowClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newDataFlowClient creates a new instance of dataFlowClient with the specified values.
//   - endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
//   - pl - the pipeline used for sending requests and handling responses.
func newDataFlowClient(endpoint string, pl runtime.Pipeline) *dataFlowClient {
	client := &dataFlowClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateDataFlow - Creates or updates a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - dataFlowName - The data flow name.
//   - dataFlow - Data flow resource definition.
//   - options - DataFlowClientBeginCreateOrUpdateDataFlowOptions contains the optional parameters for the dataFlowClient.BeginCreateOrUpdateDataFlow
//     method.
func (client *dataFlowClient) BeginCreateOrUpdateDataFlow(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowClientBeginCreateOrUpdateDataFlowOptions) (*runtime.Poller[DataFlowClientCreateOrUpdateDataFlowResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateDataFlow(ctx, dataFlowName, dataFlow, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DataFlowClientCreateOrUpdateDataFlowResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowClientCreateOrUpdateDataFlowResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdateDataFlow - Creates or updates a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *dataFlowClient) createOrUpdateDataFlow(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowClientBeginCreateOrUpdateDataFlowOptions) (*http.Response, error) {
	req, err := client.createOrUpdateDataFlowCreateRequest(ctx, dataFlowName, dataFlow, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateDataFlowCreateRequest creates the CreateOrUpdateDataFlow request.
func (client *dataFlowClient) createOrUpdateDataFlowCreateRequest(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowClientBeginCreateOrUpdateDataFlowOptions) (*policy.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, dataFlow)
}

// BeginDeleteDataFlow - Deletes a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - dataFlowName - The data flow name.
//   - options - DataFlowClientBeginDeleteDataFlowOptions contains the optional parameters for the dataFlowClient.BeginDeleteDataFlow
//     method.
func (client *dataFlowClient) BeginDeleteDataFlow(ctx context.Context, dataFlowName string, options *DataFlowClientBeginDeleteDataFlowOptions) (*runtime.Poller[DataFlowClientDeleteDataFlowResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteDataFlow(ctx, dataFlowName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DataFlowClientDeleteDataFlowResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowClientDeleteDataFlowResponse](options.ResumeToken, client.pl, nil)
	}
}

// DeleteDataFlow - Deletes a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *dataFlowClient) deleteDataFlow(ctx context.Context, dataFlowName string, options *DataFlowClientBeginDeleteDataFlowOptions) (*http.Response, error) {
	req, err := client.deleteDataFlowCreateRequest(ctx, dataFlowName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteDataFlowCreateRequest creates the DeleteDataFlow request.
func (client *dataFlowClient) deleteDataFlowCreateRequest(ctx context.Context, dataFlowName string, options *DataFlowClientBeginDeleteDataFlowOptions) (*policy.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetDataFlow - Gets a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - dataFlowName - The data flow name.
//   - options - DataFlowClientGetDataFlowOptions contains the optional parameters for the dataFlowClient.GetDataFlow method.
func (client *dataFlowClient) GetDataFlow(ctx context.Context, dataFlowName string, options *DataFlowClientGetDataFlowOptions) (DataFlowClientGetDataFlowResponse, error) {
	req, err := client.getDataFlowCreateRequest(ctx, dataFlowName, options)
	if err != nil {
		return DataFlowClientGetDataFlowResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataFlowClientGetDataFlowResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowClientGetDataFlowResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDataFlowHandleResponse(resp)
}

// getDataFlowCreateRequest creates the GetDataFlow request.
func (client *dataFlowClient) getDataFlowCreateRequest(ctx context.Context, dataFlowName string, options *DataFlowClientGetDataFlowOptions) (*policy.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDataFlowHandleResponse handles the GetDataFlow response.
func (client *dataFlowClient) getDataFlowHandleResponse(resp *http.Response) (DataFlowClientGetDataFlowResponse, error) {
	result := DataFlowClientGetDataFlowResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataFlowResource); err != nil {
		return DataFlowClientGetDataFlowResponse{}, err
	}
	return result, nil
}

// NewGetDataFlowsByWorkspacePager - Lists data flows.
//
// Generated from API version 2019-06-01-preview
//   - options - DataFlowClientGetDataFlowsByWorkspaceOptions contains the optional parameters for the dataFlowClient.NewGetDataFlowsByWorkspacePager
//     method.
func (client *dataFlowClient) NewGetDataFlowsByWorkspacePager(options *DataFlowClientGetDataFlowsByWorkspaceOptions) *runtime.Pager[DataFlowClientGetDataFlowsByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataFlowClientGetDataFlowsByWorkspaceResponse]{
		More: func(page DataFlowClientGetDataFlowsByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataFlowClientGetDataFlowsByWorkspaceResponse) (DataFlowClientGetDataFlowsByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getDataFlowsByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataFlowClientGetDataFlowsByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataFlowClientGetDataFlowsByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataFlowClientGetDataFlowsByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getDataFlowsByWorkspaceHandleResponse(resp)
		},
	})
}

// getDataFlowsByWorkspaceCreateRequest creates the GetDataFlowsByWorkspace request.
func (client *dataFlowClient) getDataFlowsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowClientGetDataFlowsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/dataflows"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDataFlowsByWorkspaceHandleResponse handles the GetDataFlowsByWorkspace response.
func (client *dataFlowClient) getDataFlowsByWorkspaceHandleResponse(resp *http.Response) (DataFlowClientGetDataFlowsByWorkspaceResponse, error) {
	result := DataFlowClientGetDataFlowsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataFlowListResponse); err != nil {
		return DataFlowClientGetDataFlowsByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameDataFlow - Renames a dataflow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - dataFlowName - The data flow name.
//   - request - proposed new name.
//   - options - DataFlowClientBeginRenameDataFlowOptions contains the optional parameters for the dataFlowClient.BeginRenameDataFlow
//     method.
func (client *dataFlowClient) BeginRenameDataFlow(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowClientBeginRenameDataFlowOptions) (*runtime.Poller[DataFlowClientRenameDataFlowResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameDataFlow(ctx, dataFlowName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DataFlowClientRenameDataFlowResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowClientRenameDataFlowResponse](options.ResumeToken, client.pl, nil)
	}
}

// RenameDataFlow - Renames a dataflow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *dataFlowClient) renameDataFlow(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowClientBeginRenameDataFlowOptions) (*http.Response, error) {
	req, err := client.renameDataFlowCreateRequest(ctx, dataFlowName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// renameDataFlowCreateRequest creates the RenameDataFlow request.
func (client *dataFlowClient) renameDataFlowCreateRequest(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowClientBeginRenameDataFlowOptions) (*policy.Request, error) {
	urlPath := "/dataflows/{dataFlowName}/rename"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}
