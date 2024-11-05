// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"corescalargroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// ScalarAzureLocationScalarServer is a fake server for instances of the corescalargroup.ScalarAzureLocationScalarClient type.
type ScalarAzureLocationScalarServer struct {
	// Get is the fake for method ScalarAzureLocationScalarClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *corescalargroup.ScalarAzureLocationScalarClientGetOptions) (resp azfake.Responder[corescalargroup.ScalarAzureLocationScalarClientGetResponse], errResp azfake.ErrorResponder)

	// Header is the fake for method ScalarAzureLocationScalarClient.Header
	// HTTP status codes to indicate success: http.StatusNoContent
	Header func(ctx context.Context, region string, options *corescalargroup.ScalarAzureLocationScalarClientHeaderOptions) (resp azfake.Responder[corescalargroup.ScalarAzureLocationScalarClientHeaderResponse], errResp azfake.ErrorResponder)

	// Post is the fake for method ScalarAzureLocationScalarClient.Post
	// HTTP status codes to indicate success: http.StatusOK
	Post func(ctx context.Context, body corescalargroup.AzureLocationModel, options *corescalargroup.ScalarAzureLocationScalarClientPostOptions) (resp azfake.Responder[corescalargroup.ScalarAzureLocationScalarClientPostResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ScalarAzureLocationScalarClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body string, options *corescalargroup.ScalarAzureLocationScalarClientPutOptions) (resp azfake.Responder[corescalargroup.ScalarAzureLocationScalarClientPutResponse], errResp azfake.ErrorResponder)

	// Query is the fake for method ScalarAzureLocationScalarClient.Query
	// HTTP status codes to indicate success: http.StatusNoContent
	Query func(ctx context.Context, region string, options *corescalargroup.ScalarAzureLocationScalarClientQueryOptions) (resp azfake.Responder[corescalargroup.ScalarAzureLocationScalarClientQueryResponse], errResp azfake.ErrorResponder)
}

// NewScalarAzureLocationScalarServerTransport creates a new instance of ScalarAzureLocationScalarServerTransport with the provided implementation.
// The returned ScalarAzureLocationScalarServerTransport instance is connected to an instance of corescalargroup.ScalarAzureLocationScalarClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScalarAzureLocationScalarServerTransport(srv *ScalarAzureLocationScalarServer) *ScalarAzureLocationScalarServerTransport {
	return &ScalarAzureLocationScalarServerTransport{srv: srv}
}

// ScalarAzureLocationScalarServerTransport connects instances of corescalargroup.ScalarAzureLocationScalarClient to instances of ScalarAzureLocationScalarServer.
// Don't use this type directly, use NewScalarAzureLocationScalarServerTransport instead.
type ScalarAzureLocationScalarServerTransport struct {
	srv *ScalarAzureLocationScalarServer
}

// Do implements the policy.Transporter interface for ScalarAzureLocationScalarServerTransport.
func (s *ScalarAzureLocationScalarServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ScalarAzureLocationScalarServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if scalarAzureLocationScalarServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = scalarAzureLocationScalarServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ScalarAzureLocationScalarClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ScalarAzureLocationScalarClient.Header":
				res.resp, res.err = s.dispatchHeader(req)
			case "ScalarAzureLocationScalarClient.Post":
				res.resp, res.err = s.dispatchPost(req)
			case "ScalarAzureLocationScalarClient.Put":
				res.resp, res.err = s.dispatchPut(req)
			case "ScalarAzureLocationScalarClient.Query":
				res.resp, res.err = s.dispatchQuery(req)
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

func (s *ScalarAzureLocationScalarServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := s.srv.Get(req.Context(), nil)
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

func (s *ScalarAzureLocationScalarServerTransport) dispatchHeader(req *http.Request) (*http.Response, error) {
	if s.srv.Header == nil {
		return nil, &nonRetriableError{errors.New("fake for method Header not implemented")}
	}
	respr, errRespr := s.srv.Header(req.Context(), getHeaderValue(req.Header, "region"), nil)
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

func (s *ScalarAzureLocationScalarServerTransport) dispatchPost(req *http.Request) (*http.Response, error) {
	if s.srv.Post == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[corescalargroup.AzureLocationModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Post(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureLocationModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScalarAzureLocationScalarServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if s.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[string](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Put(req.Context(), body, nil)
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

func (s *ScalarAzureLocationScalarServerTransport) dispatchQuery(req *http.Request) (*http.Response, error) {
	if s.srv.Query == nil {
		return nil, &nonRetriableError{errors.New("fake for method Query not implemented")}
	}
	qp := req.URL.Query()
	regionParam, err := url.QueryUnescape(qp.Get("region"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Query(req.Context(), regionParam, nil)
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

// set this to conditionally intercept incoming requests to ScalarAzureLocationScalarServerTransport
var scalarAzureLocationScalarServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
