// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DiskEncryptionSetsClient contains the methods for the DiskEncryptionSets group.
// Don't use this type directly, use NewDiskEncryptionSetsClient() instead.
type DiskEncryptionSetsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDiskEncryptionSetsClient creates a new instance of DiskEncryptionSetsClient with the specified values.
func NewDiskEncryptionSetsClient(con *armcore.Connection, subscriptionID string) *DiskEncryptionSetsClient {
	return &DiskEncryptionSetsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a disk encryption set
func (client *DiskEncryptionSetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsBeginCreateOrUpdateOptions) (DiskEncryptionSetPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return DiskEncryptionSetPollerResponse{}, err
	}
	result := DiskEncryptionSetPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DiskEncryptionSetsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return DiskEncryptionSetPollerResponse{}, err
	}
	poller := &diskEncryptionSetPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DiskEncryptionSetResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new DiskEncryptionSetPoller from the specified resume token.
// token - The value must come from a previous call to DiskEncryptionSetPoller.ResumeToken().
func (client *DiskEncryptionSetsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (DiskEncryptionSetPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("DiskEncryptionSetsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return DiskEncryptionSetPollerResponse{}, err
	}
	poller := &diskEncryptionSetPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DiskEncryptionSetPollerResponse{}, err
	}
	result := DiskEncryptionSetPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DiskEncryptionSetResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a disk encryption set
func (client *DiskEncryptionSetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiskEncryptionSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet, options *DiskEncryptionSetsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(diskEncryptionSet)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DiskEncryptionSetsClient) createOrUpdateHandleResponse(resp *azcore.Response) (DiskEncryptionSetResponse, error) {
	var val *DiskEncryptionSet
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskEncryptionSetResponse{}, err
	}
	return DiskEncryptionSetResponse{RawResponse: resp.Response, DiskEncryptionSet: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DiskEncryptionSetsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes a disk encryption set.
func (client *DiskEncryptionSetsClient) BeginDelete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, diskEncryptionSetName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DiskEncryptionSetsClient.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *DiskEncryptionSetsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("DiskEncryptionSetsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a disk encryption set.
func (client *DiskEncryptionSetsClient) deleteOperation(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiskEncryptionSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DiskEncryptionSetsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets information about a disk encryption set.
func (client *DiskEncryptionSetsClient) Get(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsGetOptions) (DiskEncryptionSetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, options)
	if err != nil {
		return DiskEncryptionSetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DiskEncryptionSetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DiskEncryptionSetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiskEncryptionSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, options *DiskEncryptionSetsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskEncryptionSetsClient) getHandleResponse(resp *azcore.Response) (DiskEncryptionSetResponse, error) {
	var val *DiskEncryptionSet
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskEncryptionSetResponse{}, err
	}
	return DiskEncryptionSetResponse{RawResponse: resp.Response, DiskEncryptionSet: val}, nil
}

// getHandleError handles the Get error response.
func (client *DiskEncryptionSetsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the disk encryption sets under a subscription.
func (client *DiskEncryptionSetsClient) List(options *DiskEncryptionSetsListOptions) DiskEncryptionSetListPager {
	return &diskEncryptionSetListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp DiskEncryptionSetListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DiskEncryptionSetList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *DiskEncryptionSetsClient) listCreateRequest(ctx context.Context, options *DiskEncryptionSetsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiskEncryptionSetsClient) listHandleResponse(resp *azcore.Response) (DiskEncryptionSetListResponse, error) {
	var val *DiskEncryptionSetList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskEncryptionSetListResponse{}, err
	}
	return DiskEncryptionSetListResponse{RawResponse: resp.Response, DiskEncryptionSetList: val}, nil
}

// listHandleError handles the List error response.
func (client *DiskEncryptionSetsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all the disk encryption sets under a resource group.
func (client *DiskEncryptionSetsClient) ListByResourceGroup(resourceGroupName string, options *DiskEncryptionSetsListByResourceGroupOptions) DiskEncryptionSetListPager {
	return &diskEncryptionSetListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp DiskEncryptionSetListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DiskEncryptionSetList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DiskEncryptionSetsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DiskEncryptionSetsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DiskEncryptionSetsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (DiskEncryptionSetListResponse, error) {
	var val *DiskEncryptionSetList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskEncryptionSetListResponse{}, err
	}
	return DiskEncryptionSetListResponse{RawResponse: resp.Response, DiskEncryptionSetList: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DiskEncryptionSetsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginUpdate - Updates (patches) a disk encryption set.
func (client *DiskEncryptionSetsClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsBeginUpdateOptions) (DiskEncryptionSetPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return DiskEncryptionSetPollerResponse{}, err
	}
	result := DiskEncryptionSetPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DiskEncryptionSetsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return DiskEncryptionSetPollerResponse{}, err
	}
	poller := &diskEncryptionSetPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DiskEncryptionSetResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new DiskEncryptionSetPoller from the specified resume token.
// token - The value must come from a previous call to DiskEncryptionSetPoller.ResumeToken().
func (client *DiskEncryptionSetsClient) ResumeUpdate(ctx context.Context, token string) (DiskEncryptionSetPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("DiskEncryptionSetsClient.Update", token, client.updateHandleError)
	if err != nil {
		return DiskEncryptionSetPollerResponse{}, err
	}
	poller := &diskEncryptionSetPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DiskEncryptionSetPollerResponse{}, err
	}
	result := DiskEncryptionSetPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DiskEncryptionSetResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates (patches) a disk encryption set.
func (client *DiskEncryptionSetsClient) update(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DiskEncryptionSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate, options *DiskEncryptionSetsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskEncryptionSetName == "" {
		return nil, errors.New("parameter diskEncryptionSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskEncryptionSetName}", url.PathEscape(diskEncryptionSetName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(diskEncryptionSet)
}

// updateHandleResponse handles the Update response.
func (client *DiskEncryptionSetsClient) updateHandleResponse(resp *azcore.Response) (DiskEncryptionSetResponse, error) {
	var val *DiskEncryptionSet
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskEncryptionSetResponse{}, err
	}
	return DiskEncryptionSetResponse{RawResponse: resp.Response, DiskEncryptionSet: val}, nil
}

// updateHandleError handles the Update error response.
func (client *DiskEncryptionSetsClient) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}