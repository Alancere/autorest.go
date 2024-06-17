// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdevopsinfrastructure

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionUsagesClient contains the methods for the Microsoft.DevOpsInfrastructure namespace.
// Don't use this type directly, use NewSubscriptionUsagesClient() instead.
type SubscriptionUsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSubscriptionUsagesClient creates a new instance of SubscriptionUsagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubscriptionUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionUsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubscriptionUsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByLocationPager - List Quota resources by subscription ID
//
// Generated from API version 2024-04-04-preview
//   - locationName - Name of the location.
//   - options - SubscriptionUsagesClientListByLocationOptions contains the optional parameters for the SubscriptionUsagesClient.NewListByLocationPager
//     method.
func (client *SubscriptionUsagesClient) NewListByLocationPager(locationName string, options *SubscriptionUsagesClientListByLocationOptions) *runtime.Pager[SubscriptionUsagesClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubscriptionUsagesClientListByLocationResponse]{
		More: func(page SubscriptionUsagesClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionUsagesClientListByLocationResponse) (SubscriptionUsagesClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SubscriptionUsagesClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, locationName, options)
			}, nil)
			if err != nil {
				return SubscriptionUsagesClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *SubscriptionUsagesClient) listByLocationCreateRequest(ctx context.Context, locationName string, _ *SubscriptionUsagesClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevOpsInfrastructure/locations/{locationName}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *SubscriptionUsagesClient) listByLocationHandleResponse(resp *http.Response) (SubscriptionUsagesClientListByLocationResponse, error) {
	result := SubscriptionUsagesClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaListResult); err != nil {
		return SubscriptionUsagesClientListByLocationResponse{}, err
	}
	return result, nil
}