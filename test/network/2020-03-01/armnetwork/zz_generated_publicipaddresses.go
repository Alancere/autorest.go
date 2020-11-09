// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PublicIPAddressesOperations contains the methods for the PublicIPAddresses group.
type PublicIPAddressesOperations interface {
	// BeginCreateOrUpdate - Creates or updates a static or dynamic public IP address.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesCreateOrUpdateOptions) (*PublicIPAddressPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (PublicIPAddressPoller, error)
	// BeginDelete - Deletes the specified public IP address.
	BeginDelete(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified public IP address in a specified resource group.
	Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesGetOptions) (*PublicIPAddressResponse, error)
	// GetVirtualMachineScaleSetPublicIPAddress - Get the specified public IP address in a virtual machine scale set.
	GetVirtualMachineScaleSetPublicIPAddress(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, publicIPAddressName string, options *PublicIPAddressesGetVirtualMachineScaleSetPublicIPAddressOptions) (*PublicIPAddressResponse, error)
	// List - Gets all public IP addresses in a resource group.
	List(resourceGroupName string, options *PublicIPAddressesListOptions) PublicIPAddressListResultPager
	// ListAll - Gets all the public IP addresses in a subscription.
	ListAll(options *PublicIPAddressesListAllOptions) PublicIPAddressListResultPager
	// ListVirtualMachineScaleSetPublicIPAddresses - Gets information about all public IP addresses on a virtual machine scale set level.
	ListVirtualMachineScaleSetPublicIPAddresses(resourceGroupName string, virtualMachineScaleSetName string, options *PublicIPAddressesListVirtualMachineScaleSetPublicIPAddressesOptions) PublicIPAddressListResultPager
	// ListVirtualMachineScaleSetVMPublicIPaddresses - Gets information about all public IP addresses in a virtual machine IP configuration in a virtual machine
	// scale set.
	ListVirtualMachineScaleSetVMPublicIPaddresses(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, options *PublicIPAddressesListVirtualMachineScaleSetVMPublicIPaddressesOptions) PublicIPAddressListResultPager
	// UpdateTags - Updates public IP address tags.
	UpdateTags(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject, options *PublicIPAddressesUpdateTagsOptions) (*PublicIPAddressResponse, error)
}

// PublicIPAddressesClient implements the PublicIPAddressesOperations interface.
// Don't use this type directly, use NewPublicIPAddressesClient() instead.
type PublicIPAddressesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPublicIPAddressesClient creates a new instance of PublicIPAddressesClient with the specified values.
func NewPublicIPAddressesClient(con *armcore.Connection, subscriptionID string) PublicIPAddressesOperations {
	return &PublicIPAddressesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *PublicIPAddressesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *PublicIPAddressesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesCreateOrUpdateOptions) (*PublicIPAddressPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, publicIPAddressName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &PublicIPAddressPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPAddressesClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &publicIPAddressPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*PublicIPAddressResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *PublicIPAddressesClient) ResumeCreateOrUpdate(token string) (PublicIPAddressPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPAddressesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &publicIPAddressPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a static or dynamic public IP address.
func (client *PublicIPAddressesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, publicIPAddressName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PublicIPAddressesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, options *PublicIPAddressesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PublicIPAddressesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*PublicIPAddressResponse, error) {
	result := PublicIPAddressResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPAddress)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PublicIPAddressesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *PublicIPAddressesClient) BeginDelete(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, publicIPAddressName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPAddressesClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *PublicIPAddressesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPAddressesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified public IP address.
func (client *PublicIPAddressesClient) Delete(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, publicIPAddressName, options)
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
func (client *PublicIPAddressesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *PublicIPAddressesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified public IP address in a specified resource group.
func (client *PublicIPAddressesClient) Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesGetOptions) (*PublicIPAddressResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, publicIPAddressName, options)
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
func (client *PublicIPAddressesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *PublicIPAddressesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *PublicIPAddressesClient) GetHandleResponse(resp *azcore.Response) (*PublicIPAddressResponse, error) {
	result := PublicIPAddressResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPAddress)
}

