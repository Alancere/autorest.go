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

// LrOSCustomHeaderClient contains the methods for the LrOSCustomHeader group.
// Don't use this type directly, use NewLrOSCustomHeaderClient() instead.
type LrOSCustomHeaderClient struct {
	con *Connection
}

// NewLrOSCustomHeaderClient creates a new instance of LrOSCustomHeaderClient with the specified values.
func NewLrOSCustomHeaderClient(con *Connection) LrOSCustomHeaderClient {
	return LrOSCustomHeaderClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client LrOSCustomHeaderClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginPost202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request,
// service returns a 202 to the initial request, with 'Location' and
// 'Retry-After' headers, Polls return a 200 with a response body after success
func (client LrOSCustomHeaderClient) BeginPost202Retry200(ctx context.Context, options *LrOSCustomHeaderPost202Retry200Options) (HTTPPollerResponse, error) {
	resp, err := client.Post202Retry200(ctx, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LrOSCustomHeaderClient.Post202Retry200", "", resp, client.post202Retry200HandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
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

// ResumePost202Retry200 creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client LrOSCustomHeaderClient) ResumePost202Retry200(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LrOSCustomHeaderClient.Post202Retry200", token, client.post202Retry200HandleError)
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
func (client LrOSCustomHeaderClient) Post202Retry200(ctx context.Context, options *LrOSCustomHeaderPost202Retry200Options) (*azcore.Response, error) {
	req, err := client.post202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.post202Retry200HandleError(resp)
	}
	return resp, nil
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client LrOSCustomHeaderClient) post202Retry200CreateRequest(ctx context.Context, options *LrOSCustomHeaderPost202Retry200Options) (*azcore.Request, error) {
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

// post202Retry200HandleError handles the Post202Retry200 error response.
func (client LrOSCustomHeaderClient) post202Retry200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginPostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// post request, service returns a 202 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client LrOSCustomHeaderClient) BeginPostAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (HTTPPollerResponse, error) {
	resp, err := client.PostAsyncRetrySucceeded(ctx, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LrOSCustomHeaderClient.PostAsyncRetrySucceeded", "", resp, client.postAsyncRetrySucceededHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
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

// ResumePostAsyncRetrySucceeded creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client LrOSCustomHeaderClient) ResumePostAsyncRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LrOSCustomHeaderClient.PostAsyncRetrySucceeded", token, client.postAsyncRetrySucceededHandleError)
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
func (client LrOSCustomHeaderClient) PostAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.postAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.postAsyncRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// postAsyncRetrySucceededCreateRequest creates the PostAsyncRetrySucceeded request.
func (client LrOSCustomHeaderClient) postAsyncRetrySucceededCreateRequest(ctx context.Context, options *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*azcore.Request, error) {
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

// postAsyncRetrySucceededHandleError handles the PostAsyncRetrySucceeded error response.
func (client LrOSCustomHeaderClient) postAsyncRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginPut201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client LrOSCustomHeaderClient) BeginPut201CreatingSucceeded200(ctx context.Context, options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (ProductPollerResponse, error) {
	resp, err := client.Put201CreatingSucceeded200(ctx, options)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LrOSCustomHeaderClient.Put201CreatingSucceeded200", "", resp, client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePut201CreatingSucceeded200 creates a new ProductPoller from the specified resume token.
// token - The value must come from a previous call to ProductPoller.ResumeToken().
func (client LrOSCustomHeaderClient) ResumePut201CreatingSucceeded200(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LrOSCustomHeaderClient.Put201CreatingSucceeded200", token, client.put201CreatingSucceeded200HandleError)
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
func (client LrOSCustomHeaderClient) Put201CreatingSucceeded200(ctx context.Context, options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*azcore.Response, error) {
	req, err := client.put201CreatingSucceeded200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.put201CreatingSucceeded200HandleError(resp)
	}
	return resp, nil
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client LrOSCustomHeaderClient) put201CreatingSucceeded200CreateRequest(ctx context.Context, options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*azcore.Request, error) {
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

// put201CreatingSucceeded200HandleResponse handles the Put201CreatingSucceeded200 response.
func (client LrOSCustomHeaderClient) put201CreatingSucceeded200HandleResponse(resp *azcore.Response) (ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Product)
	return result, err
}

// put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client LrOSCustomHeaderClient) put201CreatingSucceeded200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginPutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 200 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client LrOSCustomHeaderClient) BeginPutAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (ProductPollerResponse, error) {
	resp, err := client.PutAsyncRetrySucceeded(ctx, options)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LrOSCustomHeaderClient.PutAsyncRetrySucceeded", "", resp, client.putAsyncRetrySucceededHandleError)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePutAsyncRetrySucceeded creates a new ProductPoller from the specified resume token.
// token - The value must come from a previous call to ProductPoller.ResumeToken().
func (client LrOSCustomHeaderClient) ResumePutAsyncRetrySucceeded(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LrOSCustomHeaderClient.PutAsyncRetrySucceeded", token, client.putAsyncRetrySucceededHandleError)
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
func (client LrOSCustomHeaderClient) PutAsyncRetrySucceeded(ctx context.Context, options *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.putAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putAsyncRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// putAsyncRetrySucceededCreateRequest creates the PutAsyncRetrySucceeded request.
func (client LrOSCustomHeaderClient) putAsyncRetrySucceededCreateRequest(ctx context.Context, options *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*azcore.Request, error) {
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

// putAsyncRetrySucceededHandleResponse handles the PutAsyncRetrySucceeded response.
func (client LrOSCustomHeaderClient) putAsyncRetrySucceededHandleResponse(resp *azcore.Response) (ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Product)
	return result, err
}

// putAsyncRetrySucceededHandleError handles the PutAsyncRetrySucceeded error response.
func (client LrOSCustomHeaderClient) putAsyncRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
