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

// NetworkManagementClientOperations contains the methods for the NetworkManagementClient group.
type NetworkManagementClientOperations interface {
	// CheckDNSNameAvailability - Checks whether a domain name in the cloudapp.azure.com zone is available for use.
	CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string, options *NetworkManagementClientCheckDNSNameAvailabilityOptions) (*DNSNameAvailabilityResultResponse, error)
	// BeginDeleteBastionShareableLink - Deletes the Bastion Shareable Links for all the VMs specified in the request.
	BeginDeleteBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientDeleteBastionShareableLinkOptions) (*HTTPPollerResponse, error)
	// ResumeDeleteBastionShareableLink - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDeleteBastionShareableLink(token string) (HTTPPoller, error)
	// DisconnectActiveSessions - Returns the list of currently active sessions on the Bastion.
	DisconnectActiveSessions(resourceGroupName string, bastionHostName string, sessionIds SessionIDs, options *NetworkManagementClientDisconnectActiveSessionsOptions) BastionSessionDeleteResultPager
	// BeginGeneratevirtualwanvpnserverconfigurationvpnprofile - Generates a unique VPN profile for P2S clients for VirtualWan and associated VpnServerConfiguration
	// combination in the specified resource group.
	BeginGeneratevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWanName string, vpnClientParams VirtualWanVpnProfileParameters, options *NetworkManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*VpnProfileResponsePollerResponse, error)
	// ResumeGeneratevirtualwanvpnserverconfigurationvpnprofile - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeGeneratevirtualwanvpnserverconfigurationvpnprofile(token string) (VpnProfileResponsePoller, error)
	// BeginGetActiveSessions - Returns the list of currently active sessions on the Bastion.
	BeginGetActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string, options *NetworkManagementClientGetActiveSessionsOptions) (*BastionActiveSessionListResultPagerPollerResponse, error)
	// ResumeGetActiveSessions - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeGetActiveSessions(token string) (BastionActiveSessionListResultPagerPoller, error)
	// GetBastionShareableLink - Return the Bastion Shareable Links for all the VMs specified in the request.
	GetBastionShareableLink(resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientGetBastionShareableLinkOptions) BastionShareableLinkListResultPager
	// BeginPutBastionShareableLink - Creates a Bastion Shareable Links for all the VMs specified in the request.
	BeginPutBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientPutBastionShareableLinkOptions) (*BastionShareableLinkListResultPagerPollerResponse, error)
	// ResumePutBastionShareableLink - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePutBastionShareableLink(token string) (BastionShareableLinkListResultPagerPoller, error)
	// SupportedSecurityProviders - Gives the supported security providers for the virtual wan.
	SupportedSecurityProviders(ctx context.Context, resourceGroupName string, virtualWanName string, options *NetworkManagementClientSupportedSecurityProvidersOptions) (*VirtualWanSecurityProvidersResponse, error)
}

// NetworkManagementClient implements the NetworkManagementClientOperations interface.
// Don't use this type directly, use NewNetworkManagementClient() instead.
type NetworkManagementClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkManagementClient creates a new instance of NetworkManagementClient with the specified values.
func NewNetworkManagementClient(con *armcore.Connection, subscriptionID string) NetworkManagementClientOperations {
	return &NetworkManagementClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *NetworkManagementClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CheckDNSNameAvailability - Checks whether a domain name in the cloudapp.azure.com zone is available for use.
func (client *NetworkManagementClient) CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string, options *NetworkManagementClientCheckDNSNameAvailabilityOptions) (*DNSNameAvailabilityResultResponse, error) {
	req, err := client.CheckDNSNameAvailabilityCreateRequest(ctx, location, domainNameLabel, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CheckDNSNameAvailabilityHandleError(resp)
	}
	result, err := client.CheckDNSNameAvailabilityHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CheckDNSNameAvailabilityCreateRequest creates the CheckDNSNameAvailability request.
func (client *NetworkManagementClient) CheckDNSNameAvailabilityCreateRequest(ctx context.Context, location string, domainNameLabel string, options *NetworkManagementClientCheckDNSNameAvailabilityOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("domainNameLabel", domainNameLabel)
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// CheckDNSNameAvailabilityHandleResponse handles the CheckDNSNameAvailability response.
func (client *NetworkManagementClient) CheckDNSNameAvailabilityHandleResponse(resp *azcore.Response) (*DNSNameAvailabilityResultResponse, error) {
	result := DNSNameAvailabilityResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DNSNameAvailabilityResult)
}

// CheckDNSNameAvailabilityHandleError handles the CheckDNSNameAvailability error response.
func (client *NetworkManagementClient) CheckDNSNameAvailabilityHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *NetworkManagementClient) BeginDeleteBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientDeleteBastionShareableLinkOptions) (*HTTPPollerResponse, error) {
	resp, err := client.DeleteBastionShareableLink(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkManagementClient.DeleteBastionShareableLink", "location", resp, client.DeleteBastionShareableLinkHandleError)
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

func (client *NetworkManagementClient) ResumeDeleteBastionShareableLink(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkManagementClient.DeleteBastionShareableLink", token, client.DeleteBastionShareableLinkHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// DeleteBastionShareableLink - Deletes the Bastion Shareable Links for all the VMs specified in the request.
func (client *NetworkManagementClient) DeleteBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientDeleteBastionShareableLinkOptions) (*azcore.Response, error) {
	req, err := client.DeleteBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.DeleteBastionShareableLinkHandleError(resp)
	}
	return resp, nil
}

// DeleteBastionShareableLinkCreateRequest creates the DeleteBastionShareableLink request.
func (client *NetworkManagementClient) DeleteBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientDeleteBastionShareableLinkOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/deleteShareableLinks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bslRequest)
}

