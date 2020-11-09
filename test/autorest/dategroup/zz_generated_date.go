// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dategroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// DateOperations contains the methods for the Date group.
type DateOperations interface {
	// GetInvalidDate - Get invalid date value
	GetInvalidDate(ctx context.Context, options *DateGetInvalidDateOptions) (*TimeResponse, error)
	// GetMaxDate - Get max date value 9999-12-31
	GetMaxDate(ctx context.Context, options *DateGetMaxDateOptions) (*TimeResponse, error)
	// GetMinDate - Get min date value 0000-01-01
	GetMinDate(ctx context.Context, options *DateGetMinDateOptions) (*TimeResponse, error)
	// GetNull - Get null date value
	GetNull(ctx context.Context, options *DateGetNullOptions) (*TimeResponse, error)
	// GetOverflowDate - Get overflow date value
	GetOverflowDate(ctx context.Context, options *DateGetOverflowDateOptions) (*TimeResponse, error)
	// GetUnderflowDate - Get underflow date value
	GetUnderflowDate(ctx context.Context, options *DateGetUnderflowDateOptions) (*TimeResponse, error)
	// PutMaxDate - Put max date value 9999-12-31
	PutMaxDate(ctx context.Context, dateBody time.Time, options *DatePutMaxDateOptions) (*http.Response, error)
	// PutMinDate - Put min date value 0000-01-01
	PutMinDate(ctx context.Context, dateBody time.Time, options *DatePutMinDateOptions) (*http.Response, error)
}

// DateClient implements the DateOperations interface.
// Don't use this type directly, use NewDateClient() instead.
type DateClient struct {
	con *Connection
}

