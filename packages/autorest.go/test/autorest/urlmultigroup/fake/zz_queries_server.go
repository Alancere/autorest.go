// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/urlmultigroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// QueriesServer is a fake server for instances of the urlmultigroup.QueriesClient type.
type QueriesServer struct {
	// ArrayStringMultiEmpty is the fake for method QueriesClient.ArrayStringMultiEmpty
	// HTTP status codes to indicate success: http.StatusOK
	ArrayStringMultiEmpty func(ctx context.Context, options *urlmultigroup.QueriesClientArrayStringMultiEmptyOptions) (resp azfake.Responder[urlmultigroup.QueriesClientArrayStringMultiEmptyResponse], errResp azfake.ErrorResponder)

	// ArrayStringMultiNull is the fake for method QueriesClient.ArrayStringMultiNull
	// HTTP status codes to indicate success: http.StatusOK
	ArrayStringMultiNull func(ctx context.Context, options *urlmultigroup.QueriesClientArrayStringMultiNullOptions) (resp azfake.Responder[urlmultigroup.QueriesClientArrayStringMultiNullResponse], errResp azfake.ErrorResponder)

	// ArrayStringMultiValid is the fake for method QueriesClient.ArrayStringMultiValid
	// HTTP status codes to indicate success: http.StatusOK
	ArrayStringMultiValid func(ctx context.Context, options *urlmultigroup.QueriesClientArrayStringMultiValidOptions) (resp azfake.Responder[urlmultigroup.QueriesClientArrayStringMultiValidResponse], errResp azfake.ErrorResponder)
}

// NewQueriesServerTransport creates a new instance of QueriesServerTransport with the provided implementation.
// The returned QueriesServerTransport instance is connected to an instance of urlmultigroup.QueriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQueriesServerTransport(srv *QueriesServer) *QueriesServerTransport {
	return &QueriesServerTransport{srv: srv}
}

// QueriesServerTransport connects instances of urlmultigroup.QueriesClient to instances of QueriesServer.
// Don't use this type directly, use NewQueriesServerTransport instead.
type QueriesServerTransport struct {
	srv *QueriesServer
}

// Do implements the policy.Transporter interface for QueriesServerTransport.
func (q *QueriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return q.dispatchToMethodFake(req, method)
}

func (q *QueriesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if queriesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = queriesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "QueriesClient.ArrayStringMultiEmpty":
				res.resp, res.err = q.dispatchArrayStringMultiEmpty(req)
			case "QueriesClient.ArrayStringMultiNull":
				res.resp, res.err = q.dispatchArrayStringMultiNull(req)
			case "QueriesClient.ArrayStringMultiValid":
				res.resp, res.err = q.dispatchArrayStringMultiValid(req)
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

func (q *QueriesServerTransport) dispatchArrayStringMultiEmpty(req *http.Request) (*http.Response, error) {
	if q.srv.ArrayStringMultiEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method ArrayStringMultiEmpty not implemented")}
	}
	qp := req.URL.Query()
	arrayQueryEscaped := qp["arrayQuery"]
	arrayQueryParam := make([]string, len(arrayQueryEscaped))
	for i, v := range arrayQueryEscaped {
		u, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return nil, unescapeErr
		}
		arrayQueryParam[i] = u
	}
	var options *urlmultigroup.QueriesClientArrayStringMultiEmptyOptions
	if len(arrayQueryParam) > 0 {
		options = &urlmultigroup.QueriesClientArrayStringMultiEmptyOptions{
			ArrayQuery: arrayQueryParam,
		}
	}
	respr, errRespr := q.srv.ArrayStringMultiEmpty(req.Context(), options)
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

func (q *QueriesServerTransport) dispatchArrayStringMultiNull(req *http.Request) (*http.Response, error) {
	if q.srv.ArrayStringMultiNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method ArrayStringMultiNull not implemented")}
	}
	qp := req.URL.Query()
	arrayQueryEscaped := qp["arrayQuery"]
	arrayQueryParam := make([]string, len(arrayQueryEscaped))
	for i, v := range arrayQueryEscaped {
		u, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return nil, unescapeErr
		}
		arrayQueryParam[i] = u
	}
	var options *urlmultigroup.QueriesClientArrayStringMultiNullOptions
	if len(arrayQueryParam) > 0 {
		options = &urlmultigroup.QueriesClientArrayStringMultiNullOptions{
			ArrayQuery: arrayQueryParam,
		}
	}
	respr, errRespr := q.srv.ArrayStringMultiNull(req.Context(), options)
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

func (q *QueriesServerTransport) dispatchArrayStringMultiValid(req *http.Request) (*http.Response, error) {
	if q.srv.ArrayStringMultiValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method ArrayStringMultiValid not implemented")}
	}
	qp := req.URL.Query()
	arrayQueryEscaped := qp["arrayQuery"]
	arrayQueryParam := make([]string, len(arrayQueryEscaped))
	for i, v := range arrayQueryEscaped {
		u, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return nil, unescapeErr
		}
		arrayQueryParam[i] = u
	}
	var options *urlmultigroup.QueriesClientArrayStringMultiValidOptions
	if len(arrayQueryParam) > 0 {
		options = &urlmultigroup.QueriesClientArrayStringMultiValidOptions{
			ArrayQuery: arrayQueryParam,
		}
	}
	respr, errRespr := q.srv.ArrayStringMultiValid(req.Context(), options)
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

// set this to conditionally intercept incoming requests to QueriesServerTransport
var queriesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
