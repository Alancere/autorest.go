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
	"xmlgroup"
)

// XMLSimpleModelValueServer is a fake server for instances of the xmlgroup.XMLSimpleModelValueClient type.
type XMLSimpleModelValueServer struct {
	// Get is the fake for method XMLSimpleModelValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *xmlgroup.XMLSimpleModelValueClientGetOptions) (resp azfake.Responder[xmlgroup.XMLSimpleModelValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method XMLSimpleModelValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, input xmlgroup.SimpleModel, options *xmlgroup.XMLSimpleModelValueClientPutOptions) (resp azfake.Responder[xmlgroup.XMLSimpleModelValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewXMLSimpleModelValueServerTransport creates a new instance of XMLSimpleModelValueServerTransport with the provided implementation.
// The returned XMLSimpleModelValueServerTransport instance is connected to an instance of xmlgroup.XMLSimpleModelValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewXMLSimpleModelValueServerTransport(srv *XMLSimpleModelValueServer) *XMLSimpleModelValueServerTransport {
	return &XMLSimpleModelValueServerTransport{srv: srv}
}

// XMLSimpleModelValueServerTransport connects instances of xmlgroup.XMLSimpleModelValueClient to instances of XMLSimpleModelValueServer.
// Don't use this type directly, use NewXMLSimpleModelValueServerTransport instead.
type XMLSimpleModelValueServerTransport struct {
	srv *XMLSimpleModelValueServer
}

// Do implements the policy.Transporter interface for XMLSimpleModelValueServerTransport.
func (x *XMLSimpleModelValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return x.dispatchToMethodFake(req, method)
}

func (x *XMLSimpleModelValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if xmlSimpleModelValueServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = xmlSimpleModelValueServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "XMLSimpleModelValueClient.Get":
				res.resp, res.err = x.dispatchGet(req)
			case "XMLSimpleModelValueClient.Put":
				res.resp, res.err = x.dispatchPut(req)
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

func (x *XMLSimpleModelValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if x.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := x.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).SimpleModel, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("content-type", "application/xml")
	}
	return resp, nil
}

func (x *XMLSimpleModelValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if x.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsXML[xmlgroup.SimpleModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := x.srv.Put(req.Context(), body, nil)
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

// set this to conditionally intercept incoming requests to XMLSimpleModelValueServerTransport
var xmlSimpleModelValueServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
