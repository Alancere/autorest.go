//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GalleryApplicationVersionsClient contains the methods for the GalleryApplicationVersions group.
// Don't use this type directly, use NewGalleryApplicationVersionsClient() instead.
type GalleryApplicationVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGalleryApplicationVersionsClient creates a new instance of GalleryApplicationVersionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGalleryApplicationVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *GalleryApplicationVersionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &GalleryApplicationVersionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create or update a gallery Application Version.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the gallery Application Definition in which the Application Version is to be created.
// galleryApplicationVersionName - The name of the gallery Application Version to be created. Needs to follow semantic version
// name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit
// integer. Format: ..
// galleryApplicationVersion - Parameters supplied to the create or update gallery Application Version operation.
// options - GalleryApplicationVersionsBeginCreateOrUpdateOptions contains the optional parameters for the GalleryApplicationVersionsClient.BeginCreateOrUpdate
// method.
func (client *GalleryApplicationVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsBeginCreateOrUpdateOptions) (GalleryApplicationVersionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return GalleryApplicationVersionsCreateOrUpdatePollerResponse{}, err
	}
	result := GalleryApplicationVersionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GalleryApplicationVersionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return GalleryApplicationVersionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &GalleryApplicationVersionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a gallery Application Version.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryApplicationVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, galleryApplicationVersion)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GalleryApplicationVersionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete a gallery Application Version.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the gallery Application Definition in which the Application Version resides.
// galleryApplicationVersionName - The name of the gallery Application Version to be deleted.
// options - GalleryApplicationVersionsBeginDeleteOptions contains the optional parameters for the GalleryApplicationVersionsClient.BeginDelete
// method.
func (client *GalleryApplicationVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsBeginDeleteOptions) (GalleryApplicationVersionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return GalleryApplicationVersionsDeletePollerResponse{}, err
	}
	result := GalleryApplicationVersionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GalleryApplicationVersionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return GalleryApplicationVersionsDeletePollerResponse{}, err
	}
	result.Poller = &GalleryApplicationVersionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a gallery Application Version.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GalleryApplicationVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *GalleryApplicationVersionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves information about a gallery Application Version.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the gallery Application Definition in which the Application Version resides.
// galleryApplicationVersionName - The name of the gallery Application Version to be retrieved.
// options - GalleryApplicationVersionsGetOptions contains the optional parameters for the GalleryApplicationVersionsClient.Get
// method.
func (client *GalleryApplicationVersionsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsGetOptions) (GalleryApplicationVersionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return GalleryApplicationVersionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GalleryApplicationVersionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GalleryApplicationVersionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GalleryApplicationVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryApplicationVersionsClient) getHandleResponse(resp *http.Response) (GalleryApplicationVersionsGetResponse, error) {
	result := GalleryApplicationVersionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryApplicationVersion); err != nil {
		return GalleryApplicationVersionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *GalleryApplicationVersionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByGalleryApplication - List gallery Application Versions in a gallery Application Definition.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the Shared Application Gallery Application Definition from which the Application Versions
// are to be listed.
// options - GalleryApplicationVersionsListByGalleryApplicationOptions contains the optional parameters for the GalleryApplicationVersionsClient.ListByGalleryApplication
// method.
func (client *GalleryApplicationVersionsClient) ListByGalleryApplication(resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsListByGalleryApplicationOptions) *GalleryApplicationVersionsListByGalleryApplicationPager {
	return &GalleryApplicationVersionsListByGalleryApplicationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByGalleryApplicationCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
		},
		advancer: func(ctx context.Context, resp GalleryApplicationVersionsListByGalleryApplicationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.GalleryApplicationVersionList.NextLink)
		},
	}
}

// listByGalleryApplicationCreateRequest creates the ListByGalleryApplication request.
func (client *GalleryApplicationVersionsClient) listByGalleryApplicationCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsListByGalleryApplicationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByGalleryApplicationHandleResponse handles the ListByGalleryApplication response.
func (client *GalleryApplicationVersionsClient) listByGalleryApplicationHandleResponse(resp *http.Response) (GalleryApplicationVersionsListByGalleryApplicationResponse, error) {
	result := GalleryApplicationVersionsListByGalleryApplicationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryApplicationVersionList); err != nil {
		return GalleryApplicationVersionsListByGalleryApplicationResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByGalleryApplicationHandleError handles the ListByGalleryApplication error response.
func (client *GalleryApplicationVersionsClient) listByGalleryApplicationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Update a gallery Application Version.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - The name of the gallery Application Definition in which the Application Version is to be updated.
// galleryApplicationVersionName - The name of the gallery Application Version to be updated. Needs to follow semantic version
// name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit
// integer. Format: ..
// galleryApplicationVersion - Parameters supplied to the update gallery Application Version operation.
// options - GalleryApplicationVersionsBeginUpdateOptions contains the optional parameters for the GalleryApplicationVersionsClient.BeginUpdate
// method.
func (client *GalleryApplicationVersionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsBeginUpdateOptions) (GalleryApplicationVersionsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return GalleryApplicationVersionsUpdatePollerResponse{}, err
	}
	result := GalleryApplicationVersionsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GalleryApplicationVersionsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return GalleryApplicationVersionsUpdatePollerResponse{}, err
	}
	result.Poller = &GalleryApplicationVersionsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update a gallery Application Version.
// If the operation fails it returns the *CloudError error type.
func (client *GalleryApplicationVersionsClient) update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *GalleryApplicationVersionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, galleryApplicationVersion)
}

// updateHandleError handles the Update error response.
func (client *GalleryApplicationVersionsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
