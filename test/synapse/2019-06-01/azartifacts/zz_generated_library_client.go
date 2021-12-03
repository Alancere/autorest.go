//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type libraryClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newLibraryClient creates a new instance of libraryClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newLibraryClient(endpoint string, pl runtime.Pipeline) *libraryClient {
	client := &libraryClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Append - Append the content to the library resource created using the create operation. The maximum content size is 4MiB.
// Content larger than 4MiB must be appended in 4MiB chunks
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// content - Library file chunk.
// options - LibraryAppendOptions contains the optional parameters for the libraryClient.Append method.
func (client *libraryClient) Append(ctx context.Context, libraryName string, content io.ReadSeekCloser, options *LibraryAppendOptions) (LibraryAppendResponse, error) {
	req, err := client.appendCreateRequest(ctx, libraryName, content, options)
	if err != nil {
		return LibraryAppendResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LibraryAppendResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return LibraryAppendResponse{}, client.appendHandleError(resp)
	}
	return LibraryAppendResponse{RawResponse: resp}, nil
}

// appendCreateRequest creates the Append request.
func (client *libraryClient) appendCreateRequest(ctx context.Context, libraryName string, content io.ReadSeekCloser, options *LibraryAppendOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.XMSBlobConditionAppendpos != nil {
		req.Raw().Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*options.XMSBlobConditionAppendpos, 10))
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, req.SetBody(content, "application/octet-stream")
}

// appendHandleError handles the Append error response.
func (client *libraryClient) appendHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreate - Creates a library with the library name.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// options - LibraryBeginCreateOptions contains the optional parameters for the libraryClient.BeginCreate method.
func (client *libraryClient) BeginCreate(ctx context.Context, libraryName string, options *LibraryBeginCreateOptions) (LibraryCreatePollerResponse, error) {
	resp, err := client.create(ctx, libraryName, options)
	if err != nil {
		return LibraryCreatePollerResponse{}, err
	}
	result := LibraryCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("libraryClient.Create", resp, client.pl, client.createHandleError)
	if err != nil {
		return LibraryCreatePollerResponse{}, err
	}
	result.Poller = &LibraryCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a library with the library name.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *libraryClient) create(ctx context.Context, libraryName string, options *LibraryBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *libraryClient) createCreateRequest(ctx context.Context, libraryName string, options *LibraryBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleError handles the Create error response.
func (client *libraryClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete Library
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// options - LibraryBeginDeleteOptions contains the optional parameters for the libraryClient.BeginDelete method.
func (client *libraryClient) BeginDelete(ctx context.Context, libraryName string, options *LibraryBeginDeleteOptions) (LibraryDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, libraryName, options)
	if err != nil {
		return LibraryDeletePollerResponse{}, err
	}
	result := LibraryDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("libraryClient.Delete", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return LibraryDeletePollerResponse{}, err
	}
	result.Poller = &LibraryDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete Library
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *libraryClient) deleteOperation(ctx context.Context, libraryName string, options *LibraryBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusConflict) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *libraryClient) deleteCreateRequest(ctx context.Context, libraryName string, options *LibraryBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *libraryClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginFlush - Flush Library
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// options - LibraryBeginFlushOptions contains the optional parameters for the libraryClient.BeginFlush method.
func (client *libraryClient) BeginFlush(ctx context.Context, libraryName string, options *LibraryBeginFlushOptions) (LibraryFlushPollerResponse, error) {
	resp, err := client.flush(ctx, libraryName, options)
	if err != nil {
		return LibraryFlushPollerResponse{}, err
	}
	result := LibraryFlushPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("libraryClient.Flush", resp, client.pl, client.flushHandleError)
	if err != nil {
		return LibraryFlushPollerResponse{}, err
	}
	result.Poller = &LibraryFlushPoller{
		pt: pt,
	}
	return result, nil
}

// Flush - Flush Library
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *libraryClient) flush(ctx context.Context, libraryName string, options *LibraryBeginFlushOptions) (*http.Response, error) {
	req, err := client.flushCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.flushHandleError(resp)
	}
	return resp, nil
}

// flushCreateRequest creates the Flush request.
func (client *libraryClient) flushCreateRequest(ctx context.Context, libraryName string, options *LibraryBeginFlushOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}/flush"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// flushHandleError handles the Flush error response.
func (client *libraryClient) flushHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get Library
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
// libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
// options - LibraryGetOptions contains the optional parameters for the libraryClient.Get method.
func (client *libraryClient) Get(ctx context.Context, libraryName string, options *LibraryGetOptions) (LibraryGetResponse, error) {
	req, err := client.getCreateRequest(ctx, libraryName, options)
	if err != nil {
		return LibraryGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LibraryGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return LibraryGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *libraryClient) getCreateRequest(ctx context.Context, libraryName string, options *LibraryGetOptions) (*policy.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *libraryClient) getHandleResponse(resp *http.Response) (LibraryGetResponse, error) {
	result := LibraryGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryResource); err != nil {
		return LibraryGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *libraryClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetOperationResult - Get Operation result for Library
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
// operationID - operation id for which status is requested
// options - LibraryGetOperationResultOptions contains the optional parameters for the libraryClient.GetOperationResult method.
func (client *libraryClient) GetOperationResult(ctx context.Context, operationID string, options *LibraryGetOperationResultOptions) (LibraryGetOperationResultResponse, error) {
	req, err := client.getOperationResultCreateRequest(ctx, operationID, options)
	if err != nil {
		return LibraryGetOperationResultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LibraryGetOperationResultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return LibraryGetOperationResultResponse{}, client.getOperationResultHandleError(resp)
	}
	return client.getOperationResultHandleResponse(resp)
}

// getOperationResultCreateRequest creates the GetOperationResult request.
func (client *libraryClient) getOperationResultCreateRequest(ctx context.Context, operationID string, options *LibraryGetOperationResultOptions) (*policy.Request, error) {
	urlPath := "/libraryOperationResults/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOperationResultHandleResponse handles the GetOperationResult response.
func (client *libraryClient) getOperationResultHandleResponse(resp *http.Response) (LibraryGetOperationResultResponse, error) {
	result := LibraryGetOperationResultResponse{RawResponse: resp}
	switch resp.StatusCode {
	case http.StatusOK:
		var val LibraryResource
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return LibraryGetOperationResultResponse{}, runtime.NewResponseError(err, resp)
		}
		result.Value = val
	case http.StatusAccepted:
		var val OperationResult
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return LibraryGetOperationResultResponse{}, runtime.NewResponseError(err, resp)
		}
		result.Value = val
	default:
		return LibraryGetOperationResultResponse{}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)
	}
	return result, nil
}

// getOperationResultHandleError handles the GetOperationResult error response.
func (client *libraryClient) getOperationResultHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists Library.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
// options - LibraryListOptions contains the optional parameters for the libraryClient.List method.
func (client *libraryClient) List(options *LibraryListOptions) *LibraryListPager {
	return &LibraryListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp LibraryListResponseEnvelope) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LibraryListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *libraryClient) listCreateRequest(ctx context.Context, options *LibraryListOptions) (*policy.Request, error) {
	urlPath := "/libraries"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *libraryClient) listHandleResponse(resp *http.Response) (LibraryListResponseEnvelope, error) {
	result := LibraryListResponseEnvelope{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryListResponse); err != nil {
		return LibraryListResponseEnvelope{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *libraryClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
