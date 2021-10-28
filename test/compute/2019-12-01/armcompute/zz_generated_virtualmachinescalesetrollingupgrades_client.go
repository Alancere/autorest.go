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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineScaleSetRollingUpgradesClient contains the methods for the VirtualMachineScaleSetRollingUpgrades group.
// Don't use this type directly, use NewVirtualMachineScaleSetRollingUpgradesClient() instead.
type VirtualMachineScaleSetRollingUpgradesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualMachineScaleSetRollingUpgradesClient creates a new instance of VirtualMachineScaleSetRollingUpgradesClient with the specified values.
func NewVirtualMachineScaleSetRollingUpgradesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineScaleSetRollingUpgradesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VirtualMachineScaleSetRollingUpgradesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCancel - Cancels the current virtual machine scale set rolling upgrade.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginCancel(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginCancelOptions) (VirtualMachineScaleSetRollingUpgradesCancelPollerResponse, error) {
	resp, err := client.cancel(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesCancelPollerResponse{}, err
	}
	result := VirtualMachineScaleSetRollingUpgradesCancelPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.Cancel", "", resp, client.pl, client.cancelHandleError)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesCancelPollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetRollingUpgradesCancelPoller{
		pt: pt,
	}
	return result, nil
}

// Cancel - Cancels the current virtual machine scale set rolling upgrade.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) cancel(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginCancelOptions) (*http.Response, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.cancelHandleError(resp)
	}
	return resp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetLatest - Gets the status of the latest virtual machine scale set rolling upgrade.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) GetLatest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesGetLatestOptions) (VirtualMachineScaleSetRollingUpgradesGetLatestResponse, error) {
	req, err := client.getLatestCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesGetLatestResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesGetLatestResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineScaleSetRollingUpgradesGetLatestResponse{}, client.getLatestHandleError(resp)
	}
	return client.getLatestHandleResponse(resp)
}

// getLatestCreateRequest creates the GetLatest request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) getLatestCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesGetLatestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/latest"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getLatestHandleResponse handles the GetLatest response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) getLatestHandleResponse(resp *http.Response) (VirtualMachineScaleSetRollingUpgradesGetLatestResponse, error) {
	result := VirtualMachineScaleSetRollingUpgradesGetLatestResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RollingUpgradeStatusInfo); err != nil {
		return VirtualMachineScaleSetRollingUpgradesGetLatestResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getLatestHandleError handles the GetLatest error response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) getLatestHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginStartExtensionUpgrade - Starts a rolling upgrade to move all extensions for all virtual machine scale set instances to the latest available extension
// version. Instances which are already running the latest extension versions
// are not affected.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginStartExtensionUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartExtensionUpgradeOptions) (VirtualMachineScaleSetRollingUpgradesStartExtensionUpgradePollerResponse, error) {
	resp, err := client.startExtensionUpgrade(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesStartExtensionUpgradePollerResponse{}, err
	}
	result := VirtualMachineScaleSetRollingUpgradesStartExtensionUpgradePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.StartExtensionUpgrade", "", resp, client.pl, client.startExtensionUpgradeHandleError)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesStartExtensionUpgradePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetRollingUpgradesStartExtensionUpgradePoller{
		pt: pt,
	}
	return result, nil
}

// StartExtensionUpgrade - Starts a rolling upgrade to move all extensions for all virtual machine scale set instances to the latest available extension
// version. Instances which are already running the latest extension versions
// are not affected.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startExtensionUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartExtensionUpgradeOptions) (*http.Response, error) {
	req, err := client.startExtensionUpgradeCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.startExtensionUpgradeHandleError(resp)
	}
	return resp, nil
}

// startExtensionUpgradeCreateRequest creates the StartExtensionUpgrade request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startExtensionUpgradeCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartExtensionUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensionRollingUpgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// startExtensionUpgradeHandleError handles the StartExtensionUpgrade error response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startExtensionUpgradeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginStartOSUpgrade - Starts a rolling upgrade to move all virtual machine scale set instances to the latest available Platform Image OS version. Instances
// which are already running the latest available OS version are not
// affected.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginStartOSUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartOSUpgradeOptions) (VirtualMachineScaleSetRollingUpgradesStartOSUpgradePollerResponse, error) {
	resp, err := client.startOSUpgrade(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesStartOSUpgradePollerResponse{}, err
	}
	result := VirtualMachineScaleSetRollingUpgradesStartOSUpgradePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.StartOSUpgrade", "", resp, client.pl, client.startOSUpgradeHandleError)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesStartOSUpgradePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetRollingUpgradesStartOSUpgradePoller{
		pt: pt,
	}
	return result, nil
}

// StartOSUpgrade - Starts a rolling upgrade to move all virtual machine scale set instances to the latest available Platform Image OS version. Instances
// which are already running the latest available OS version are not
// affected.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startOSUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartOSUpgradeOptions) (*http.Response, error) {
	req, err := client.startOSUpgradeCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.startOSUpgradeHandleError(resp)
	}
	return resp, nil
}

// startOSUpgradeCreateRequest creates the StartOSUpgrade request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startOSUpgradeCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartOSUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/osRollingUpgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// startOSUpgradeHandleError handles the StartOSUpgrade error response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startOSUpgradeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
