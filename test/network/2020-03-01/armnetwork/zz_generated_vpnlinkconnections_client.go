//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// VPNLinkConnectionsClient contains the methods for the VPNLinkConnections group.
// Don't use this type directly, use NewVPNLinkConnectionsClient() instead.
type VPNLinkConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNLinkConnectionsClient creates a new instance of VPNLinkConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNLinkConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VPNLinkConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &VPNLinkConnectionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// ListByVPNConnection - Retrieves all vpn site link connections for a particular virtual wan vpn gateway vpn connection.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// connectionName - The name of the vpn connection.
// options - VPNLinkConnectionsListByVPNConnectionOptions contains the optional parameters for the VPNLinkConnectionsClient.ListByVPNConnection
// method.
func (client *VPNLinkConnectionsClient) ListByVPNConnection(resourceGroupName string, gatewayName string, connectionName string, options *VPNLinkConnectionsListByVPNConnectionOptions) *VPNLinkConnectionsListByVPNConnectionPager {
	return &VPNLinkConnectionsListByVPNConnectionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVPNConnectionCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
		},
		advancer: func(ctx context.Context, resp VPNLinkConnectionsListByVPNConnectionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNSiteLinkConnectionsResult.NextLink)
		},
	}
}

// listByVPNConnectionCreateRequest creates the ListByVPNConnection request.
func (client *VPNLinkConnectionsClient) listByVPNConnectionCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNLinkConnectionsListByVPNConnectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVPNConnectionHandleResponse handles the ListByVPNConnection response.
func (client *VPNLinkConnectionsClient) listByVPNConnectionHandleResponse(resp *http.Response) (VPNLinkConnectionsListByVPNConnectionResponse, error) {
	result := VPNLinkConnectionsListByVPNConnectionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNSiteLinkConnectionsResult); err != nil {
		return VPNLinkConnectionsListByVPNConnectionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByVPNConnectionHandleError handles the ListByVPNConnection error response.
func (client *VPNLinkConnectionsClient) listByVPNConnectionHandleError(resp *http.Response) error {
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
