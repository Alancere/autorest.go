// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"accessgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// AccessPublicOperationServer is a fake server for instances of the accessgroup.AccessPublicOperationClient type.
type AccessPublicOperationServer struct {
	// NoDecoratorInPublic is the fake for method AccessPublicOperationClient.NoDecoratorInPublic
	// HTTP status codes to indicate success: http.StatusOK
	NoDecoratorInPublic func(ctx context.Context, name string, options *accessgroup.AccessPublicOperationClientNoDecoratorInPublicOptions) (resp azfake.Responder[accessgroup.AccessPublicOperationClientNoDecoratorInPublicResponse], errResp azfake.ErrorResponder)

	// PublicDecoratorInPublic is the fake for method AccessPublicOperationClient.PublicDecoratorInPublic
	// HTTP status codes to indicate success: http.StatusOK
	PublicDecoratorInPublic func(ctx context.Context, name string, options *accessgroup.AccessPublicOperationClientPublicDecoratorInPublicOptions) (resp azfake.Responder[accessgroup.AccessPublicOperationClientPublicDecoratorInPublicResponse], errResp azfake.ErrorResponder)
}

// NewAccessPublicOperationServerTransport creates a new instance of AccessPublicOperationServerTransport with the provided implementation.
// The returned AccessPublicOperationServerTransport instance is connected to an instance of accessgroup.AccessPublicOperationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessPublicOperationServerTransport(srv *AccessPublicOperationServer) *AccessPublicOperationServerTransport {
	return &AccessPublicOperationServerTransport{srv: srv}
}

// AccessPublicOperationServerTransport connects instances of accessgroup.AccessPublicOperationClient to instances of AccessPublicOperationServer.
// Don't use this type directly, use NewAccessPublicOperationServerTransport instead.
type AccessPublicOperationServerTransport struct {
	srv *AccessPublicOperationServer
}

// Do implements the policy.Transporter interface for AccessPublicOperationServerTransport.
func (a *AccessPublicOperationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AccessPublicOperationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if accessPublicOperationServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = accessPublicOperationServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AccessPublicOperationClient.NoDecoratorInPublic":
				res.resp, res.err = a.dispatchNoDecoratorInPublic(req)
			case "AccessPublicOperationClient.PublicDecoratorInPublic":
				res.resp, res.err = a.dispatchPublicDecoratorInPublic(req)
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

func (a *AccessPublicOperationServerTransport) dispatchNoDecoratorInPublic(req *http.Request) (*http.Response, error) {
	if a.srv.NoDecoratorInPublic == nil {
		return nil, &nonRetriableError{errors.New("fake for method NoDecoratorInPublic not implemented")}
	}
	qp := req.URL.Query()
	nameParam, err := url.QueryUnescape(qp.Get("name"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.NoDecoratorInPublic(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NoDecoratorModelInPublic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessPublicOperationServerTransport) dispatchPublicDecoratorInPublic(req *http.Request) (*http.Response, error) {
	if a.srv.PublicDecoratorInPublic == nil {
		return nil, &nonRetriableError{errors.New("fake for method PublicDecoratorInPublic not implemented")}
	}
	qp := req.URL.Query()
	nameParam, err := url.QueryUnescape(qp.Get("name"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.PublicDecoratorInPublic(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicDecoratorModelInPublic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AccessPublicOperationServerTransport
var accessPublicOperationServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
