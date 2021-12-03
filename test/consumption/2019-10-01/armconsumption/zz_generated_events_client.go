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

// EventsClient contains the methods for the Events group.
// Don't use this type directly, use NewEventsClient() instead.
type EventsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewEventsClient creates a new instance of EventsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEventsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *EventsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &EventsClient{
		host: string(cp.Host),
		pl:   armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// List - Lists the events by billingAccountId and billingProfileId for given start and end date.
// If the operation fails it returns the *ErrorResponse error type.
// startDate - Start date
// endDate - End date
// scope - The scope associated with events operations. This includes '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfile/{billingProfileId}'
// for Billing Profile scope, and
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// options - EventsListOptions contains the optional parameters for the EventsClient.List method.
func (client *EventsClient) List(startDate string, endDate string, scope string, options *EventsListOptions) *EventsListPager {
	return &EventsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, startDate, endDate, scope, options)
		},
		advancer: func(ctx context.Context, resp EventsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Events.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *EventsClient) listCreateRequest(ctx context.Context, startDate string, endDate string, scope string, options *EventsListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/events"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	reqQP.Set("startDate", startDate)
	reqQP.Set("endDate", endDate)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EventsClient) listHandleResponse(resp *http.Response) (EventsListResponse, error) {
	result := EventsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *EventsClient) listHandleError(resp *http.Response) error {
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
