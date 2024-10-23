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

// XMLModelWithTextValueServer is a fake server for instances of the xmlgroup.XMLModelWithTextValueClient type.
type XMLModelWithTextValueServer struct {
	// Get is the fake for method XMLModelWithTextValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *xmlgroup.XMLModelWithTextValueClientGetOptions) (resp azfake.Responder[xmlgroup.XMLModelWithTextValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method XMLModelWithTextValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, input xmlgroup.ModelWithText, options *xmlgroup.XMLModelWithTextValueClientPutOptions) (resp azfake.Responder[xmlgroup.XMLModelWithTextValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewXMLModelWithTextValueServerTransport creates a new instance of XMLModelWithTextValueServerTransport with the provided implementation.
// The returned XMLModelWithTextValueServerTransport instance is connected to an instance of xmlgroup.XMLModelWithTextValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewXMLModelWithTextValueServerTransport(srv *XMLModelWithTextValueServer) *XMLModelWithTextValueServerTransport {
	return &XMLModelWithTextValueServerTransport{srv: srv}
}

// XMLModelWithTextValueServerTransport connects instances of xmlgroup.XMLModelWithTextValueClient to instances of XMLModelWithTextValueServer.
// Don't use this type directly, use NewXMLModelWithTextValueServerTransport instead.
type XMLModelWithTextValueServerTransport struct {
	srv *XMLModelWithTextValueServer
}

// Do implements the policy.Transporter interface for XMLModelWithTextValueServerTransport.
func (x *XMLModelWithTextValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return x.dispatchToMethodFake(req, method)
}

func (x *XMLModelWithTextValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "XMLModelWithTextValueClient.Get":
			res.resp, res.err = x.dispatchGet(req)
		case "XMLModelWithTextValueClient.Put":
			res.resp, res.err = x.dispatchPut(req)
		default:
			res.err = fmt.Errorf("unhandled API %s", method)
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

func (x *XMLModelWithTextValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsXML(respContent, server.GetResponse(respr).ModelWithText, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentType; val != nil {
		resp.Header.Set("content-type", "application/xml")
	}
	return resp, nil
}

func (x *XMLModelWithTextValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if x.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsXML[xmlgroup.ModelWithText](req)
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
