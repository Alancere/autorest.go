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

// LegacyCreateResourcePollViaOperationLocationClient - POST to create resource, poll URL via operation-location header.
// Don't use this type directly, use [LegacyClient.NewLegacyCreateResourcePollViaOperationLocationClient] instead.
type LegacyCreateResourcePollViaOperationLocationClient struct {
	internal *azcore.Client
}

// BeginCreateJob - Creates a Job
//   - options - LegacyCreateResourcePollViaOperationLocationClientBeginCreateJobOptions contains the optional parameters for
//     the LegacyCreateResourcePollViaOperationLocationClient.CreateJob method.
func (client *LegacyCreateResourcePollViaOperationLocationClient) BeginCreateJob(ctx context.Context, jobData JobData, options *LegacyCreateResourcePollViaOperationLocationClientBeginCreateJobOptions) (*runtime.Poller[LegacyCreateResourcePollViaOperationLocationClientCreateJobResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createJob(ctx, jobData, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LegacyCreateResourcePollViaOperationLocationClientCreateJobResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LegacyCreateResourcePollViaOperationLocationClientCreateJobResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateJob - Creates a Job
func (client *LegacyCreateResourcePollViaOperationLocationClient) createJob(ctx context.Context, jobData JobData, options *LegacyCreateResourcePollViaOperationLocationClientBeginCreateJobOptions) (*http.Response, error) {
	var err error
	const operationName = "LegacyCreateResourcePollViaOperationLocationClient.BeginCreateJob"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createJobCreateRequest(ctx, jobData, options)
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
func (client *LegacyCreateResourcePollViaOperationLocationClient) createJobCreateRequest(ctx context.Context, jobData JobData, options *LegacyCreateResourcePollViaOperationLocationClientBeginCreateJobOptions) (*policy.Request, error) {
	urlPath := "/azure/core/lro/rpc/legacy/create-resource-poll-via-operation-location/jobs"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, jobData); err != nil {
		return nil, err
	}
	return req, nil
}

// GetJob - Poll a Job
//   - jobID - A processing job identifier.
//   - options - LegacyCreateResourcePollViaOperationLocationClientGetJobOptions contains the optional parameters for the LegacyCreateResourcePollViaOperationLocationClient.GetJob
//     method.
func (client *LegacyCreateResourcePollViaOperationLocationClient) GetJob(ctx context.Context, jobID string, options *LegacyCreateResourcePollViaOperationLocationClientGetJobOptions) (LegacyCreateResourcePollViaOperationLocationClientGetJobResponse, error) {
	var err error
	const operationName = "LegacyCreateResourcePollViaOperationLocationClient.GetJob"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getJobCreateRequest(ctx, jobID, options)
	if err != nil {
		return LegacyCreateResourcePollViaOperationLocationClientGetJobResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LegacyCreateResourcePollViaOperationLocationClientGetJobResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LegacyCreateResourcePollViaOperationLocationClientGetJobResponse{}, err
	}
	resp, err := client.getJobHandleResponse(httpResp)
	return resp, err
}

// getJobCreateRequest creates the GetJob request.
func (client *LegacyCreateResourcePollViaOperationLocationClient) getJobCreateRequest(ctx context.Context, jobID string, options *LegacyCreateResourcePollViaOperationLocationClientGetJobOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getJobHandleResponse handles the GetJob response.
func (client *LegacyCreateResourcePollViaOperationLocationClient) getJobHandleResponse(resp *http.Response) (LegacyCreateResourcePollViaOperationLocationClientGetJobResponse, error) {
	result := LegacyCreateResourcePollViaOperationLocationClientGetJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobResult); err != nil {
		return LegacyCreateResourcePollViaOperationLocationClientGetJobResponse{}, err
	}
	return result, nil
}
