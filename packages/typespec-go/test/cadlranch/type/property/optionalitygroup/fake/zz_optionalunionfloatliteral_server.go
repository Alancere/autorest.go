// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"optionalitygroup"
)

// OptionalUnionFloatLiteralServer is a fake server for instances of the optionalitygroup.OptionalUnionFloatLiteralClient type.
type OptionalUnionFloatLiteralServer struct {
	// GetAll is the fake for method OptionalUnionFloatLiteralClient.GetAll
	// HTTP status codes to indicate success: http.StatusOK
	GetAll func(ctx context.Context, options *optionalitygroup.OptionalUnionFloatLiteralClientGetAllOptions) (resp azfake.Responder[optionalitygroup.OptionalUnionFloatLiteralClientGetAllResponse], errResp azfake.ErrorResponder)

	// GetDefault is the fake for method OptionalUnionFloatLiteralClient.GetDefault
	// HTTP status codes to indicate success: http.StatusOK
	GetDefault func(ctx context.Context, options *optionalitygroup.OptionalUnionFloatLiteralClientGetDefaultOptions) (resp azfake.Responder[optionalitygroup.OptionalUnionFloatLiteralClientGetDefaultResponse], errResp azfake.ErrorResponder)

	// PutAll is the fake for method OptionalUnionFloatLiteralClient.PutAll
	// HTTP status codes to indicate success: http.StatusNoContent
	PutAll func(ctx context.Context, body optionalitygroup.UnionFloatLiteralProperty, options *optionalitygroup.OptionalUnionFloatLiteralClientPutAllOptions) (resp azfake.Responder[optionalitygroup.OptionalUnionFloatLiteralClientPutAllResponse], errResp azfake.ErrorResponder)

	// PutDefault is the fake for method OptionalUnionFloatLiteralClient.PutDefault
	// HTTP status codes to indicate success: http.StatusNoContent
	PutDefault func(ctx context.Context, body optionalitygroup.UnionFloatLiteralProperty, options *optionalitygroup.OptionalUnionFloatLiteralClientPutDefaultOptions) (resp azfake.Responder[optionalitygroup.OptionalUnionFloatLiteralClientPutDefaultResponse], errResp azfake.ErrorResponder)
}

// NewOptionalUnionFloatLiteralServerTransport creates a new instance of OptionalUnionFloatLiteralServerTransport with the provided implementation.
// The returned OptionalUnionFloatLiteralServerTransport instance is connected to an instance of optionalitygroup.OptionalUnionFloatLiteralClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOptionalUnionFloatLiteralServerTransport(srv *OptionalUnionFloatLiteralServer) *OptionalUnionFloatLiteralServerTransport {
	return &OptionalUnionFloatLiteralServerTransport{srv: srv}
}

// OptionalUnionFloatLiteralServerTransport connects instances of optionalitygroup.OptionalUnionFloatLiteralClient to instances of OptionalUnionFloatLiteralServer.
// Don't use this type directly, use NewOptionalUnionFloatLiteralServerTransport instead.
type OptionalUnionFloatLiteralServerTransport struct {
	srv *OptionalUnionFloatLiteralServer
}

// Do implements the policy.Transporter interface for OptionalUnionFloatLiteralServerTransport.
func (o *OptionalUnionFloatLiteralServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OptionalUnionFloatLiteralServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "OptionalUnionFloatLiteralClient.GetAll":
		resp, err = o.dispatchGetAll(req)
	case "OptionalUnionFloatLiteralClient.GetDefault":
		resp, err = o.dispatchGetDefault(req)
	case "OptionalUnionFloatLiteralClient.PutAll":
		resp, err = o.dispatchPutAll(req)
	case "OptionalUnionFloatLiteralClient.PutDefault":
		resp, err = o.dispatchPutDefault(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (o *OptionalUnionFloatLiteralServerTransport) dispatchGetAll(req *http.Request) (*http.Response, error) {
	if o.srv.GetAll == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAll not implemented")}
	}
	respr, errRespr := o.srv.GetAll(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UnionFloatLiteralProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OptionalUnionFloatLiteralServerTransport) dispatchGetDefault(req *http.Request) (*http.Response, error) {
	if o.srv.GetDefault == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDefault not implemented")}
	}
	respr, errRespr := o.srv.GetDefault(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UnionFloatLiteralProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OptionalUnionFloatLiteralServerTransport) dispatchPutAll(req *http.Request) (*http.Response, error) {
	if o.srv.PutAll == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutAll not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalitygroup.UnionFloatLiteralProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.PutAll(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OptionalUnionFloatLiteralServerTransport) dispatchPutDefault(req *http.Request) (*http.Response, error) {
	if o.srv.PutDefault == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutDefault not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalitygroup.UnionFloatLiteralProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.PutDefault(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
