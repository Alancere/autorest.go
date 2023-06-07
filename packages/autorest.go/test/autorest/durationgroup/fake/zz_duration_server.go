//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/durationgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DurationServer is a fake server for instances of the durationgroup.DurationClient type.
type DurationServer struct {
	// GetInvalid is the fake for method DurationClient.GetInvalid
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalid func(ctx context.Context, options *durationgroup.DurationClientGetInvalidOptions) (resp azfake.Responder[durationgroup.DurationClientGetInvalidResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method DurationClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *durationgroup.DurationClientGetNullOptions) (resp azfake.Responder[durationgroup.DurationClientGetNullResponse], errResp azfake.ErrorResponder)

	// GetPositiveDuration is the fake for method DurationClient.GetPositiveDuration
	// HTTP status codes to indicate success: http.StatusOK
	GetPositiveDuration func(ctx context.Context, options *durationgroup.DurationClientGetPositiveDurationOptions) (resp azfake.Responder[durationgroup.DurationClientGetPositiveDurationResponse], errResp azfake.ErrorResponder)

	// PutPositiveDuration is the fake for method DurationClient.PutPositiveDuration
	// HTTP status codes to indicate success: http.StatusOK
	PutPositiveDuration func(ctx context.Context, durationBody string, options *durationgroup.DurationClientPutPositiveDurationOptions) (resp azfake.Responder[durationgroup.DurationClientPutPositiveDurationResponse], errResp azfake.ErrorResponder)
}

// NewDurationServerTransport creates a new instance of DurationServerTransport with the provided implementation.
// The returned DurationServerTransport instance is connected to an instance of durationgroup.DurationClient by way of the
// undefined.Transporter field.
func NewDurationServerTransport(srv *DurationServer) *DurationServerTransport {
	return &DurationServerTransport{srv: srv}
}

// DurationServerTransport connects instances of durationgroup.DurationClient to instances of DurationServer.
// Don't use this type directly, use NewDurationServerTransport instead.
type DurationServerTransport struct {
	srv *DurationServer
}

// Do implements the policy.Transporter interface for DurationServerTransport.
func (d *DurationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DurationClient.GetInvalid":
		resp, err = d.dispatchGetInvalid(req)
	case "DurationClient.GetNull":
		resp, err = d.dispatchGetNull(req)
	case "DurationClient.GetPositiveDuration":
		resp, err = d.dispatchGetPositiveDuration(req)
	case "DurationClient.PutPositiveDuration":
		resp, err = d.dispatchPutPositiveDuration(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DurationServerTransport) dispatchGetInvalid(req *http.Request) (*http.Response, error) {
	if d.srv.GetInvalid == nil {
		return nil, &nonRetriableError{errors.New("method GetInvalid not implemented")}
	}
	respr, errRespr := d.srv.GetInvalid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DurationServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if d.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("method GetNull not implemented")}
	}
	respr, errRespr := d.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DurationServerTransport) dispatchGetPositiveDuration(req *http.Request) (*http.Response, error) {
	if d.srv.GetPositiveDuration == nil {
		return nil, &nonRetriableError{errors.New("method GetPositiveDuration not implemented")}
	}
	respr, errRespr := d.srv.GetPositiveDuration(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DurationServerTransport) dispatchPutPositiveDuration(req *http.Request) (*http.Response, error) {
	if d.srv.PutPositiveDuration == nil {
		return nil, &nonRetriableError{errors.New("method PutPositiveDuration not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[string](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PutPositiveDuration(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}