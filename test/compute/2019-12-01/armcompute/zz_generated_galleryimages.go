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

// GalleryImagesOperations contains the methods for the GalleryImages group.
type GalleryImagesOperations interface {
	// BeginCreateOrUpdate - Create or update a gallery Image Definition.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImage, options *GalleryImagesCreateOrUpdateOptions) (*GalleryImagePollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (GalleryImagePoller, error)
	// BeginDelete - Delete a gallery image.
	BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves information about a gallery Image Definition.
	Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesGetOptions) (*GalleryImageResponse, error)
	// ListByGallery - List gallery Image Definitions in a gallery.
	ListByGallery(resourceGroupName string, galleryName string, options *GalleryImagesListByGalleryOptions) GalleryImageListPager
	// BeginUpdate - Update a gallery Image Definition.
	BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImageUpdate, options *GalleryImagesUpdateOptions) (*GalleryImagePollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (GalleryImagePoller, error)
}

// GalleryImagesClient implements the GalleryImagesOperations interface.
// Don't use this type directly, use NewGalleryImagesClient() instead.
type GalleryImagesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewGalleryImagesClient creates a new instance of GalleryImagesClient with the specified values.
func NewGalleryImagesClient(con *armcore.Connection, subscriptionID string) GalleryImagesOperations {
	return &GalleryImagesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *GalleryImagesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *GalleryImagesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImage, options *GalleryImagesCreateOrUpdateOptions) (*GalleryImagePollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, galleryName, galleryImageName, galleryImage, options)
	if err != nil {
		return nil, err
	}
	result := &GalleryImagePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryImagesClient.CreateOrUpdate", "", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &galleryImagePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*GalleryImageResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *GalleryImagesClient) ResumeCreateOrUpdate(token string) (GalleryImagePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryImagesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryImagePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a gallery Image Definition.
func (client *GalleryImagesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImage, options *GalleryImagesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, galleryImage, options)
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
func (client *GalleryImagesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImage, options *GalleryImagesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryImage)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GalleryImagesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*GalleryImageResponse, error) {
	result := GalleryImageResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryImage)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GalleryImagesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *GalleryImagesClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, galleryName, galleryImageName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryImagesClient.Delete", "", resp, client.DeleteHandleError)
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

func (client *GalleryImagesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryImagesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Delete a gallery image.
func (client *GalleryImagesClient) Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, options)
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
func (client *GalleryImagesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
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
func (client *GalleryImagesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves information about a gallery Image Definition.
func (client *GalleryImagesClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesGetOptions) (*GalleryImageResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, options)
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
func (client *GalleryImagesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImagesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
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

// GetHandleResponse handles the Get response.
func (client *GalleryImagesClient) GetHandleResponse(resp *azcore.Response) (*GalleryImageResponse, error) {
	result := GalleryImageResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryImage)
}

// GetHandleError handles the Get error response.
func (client *GalleryImagesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByGallery - List gallery Image Definitions in a gallery.
func (client *GalleryImagesClient) ListByGallery(resourceGroupName string, galleryName string, options *GalleryImagesListByGalleryOptions) GalleryImageListPager {
	return &galleryImageListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByGalleryCreateRequest(ctx, resourceGroupName, galleryName, options)
		},
		responder: client.ListByGalleryHandleResponse,
		errorer:   client.ListByGalleryHandleError,
		advancer: func(ctx context.Context, resp *GalleryImageListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.GalleryImageList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByGalleryCreateRequest creates the ListByGallery request.
func (client *GalleryImagesClient) ListByGalleryCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleryImagesListByGalleryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images"
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

// ListByGalleryHandleResponse handles the ListByGallery response.
func (client *GalleryImagesClient) ListByGalleryHandleResponse(resp *azcore.Response) (*GalleryImageListResponse, error) {
	result := GalleryImageListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryImageList)
}

// ListByGalleryHandleError handles the ListByGallery error response.
func (client *GalleryImagesClient) ListByGalleryHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *GalleryImagesClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImageUpdate, options *GalleryImagesUpdateOptions) (*GalleryImagePollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, galleryName, galleryImageName, galleryImage, options)
	if err != nil {
		return nil, err
	}
	result := &GalleryImagePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryImagesClient.Update", "", resp, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &galleryImagePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*GalleryImageResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *GalleryImagesClient) ResumeUpdate(token string) (GalleryImagePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryImagesClient.Update", token, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryImagePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - Update a gallery Image Definition.
func (client *GalleryImagesClient) Update(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImageUpdate, options *GalleryImagesUpdateOptions) (*azcore.Response, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, galleryImage, options)
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
func (client *GalleryImagesClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImage GalleryImageUpdate, options *GalleryImagesUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryImage)
}

// UpdateHandleResponse handles the Update response.
func (client *GalleryImagesClient) UpdateHandleResponse(resp *azcore.Response) (*GalleryImageResponse, error) {
	result := GalleryImageResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryImage)
}

// UpdateHandleError handles the Update error response.
func (client *GalleryImagesClient) UpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
