// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// BoolClient contains the methods for the Bool group.
// Don't use this type directly, use NewBoolClient() instead.
type BoolClient struct {
	con *Connection
}

// NewBoolClient creates a new instance of BoolClient with the specified values.
func NewBoolClient(con *Connection) BoolClient {
	return BoolClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client BoolClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetFalse - Get false Boolean value
func (client BoolClient) GetFalse(ctx context.Context, options *BoolGetFalseOptions) (BoolResponse, error) {
	req, err := client.getFalseCreateRequest(ctx, options)
	if err != nil {
		return BoolResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BoolResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolResponse{}, client.getFalseHandleError(resp)
	}
	result, err := client.getFalseHandleResponse(resp)
	if err != nil {
		return BoolResponse{}, err
	}
	return result, nil
}

// getFalseCreateRequest creates the GetFalse request.
func (client BoolClient) getFalseCreateRequest(ctx context.Context, options *BoolGetFalseOptions) (*azcore.Request, error) {
	urlPath := "/bool/false"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getFalseHandleResponse handles the GetFalse response.
func (client BoolClient) getFalseHandleResponse(resp *azcore.Response) (BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// getFalseHandleError handles the GetFalse error response.
func (client BoolClient) getFalseHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetInvalid - Get invalid Boolean value
func (client BoolClient) GetInvalid(ctx context.Context, options *BoolGetInvalidOptions) (BoolResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return BoolResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BoolResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolResponse{}, client.getInvalidHandleError(resp)
	}
	result, err := client.getInvalidHandleResponse(resp)
	if err != nil {
		return BoolResponse{}, err
	}
	return result, nil
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client BoolClient) getInvalidCreateRequest(ctx context.Context, options *BoolGetInvalidOptions) (*azcore.Request, error) {
	urlPath := "/bool/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client BoolClient) getInvalidHandleResponse(resp *azcore.Response) (BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// getInvalidHandleError handles the GetInvalid error response.
func (client BoolClient) getInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get null Boolean value
func (client BoolClient) GetNull(ctx context.Context, options *BoolGetNullOptions) (BoolResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return BoolResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BoolResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolResponse{}, client.getNullHandleError(resp)
	}
	result, err := client.getNullHandleResponse(resp)
	if err != nil {
		return BoolResponse{}, err
	}
	return result, nil
}

// getNullCreateRequest creates the GetNull request.
func (client BoolClient) getNullCreateRequest(ctx context.Context, options *BoolGetNullOptions) (*azcore.Request, error) {
	urlPath := "/bool/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client BoolClient) getNullHandleResponse(resp *azcore.Response) (BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// getNullHandleError handles the GetNull error response.
func (client BoolClient) getNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetTrue - Get true Boolean value
func (client BoolClient) GetTrue(ctx context.Context, options *BoolGetTrueOptions) (BoolResponse, error) {
	req, err := client.getTrueCreateRequest(ctx, options)
	if err != nil {
		return BoolResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BoolResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolResponse{}, client.getTrueHandleError(resp)
	}
	result, err := client.getTrueHandleResponse(resp)
	if err != nil {
		return BoolResponse{}, err
	}
	return result, nil
}

// getTrueCreateRequest creates the GetTrue request.
func (client BoolClient) getTrueCreateRequest(ctx context.Context, options *BoolGetTrueOptions) (*azcore.Request, error) {
	urlPath := "/bool/true"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTrueHandleResponse handles the GetTrue response.
func (client BoolClient) getTrueHandleResponse(resp *azcore.Response) (BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// getTrueHandleError handles the GetTrue error response.
func (client BoolClient) getTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutFalse - Set Boolean value false
func (client BoolClient) PutFalse(ctx context.Context, options *BoolPutFalseOptions) (*http.Response, error) {
	req, err := client.putFalseCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putFalseHandleError(resp)
	}
	return resp.Response, nil
}

// putFalseCreateRequest creates the PutFalse request.
func (client BoolClient) putFalseCreateRequest(ctx context.Context, options *BoolPutFalseOptions) (*azcore.Request, error) {
	urlPath := "/bool/false"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(false)
}

// putFalseHandleError handles the PutFalse error response.
func (client BoolClient) putFalseHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutTrue - Set Boolean value true
func (client BoolClient) PutTrue(ctx context.Context, options *BoolPutTrueOptions) (*http.Response, error) {
	req, err := client.putTrueCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putTrueHandleError(resp)
	}
	return resp.Response, nil
}

// putTrueCreateRequest creates the PutTrue request.
func (client BoolClient) putTrueCreateRequest(ctx context.Context, options *BoolPutTrueOptions) (*azcore.Request, error) {
	urlPath := "/bool/true"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// putTrueHandleError handles the PutTrue error response.
func (client BoolClient) putTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
