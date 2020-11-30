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

// LroRetrysClient contains the methods for the LroRetrys group.
// Don't use this type directly, use NewLroRetrysClient() instead.
type LroRetrysClient struct {
	con *Connection
}

// NewLroRetrysClient creates a new instance of LroRetrysClient with the specified values.
func NewLroRetrysClient(con *Connection) LroRetrysClient {
	return LroRetrysClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client LroRetrysClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginDelete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last
// poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client LroRetrysClient) BeginDelete202Retry200(ctx context.Context, options *LroRetrysDelete202Retry200Options) (HTTPPollerResponse, error) {
	resp, err := client.Delete202Retry200(ctx, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LroRetrysClient.Delete202Retry200", "", resp, client.delete202Retry200HandleError)
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

// ResumeDelete202Retry200 creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client LroRetrysClient) ResumeDelete202Retry200(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.Delete202Retry200", token, client.delete202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll
// returns a ‘200’ with ProvisioningState=’Succeeded’
func (client LroRetrysClient) Delete202Retry200(ctx context.Context, options *LroRetrysDelete202Retry200Options) (*azcore.Response, error) {
	req, err := client.delete202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.delete202Retry200HandleError(resp)
	}
	return resp, nil
}

// delete202Retry200CreateRequest creates the Delete202Retry200 request.
func (client LroRetrysClient) delete202Retry200CreateRequest(ctx context.Context, options *LroRetrysDelete202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/202/retry/200"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// delete202Retry200HandleError handles the Delete202Retry200 error response.
func (client LroRetrysClient) delete202Retry200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated
// in the Azure-AsyncOperation header for operation status
func (client LroRetrysClient) BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysDeleteAsyncRelativeRetrySucceededOptions) (HTTPPollerResponse, error) {
	resp, err := client.DeleteAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LroRetrysClient.DeleteAsyncRelativeRetrySucceeded", "", resp, client.deleteAsyncRelativeRetrySucceededHandleError)
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

// ResumeDeleteAsyncRelativeRetrySucceeded creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client LroRetrysClient) ResumeDeleteAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.DeleteAsyncRelativeRetrySucceeded", token, client.deleteAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// DeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated
// in the Azure-AsyncOperation header for operation status
func (client LroRetrysClient) DeleteAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysDeleteAsyncRelativeRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.deleteAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.deleteAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// deleteAsyncRelativeRetrySucceededCreateRequest creates the DeleteAsyncRelativeRetrySucceeded request.
func (client LroRetrysClient) deleteAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LroRetrysDeleteAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/deleteasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAsyncRelativeRetrySucceededHandleError handles the DeleteAsyncRelativeRetrySucceeded error response.
func (client LroRetrysClient) deleteAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a 202 to the initial request, with an entity
// that contains ProvisioningState=’Accepted’. Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
func (client LroRetrysClient) BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context, options *LroRetrysDeleteProvisioning202Accepted200SucceededOptions) (ProductPollerResponse, error) {
	resp, err := client.DeleteProvisioning202Accepted200Succeeded(ctx, options)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LroRetrysClient.DeleteProvisioning202Accepted200Succeeded", "", resp, client.deleteProvisioning202Accepted200SucceededHandleError)
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

// ResumeDeleteProvisioning202Accepted200Succeeded creates a new ProductPoller from the specified resume token.
// token - The value must come from a previous call to ProductPoller.ResumeToken().
func (client LroRetrysClient) ResumeDeleteProvisioning202Accepted200Succeeded(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.DeleteProvisioning202Accepted200Succeeded", token, client.deleteProvisioning202Accepted200SucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// DeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a 202 to the initial request, with an entity that
// contains ProvisioningState=’Accepted’. Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
func (client LroRetrysClient) DeleteProvisioning202Accepted200Succeeded(ctx context.Context, options *LroRetrysDeleteProvisioning202Accepted200SucceededOptions) (*azcore.Response, error) {
	req, err := client.deleteProvisioning202Accepted200SucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.deleteProvisioning202Accepted200SucceededHandleError(resp)
	}
	return resp, nil
}

// deleteProvisioning202Accepted200SucceededCreateRequest creates the DeleteProvisioning202Accepted200Succeeded request.
func (client LroRetrysClient) deleteProvisioning202Accepted200SucceededCreateRequest(ctx context.Context, options *LroRetrysDeleteProvisioning202Accepted200SucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteProvisioning202Accepted200SucceededHandleResponse handles the DeleteProvisioning202Accepted200Succeeded response.
func (client LroRetrysClient) deleteProvisioning202Accepted200SucceededHandleResponse(resp *azcore.Response) (ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Product)
	return result, err
}

// deleteProvisioning202Accepted200SucceededHandleError handles the DeleteProvisioning202Accepted200Succeeded error response.
func (client LroRetrysClient) deleteProvisioning202Accepted200SucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginPost202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers,
// Polls return a 200 with a response body after success
func (client LroRetrysClient) BeginPost202Retry200(ctx context.Context, options *LroRetrysPost202Retry200Options) (HTTPPollerResponse, error) {
	resp, err := client.Post202Retry200(ctx, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LroRetrysClient.Post202Retry200", "", resp, client.post202Retry200HandleError)
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
func (client LroRetrysClient) ResumePost202Retry200(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.Post202Retry200", token, client.post202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Post202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls
// return a 200 with a response body after success
func (client LroRetrysClient) Post202Retry200(ctx context.Context, options *LroRetrysPost202Retry200Options) (*azcore.Response, error) {
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
func (client LroRetrysClient) post202Retry200CreateRequest(ctx context.Context, options *LroRetrysPost202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/post/202/retry/200"
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
func (client LroRetrysClient) post202Retry200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginPostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains
// ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
func (client LroRetrysClient) BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysPostAsyncRelativeRetrySucceededOptions) (HTTPPollerResponse, error) {
	resp, err := client.PostAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LroRetrysClient.PostAsyncRelativeRetrySucceeded", "", resp, client.postAsyncRelativeRetrySucceededHandleError)
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

// ResumePostAsyncRelativeRetrySucceeded creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client LroRetrysClient) ResumePostAsyncRelativeRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.PostAsyncRelativeRetrySucceeded", token, client.postAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// PostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
func (client LroRetrysClient) PostAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.postAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.postAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// postAsyncRelativeRetrySucceededCreateRequest creates the PostAsyncRelativeRetrySucceeded request.
func (client LroRetrysClient) postAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/postasync/retry/succeeded"
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

// postAsyncRelativeRetrySucceededHandleError handles the PostAsyncRelativeRetrySucceeded error response.
func (client LroRetrysClient) postAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginPut201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
func (client LroRetrysClient) BeginPut201CreatingSucceeded200(ctx context.Context, options *LroRetrysPut201CreatingSucceeded200Options) (ProductPollerResponse, error) {
	resp, err := client.Put201CreatingSucceeded200(ctx, options)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LroRetrysClient.Put201CreatingSucceeded200", "", resp, client.put201CreatingSucceeded200HandleError)
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
func (client LroRetrysClient) ResumePut201CreatingSucceeded200(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.Put201CreatingSucceeded200", token, client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Put201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
func (client LroRetrysClient) Put201CreatingSucceeded200(ctx context.Context, options *LroRetrysPut201CreatingSucceeded200Options) (*azcore.Response, error) {
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
func (client LroRetrysClient) put201CreatingSucceeded200CreateRequest(ctx context.Context, options *LroRetrysPut201CreatingSucceeded200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/put/201/creating/succeeded/200"
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
func (client LroRetrysClient) put201CreatingSucceeded200HandleResponse(resp *azcore.Response) (ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Product)
	return result, err
}

// put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client LroRetrysClient) put201CreatingSucceeded200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginPutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains
// ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
func (client LroRetrysClient) BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysPutAsyncRelativeRetrySucceededOptions) (ProductPollerResponse, error) {
	resp, err := client.PutAsyncRelativeRetrySucceeded(ctx, options)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LroRetrysClient.PutAsyncRelativeRetrySucceeded", "", resp, client.putAsyncRelativeRetrySucceededHandleError)
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

// ResumePutAsyncRelativeRetrySucceeded creates a new ProductPoller from the specified resume token.
// token - The value must come from a previous call to ProductPoller.ResumeToken().
func (client LroRetrysClient) ResumePutAsyncRelativeRetrySucceeded(token string) (ProductPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LroRetrysClient.PutAsyncRelativeRetrySucceeded", token, client.putAsyncRelativeRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// PutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’.
// Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
func (client LroRetrysClient) PutAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.putAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putAsyncRelativeRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// putAsyncRelativeRetrySucceededCreateRequest creates the PutAsyncRelativeRetrySucceeded request.
func (client LroRetrysClient) putAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/putasync/retry/succeeded"
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

// putAsyncRelativeRetrySucceededHandleResponse handles the PutAsyncRelativeRetrySucceeded response.
func (client LroRetrysClient) putAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (ProductResponse, error) {
	result := ProductResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Product)
	return result, err
}

// putAsyncRelativeRetrySucceededHandleError handles the PutAsyncRelativeRetrySucceeded error response.
func (client LroRetrysClient) putAsyncRelativeRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