// DeleteBastionShareableLinkHandleError handles the DeleteBastionShareableLink error response.
func (client *NetworkManagementClient) DeleteBastionShareableLinkHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DisconnectActiveSessions - Returns the list of currently active sessions on the Bastion.
func (client *NetworkManagementClient) DisconnectActiveSessions(resourceGroupName string, bastionHostName string, sessionIds SessionIDs, options *NetworkManagementClientDisconnectActiveSessionsOptions) BastionSessionDeleteResultPager {
	return &bastionSessionDeleteResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.DisconnectActiveSessionsCreateRequest(ctx, resourceGroupName, bastionHostName, sessionIds, options)
		},
		responder: client.DisconnectActiveSessionsHandleResponse,
		errorer:   client.DisconnectActiveSessionsHandleError,
		advancer: func(ctx context.Context, resp *BastionSessionDeleteResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.BastionSessionDeleteResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// DisconnectActiveSessionsCreateRequest creates the DisconnectActiveSessions request.
func (client *NetworkManagementClient) DisconnectActiveSessionsCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, sessionIds SessionIDs, options *NetworkManagementClientDisconnectActiveSessionsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/disconnectActiveSessions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sessionIds)
}

// DisconnectActiveSessionsHandleResponse handles the DisconnectActiveSessions response.
func (client *NetworkManagementClient) DisconnectActiveSessionsHandleResponse(resp *azcore.Response) (*BastionSessionDeleteResultResponse, error) {
	result := BastionSessionDeleteResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BastionSessionDeleteResult)
}

// DisconnectActiveSessionsHandleError handles the DisconnectActiveSessions error response.
func (client *NetworkManagementClient) DisconnectActiveSessionsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *NetworkManagementClient) BeginGeneratevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWanName string, vpnClientParams VirtualWanVpnProfileParameters, options *NetworkManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*VpnProfileResponsePollerResponse, error) {
	resp, err := client.Generatevirtualwanvpnserverconfigurationvpnprofile(ctx, resourceGroupName, virtualWanName, vpnClientParams, options)
	if err != nil {
		return nil, err
	}
	result := &VpnProfileResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkManagementClient.Generatevirtualwanvpnserverconfigurationvpnprofile", "location", resp, client.GeneratevirtualwanvpnserverconfigurationvpnprofileHandleError)
	if err != nil {
		return nil, err
	}
	poller := &vpnProfileResponsePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VpnProfileResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *NetworkManagementClient) ResumeGeneratevirtualwanvpnserverconfigurationvpnprofile(token string) (VpnProfileResponsePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkManagementClient.Generatevirtualwanvpnserverconfigurationvpnprofile", token, client.GeneratevirtualwanvpnserverconfigurationvpnprofileHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnProfileResponsePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Generatevirtualwanvpnserverconfigurationvpnprofile - Generates a unique VPN profile for P2S clients for VirtualWan and associated VpnServerConfiguration
// combination in the specified resource group.
func (client *NetworkManagementClient) Generatevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWanName string, vpnClientParams VirtualWanVpnProfileParameters, options *NetworkManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*azcore.Response, error) {
	req, err := client.GeneratevirtualwanvpnserverconfigurationvpnprofileCreateRequest(ctx, resourceGroupName, virtualWanName, vpnClientParams, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.GeneratevirtualwanvpnserverconfigurationvpnprofileHandleError(resp)
	}
	return resp, nil
}

