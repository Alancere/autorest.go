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
)

// BgpServiceCommunitiesOperations contains the methods for the BgpServiceCommunities group.
type BgpServiceCommunitiesOperations interface {
	// List - Gets all the available bgp service communities.
	List(options *BgpServiceCommunitiesListOptions) BgpServiceCommunityListResultPager
}

// BgpServiceCommunitiesClient implements the BgpServiceCommunitiesOperations interface.
// Don't use this type directly, use NewBgpServiceCommunitiesClient() instead.
type BgpServiceCommunitiesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewBgpServiceCommunitiesClient creates a new instance of BgpServiceCommunitiesClient with the specified values.
func NewBgpServiceCommunitiesClient(con *armcore.Connection, subscriptionID string) BgpServiceCommunitiesOperations {
	return &BgpServiceCommunitiesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *BgpServiceCommunitiesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// List - Gets all the available bgp service communities.
func (client *BgpServiceCommunitiesClient) List(options *BgpServiceCommunitiesListOptions) BgpServiceCommunityListResultPager {
	return &bgpServiceCommunityListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *BgpServiceCommunityListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.BgpServiceCommunityListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *BgpServiceCommunitiesClient) ListCreateRequest(ctx context.Context, options *BgpServiceCommunitiesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// ListHandleResponse handles the List response.
func (client *BgpServiceCommunitiesClient) ListHandleResponse(resp *azcore.Response) (*BgpServiceCommunityListResultResponse, error) {
	result := BgpServiceCommunityListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BgpServiceCommunityListResult)
}

// ListHandleError handles the List error response.
func (client *BgpServiceCommunitiesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
