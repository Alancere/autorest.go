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
	"naminggroup"
	"net/http"
)

// NamingModelServer is a fake server for instances of the naminggroup.NamingModelClient type.
type NamingModelServer struct {
	// Client is the fake for method NamingModelClient.Client
	// HTTP status codes to indicate success: http.StatusNoContent
	Client func(ctx context.Context, clientModel naminggroup.ClientModel, options *naminggroup.NamingModelClientClientOptions) (resp azfake.Responder[naminggroup.NamingModelClientClientResponse], errResp azfake.ErrorResponder)

	// Language is the fake for method NamingModelClient.Language
	// HTTP status codes to indicate success: http.StatusNoContent
	Language func(ctx context.Context, goModel naminggroup.GoModel, options *naminggroup.NamingModelClientLanguageOptions) (resp azfake.Responder[naminggroup.NamingModelClientLanguageResponse], errResp azfake.ErrorResponder)
}

// NewNamingModelServerTransport creates a new instance of NamingModelServerTransport with the provided implementation.
// The returned NamingModelServerTransport instance is connected to an instance of naminggroup.NamingModelClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNamingModelServerTransport(srv *NamingModelServer) *NamingModelServerTransport {
	return &NamingModelServerTransport{srv: srv}
}

// NamingModelServerTransport connects instances of naminggroup.NamingModelClient to instances of NamingModelServer.
// Don't use this type directly, use NewNamingModelServerTransport instead.
type NamingModelServerTransport struct {
	srv *NamingModelServer
}

// Do implements the policy.Transporter interface for NamingModelServerTransport.
func (n *NamingModelServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NamingModelServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "NamingModelClient.Client":
		resp, err = n.dispatchClient(req)
	case "NamingModelClient.Language":
		resp, err = n.dispatchLanguage(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (n *NamingModelServerTransport) dispatchClient(req *http.Request) (*http.Response, error) {
	if n.srv.Client == nil {
		return nil, &nonRetriableError{errors.New("fake for method Client not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[naminggroup.ClientModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Client(req.Context(), body, nil)
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

func (n *NamingModelServerTransport) dispatchLanguage(req *http.Request) (*http.Response, error) {
	if n.srv.Language == nil {
		return nil, &nonRetriableError{errors.New("fake for method Language not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[naminggroup.GoModel](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Language(req.Context(), body, nil)
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