// GeneratevirtualwanvpnserverconfigurationvpnprofileCreateRequest creates the Generatevirtualwanvpnserverconfigurationvpnprofile request.
func (client *NetworkManagementClient) GeneratevirtualwanvpnserverconfigurationvpnprofileCreateRequest(ctx context.Context, resourceGroupName string, virtualWanName string, vpnClientParams VirtualWanVpnProfileParameters, options *NetworkManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/GenerateVpnProfile"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWanName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnClientParams)
}

// GeneratevirtualwanvpnserverconfigurationvpnprofileHandleResponse handles the Generatevirtualwanvpnserverconfigurationvpnprofile response.
func (client *NetworkManagementClient) GeneratevirtualwanvpnserverconfigurationvpnprofileHandleResponse(resp *azcore.Response) (*VpnProfileResponseResponse, error) {
	result := VpnProfileResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnProfileResponse)
}

// GeneratevirtualwanvpnserverconfigurationvpnprofileHandleError handles the Generatevirtualwanvpnserverconfigurationvpnprofile error response.
func (client *NetworkManagementClient) GeneratevirtualwanvpnserverconfigurationvpnprofileHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *NetworkManagementClient) BeginGetActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string, options *NetworkManagementClientGetActiveSessionsOptions) (*BastionActiveSessionListResultPagerPollerResponse, error) {
	resp, err := client.GetActiveSessions(ctx, resourceGroupName, bastionHostName, options)
	if err != nil {
		return nil, err
	}
	result := &BastionActiveSessionListResultPagerPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkManagementClient.GetActiveSessions", "location", resp, client.GetActiveSessionsHandleError)
	if err != nil {
		return nil, err
	}
	poller := &bastionActiveSessionListResultPagerPoller{
		pt: pt,
		errHandler: func(resp *azcore.Response) error {
			if resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
				return nil
			}
			return client.GetActiveSessionsHandleError(resp)
		},
		respHandler: func(resp *azcore.Response) (*BastionActiveSessionListResultResponse, error) {
			result := BastionActiveSessionListResultResponse{RawResponse: resp.Response}
			return &result, resp.UnmarshalAsJSON(&result.BastionActiveSessionListResult)
		},
		statusCodes: []int{http.StatusOK, http.StatusAccepted, http.StatusNoContent},
		pipeline:    client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (BastionActiveSessionListResultPager, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *NetworkManagementClient) ResumeGetActiveSessions(token string) (BastionActiveSessionListResultPagerPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkManagementClient.GetActiveSessions", token, client.GetActiveSessionsHandleError)
	if err != nil {
		return nil, err
	}
	return &bastionActiveSessionListResultPagerPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// GetActiveSessions - Returns the list of currently active sessions on the Bastion.
func (client *NetworkManagementClient) GetActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string, options *NetworkManagementClientGetActiveSessionsOptions) (*azcore.Response, error) {
	req, err := client.GetActiveSessionsCreateRequest(ctx, resourceGroupName, bastionHostName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.GetActiveSessionsHandleError(resp)
	}
	return resp, nil
}

// GetActiveSessionsCreateRequest creates the GetActiveSessions request.
func (client *NetworkManagementClient) GetActiveSessionsCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, options *NetworkManagementClientGetActiveSessionsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/getActiveSessions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetActiveSessionsHandleResponse handles the GetActiveSessions response.
func (client *NetworkManagementClient) GetActiveSessionsHandleResponse(resp *azcore.Response) (*BastionActiveSessionListResultResponse, error) {
	result := BastionActiveSessionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BastionActiveSessionListResult)
}

// GetActiveSessionsHandleError handles the GetActiveSessions error response.
func (client *NetworkManagementClient) GetActiveSessionsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetBastionShareableLink - Return the Bastion Shareable Links for all the VMs specified in the request.
func (client *NetworkManagementClient) GetBastionShareableLink(resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientGetBastionShareableLinkOptions) BastionShareableLinkListResultPager {
	return &bastionShareableLinkListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
		},
		responder: client.GetBastionShareableLinkHandleResponse,
		errorer:   client.GetBastionShareableLinkHandleError,
		advancer: func(ctx context.Context, resp *BastionShareableLinkListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.BastionShareableLinkListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// GetBastionShareableLinkCreateRequest creates the GetBastionShareableLink request.
func (client *NetworkManagementClient) GetBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientGetBastionShareableLinkOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/getShareableLinks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bslRequest)
}