// NewDateClient creates a new instance of DateClient with the specified values.
func NewDateClient(con *Connection) DateOperations {
	return &DateClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *DateClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetInvalidDate - Get invalid date value
func (client *DateClient) GetInvalidDate(ctx context.Context, options *DateGetInvalidDateOptions) (*TimeResponse, error) {
	req, err := client.GetInvalidDateCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetInvalidDateHandleError(resp)
	}
	result, err := client.GetInvalidDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInvalidDateCreateRequest creates the GetInvalidDate request.
func (client *DateClient) GetInvalidDateCreateRequest(ctx context.Context, options *DateGetInvalidDateOptions) (*azcore.Request, error) {
	urlPath := "/date/invaliddate"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetInvalidDateHandleResponse handles the GetInvalidDate response.
func (client *DateClient) GetInvalidDateHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *dateType
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetInvalidDateHandleError handles the GetInvalidDate error response.
func (client *DateClient) GetInvalidDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetMaxDate - Get max date value 9999-12-31
func (client *DateClient) GetMaxDate(ctx context.Context, options *DateGetMaxDateOptions) (*TimeResponse, error) {
	req, err := client.GetMaxDateCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetMaxDateHandleError(resp)
	}
	result, err := client.GetMaxDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMaxDateCreateRequest creates the GetMaxDate request.
func (client *DateClient) GetMaxDateCreateRequest(ctx context.Context, options *DateGetMaxDateOptions) (*azcore.Request, error) {
	urlPath := "/date/max"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetMaxDateHandleResponse handles the GetMaxDate response.
func (client *DateClient) GetMaxDateHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *dateType
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetMaxDateHandleError handles the GetMaxDate error response.
func (client *DateClient) GetMaxDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetMinDate - Get min date value 0000-01-01
func (client *DateClient) GetMinDate(ctx context.Context, options *DateGetMinDateOptions) (*TimeResponse, error) {
	req, err := client.GetMinDateCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetMinDateHandleError(resp)
	}
	result, err := client.GetMinDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMinDateCreateRequest creates the GetMinDate request.
func (client *DateClient) GetMinDateCreateRequest(ctx context.Context, options *DateGetMinDateOptions) (*azcore.Request, error) {
	urlPath := "/date/min"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetMinDateHandleResponse handles the GetMinDate response.
func (client *DateClient) GetMinDateHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *dateType
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetMinDateHandleError handles the GetMinDate error response.
func (client *DateClient) GetMinDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get null date value
func (client *DateClient) GetNull(ctx context.Context, options *DateGetNullOptions) (*TimeResponse, error) {
	req, err := client.GetNullCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNullHandleError(resp)
	}
	result, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNullCreateRequest creates the GetNull request.
func (client *DateClient) GetNullCreateRequest(ctx context.Context, options *DateGetNullOptions) (*azcore.Request, error) {
	urlPath := "/date/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNullHandleResponse handles the GetNull response.
func (client *DateClient) GetNullHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *dateType
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetNullHandleError handles the GetNull error response.
func (client *DateClient) GetNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetOverflowDate - Get overflow date value
func (client *DateClient) GetOverflowDate(ctx context.Context, options *DateGetOverflowDateOptions) (*TimeResponse, error) {
	req, err := client.GetOverflowDateCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetOverflowDateHandleError(resp)
	}
	result, err := client.GetOverflowDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOverflowDateCreateRequest creates the GetOverflowDate request.
func (client *DateClient) GetOverflowDateCreateRequest(ctx context.Context, options *DateGetOverflowDateOptions) (*azcore.Request, error) {
	urlPath := "/date/overflowdate"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetOverflowDateHandleResponse handles the GetOverflowDate response.
func (client *DateClient) GetOverflowDateHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *dateType
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetOverflowDateHandleError handles the GetOverflowDate error response.
func (client *DateClient) GetOverflowDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetUnderflowDate - Get underflow date value
func (client *DateClient) GetUnderflowDate(ctx context.Context, options *DateGetUnderflowDateOptions) (*TimeResponse, error) {
	req, err := client.GetUnderflowDateCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetUnderflowDateHandleError(resp)
	}
	result, err := client.GetUnderflowDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUnderflowDateCreateRequest creates the GetUnderflowDate request.
func (client *DateClient) GetUnderflowDateCreateRequest(ctx context.Context, options *DateGetUnderflowDateOptions) (*azcore.Request, error) {
	urlPath := "/date/underflowdate"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetUnderflowDateHandleResponse handles the GetUnderflowDate response.
func (client *DateClient) GetUnderflowDateHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *dateType
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetUnderflowDateHandleError handles the GetUnderflowDate error response.
func (client *DateClient) GetUnderflowDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutMaxDate - Put max date value 9999-12-31
func (client *DateClient) PutMaxDate(ctx context.Context, dateBody time.Time, options *DatePutMaxDateOptions) (*http.Response, error) {
	req, err := client.PutMaxDateCreateRequest(ctx, dateBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutMaxDateHandleError(resp)
	}
	return resp.Response, nil
}

// PutMaxDateCreateRequest creates the PutMaxDate request.
func (client *DateClient) PutMaxDateCreateRequest(ctx context.Context, dateBody time.Time, options *DatePutMaxDateOptions) (*azcore.Request, error) {
	urlPath := "/date/max"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(dateType(dateBody))
}

// PutMaxDateHandleError handles the PutMaxDate error response.
func (client *DateClient) PutMaxDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutMinDate - Put min date value 0000-01-01
func (client *DateClient) PutMinDate(ctx context.Context, dateBody time.Time, options *DatePutMinDateOptions) (*http.Response, error) {
	req, err := client.PutMinDateCreateRequest(ctx, dateBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutMinDateHandleError(resp)
	}
	return resp.Response, nil
}

// PutMinDateCreateRequest creates the PutMinDate request.
func (client *DateClient) PutMinDateCreateRequest(ctx context.Context, dateBody time.Time, options *DatePutMinDateOptions) (*azcore.Request, error) {
	urlPath := "/date/min"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(dateType(dateBody))
}

// PutMinDateHandleError handles the PutMinDate error response.
func (client *DateClient) PutMinDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