// GetHandleError handles the Get error response.
func (client *PublicIPAddressesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetVirtualMachineScaleSetPublicIPAddress - Get the specified public IP address in a virtual machine scale set.
func (client *PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddress(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, publicIPAddressName string, options *PublicIPAddressesGetVirtualMachineScaleSetPublicIPAddressOptions) (*PublicIPAddressResponse, error) {
	req, err := client.GetVirtualMachineScaleSetPublicIPAddressCreateRequest(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, ipConfigurationName, publicIPAddressName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetVirtualMachineScaleSetPublicIPAddressHandleError(resp)
	}
	result, err := client.GetVirtualMachineScaleSetPublicIPAddressHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetVirtualMachineScaleSetPublicIPAddressCreateRequest creates the GetVirtualMachineScaleSetPublicIPAddress request.
func (client *PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddressCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, publicIPAddressName string, options *PublicIPAddressesGetVirtualMachineScaleSetPublicIPAddressOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineScaleSetName}", url.PathEscape(virtualMachineScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualmachineIndex}", url.PathEscape(virtualmachineIndex))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-10-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetVirtualMachineScaleSetPublicIPAddressHandleResponse handles the GetVirtualMachineScaleSetPublicIPAddress response.
func (client *PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddressHandleResponse(resp *azcore.Response) (*PublicIPAddressResponse, error) {
	result := PublicIPAddressResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPAddress)
}

// GetVirtualMachineScaleSetPublicIPAddressHandleError handles the GetVirtualMachineScaleSetPublicIPAddress error response.
func (client *PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddressHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all public IP addresses in a resource group.
func (client *PublicIPAddressesClient) List(resourceGroupName string, options *PublicIPAddressesListOptions) PublicIPAddressListResultPager {
	return &publicIPAddressListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *PublicIPAddressListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPAddressListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *PublicIPAddressesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, options *PublicIPAddressesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *PublicIPAddressesClient) ListHandleResponse(resp *azcore.Response) (*PublicIPAddressListResultResponse, error) {
	result := PublicIPAddressListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPAddressListResult)
}

// ListHandleError handles the List error response.
func (client *PublicIPAddressesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the public IP addresses in a subscription.
func (client *PublicIPAddressesClient) ListAll(options *PublicIPAddressesListAllOptions) PublicIPAddressListResultPager {
	return &publicIPAddressListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx, options)
		},
		responder: client.ListAllHandleResponse,
		errorer:   client.ListAllHandleError,
		advancer: func(ctx context.Context, resp *PublicIPAddressListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPAddressListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *PublicIPAddressesClient) ListAllCreateRequest(ctx context.Context, options *PublicIPAddressesListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListAllHandleResponse handles the ListAll response.
func (client *PublicIPAddressesClient) ListAllHandleResponse(resp *azcore.Response) (*PublicIPAddressListResultResponse, error) {
	result := PublicIPAddressListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPAddressListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *PublicIPAddressesClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListVirtualMachineScaleSetPublicIPAddresses - Gets information about all public IP addresses on a virtual machine scale set level.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddresses(resourceGroupName string, virtualMachineScaleSetName string, options *PublicIPAddressesListVirtualMachineScaleSetPublicIPAddressesOptions) PublicIPAddressListResultPager {
	return &publicIPAddressListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListVirtualMachineScaleSetPublicIPAddressesCreateRequest(ctx, resourceGroupName, virtualMachineScaleSetName, options)
		},
		responder: client.ListVirtualMachineScaleSetPublicIPAddressesHandleResponse,
		errorer:   client.ListVirtualMachineScaleSetPublicIPAddressesHandleError,
		advancer: func(ctx context.Context, resp *PublicIPAddressListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPAddressListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListVirtualMachineScaleSetPublicIPAddressesCreateRequest creates the ListVirtualMachineScaleSetPublicIPAddresses request.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddressesCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, options *PublicIPAddressesListVirtualMachineScaleSetPublicIPAddressesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineScaleSetName}", url.PathEscape(virtualMachineScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListVirtualMachineScaleSetPublicIPAddressesHandleResponse handles the ListVirtualMachineScaleSetPublicIPAddresses response.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddressesHandleResponse(resp *azcore.Response) (*PublicIPAddressListResultResponse, error) {
	result := PublicIPAddressListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPAddressListResult)
}

// ListVirtualMachineScaleSetPublicIPAddressesHandleError handles the ListVirtualMachineScaleSetPublicIPAddresses error response.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddressesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListVirtualMachineScaleSetVMPublicIPaddresses - Gets information about all public IP addresses in a virtual machine IP configuration in a virtual machine
// scale set.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPaddresses(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, options *PublicIPAddressesListVirtualMachineScaleSetVMPublicIPaddressesOptions) PublicIPAddressListResultPager {
	return &publicIPAddressListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListVirtualMachineScaleSetVMPublicIPaddressesCreateRequest(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, ipConfigurationName, options)
		},
		responder: client.ListVirtualMachineScaleSetVMPublicIPaddressesHandleResponse,
		errorer:   client.ListVirtualMachineScaleSetVMPublicIPaddressesHandleError,
		advancer: func(ctx context.Context, resp *PublicIPAddressListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPAddressListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListVirtualMachineScaleSetVMPublicIPaddressesCreateRequest creates the ListVirtualMachineScaleSetVMPublicIPaddresses request.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPaddressesCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, options *PublicIPAddressesListVirtualMachineScaleSetVMPublicIPaddressesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineScaleSetName}", url.PathEscape(virtualMachineScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualmachineIndex}", url.PathEscape(virtualmachineIndex))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListVirtualMachineScaleSetVMPublicIPaddressesHandleResponse handles the ListVirtualMachineScaleSetVMPublicIPaddresses response.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPaddressesHandleResponse(resp *azcore.Response) (*PublicIPAddressListResultResponse, error) {
	result := PublicIPAddressListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPAddressListResult)
}

// ListVirtualMachineScaleSetVMPublicIPaddressesHandleError handles the ListVirtualMachineScaleSetVMPublicIPaddresses error response.
func (client *PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPaddressesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates public IP address tags.
func (client *PublicIPAddressesClient) UpdateTags(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject, options *PublicIPAddressesUpdateTagsOptions) (*PublicIPAddressResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, publicIPAddressName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *PublicIPAddressesClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject, options *PublicIPAddressesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpAddressName}", url.PathEscape(publicIPAddressName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *PublicIPAddressesClient) UpdateTagsHandleResponse(resp *azcore.Response) (*PublicIPAddressResponse, error) {
	result := PublicIPAddressResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPAddress)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *PublicIPAddressesClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
