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
	"net/url"
	"scalargroup"
	"strconv"
)

// ScalarDecimal128TypeServer is a fake server for instances of the scalargroup.ScalarDecimal128TypeClient type.
type ScalarDecimal128TypeServer struct {
	// RequestBody is the fake for method ScalarDecimal128TypeClient.RequestBody
	// HTTP status codes to indicate success: http.StatusNoContent
	RequestBody func(ctx context.Context, body float64, options *scalargroup.ScalarDecimal128TypeClientRequestBodyOptions) (resp azfake.Responder[scalargroup.ScalarDecimal128TypeClientRequestBodyResponse], errResp azfake.ErrorResponder)

	// RequestParameter is the fake for method ScalarDecimal128TypeClient.RequestParameter
	// HTTP status codes to indicate success: http.StatusNoContent
	RequestParameter func(ctx context.Context, value float64, options *scalargroup.ScalarDecimal128TypeClientRequestParameterOptions) (resp azfake.Responder[scalargroup.ScalarDecimal128TypeClientRequestParameterResponse], errResp azfake.ErrorResponder)

	// ResponseBody is the fake for method ScalarDecimal128TypeClient.ResponseBody
	// HTTP status codes to indicate success: http.StatusOK
	ResponseBody func(ctx context.Context, options *scalargroup.ScalarDecimal128TypeClientResponseBodyOptions) (resp azfake.Responder[scalargroup.ScalarDecimal128TypeClientResponseBodyResponse], errResp azfake.ErrorResponder)
}

// NewScalarDecimal128TypeServerTransport creates a new instance of ScalarDecimal128TypeServerTransport with the provided implementation.
// The returned ScalarDecimal128TypeServerTransport instance is connected to an instance of scalargroup.ScalarDecimal128TypeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScalarDecimal128TypeServerTransport(srv *ScalarDecimal128TypeServer) *ScalarDecimal128TypeServerTransport {
	return &ScalarDecimal128TypeServerTransport{srv: srv}
}

// ScalarDecimal128TypeServerTransport connects instances of scalargroup.ScalarDecimal128TypeClient to instances of ScalarDecimal128TypeServer.
// Don't use this type directly, use NewScalarDecimal128TypeServerTransport instead.
type ScalarDecimal128TypeServerTransport struct {
	srv *ScalarDecimal128TypeServer
}

// Do implements the policy.Transporter interface for ScalarDecimal128TypeServerTransport.
func (s *ScalarDecimal128TypeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ScalarDecimal128TypeServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if scalarDecimal128TypeServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = scalarDecimal128TypeServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ScalarDecimal128TypeClient.RequestBody":
				res.resp, res.err = s.dispatchRequestBody(req)
			case "ScalarDecimal128TypeClient.RequestParameter":
				res.resp, res.err = s.dispatchRequestParameter(req)
			case "ScalarDecimal128TypeClient.ResponseBody":
				res.resp, res.err = s.dispatchResponseBody(req)
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

func (s *ScalarDecimal128TypeServerTransport) dispatchRequestBody(req *http.Request) (*http.Response, error) {
	if s.srv.RequestBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method RequestBody not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[float64](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.RequestBody(req.Context(), body, nil)
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

func (s *ScalarDecimal128TypeServerTransport) dispatchRequestParameter(req *http.Request) (*http.Response, error) {
	if s.srv.RequestParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method RequestParameter not implemented")}
	}
	qp := req.URL.Query()
	valueUnescaped, err := url.QueryUnescape(qp.Get("value"))
	if err != nil {
		return nil, err
	}
	valueParam, err := strconv.ParseFloat(valueUnescaped, 64)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.RequestParameter(req.Context(), valueParam, nil)
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

func (s *ScalarDecimal128TypeServerTransport) dispatchResponseBody(req *http.Request) (*http.Response, error) {
	if s.srv.ResponseBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method ResponseBody not implemented")}
	}
	respr, errRespr := s.srv.ResponseBody(req.Context(), nil)
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

// set this to conditionally intercept incoming requests to ScalarDecimal128TypeServerTransport
var scalarDecimal128TypeServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
