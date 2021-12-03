//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BandwidthSchedulesClient contains the methods for the BandwidthSchedules group.
// Don't use this type directly, use NewBandwidthSchedulesClient() instead.
type BandwidthSchedulesClient struct {
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBandwidthSchedulesClient creates a new instance of BandwidthSchedulesClient with the specified values.
// subscriptionID - The subscription ID.
// options - pass nil to accept the default values.
func NewBandwidthSchedulesClient(subscriptionID string, options *azcore.ClientOptions) *BandwidthSchedulesClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &BandwidthSchedulesClient{
		subscriptionID: subscriptionID,
		pl:             runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a bandwidth schedule.
// If the operation fails it returns the *CloudError error type.
// deviceName - The device name.
// name - The bandwidth schedule name which needs to be added/updated.
// resourceGroupName - The resource group name.
// parameters - The bandwidth schedule to be added or updated.
// options - BandwidthSchedulesBeginCreateOrUpdateOptions contains the optional parameters for the BandwidthSchedulesClient.BeginCreateOrUpdate
// method.
func (client *BandwidthSchedulesClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, parameters BandwidthSchedule, options *BandwidthSchedulesBeginCreateOrUpdateOptions) (BandwidthSchedulesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, name, resourceGroupName, parameters, options)
	if err != nil {
		return BandwidthSchedulesCreateOrUpdatePollerResponse{}, err
	}
	result := BandwidthSchedulesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BandwidthSchedulesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return BandwidthSchedulesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &BandwidthSchedulesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a bandwidth schedule.
// If the operation fails it returns the *CloudError error type.
func (client *BandwidthSchedulesClient) createOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, parameters BandwidthSchedule, options *BandwidthSchedulesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, name, resourceGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BandwidthSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, parameters BandwidthSchedule, options *BandwidthSchedulesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *BandwidthSchedulesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the specified bandwidth schedule.
// If the operation fails it returns the *CloudError error type.
// deviceName - The device name.
// name - The bandwidth schedule name.
// resourceGroupName - The resource group name.
// options - BandwidthSchedulesBeginDeleteOptions contains the optional parameters for the BandwidthSchedulesClient.BeginDelete
// method.
func (client *BandwidthSchedulesClient) BeginDelete(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesBeginDeleteOptions) (BandwidthSchedulesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return BandwidthSchedulesDeletePollerResponse{}, err
	}
	result := BandwidthSchedulesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BandwidthSchedulesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return BandwidthSchedulesDeletePollerResponse{}, err
	}
	result.Poller = &BandwidthSchedulesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified bandwidth schedule.
// If the operation fails it returns the *CloudError error type.
func (client *BandwidthSchedulesClient) deleteOperation(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, name, resourceGroupName, options)
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
func (client *BandwidthSchedulesClient) deleteCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *BandwidthSchedulesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the properties of the specified bandwidth schedule.
// If the operation fails it returns the *CloudError error type.
// deviceName - The device name.
// name - The bandwidth schedule name.
// resourceGroupName - The resource group name.
// options - BandwidthSchedulesGetOptions contains the optional parameters for the BandwidthSchedulesClient.Get method.
func (client *BandwidthSchedulesClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesGetOptions) (BandwidthSchedulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return BandwidthSchedulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BandwidthSchedulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BandwidthSchedulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BandwidthSchedulesClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BandwidthSchedulesClient) getHandleResponse(resp *http.Response) (BandwidthSchedulesGetResponse, error) {
	result := BandwidthSchedulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BandwidthSchedule); err != nil {
		return BandwidthSchedulesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *BandwidthSchedulesClient) getHandleError(resp *http.Response) error {
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

// ListByDataBoxEdgeDevice - Gets all the bandwidth schedules for a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
// deviceName - The device name.
// resourceGroupName - The resource group name.
// options - BandwidthSchedulesListByDataBoxEdgeDeviceOptions contains the optional parameters for the BandwidthSchedulesClient.ListByDataBoxEdgeDevice
// method.
func (client *BandwidthSchedulesClient) ListByDataBoxEdgeDevice(deviceName string, resourceGroupName string, options *BandwidthSchedulesListByDataBoxEdgeDeviceOptions) *BandwidthSchedulesListByDataBoxEdgeDevicePager {
	return &BandwidthSchedulesListByDataBoxEdgeDevicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp BandwidthSchedulesListByDataBoxEdgeDeviceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BandwidthSchedulesList.NextLink)
		},
	}
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *BandwidthSchedulesClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *BandwidthSchedulesListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *BandwidthSchedulesClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (BandwidthSchedulesListByDataBoxEdgeDeviceResponse, error) {
	result := BandwidthSchedulesListByDataBoxEdgeDeviceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BandwidthSchedulesList); err != nil {
		return BandwidthSchedulesListByDataBoxEdgeDeviceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByDataBoxEdgeDeviceHandleError handles the ListByDataBoxEdgeDevice error response.
func (client *BandwidthSchedulesClient) listByDataBoxEdgeDeviceHandleError(resp *http.Response) error {
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