// GetBastionShareableLinkHandleResponse handles the GetBastionShareableLink response.
func (client *NetworkManagementClient) GetBastionShareableLinkHandleResponse(resp *azcore.Response) (*BastionShareableLinkListResultResponse, error) {
	result := BastionShareableLinkListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BastionShareableLinkListResult)
}

// GetBastionShareableLinkHandleError handles the GetBastionShareableLink error response.
func (client *NetworkManagementClient) GetBastionShareableLinkHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *NetworkManagementClient) BeginPutBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientPutBastionShareableLinkOptions) (*BastionShareableLinkListResultPagerPollerResponse, error) {
	resp, err := client.PutBastionShareableLink(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return nil, err
	}
	result := &BastionShareableLinkListResultPagerPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkManagementClient.PutBastionShareableLink", "location", resp, client.PutBastionShareableLinkHandleError)
	if err != nil {
		return nil, err
	}
	poller := &bastionShareableLinkListResultPagerPoller{
		pt: pt,
		errHandler: func(resp *azcore.Response) error {
			if resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
				return nil
			}
			return client.PutBastionShareableLinkHandleError(resp)
		},
		respHandler: func(resp *azcore.Response) (*BastionShareableLinkListResultResponse, error) {
			result := BastionShareableLinkListResultResponse{RawResponse: resp.Response}
			return &result, resp.UnmarshalAsJSON(&result.BastionShareableLinkListResult)
		},
		statusCodes: []int{http.StatusOK, http.StatusAccepted, http.StatusNoContent},
		pipeline:    client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (BastionShareableLinkListResultPager, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *NetworkManagementClient) ResumePutBastionShareableLink(token string) (BastionShareableLinkListResultPagerPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkManagementClient.PutBastionShareableLink", token, client.PutBastionShareableLinkHandleError)
	if err != nil {
		return nil, err
	}
	return &bastionShareableLinkListResultPagerPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// PutBastionShareableLink - Creates a Bastion Shareable Links for all the VMs specified in the request.
func (client *NetworkManagementClient) PutBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientPutBastionShareableLinkOptions) (*azcore.Response, error) {
	req, err := client.PutBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.PutBastionShareableLinkHandleError(resp)
	}
	return resp, nil
}

// PutBastionShareableLinkCreateRequest creates the PutBastionShareableLink request.
func (client *NetworkManagementClient) PutBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientPutBastionShareableLinkOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/createShareableLinks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bslRequest)
}

// PutBastionShareableLinkHandleResponse handles the PutBastionShareableLink response.
func (client *NetworkManagementClient) PutBastionShareableLinkHandleResponse(resp *azcore.Response) (*BastionShareableLinkListResultResponse, error) {
	result := BastionShareableLinkListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BastionShareableLinkListResult)
}

// PutBastionShareableLinkHandleError handles the PutBastionShareableLink error response.
func (client *NetworkManagementClient) PutBastionShareableLinkHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SupportedSecurityProviders - Gives the supported security providers for the virtual wan.
func (client *NetworkManagementClient) SupportedSecurityProviders(ctx context.Context, resourceGroupName string, virtualWanName string, options *NetworkManagementClientSupportedSecurityProvidersOptions) (*VirtualWanSecurityProvidersResponse, error) {
	req, err := client.SupportedSecurityProvidersCreateRequest(ctx, resourceGroupName, virtualWanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.SupportedSecurityProvidersHandleError(resp)
	}
	result, err := client.SupportedSecurityProvidersHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SupportedSecurityProvidersCreateRequest creates the SupportedSecurityProviders request.
func (client *NetworkManagementClient) SupportedSecurityProvidersCreateRequest(ctx context.Context, resourceGroupName string, virtualWanName string, options *NetworkManagementClientSupportedSecurityProvidersOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWanName))
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

// SupportedSecurityProvidersHandleResponse handles the SupportedSecurityProviders response.
func (client *NetworkManagementClient) SupportedSecurityProvidersHandleResponse(resp *azcore.Response) (*VirtualWanSecurityProvidersResponse, error) {
	result := VirtualWanSecurityProvidersResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualWanSecurityProviders)
}

// SupportedSecurityProvidersHandleError handles the SupportedSecurityProviders error response.
func (client *NetworkManagementClient) SupportedSecurityProvidersHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
