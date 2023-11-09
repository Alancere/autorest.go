//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SparkConfigurationClient contains the methods for the SparkConfiguration group.
// Don't use this type directly, use a constructor function instead.
type SparkConfigurationClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateOrUpdateSparkConfiguration - Creates or updates a sparkconfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - sparkConfigurationName - The spark Configuration name.
//   - sparkConfiguration - SparkConfiguration resource definition.
//   - options - SparkConfigurationClientBeginCreateOrUpdateSparkConfigurationOptions contains the optional parameters for the
//     SparkConfigurationClient.BeginCreateOrUpdateSparkConfiguration method.
func (client *SparkConfigurationClient) BeginCreateOrUpdateSparkConfiguration(ctx context.Context, sparkConfigurationName string, sparkConfiguration SparkConfigurationResource, options *SparkConfigurationClientBeginCreateOrUpdateSparkConfigurationOptions) (*runtime.Poller[SparkConfigurationClientCreateOrUpdateSparkConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateSparkConfiguration(ctx, sparkConfigurationName, sparkConfiguration, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SparkConfigurationClientCreateOrUpdateSparkConfigurationResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SparkConfigurationClientCreateOrUpdateSparkConfigurationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateSparkConfiguration - Creates or updates a sparkconfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
func (client *SparkConfigurationClient) createOrUpdateSparkConfiguration(ctx context.Context, sparkConfigurationName string, sparkConfiguration SparkConfigurationResource, options *SparkConfigurationClientBeginCreateOrUpdateSparkConfigurationOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateSparkConfigurationCreateRequest(ctx, sparkConfigurationName, sparkConfiguration, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateSparkConfigurationCreateRequest creates the CreateOrUpdateSparkConfiguration request.
func (client *SparkConfigurationClient) createOrUpdateSparkConfigurationCreateRequest(ctx context.Context, sparkConfigurationName string, sparkConfiguration SparkConfigurationResource, options *SparkConfigurationClientBeginCreateOrUpdateSparkConfigurationOptions) (*policy.Request, error) {
	urlPath := "/sparkconfigurations/{sparkConfigurationName}"
	if sparkConfigurationName == "" {
		return nil, errors.New("parameter sparkConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkConfigurationName}", url.PathEscape(sparkConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sparkConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeleteSparkConfiguration - Deletes a sparkConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - sparkConfigurationName - The spark Configuration name.
//   - options - SparkConfigurationClientBeginDeleteSparkConfigurationOptions contains the optional parameters for the SparkConfigurationClient.BeginDeleteSparkConfiguration
//     method.
func (client *SparkConfigurationClient) BeginDeleteSparkConfiguration(ctx context.Context, sparkConfigurationName string, options *SparkConfigurationClientBeginDeleteSparkConfigurationOptions) (*runtime.Poller[SparkConfigurationClientDeleteSparkConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteSparkConfiguration(ctx, sparkConfigurationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SparkConfigurationClientDeleteSparkConfigurationResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SparkConfigurationClientDeleteSparkConfigurationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DeleteSparkConfiguration - Deletes a sparkConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
func (client *SparkConfigurationClient) deleteSparkConfiguration(ctx context.Context, sparkConfigurationName string, options *SparkConfigurationClientBeginDeleteSparkConfigurationOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteSparkConfigurationCreateRequest(ctx, sparkConfigurationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteSparkConfigurationCreateRequest creates the DeleteSparkConfiguration request.
func (client *SparkConfigurationClient) deleteSparkConfigurationCreateRequest(ctx context.Context, sparkConfigurationName string, options *SparkConfigurationClientBeginDeleteSparkConfigurationOptions) (*policy.Request, error) {
	urlPath := "/sparkconfigurations/{sparkConfigurationName}"
	if sparkConfigurationName == "" {
		return nil, errors.New("parameter sparkConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkConfigurationName}", url.PathEscape(sparkConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetSparkConfiguration - Gets a sparkConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - sparkConfigurationName - The spark Configuration name.
//   - options - SparkConfigurationClientGetSparkConfigurationOptions contains the optional parameters for the SparkConfigurationClient.GetSparkConfiguration
//     method.
func (client *SparkConfigurationClient) GetSparkConfiguration(ctx context.Context, sparkConfigurationName string, options *SparkConfigurationClientGetSparkConfigurationOptions) (SparkConfigurationClientGetSparkConfigurationResponse, error) {
	var err error
	req, err := client.getSparkConfigurationCreateRequest(ctx, sparkConfigurationName, options)
	if err != nil {
		return SparkConfigurationClientGetSparkConfigurationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SparkConfigurationClientGetSparkConfigurationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(httpResp)
		return SparkConfigurationClientGetSparkConfigurationResponse{}, err
	}
	resp, err := client.getSparkConfigurationHandleResponse(httpResp)
	return resp, err
}

// getSparkConfigurationCreateRequest creates the GetSparkConfiguration request.
func (client *SparkConfigurationClient) getSparkConfigurationCreateRequest(ctx context.Context, sparkConfigurationName string, options *SparkConfigurationClientGetSparkConfigurationOptions) (*policy.Request, error) {
	urlPath := "/sparkconfigurations/{sparkConfigurationName}"
	if sparkConfigurationName == "" {
		return nil, errors.New("parameter sparkConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkConfigurationName}", url.PathEscape(sparkConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSparkConfigurationHandleResponse handles the GetSparkConfiguration response.
func (client *SparkConfigurationClient) getSparkConfigurationHandleResponse(resp *http.Response) (SparkConfigurationClientGetSparkConfigurationResponse, error) {
	result := SparkConfigurationClientGetSparkConfigurationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkConfigurationResource); err != nil {
		return SparkConfigurationClientGetSparkConfigurationResponse{}, err
	}
	return result, nil
}

// NewGetSparkConfigurationsByWorkspacePager - Lists sparkconfigurations.
//
// Generated from API version 2021-06-01-preview
//   - options - SparkConfigurationClientGetSparkConfigurationsByWorkspaceOptions contains the optional parameters for the SparkConfigurationClient.NewGetSparkConfigurationsByWorkspacePager
//     method.
func (client *SparkConfigurationClient) NewGetSparkConfigurationsByWorkspacePager(options *SparkConfigurationClientGetSparkConfigurationsByWorkspaceOptions) *runtime.Pager[SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse]{
		More: func(page SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse) (SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getSparkConfigurationsByWorkspaceCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse{}, err
			}
			return client.getSparkConfigurationsByWorkspaceHandleResponse(resp)
		},
	})
}

// getSparkConfigurationsByWorkspaceCreateRequest creates the GetSparkConfigurationsByWorkspace request.
func (client *SparkConfigurationClient) getSparkConfigurationsByWorkspaceCreateRequest(ctx context.Context, options *SparkConfigurationClientGetSparkConfigurationsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/sparkconfigurations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSparkConfigurationsByWorkspaceHandleResponse handles the GetSparkConfigurationsByWorkspace response.
func (client *SparkConfigurationClient) getSparkConfigurationsByWorkspaceHandleResponse(resp *http.Response) (SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse, error) {
	result := SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkConfigurationListResponse); err != nil {
		return SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameSparkConfiguration - Renames a sparkConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - sparkConfigurationName - The spark Configuration name.
//   - request - proposed new name.
//   - options - SparkConfigurationClientBeginRenameSparkConfigurationOptions contains the optional parameters for the SparkConfigurationClient.BeginRenameSparkConfiguration
//     method.
func (client *SparkConfigurationClient) BeginRenameSparkConfiguration(ctx context.Context, sparkConfigurationName string, request ArtifactRenameRequest, options *SparkConfigurationClientBeginRenameSparkConfigurationOptions) (*runtime.Poller[SparkConfigurationClientRenameSparkConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameSparkConfiguration(ctx, sparkConfigurationName, request, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SparkConfigurationClientRenameSparkConfigurationResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SparkConfigurationClientRenameSparkConfigurationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RenameSparkConfiguration - Renames a sparkConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
func (client *SparkConfigurationClient) renameSparkConfiguration(ctx context.Context, sparkConfigurationName string, request ArtifactRenameRequest, options *SparkConfigurationClientBeginRenameSparkConfigurationOptions) (*http.Response, error) {
	var err error
	req, err := client.renameSparkConfigurationCreateRequest(ctx, sparkConfigurationName, request, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// renameSparkConfigurationCreateRequest creates the RenameSparkConfiguration request.
func (client *SparkConfigurationClient) renameSparkConfigurationCreateRequest(ctx context.Context, sparkConfigurationName string, request ArtifactRenameRequest, options *SparkConfigurationClientBeginRenameSparkConfigurationOptions) (*policy.Request, error) {
	urlPath := "/sparkconfigurations/{sparkConfigurationName}/rename"
	if sparkConfigurationName == "" {
		return nil, errors.New("parameter sparkConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sparkConfigurationName}", url.PathEscape(sparkConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, request); err != nil {
		return nil, err
	}
	return req, nil
}
