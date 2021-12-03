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

// TriggersClient contains the methods for the Triggers group.
// Don't use this type directly, use NewTriggersClient() instead.
type TriggersClient struct {
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTriggersClient creates a new instance of TriggersClient with the specified values.
// subscriptionID - The subscription ID.
// options - pass nil to accept the default values.
func NewTriggersClient(subscriptionID string, options *azcore.ClientOptions) *TriggersClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &TriggersClient{
		subscriptionID: subscriptionID,
		pl:             runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a trigger.
// If the operation fails it returns the *CloudError error type.
// deviceName - Creates or updates a trigger
// name - The trigger name.
// resourceGroupName - The resource group name.
// trigger - The trigger.
// options - TriggersBeginCreateOrUpdateOptions contains the optional parameters for the TriggersClient.BeginCreateOrUpdate
// method.
func (client *TriggersClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, trigger TriggerClassification, options *TriggersBeginCreateOrUpdateOptions) (TriggersCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, name, resourceGroupName, trigger, options)
	if err != nil {
		return TriggersCreateOrUpdatePollerResponse{}, err
	}
	result := TriggersCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TriggersClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return TriggersCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &TriggersCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a trigger.
// If the operation fails it returns the *CloudError error type.
func (client *TriggersClient) createOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, trigger TriggerClassification, options *TriggersBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, name, resourceGroupName, trigger, options)
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
func (client *TriggersClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, trigger TriggerClassification, options *TriggersBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}"
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
	return req, runtime.MarshalAsJSON(req, trigger)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *TriggersClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the trigger on the gateway device.
// If the operation fails it returns the *CloudError error type.
// deviceName - The device name.
// name - The trigger name.
// resourceGroupName - The resource group name.
// options - TriggersBeginDeleteOptions contains the optional parameters for the TriggersClient.BeginDelete method.
func (client *TriggersClient) BeginDelete(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersBeginDeleteOptions) (TriggersDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return TriggersDeletePollerResponse{}, err
	}
	result := TriggersDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TriggersClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return TriggersDeletePollerResponse{}, err
	}
	result.Poller = &TriggersDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the trigger on the gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *TriggersClient) deleteOperation(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersBeginDeleteOptions) (*http.Response, error) {
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
func (client *TriggersClient) deleteCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}"
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
func (client *TriggersClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get a specific trigger by name.
// If the operation fails it returns the *CloudError error type.
// deviceName - The device name.
// name - The trigger name.
// resourceGroupName - The resource group name.
// options - TriggersGetOptions contains the optional parameters for the TriggersClient.Get method.
func (client *TriggersClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersGetOptions) (TriggersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return TriggersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TriggersClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *TriggersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}"
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
func (client *TriggersClient) getHandleResponse(resp *http.Response) (TriggersGetResponse, error) {
	result := TriggersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return TriggersGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TriggersClient) getHandleError(resp *http.Response) error {
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

// ListByDataBoxEdgeDevice - Lists all the triggers configured in the device.
// If the operation fails it returns the *CloudError error type.
// deviceName - The device name.
// resourceGroupName - The resource group name.
// options - TriggersListByDataBoxEdgeDeviceOptions contains the optional parameters for the TriggersClient.ListByDataBoxEdgeDevice
// method.
func (client *TriggersClient) ListByDataBoxEdgeDevice(deviceName string, resourceGroupName string, options *TriggersListByDataBoxEdgeDeviceOptions) *TriggersListByDataBoxEdgeDevicePager {
	return &TriggersListByDataBoxEdgeDevicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp TriggersListByDataBoxEdgeDeviceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TriggerList.NextLink)
		},
	}
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *TriggersClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *TriggersListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *TriggersClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (TriggersListByDataBoxEdgeDeviceResponse, error) {
	result := TriggersListByDataBoxEdgeDeviceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerList); err != nil {
		return TriggersListByDataBoxEdgeDeviceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByDataBoxEdgeDeviceHandleError handles the ListByDataBoxEdgeDevice error response.
func (client *TriggersClient) listByDataBoxEdgeDeviceHandleError(resp *http.Response) error {
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
