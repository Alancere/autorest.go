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
//   - endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
//   - pl - the pipeline used for sending requests and handling responses.
func newLibraryClient(endpoint string, pl runtime.Pipeline) *libraryClient {
	client := &libraryClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Append - Append the content to the library resource created using the create operation. The maximum content size is 4MiB.
// Content larger than 4MiB must be appended in 4MiB chunks
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
//   - content - Library file chunk.
//   - options - LibraryClientAppendOptions contains the optional parameters for the libraryClient.Append method.
func (client *libraryClient) Append(ctx context.Context, libraryName string, content io.ReadSeekCloser, options *LibraryClientAppendOptions) (LibraryClientAppendResponse, error) {
	req, err := client.appendCreateRequest(ctx, libraryName, content, options)
	if err != nil {
		return LibraryClientAppendResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LibraryClientAppendResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return LibraryClientAppendResponse{}, runtime.NewResponseError(resp)
	}
	return LibraryClientAppendResponse{}, nil
}

// appendCreateRequest creates the Append request.
func (client *libraryClient) appendCreateRequest(ctx context.Context, libraryName string, content io.ReadSeekCloser, options *LibraryClientAppendOptions) (*policy.Request, error) {
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
		req.Raw().Header["x-ms-blob-condition-appendpos"] = []string{strconv.FormatInt(*options.XMSBlobConditionAppendpos, 10)}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, req.SetBody(content, "application/octet-stream")
}

// BeginCreate - Creates a library with the library name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
//   - options - LibraryClientBeginCreateOptions contains the optional parameters for the libraryClient.BeginCreate method.
func (client *libraryClient) BeginCreate(ctx context.Context, libraryName string, options *LibraryClientBeginCreateOptions) (*runtime.Poller[LibraryClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, libraryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LibraryClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LibraryClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a library with the library name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *libraryClient) create(ctx context.Context, libraryName string, options *LibraryClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, libraryName, options)
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

// createCreateRequest creates the Create request.
func (client *libraryClient) createCreateRequest(ctx context.Context, libraryName string, options *LibraryClientBeginCreateOptions) (*policy.Request, error) {
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Delete Library
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
//   - options - LibraryClientBeginDeleteOptions contains the optional parameters for the libraryClient.BeginDelete method.
func (client *libraryClient) BeginDelete(ctx context.Context, libraryName string, options *LibraryClientBeginDeleteOptions) (*runtime.Poller[LibraryClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, libraryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LibraryClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LibraryClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete Library
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *libraryClient) deleteOperation(ctx context.Context, libraryName string, options *LibraryClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusConflict) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *libraryClient) deleteCreateRequest(ctx context.Context, libraryName string, options *LibraryClientBeginDeleteOptions) (*policy.Request, error) {
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginFlush - Flush Library
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
//   - options - LibraryClientBeginFlushOptions contains the optional parameters for the libraryClient.BeginFlush method.
func (client *libraryClient) BeginFlush(ctx context.Context, libraryName string, options *LibraryClientBeginFlushOptions) (*runtime.Poller[LibraryClientFlushResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.flush(ctx, libraryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LibraryClientFlushResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LibraryClientFlushResponse](options.ResumeToken, client.pl, nil)
	}
}

// Flush - Flush Library
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *libraryClient) flush(ctx context.Context, libraryName string, options *LibraryClientBeginFlushOptions) (*http.Response, error) {
	req, err := client.flushCreateRequest(ctx, libraryName, options)
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

// flushCreateRequest creates the Flush request.
func (client *libraryClient) flushCreateRequest(ctx context.Context, libraryName string, options *LibraryClientBeginFlushOptions) (*policy.Request, error) {
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get Library
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - libraryName - file name to upload. Minimum length of the filename should be 1 excluding the extension length.
//   - options - LibraryClientGetOptions contains the optional parameters for the libraryClient.Get method.
func (client *libraryClient) Get(ctx context.Context, libraryName string, options *LibraryClientGetOptions) (LibraryClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, libraryName, options)
	if err != nil {
		return LibraryClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LibraryClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return LibraryClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *libraryClient) getCreateRequest(ctx context.Context, libraryName string, options *LibraryClientGetOptions) (*policy.Request, error) {
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *libraryClient) getHandleResponse(resp *http.Response) (LibraryClientGetResponse, error) {
	result := LibraryClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryResource); err != nil {
		return LibraryClientGetResponse{}, err
	}
	return result, nil
}

// GetOperationResult - Get Operation result for Library
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - operationID - operation id for which status is requested
//   - options - LibraryClientGetOperationResultOptions contains the optional parameters for the libraryClient.GetOperationResult
//     method.
func (client *libraryClient) GetOperationResult(ctx context.Context, operationID string, options *LibraryClientGetOperationResultOptions) (LibraryClientGetOperationResultResponse, error) {
	req, err := client.getOperationResultCreateRequest(ctx, operationID, options)
	if err != nil {
		return LibraryClientGetOperationResultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LibraryClientGetOperationResultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return LibraryClientGetOperationResultResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationResultHandleResponse(resp)
}

// getOperationResultCreateRequest creates the GetOperationResult request.
func (client *libraryClient) getOperationResultCreateRequest(ctx context.Context, operationID string, options *LibraryClientGetOperationResultOptions) (*policy.Request, error) {
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationResultHandleResponse handles the GetOperationResult response.
func (client *libraryClient) getOperationResultHandleResponse(resp *http.Response) (LibraryClientGetOperationResultResponse, error) {
	result := LibraryClientGetOperationResultResponse{}
	switch resp.StatusCode {
	case http.StatusOK:
		var val LibraryResource
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return LibraryClientGetOperationResultResponse{}, err
		}
		result.Value = val
	case http.StatusAccepted:
		var val OperationResult
		if err := runtime.UnmarshalAsJSON(resp, &val); err != nil {
			return LibraryClientGetOperationResultResponse{}, err
		}
		result.Value = val
	default:
		return LibraryClientGetOperationResultResponse{}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)
	}
	return result, nil
}

// NewListPager - Lists Library.
//
// Generated from API version 2019-06-01-preview
//   - options - LibraryClientListOptions contains the optional parameters for the libraryClient.NewListPager method.
func (client *libraryClient) NewListPager(options *LibraryClientListOptions) *runtime.Pager[LibraryClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LibraryClientListResponse]{
		More: func(page LibraryClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LibraryClientListResponse) (LibraryClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LibraryClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LibraryClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LibraryClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *libraryClient) listCreateRequest(ctx context.Context, options *LibraryClientListOptions) (*policy.Request, error) {
	urlPath := "/libraries"
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

// listHandleResponse handles the List response.
func (client *libraryClient) listHandleResponse(resp *http.Response) (LibraryClientListResponse, error) {
	result := LibraryClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryListResponse); err != nil {
		return LibraryClientListResponse{}, err
	}
	return result, nil
}