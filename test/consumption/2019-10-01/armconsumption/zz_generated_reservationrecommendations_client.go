//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ReservationRecommendationsClient contains the methods for the ReservationRecommendations group.
// Don't use this type directly, use NewReservationRecommendationsClient() instead.
type ReservationRecommendationsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewReservationRecommendationsClient creates a new instance of ReservationRecommendationsClient with the specified values.
func NewReservationRecommendationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ReservationRecommendationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ReservationRecommendationsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - List of recommendations for purchasing reserved instances.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReservationRecommendationsClient) List(scope string, options *ReservationRecommendationsListOptions) *ReservationRecommendationsListPager {
	return &ReservationRecommendationsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp ReservationRecommendationsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationRecommendationsListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ReservationRecommendationsClient) listCreateRequest(ctx context.Context, scope string, options *ReservationRecommendationsListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/reservationRecommendations"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReservationRecommendationsClient) listHandleResponse(resp *http.Response) (ReservationRecommendationsListResponse, error) {
	result := ReservationRecommendationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationRecommendationsListResult); err != nil {
		return ReservationRecommendationsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ReservationRecommendationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
