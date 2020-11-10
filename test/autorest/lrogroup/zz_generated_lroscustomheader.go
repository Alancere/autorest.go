// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// LrOSCustomHeaderOperations contains the methods for the LrOSCustomHeader group.
type LrOSCustomHeaderOperations interface {
	// BeginPost202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request,
	// service returns a 202 to the initial request, with 'Location' and
	// 'Retry-After' headers, Polls return a 200 with a response body after success
	BeginPost202Retry200(ctx context.Context, options *LrOSCustomHeaderPost202Retry200Options) (*HTTPPollerResponse, error)
	// ResumePost202Retry200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePost202Retry200(token string) (HTTPPoller, error)
	// BeginPostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
	// post request, service returns a 202 to the initial request, with an entity that
	// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPostAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*HTTPPollerResponse, error)
	// ResumePostAsyncRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePostAsyncRetrySucceeded(token string) (HTTPPoller, error)
	// BeginPut201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
	// put request, service returns a 201 to the initial request, with an entity that
	// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginPut201CreatingSucceeded200(ctx context.Context, options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*ProductPollerResponse, error)
	// ResumePut201CreatingSucceeded200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePut201CreatingSucceeded200(token string) (ProductPoller, error)
	// BeginPutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
	// put request, service returns a 200 to the initial request, with an entity that
	// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPutAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*ProductPollerResponse, error)
	// ResumePutAsyncRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePutAsyncRetrySucceeded(token string) (ProductPoller, error)
}

// LrOSCustomHeaderClient implements the LrOSCustomHeaderOperations interface.
// Don't use this type directly, use NewLrOSCustomHeaderClient() instead.
type LrOSCustomHeaderClient struct {
	con *Connection
}

// NewLrOSCustomHeaderClient creates a new instance of LrOSCustomHeaderClient with the specified values.
func NewLrOSCustomHeaderClient(con *Connection) LrOSCustomHeaderOperations {
	return &LrOSCustomHeaderClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *LrOSCustomHeaderClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *LrOSCustomHeaderClient) BeginPost202Retry200(ctx context.Context, options *LrOSCustomHeaderPost202Retry200Options) (*HTTPPollerResponse, error) {
	resp, err := client.Post202Retry200(ctx, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LrOSCustomHeaderClient.Post202Retry200", "", resp, client.Post202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LrOSCustomHeaderClient) ResumePost202Retry200(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LrOSCustomHeaderClient.Post202Retry200", token, client.Post202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Post202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request,
// service returns a 202 to the initial request, with 'Location' and
// 'Retry-After' headers, Polls return a 200 with a response body after success
func (client *LrOSCustomHeaderClient) Post202Retry200(ctx context.Context, options *LrOSCustomHeaderPost202Retry200Options) (*azcore.Response, error) {
	req, err := client.Post202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.Post202Retry200HandleError(resp)
	}
	return resp, nil
}

// Post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *LrOSCustomHeaderClient) Post202Retry200CreateRequest(ctx context.Context, options *LrOSCustomHeaderPost202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/customheader/post/202/retry/200"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Product)
	}
	return req, nil
}

// Post202Retry200HandleError handles the Post202Retry200 error response.
func (client *LrOSCustomHeaderClient) Post202Retry200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *LrOSCustomHeaderClient) BeginPostAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*HTTPPollerResponse, error) {
	resp, err := client.PostAsyncRetrySucceeded(ctx, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LrOSCustomHeaderClient.PostAsyncRetrySucceeded", "", resp, client.PostAsyncRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LrOSCustomHeaderClient) ResumePostAsyncRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LrOSCustomHeaderClient.PostAsyncRetrySucceeded", token, client.PostAsyncRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// PostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post
// request, service returns a 202 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *LrOSCustomHeaderClient) PostAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.PostAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.PostAsyncRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// PostAsyncRetrySucceededCreateRequest creates the PostAsyncRetrySucceeded request.
func (client *LrOSCustomHeaderClient) PostAsyncRetrySucceededCreateRequest(ctx context.Context, options *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/customheader/postasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Product)
	}
	return req, nil
}

// PostAsyncRetrySucceededHandleError handles the PostAsyncRetrySucceeded error response.
func (client *LrOSCustomHeaderClient) PostAsyncRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *LrOSCustomHeaderClient) BeginPut201CreatingSucceeded200(ctx context.Context, options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*ProductPollerResponse, error) {
	resp, err := client.Put201CreatingSucceeded200(ctx, options)
	if err != nil {
		return nil, err
	}
	result := &ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LrOSCustomHeaderClient.Put201CreatingSucceeded200", "", resp, client.Put201CreatingSucceeded200HandleError)
	if err != nil {
		return nil, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LrOSCustomHeaderClient) ResumePut201CreatingSucceeded200(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LrOSCustomHeaderClient.Put201CreatingSucceeded200", token, client.Put201CreatingSucceeded200HandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Put201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *LrOSCustomHeaderClient) Put201CreatingSucceeded200(ctx context.Context, options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*azcore.Response, error) {
	req, err := client.Put201CreatingSucceeded200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.Put201CreatingSucceeded200HandleError(resp)
	}
	return resp, nil
}

// Put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *LrOSCustomHeaderClient) Put201CreatingSucceeded200CreateRequest(ctx context.Context, options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*azcore.Request, error) {
	urlPath := "/lro/customheader/put/201/creating/succeeded/200"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Product)
	}
	return req, nil
}

// Put201CreatingSucceeded200HandleResponse handles the Put201CreatingSucceeded200 response.
func (client *LrOSCustomHeaderClient) Put201CreatingSucceeded200HandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// Put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client *LrOSCustomHeaderClient) Put201CreatingSucceeded200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *LrOSCustomHeaderClient) BeginPutAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*ProductPollerResponse, error) {
	resp, err := client.PutAsyncRetrySucceeded(ctx, options)
	if err != nil {
		return nil, err
	}
	result := &ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LrOSCustomHeaderClient.PutAsyncRetrySucceeded", "", resp, client.PutAsyncRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LrOSCustomHeaderClient) ResumePutAsyncRetrySucceeded(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LrOSCustomHeaderClient.PutAsyncRetrySucceeded", token, client.PutAsyncRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// PutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put
// request, service returns a 200 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *LrOSCustomHeaderClient) PutAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.PutAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutAsyncRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// PutAsyncRetrySucceededCreateRequest creates the PutAsyncRetrySucceeded request.
func (client *LrOSCustomHeaderClient) PutAsyncRetrySucceededCreateRequest(ctx context.Context, options *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/customheader/putasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Product)
	}
	return req, nil
}

// PutAsyncRetrySucceededHandleResponse handles the PutAsyncRetrySucceeded response.
func (client *LrOSCustomHeaderClient) PutAsyncRetrySucceededHandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// PutAsyncRetrySucceededHandleError handles the PutAsyncRetrySucceeded error response.
func (client *LrOSCustomHeaderClient) PutAsyncRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
