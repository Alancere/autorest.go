// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// VPNServerConfigurationsAssociatedWithVirtualWanClient contains the methods for the VPNServerConfigurationsAssociatedWithVirtualWan group.
// Don't use this type directly, use NewVPNServerConfigurationsAssociatedWithVirtualWanClient() instead.
type VPNServerConfigurationsAssociatedWithVirtualWanClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVPNServerConfigurationsAssociatedWithVirtualWanClient creates a new instance of VPNServerConfigurationsAssociatedWithVirtualWanClient with the specified values.
func NewVPNServerConfigurationsAssociatedWithVirtualWanClient(con *armcore.Connection, subscriptionID string) *VPNServerConfigurationsAssociatedWithVirtualWanClient {
	return &VPNServerConfigurationsAssociatedWithVirtualWanClient{con: con, subscriptionID: subscriptionID}
}

// BeginList - Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) BeginList(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (VPNServerConfigurationsResponsePollerResponse, error) {
	resp, err := client.list(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return VPNServerConfigurationsResponsePollerResponse{}, err
	}
	result := VPNServerConfigurationsResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VPNServerConfigurationsAssociatedWithVirtualWanClient.List", "location", resp, client.listHandleError)
	if err != nil {
		return VPNServerConfigurationsResponsePollerResponse{}, err
	}
	poller := &vpnServerConfigurationsResponsePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VPNServerConfigurationsResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeList creates a new VPNServerConfigurationsResponsePoller from the specified resume token.
// token - The value must come from a previous call to VPNServerConfigurationsResponsePoller.ResumeToken().
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) ResumeList(token string) (VPNServerConfigurationsResponsePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VPNServerConfigurationsAssociatedWithVirtualWanClient.List", token, client.listHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnServerConfigurationsResponsePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// List - Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) list(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (*azcore.Response, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, virtualWANName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listHandleError(resp)
	}
	return resp, nil
}

// listCreateRequest creates the List request.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, options *VPNServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnServerConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWANName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) listHandleResponse(resp *azcore.Response) (VPNServerConfigurationsResponseResponse, error) {
	var val *VPNServerConfigurationsResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNServerConfigurationsResponseResponse{}, err
	}
	return VPNServerConfigurationsResponseResponse{RawResponse: resp.Response, VPNServerConfigurationsResponse: val}, nil
}

// listHandleError handles the List error response.
func (client *VPNServerConfigurationsAssociatedWithVirtualWanClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
