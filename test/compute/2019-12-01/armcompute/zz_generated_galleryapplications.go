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

// GalleryApplicationsClient contains the methods for the GalleryApplications group.
// Don't use this type directly, use NewGalleryApplicationsClient() instead.
type GalleryApplicationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewGalleryApplicationsClient creates a new instance of GalleryApplicationsClient with the specified values.
func NewGalleryApplicationsClient(con *armcore.Connection, subscriptionID string) GalleryApplicationsClient {
	return GalleryApplicationsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client GalleryApplicationsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Create or update a gallery Application Definition.
func (client GalleryApplicationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsCreateOrUpdateOptions) (GalleryApplicationPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return GalleryApplicationPollerResponse{}, err
	}
	result := GalleryApplicationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return GalleryApplicationPollerResponse{}, err
	}
	poller := &galleryApplicationPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new GalleryApplicationPoller from the specified resume token.
// token - The value must come from a previous call to GalleryApplicationPoller.ResumeToken().
func (client GalleryApplicationsClient) ResumeCreateOrUpdate(token string) (GalleryApplicationPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryApplicationPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a gallery Application Definition.
func (client GalleryApplicationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client GalleryApplicationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplication, options *GalleryApplicationsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryApplication)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client GalleryApplicationsClient) createOrUpdateHandleResponse(resp *azcore.Response) (GalleryApplicationResponse, error) {
	result := GalleryApplicationResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.GalleryApplication)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client GalleryApplicationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Delete a gallery Application.
func (client GalleryApplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationsClient.Delete", "", resp, client.deleteHandleError)
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

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client GalleryApplicationsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Delete a gallery Application.
func (client GalleryApplicationsClient) Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client GalleryApplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
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

// deleteHandleError handles the Delete error response.
func (client GalleryApplicationsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves information about a gallery Application Definition.
func (client GalleryApplicationsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsGetOptions) (GalleryApplicationResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
	if err != nil {
		return GalleryApplicationResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return GalleryApplicationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return GalleryApplicationResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return GalleryApplicationResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client GalleryApplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
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

// getHandleResponse handles the Get response.
func (client GalleryApplicationsClient) getHandleResponse(resp *azcore.Response) (GalleryApplicationResponse, error) {
	result := GalleryApplicationResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.GalleryApplication)
	return result, err
}

// getHandleError handles the Get error response.
func (client GalleryApplicationsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByGallery - List gallery Application Definitions in a gallery.
func (client GalleryApplicationsClient) ListByGallery(resourceGroupName string, galleryName string, options *GalleryApplicationsListByGalleryOptions) GalleryApplicationListPager {
	return &galleryApplicationListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByGalleryCreateRequest(ctx, resourceGroupName, galleryName, options)
		},
		responder: client.listByGalleryHandleResponse,
		errorer:   client.listByGalleryHandleError,
		advancer: func(ctx context.Context, resp GalleryApplicationListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.GalleryApplicationList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByGalleryCreateRequest creates the ListByGallery request.
func (client GalleryApplicationsClient) listByGalleryCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleryApplicationsListByGalleryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
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

// listByGalleryHandleResponse handles the ListByGallery response.
func (client GalleryApplicationsClient) listByGalleryHandleResponse(resp *azcore.Response) (GalleryApplicationListResponse, error) {
	result := GalleryApplicationListResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.GalleryApplicationList)
	return result, err
}

// listByGalleryHandleError handles the ListByGallery error response.
func (client GalleryApplicationsClient) listByGalleryHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginUpdate - Update a gallery Application Definition.
func (client GalleryApplicationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsUpdateOptions) (GalleryApplicationPollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return GalleryApplicationPollerResponse{}, err
	}
	result := GalleryApplicationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return GalleryApplicationPollerResponse{}, err
	}
	poller := &galleryApplicationPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new GalleryApplicationPoller from the specified resume token.
// token - The value must come from a previous call to GalleryApplicationPoller.ResumeToken().
func (client GalleryApplicationsClient) ResumeUpdate(token string) (GalleryApplicationPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationsClient.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryApplicationPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - Update a gallery Application Definition.
func (client GalleryApplicationsClient) Update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplication, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client GalleryApplicationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication GalleryApplicationUpdate, options *GalleryApplicationsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryApplication)
}

// updateHandleResponse handles the Update response.
func (client GalleryApplicationsClient) updateHandleResponse(resp *azcore.Response) (GalleryApplicationResponse, error) {
	result := GalleryApplicationResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.GalleryApplication)
	return result, err
}

// updateHandleError handles the Update error response.
func (client GalleryApplicationsClient) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
