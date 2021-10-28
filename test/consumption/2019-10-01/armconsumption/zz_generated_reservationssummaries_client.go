//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

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

// ReservationsSummariesClient contains the methods for the ReservationsSummaries group.
// Don't use this type directly, use NewReservationsSummariesClient() instead.
type ReservationsSummariesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewReservationsSummariesClient creates a new instance of ReservationsSummariesClient with the specified values.
func NewReservationsSummariesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ReservationsSummariesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ReservationsSummariesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Lists the reservations summaries for the defined scope daily or monthly grain.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReservationsSummariesClient) List(scope string, grain Datagrain, options *ReservationsSummariesListOptions) *ReservationsSummariesListPager {
	return &ReservationsSummariesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, grain, options)
		},
		advancer: func(ctx context.Context, resp ReservationsSummariesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationSummariesListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ReservationsSummariesClient) listCreateRequest(ctx context.Context, scope string, grain Datagrain, options *ReservationsSummariesListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/reservationSummaries"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("grain", string(grain))
	if options != nil && options.StartDate != nil {
		reqQP.Set("startDate", *options.StartDate)
	}
	if options != nil && options.EndDate != nil {
		reqQP.Set("endDate", *options.EndDate)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ReservationID != nil {
		reqQP.Set("reservationId", *options.ReservationID)
	}
	if options != nil && options.ReservationOrderID != nil {
		reqQP.Set("reservationOrderId", *options.ReservationOrderID)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReservationsSummariesClient) listHandleResponse(resp *http.Response) (ReservationsSummariesListResponse, error) {
	result := ReservationsSummariesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationSummariesListResult); err != nil {
		return ReservationsSummariesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ReservationsSummariesClient) listHandleError(resp *http.Response) error {
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

// ListByReservationOrder - Lists the reservations summaries for daily or monthly grain.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReservationsSummariesClient) ListByReservationOrder(reservationOrderID string, grain Datagrain, options *ReservationsSummariesListByReservationOrderOptions) *ReservationsSummariesListByReservationOrderPager {
	return &ReservationsSummariesListByReservationOrderPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByReservationOrderCreateRequest(ctx, reservationOrderID, grain, options)
		},
		advancer: func(ctx context.Context, resp ReservationsSummariesListByReservationOrderResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationSummariesListResult.NextLink)
		},
	}
}

// listByReservationOrderCreateRequest creates the ListByReservationOrder request.
func (client *ReservationsSummariesClient) listByReservationOrderCreateRequest(ctx context.Context, reservationOrderID string, grain Datagrain, options *ReservationsSummariesListByReservationOrderOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationSummaries"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("grain", string(grain))
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByReservationOrderHandleResponse handles the ListByReservationOrder response.
func (client *ReservationsSummariesClient) listByReservationOrderHandleResponse(resp *http.Response) (ReservationsSummariesListByReservationOrderResponse, error) {
	result := ReservationsSummariesListByReservationOrderResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationSummariesListResult); err != nil {
		return ReservationsSummariesListByReservationOrderResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByReservationOrderHandleError handles the ListByReservationOrder error response.
func (client *ReservationsSummariesClient) listByReservationOrderHandleError(resp *http.Response) error {
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

// ListByReservationOrderAndReservation - Lists the reservations summaries for daily or monthly grain.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReservationsSummariesClient) ListByReservationOrderAndReservation(reservationOrderID string, reservationID string, grain Datagrain, options *ReservationsSummariesListByReservationOrderAndReservationOptions) *ReservationsSummariesListByReservationOrderAndReservationPager {
	return &ReservationsSummariesListByReservationOrderAndReservationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByReservationOrderAndReservationCreateRequest(ctx, reservationOrderID, reservationID, grain, options)
		},
		advancer: func(ctx context.Context, resp ReservationsSummariesListByReservationOrderAndReservationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationSummariesListResult.NextLink)
		},
	}
}

// listByReservationOrderAndReservationCreateRequest creates the ListByReservationOrderAndReservation request.
func (client *ReservationsSummariesClient) listByReservationOrderAndReservationCreateRequest(ctx context.Context, reservationOrderID string, reservationID string, grain Datagrain, options *ReservationsSummariesListByReservationOrderAndReservationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationSummaries"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("grain", string(grain))
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByReservationOrderAndReservationHandleResponse handles the ListByReservationOrderAndReservation response.
func (client *ReservationsSummariesClient) listByReservationOrderAndReservationHandleResponse(resp *http.Response) (ReservationsSummariesListByReservationOrderAndReservationResponse, error) {
	result := ReservationsSummariesListByReservationOrderAndReservationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationSummariesListResult); err != nil {
		return ReservationsSummariesListByReservationOrderAndReservationResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByReservationOrderAndReservationHandleError handles the ListByReservationOrderAndReservation error response.
func (client *ReservationsSummariesClient) listByReservationOrderAndReservationHandleError(resp *http.Response) error {
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
