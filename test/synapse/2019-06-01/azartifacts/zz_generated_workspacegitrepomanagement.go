// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

type workspaceGitRepoManagementClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *workspaceGitRepoManagementClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetGitHubAccessToken - Get the GitHub access token.
func (client *workspaceGitRepoManagementClient) GetGitHubAccessToken(ctx context.Context, gitHubAccessTokenRequest GitHubAccessTokenRequest, options *WorkspaceGitRepoManagementGetGitHubAccessTokenOptions) (*GitHubAccessTokenResponseResponse, error) {
	req, err := client.GetGitHubAccessTokenCreateRequest(ctx, gitHubAccessTokenRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetGitHubAccessTokenHandleError(resp)
	}
	result, err := client.GetGitHubAccessTokenHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetGitHubAccessTokenCreateRequest creates the GetGitHubAccessToken request.
func (client *workspaceGitRepoManagementClient) GetGitHubAccessTokenCreateRequest(ctx context.Context, gitHubAccessTokenRequest GitHubAccessTokenRequest, options *WorkspaceGitRepoManagementGetGitHubAccessTokenOptions) (*azcore.Request, error) {
	urlPath := "/getGitHubAccessToken"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.ClientRequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.ClientRequestId)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(gitHubAccessTokenRequest)
}

// GetGitHubAccessTokenHandleResponse handles the GetGitHubAccessToken response.
func (client *workspaceGitRepoManagementClient) GetGitHubAccessTokenHandleResponse(resp *azcore.Response) (*GitHubAccessTokenResponseResponse, error) {
	result := GitHubAccessTokenResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GitHubAccessTokenResponse)
}

// GetGitHubAccessTokenHandleError handles the GetGitHubAccessToken error response.
func (client *workspaceGitRepoManagementClient) GetGitHubAccessTokenHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
