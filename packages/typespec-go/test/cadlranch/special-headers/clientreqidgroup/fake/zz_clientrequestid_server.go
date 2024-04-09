// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"clientreqidgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ClientRequestIdServer is a fake server for instances of the clientreqidgroup.ClientRequestIdClient type.
type ClientRequestIdServer struct {
	// Get is the fake for method ClientRequestIdClient.Get
	// HTTP status codes to indicate success: http.StatusNoContent
	Get func(ctx context.Context, options *clientreqidgroup.ClientRequestIdClientGetOptions) (resp azfake.Responder[clientreqidgroup.ClientRequestIdClientGetResponse], errResp azfake.ErrorResponder)
}

// NewClientRequestIdServerTransport creates a new instance of ClientRequestIdServerTransport with the provided implementation.
// The returned ClientRequestIdServerTransport instance is connected to an instance of clientreqidgroup.ClientRequestIdClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClientRequestIdServerTransport(srv *ClientRequestIdServer) *ClientRequestIdServerTransport {
	return &ClientRequestIdServerTransport{srv: srv}
}

// ClientRequestIdServerTransport connects instances of clientreqidgroup.ClientRequestIdClient to instances of ClientRequestIdServer.
// Don't use this type directly, use NewClientRequestIdServerTransport instead.
type ClientRequestIdServerTransport struct {
	srv *ClientRequestIdServer
}

// Do implements the policy.Transporter interface for ClientRequestIdServerTransport.
func (c *ClientRequestIdServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ClientRequestIdServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ClientRequestIdClient.Get":
		resp, err = c.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (c *ClientRequestIdServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "client-request-id"))
	var options *clientreqidgroup.ClientRequestIdClientGetOptions
	if clientRequestIDParam != nil {
		options = &clientreqidgroup.ClientRequestIdClientGetOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := c.srv.Get(req.Context(), options)
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