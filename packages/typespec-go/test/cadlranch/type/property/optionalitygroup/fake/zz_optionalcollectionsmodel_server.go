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

// OptionalCollectionsModelServer is a fake server for instances of the optionalitygroup.OptionalCollectionsModelClient type.
type OptionalCollectionsModelServer struct {
	// GetAll is the fake for method OptionalCollectionsModelClient.GetAll
	// HTTP status codes to indicate success: http.StatusOK
	GetAll func(ctx context.Context, options *optionalitygroup.OptionalCollectionsModelClientGetAllOptions) (resp azfake.Responder[optionalitygroup.OptionalCollectionsModelClientGetAllResponse], errResp azfake.ErrorResponder)

	// GetDefault is the fake for method OptionalCollectionsModelClient.GetDefault
	// HTTP status codes to indicate success: http.StatusOK
	GetDefault func(ctx context.Context, options *optionalitygroup.OptionalCollectionsModelClientGetDefaultOptions) (resp azfake.Responder[optionalitygroup.OptionalCollectionsModelClientGetDefaultResponse], errResp azfake.ErrorResponder)

	// PutAll is the fake for method OptionalCollectionsModelClient.PutAll
	// HTTP status codes to indicate success: http.StatusNoContent
	PutAll func(ctx context.Context, body optionalitygroup.CollectionsModelProperty, options *optionalitygroup.OptionalCollectionsModelClientPutAllOptions) (resp azfake.Responder[optionalitygroup.OptionalCollectionsModelClientPutAllResponse], errResp azfake.ErrorResponder)

	// PutDefault is the fake for method OptionalCollectionsModelClient.PutDefault
	// HTTP status codes to indicate success: http.StatusNoContent
	PutDefault func(ctx context.Context, body optionalitygroup.CollectionsModelProperty, options *optionalitygroup.OptionalCollectionsModelClientPutDefaultOptions) (resp azfake.Responder[optionalitygroup.OptionalCollectionsModelClientPutDefaultResponse], errResp azfake.ErrorResponder)
}

// NewOptionalCollectionsModelServerTransport creates a new instance of OptionalCollectionsModelServerTransport with the provided implementation.
// The returned OptionalCollectionsModelServerTransport instance is connected to an instance of optionalitygroup.OptionalCollectionsModelClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOptionalCollectionsModelServerTransport(srv *OptionalCollectionsModelServer) *OptionalCollectionsModelServerTransport {
	return &OptionalCollectionsModelServerTransport{srv: srv}
}

// OptionalCollectionsModelServerTransport connects instances of optionalitygroup.OptionalCollectionsModelClient to instances of OptionalCollectionsModelServer.
// Don't use this type directly, use NewOptionalCollectionsModelServerTransport instead.
type OptionalCollectionsModelServerTransport struct {
	srv *OptionalCollectionsModelServer
}

// Do implements the policy.Transporter interface for OptionalCollectionsModelServerTransport.
func (o *OptionalCollectionsModelServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OptionalCollectionsModelServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "OptionalCollectionsModelClient.GetAll":
		resp, err = o.dispatchGetAll(req)
	case "OptionalCollectionsModelClient.GetDefault":
		resp, err = o.dispatchGetDefault(req)
	case "OptionalCollectionsModelClient.PutAll":
		resp, err = o.dispatchPutAll(req)
	case "OptionalCollectionsModelClient.PutDefault":
		resp, err = o.dispatchPutDefault(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (o *OptionalCollectionsModelServerTransport) dispatchGetAll(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CollectionsModelProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OptionalCollectionsModelServerTransport) dispatchGetDefault(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CollectionsModelProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OptionalCollectionsModelServerTransport) dispatchPutAll(req *http.Request) (*http.Response, error) {
	if o.srv.PutAll == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutAll not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalitygroup.CollectionsModelProperty](req)
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

func (o *OptionalCollectionsModelServerTransport) dispatchPutDefault(req *http.Request) (*http.Response, error) {
	if o.srv.PutDefault == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutDefault not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalitygroup.CollectionsModelProperty](req)
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
