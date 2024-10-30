// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/httpinfrastructuregroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPFailureServer is a fake server for instances of the httpinfrastructuregroup.HTTPFailureClient type.
type HTTPFailureServer struct {
	// GetEmptyError is the fake for method HTTPFailureClient.GetEmptyError
	// HTTP status codes to indicate success: http.StatusOK
	GetEmptyError func(ctx context.Context, options *httpinfrastructuregroup.HTTPFailureClientGetEmptyErrorOptions) (resp azfake.Responder[httpinfrastructuregroup.HTTPFailureClientGetEmptyErrorResponse], errResp azfake.ErrorResponder)

	// GetNoModelEmpty is the fake for method HTTPFailureClient.GetNoModelEmpty
	// HTTP status codes to indicate success: http.StatusOK
	GetNoModelEmpty func(ctx context.Context, options *httpinfrastructuregroup.HTTPFailureClientGetNoModelEmptyOptions) (resp azfake.Responder[httpinfrastructuregroup.HTTPFailureClientGetNoModelEmptyResponse], errResp azfake.ErrorResponder)

	// GetNoModelError is the fake for method HTTPFailureClient.GetNoModelError
	// HTTP status codes to indicate success: http.StatusOK
	GetNoModelError func(ctx context.Context, options *httpinfrastructuregroup.HTTPFailureClientGetNoModelErrorOptions) (resp azfake.Responder[httpinfrastructuregroup.HTTPFailureClientGetNoModelErrorResponse], errResp azfake.ErrorResponder)
}

// NewHTTPFailureServerTransport creates a new instance of HTTPFailureServerTransport with the provided implementation.
// The returned HTTPFailureServerTransport instance is connected to an instance of httpinfrastructuregroup.HTTPFailureClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHTTPFailureServerTransport(srv *HTTPFailureServer) *HTTPFailureServerTransport {
	return &HTTPFailureServerTransport{srv: srv}
}

// HTTPFailureServerTransport connects instances of httpinfrastructuregroup.HTTPFailureClient to instances of HTTPFailureServer.
// Don't use this type directly, use NewHTTPFailureServerTransport instead.
type HTTPFailureServerTransport struct {
	srv *HTTPFailureServer
}

// Do implements the policy.Transporter interface for HTTPFailureServerTransport.
func (h *HTTPFailureServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HTTPFailureServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if httpFailureServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = httpFailureServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "HTTPFailureClient.GetEmptyError":
				res.resp, res.err = h.dispatchGetEmptyError(req)
			case "HTTPFailureClient.GetNoModelEmpty":
				res.resp, res.err = h.dispatchGetNoModelEmpty(req)
			case "HTTPFailureClient.GetNoModelError":
				res.resp, res.err = h.dispatchGetNoModelError(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (h *HTTPFailureServerTransport) dispatchGetEmptyError(req *http.Request) (*http.Response, error) {
	if h.srv.GetEmptyError == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEmptyError not implemented")}
	}
	respr, errRespr := h.srv.GetEmptyError(req.Context(), nil)
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

func (h *HTTPFailureServerTransport) dispatchGetNoModelEmpty(req *http.Request) (*http.Response, error) {
	if h.srv.GetNoModelEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNoModelEmpty not implemented")}
	}
	respr, errRespr := h.srv.GetNoModelEmpty(req.Context(), nil)
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

func (h *HTTPFailureServerTransport) dispatchGetNoModelError(req *http.Request) (*http.Response, error) {
	if h.srv.GetNoModelError == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNoModelError not implemented")}
	}
	respr, errRespr := h.srv.GetNoModelError(req.Context(), nil)
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

// set this to conditionally intercept incoming requests to HTTPFailureServerTransport
var httpFailureServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
