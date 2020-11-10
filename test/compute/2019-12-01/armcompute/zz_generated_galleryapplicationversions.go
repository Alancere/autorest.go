// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GalleryApplicationVersionsOperations contains the methods for the GalleryApplicationVersions group.
type GalleryApplicationVersionsOperations interface {
	// BeginCreateOrUpdate - Create or update a gallery Application Version.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsCreateOrUpdateOptions) (*GalleryApplicationVersionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (GalleryApplicationVersionPoller, error)
	// BeginDelete - Delete a gallery Application Version.
	BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves information about a gallery Application Version.
	Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsGetOptions) (*GalleryApplicationVersionResponse, error)
	// ListByGalleryApplication - List gallery Application Versions in a gallery Application Definition.
	ListByGalleryApplication(resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsListByGalleryApplicationOptions) GalleryApplicationVersionListPager
	// BeginUpdate - Update a gallery Application Version.
	BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsUpdateOptions) (*GalleryApplicationVersionPollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (GalleryApplicationVersionPoller, error)
}

// GalleryApplicationVersionsClient implements the GalleryApplicationVersionsOperations interface.
// Don't use this type directly, use NewGalleryApplicationVersionsClient() instead.
type GalleryApplicationVersionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewGalleryApplicationVersionsClient creates a new instance of GalleryApplicationVersionsClient with the specified values.
func NewGalleryApplicationVersionsClient(con *armcore.Connection, subscriptionID string) GalleryApplicationVersionsOperations {
	return &GalleryApplicationVersionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *GalleryApplicationVersionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *GalleryApplicationVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsCreateOrUpdateOptions) (*GalleryApplicationVersionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	result := &GalleryApplicationVersionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationVersionsClient.CreateOrUpdate", "", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &galleryApplicationVersionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*GalleryApplicationVersionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *GalleryApplicationVersionsClient) ResumeCreateOrUpdate(token string) (GalleryApplicationVersionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationVersionsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryApplicationVersionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a gallery Application Version.
func (client *GalleryApplicationVersionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryApplicationVersionsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryApplicationVersion)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GalleryApplicationVersionsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*GalleryApplicationVersionResponse, error) {
	result := GalleryApplicationVersionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryApplicationVersion)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GalleryApplicationVersionsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *GalleryApplicationVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationVersionsClient.Delete", "", resp, client.DeleteHandleError)
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

func (client *GalleryApplicationVersionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationVersionsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Delete a gallery Application Version.
func (client *GalleryApplicationVersionsClient) Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *GalleryApplicationVersionsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *GalleryApplicationVersionsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves information about a gallery Application Version.
func (client *GalleryApplicationVersionsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsGetOptions) (*GalleryApplicationVersionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *GalleryApplicationVersionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", string(*options.Expand))
	}
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *GalleryApplicationVersionsClient) GetHandleResponse(resp *azcore.Response) (*GalleryApplicationVersionResponse, error) {
	result := GalleryApplicationVersionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryApplicationVersion)
}

// GetHandleError handles the Get error response.
func (client *GalleryApplicationVersionsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByGalleryApplication - List gallery Application Versions in a gallery Application Definition.
func (client *GalleryApplicationVersionsClient) ListByGalleryApplication(resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsListByGalleryApplicationOptions) GalleryApplicationVersionListPager {
	return &galleryApplicationVersionListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByGalleryApplicationCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
		},
		responder: client.ListByGalleryApplicationHandleResponse,
		errorer:   client.ListByGalleryApplicationHandleError,
		advancer: func(ctx context.Context, resp *GalleryApplicationVersionListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.GalleryApplicationVersionList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByGalleryApplicationCreateRequest creates the ListByGalleryApplication request.
func (client *GalleryApplicationVersionsClient) ListByGalleryApplicationCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsListByGalleryApplicationOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByGalleryApplicationHandleResponse handles the ListByGalleryApplication response.
func (client *GalleryApplicationVersionsClient) ListByGalleryApplicationHandleResponse(resp *azcore.Response) (*GalleryApplicationVersionListResponse, error) {
	result := GalleryApplicationVersionListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryApplicationVersionList)
}

// ListByGalleryApplicationHandleError handles the ListByGalleryApplication error response.
func (client *GalleryApplicationVersionsClient) ListByGalleryApplicationHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *GalleryApplicationVersionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsUpdateOptions) (*GalleryApplicationVersionPollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	result := &GalleryApplicationVersionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationVersionsClient.Update", "", resp, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &galleryApplicationVersionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*GalleryApplicationVersionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *GalleryApplicationVersionsClient) ResumeUpdate(token string) (GalleryApplicationVersionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationVersionsClient.Update", token, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryApplicationVersionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - Update a gallery Application Version.
func (client *GalleryApplicationVersionsClient) Update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsUpdateOptions) (*azcore.Response, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	return resp, nil
}

// UpdateCreateRequest creates the Update request.
func (client *GalleryApplicationVersionsClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryApplicationVersion)
}

// UpdateHandleResponse handles the Update response.
func (client *GalleryApplicationVersionsClient) UpdateHandleResponse(resp *azcore.Response) (*GalleryApplicationVersionResponse, error) {
	result := GalleryApplicationVersionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryApplicationVersion)
}

// UpdateHandleError handles the Update error response.
func (client *GalleryApplicationVersionsClient) UpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
