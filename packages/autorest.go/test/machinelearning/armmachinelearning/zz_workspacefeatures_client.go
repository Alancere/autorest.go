//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspaceFeaturesClient contains the methods for the WorkspaceFeatures group.
// Don't use this type directly, use NewWorkspaceFeaturesClient() instead.
type WorkspaceFeaturesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceFeaturesClient creates a new instance of WorkspaceFeaturesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceFeaturesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceFeaturesClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspaceFeaturesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceFeaturesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists all enabled features for a workspace
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - options - WorkspaceFeaturesClientListOptions contains the optional parameters for the WorkspaceFeaturesClient.NewListPager
//     method.
func (client *WorkspaceFeaturesClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceFeaturesClientListOptions) *runtime.Pager[WorkspaceFeaturesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceFeaturesClientListResponse]{
		More: func(page WorkspaceFeaturesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceFeaturesClientListResponse) (WorkspaceFeaturesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceFeaturesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkspaceFeaturesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceFeaturesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceFeaturesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceFeaturesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/features"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceFeaturesClient) listHandleResponse(resp *http.Response) (WorkspaceFeaturesClientListResponse, error) {
	result := WorkspaceFeaturesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListAmlUserFeatureResult); err != nil {
		return WorkspaceFeaturesClientListResponse{}, err
	}
	return result, nil
}
