// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package lrolegacygroup

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

// CreateResourcePollViaOperationLocationClient - POST to create resource, poll URL via operation-location header.
// Don't use this type directly, use [LegacyClient.NewCreateResourcePollViaOperationLocationClient] instead.
type CreateResourcePollViaOperationLocationClient struct {
	internal   *azcore.Client
	apiVersion string
}

// BeginCreateJob - Creates a Job
//   - options - CreateResourcePollViaOperationLocationClientCreateJobOptions contains the optional parameters for the CreateResourcePollViaOperationLocationClient.CreateJob
//     method.
func (client *CreateResourcePollViaOperationLocationClient) BeginCreateJob(ctx context.Context, body JobData, options *CreateResourcePollViaOperationLocationClientCreateJobOptions) (*runtime.Poller[CreateResourcePollViaOperationLocationClientCreateJobResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createJob(ctx, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[CreateResourcePollViaOperationLocationClientCreateJobResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[CreateResourcePollViaOperationLocationClientCreateJobResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateJob - Creates a Job
func (client *CreateResourcePollViaOperationLocationClient) createJob(ctx context.Context, body JobData, options *CreateResourcePollViaOperationLocationClientCreateJobOptions) (*http.Response, error) {
	var err error
	req, err := client.createJobCreateRequest(ctx, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createJobCreateRequest creates the CreateJob request.
func (client *CreateResourcePollViaOperationLocationClient) createJobCreateRequest(ctx context.Context, body JobData, options *CreateResourcePollViaOperationLocationClientCreateJobOptions) (*policy.Request, error) {
	urlPath := "/azure/core/lro/rpc/legacy/create-resource-poll-via-operation-location/jobs"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// GetJob - Poll a Job
//   - jobID - A processing job identifier.
//   - options - CreateResourcePollViaOperationLocationClientGetJobOptions contains the optional parameters for the CreateResourcePollViaOperationLocationClient.GetJob
//     method.
func (client *CreateResourcePollViaOperationLocationClient) GetJob(ctx context.Context, jobID string, options *CreateResourcePollViaOperationLocationClientGetJobOptions) (CreateResourcePollViaOperationLocationClientGetJobResponse, error) {
	var err error
	req, err := client.getJobCreateRequest(ctx, jobID, options)
	if err != nil {
		return CreateResourcePollViaOperationLocationClientGetJobResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CreateResourcePollViaOperationLocationClientGetJobResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CreateResourcePollViaOperationLocationClientGetJobResponse{}, err
	}
	resp, err := client.getJobHandleResponse(httpResp)
	return resp, err
}

// getJobCreateRequest creates the GetJob request.
func (client *CreateResourcePollViaOperationLocationClient) getJobCreateRequest(ctx context.Context, jobID string, options *CreateResourcePollViaOperationLocationClientGetJobOptions) (*policy.Request, error) {
	urlPath := "/azure/core/lro/rpc/legacy/create-resource-poll-via-operation-location/jobs/{jobId}"
	if jobID == "" {
		return nil, errors.New("parameter jobID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", client.apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getJobHandleResponse handles the GetJob response.
func (client *CreateResourcePollViaOperationLocationClient) getJobHandleResponse(resp *http.Response) (CreateResourcePollViaOperationLocationClientGetJobResponse, error) {
	result := CreateResourcePollViaOperationLocationClientGetJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobResult); err != nil {
		return CreateResourcePollViaOperationLocationClientGetJobResponse{}, err
	}
	return result, nil
}