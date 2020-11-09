// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionInCredentialsOperations contains the methods for the SubscriptionInCredentials group.
type SubscriptionInCredentialsOperations interface {
	// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456'
	// to succeed
	PostMethodGlobalNotProvidedValid(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalNotProvidedValidOptions) (*http.Response, error)
	// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to null, and client-side validation
	// should prevent you from making this call
	PostMethodGlobalNull(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalNullOptions) (*http.Response, error)
	// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostMethodGlobalValid(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalValidOptions) (*http.Response, error)
	// PostPathGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostPathGlobalValid(ctx context.Context, options *SubscriptionInCredentialsPostPathGlobalValidOptions) (*http.Response, error)
	// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostSwaggerGlobalValid(ctx context.Context, options *SubscriptionInCredentialsPostSwaggerGlobalValidOptions) (*http.Response, error)
}

// SubscriptionInCredentialsClient implements the SubscriptionInCredentialsOperations interface.
// Don't use this type directly, use NewSubscriptionInCredentialsClient() instead.
type SubscriptionInCredentialsClient struct {
	con            *Connection
	subscriptionID string
}

// NewSubscriptionInCredentialsClient creates a new instance of SubscriptionInCredentialsClient with the specified values.
func NewSubscriptionInCredentialsClient(con *Connection, subscriptionID string) SubscriptionInCredentialsOperations {
	return &SubscriptionInCredentialsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *SubscriptionInCredentialsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456'
// to succeed
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNotProvidedValid(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalNotProvidedValidOptions) (*http.Response, error) {
	req, err := client.PostMethodGlobalNotProvidedValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMethodGlobalNotProvidedValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostMethodGlobalNotProvidedValidCreateRequest creates the PostMethodGlobalNotProvidedValid request.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNotProvidedValidCreateRequest(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalNotProvidedValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/globalNotProvided/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2015-07-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostMethodGlobalNotProvidedValidHandleError handles the PostMethodGlobalNotProvidedValid error response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNotProvidedValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to null, and client-side validation
// should prevent you from making this call
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNull(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalNullOptions) (*http.Response, error) {
	req, err := client.PostMethodGlobalNullCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMethodGlobalNullHandleError(resp)
	}
	return resp.Response, nil
}

// PostMethodGlobalNullCreateRequest creates the PostMethodGlobalNull request.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNullCreateRequest(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalNullOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/global/null/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostMethodGlobalNullHandleError handles the PostMethodGlobalNull error response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *SubscriptionInCredentialsClient) PostMethodGlobalValid(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalValidOptions) (*http.Response, error) {
	req, err := client.PostMethodGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMethodGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostMethodGlobalValidCreateRequest creates the PostMethodGlobalValid request.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalValidCreateRequest(ctx context.Context, options *SubscriptionInCredentialsPostMethodGlobalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostMethodGlobalValidHandleError handles the PostMethodGlobalValid error response.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostPathGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *SubscriptionInCredentialsClient) PostPathGlobalValid(ctx context.Context, options *SubscriptionInCredentialsPostPathGlobalValidOptions) (*http.Response, error) {
	req, err := client.PostPathGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostPathGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostPathGlobalValidCreateRequest creates the PostPathGlobalValid request.
func (client *SubscriptionInCredentialsClient) PostPathGlobalValidCreateRequest(ctx context.Context, options *SubscriptionInCredentialsPostPathGlobalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/path/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostPathGlobalValidHandleError handles the PostPathGlobalValid error response.
func (client *SubscriptionInCredentialsClient) PostPathGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *SubscriptionInCredentialsClient) PostSwaggerGlobalValid(ctx context.Context, options *SubscriptionInCredentialsPostSwaggerGlobalValidOptions) (*http.Response, error) {
	req, err := client.PostSwaggerGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostSwaggerGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// PostSwaggerGlobalValidCreateRequest creates the PostSwaggerGlobalValid request.
func (client *SubscriptionInCredentialsClient) PostSwaggerGlobalValidCreateRequest(ctx context.Context, options *SubscriptionInCredentialsPostSwaggerGlobalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/swagger/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostSwaggerGlobalValidHandleError handles the PostSwaggerGlobalValid error response.
func (client *SubscriptionInCredentialsClient) PostSwaggerGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
