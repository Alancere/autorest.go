// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"azurepagegroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// PageServer is a fake server for instances of the azurepagegroup.PageClient type.
type PageServer struct {
	// PageTwoModelsAsPageItemServer contains the fakes for client PageTwoModelsAsPageItemClient
	PageTwoModelsAsPageItemServer PageTwoModelsAsPageItemServer

	// NewListWithCustomPageModelPager is the fake for method PageClient.NewListWithCustomPageModelPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWithCustomPageModelPager func(options *azurepagegroup.PageClientListWithCustomPageModelOptions) (resp azfake.PagerResponder[azurepagegroup.PageClientListWithCustomPageModelResponse])

	// NewListWithPagePager is the fake for method PageClient.NewListWithPagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWithPagePager func(options *azurepagegroup.PageClientListWithPageOptions) (resp azfake.PagerResponder[azurepagegroup.PageClientListWithPageResponse])

	// NewListWithParametersPager is the fake for method PageClient.NewListWithParametersPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWithParametersPager func(bodyInput azurepagegroup.ListItemInputBody, options *azurepagegroup.PageClientListWithParametersOptions) (resp azfake.PagerResponder[azurepagegroup.PageClientListWithParametersResponse])
}

// NewPageServerTransport creates a new instance of PageServerTransport with the provided implementation.
// The returned PageServerTransport instance is connected to an instance of azurepagegroup.PageClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPageServerTransport(srv *PageServer) *PageServerTransport {
	return &PageServerTransport{
		srv:                             srv,
		newListWithCustomPageModelPager: newTracker[azfake.PagerResponder[azurepagegroup.PageClientListWithCustomPageModelResponse]](),
		newListWithPagePager:            newTracker[azfake.PagerResponder[azurepagegroup.PageClientListWithPageResponse]](),
		newListWithParametersPager:      newTracker[azfake.PagerResponder[azurepagegroup.PageClientListWithParametersResponse]](),
	}
}

// PageServerTransport connects instances of azurepagegroup.PageClient to instances of PageServer.
// Don't use this type directly, use NewPageServerTransport instead.
type PageServerTransport struct {
	srv                             *PageServer
	trMu                            sync.Mutex
	trPageTwoModelsAsPageItemServer *PageTwoModelsAsPageItemServerTransport
	newListWithCustomPageModelPager *tracker[azfake.PagerResponder[azurepagegroup.PageClientListWithCustomPageModelResponse]]
	newListWithPagePager            *tracker[azfake.PagerResponder[azurepagegroup.PageClientListWithPageResponse]]
	newListWithParametersPager      *tracker[azfake.PagerResponder[azurepagegroup.PageClientListWithParametersResponse]]
}

// Do implements the policy.Transporter interface for PageServerTransport.
func (p *PageServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	if client := method[:strings.Index(method, ".")]; client != "PageClient" {
		return p.dispatchToClientFake(req, client)
	}
	return p.dispatchToMethodFake(req, method)
}

func (p *PageServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "PageTwoModelsAsPageItemClient":
		initServer(&p.trMu, &p.trPageTwoModelsAsPageItemServer, func() *PageTwoModelsAsPageItemServerTransport {
			return NewPageTwoModelsAsPageItemServerTransport(&p.srv.PageTwoModelsAsPageItemServer)
		})
		resp, err = p.trPageTwoModelsAsPageItemServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

func (p *PageServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "PageClient.NewListWithCustomPageModelPager":
		resp, err = p.dispatchNewListWithCustomPageModelPager(req)
	case "PageClient.NewListWithPagePager":
		resp, err = p.dispatchNewListWithPagePager(req)
	case "PageClient.NewListWithParametersPager":
		resp, err = p.dispatchNewListWithParametersPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (p *PageServerTransport) dispatchNewListWithCustomPageModelPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListWithCustomPageModelPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWithCustomPageModelPager not implemented")}
	}
	newListWithCustomPageModelPager := p.newListWithCustomPageModelPager.get(req)
	if newListWithCustomPageModelPager == nil {
		resp := p.srv.NewListWithCustomPageModelPager(nil)
		newListWithCustomPageModelPager = &resp
		p.newListWithCustomPageModelPager.add(req, newListWithCustomPageModelPager)
		server.PagerResponderInjectNextLinks(newListWithCustomPageModelPager, req, func(page *azurepagegroup.PageClientListWithCustomPageModelResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWithCustomPageModelPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListWithCustomPageModelPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWithCustomPageModelPager) {
		p.newListWithCustomPageModelPager.remove(req)
	}
	return resp, nil
}

func (p *PageServerTransport) dispatchNewListWithPagePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListWithPagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWithPagePager not implemented")}
	}
	newListWithPagePager := p.newListWithPagePager.get(req)
	if newListWithPagePager == nil {
		resp := p.srv.NewListWithPagePager(nil)
		newListWithPagePager = &resp
		p.newListWithPagePager.add(req, newListWithPagePager)
		server.PagerResponderInjectNextLinks(newListWithPagePager, req, func(page *azurepagegroup.PageClientListWithPageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWithPagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListWithPagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWithPagePager) {
		p.newListWithPagePager.remove(req)
	}
	return resp, nil
}

func (p *PageServerTransport) dispatchNewListWithParametersPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListWithParametersPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWithParametersPager not implemented")}
	}
	newListWithParametersPager := p.newListWithParametersPager.get(req)
	if newListWithParametersPager == nil {
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[azurepagegroup.ListItemInputBody](req)
		if err != nil {
			return nil, err
		}
		anotherUnescaped, err := url.QueryUnescape(qp.Get("another"))
		if err != nil {
			return nil, err
		}
		anotherParam := getOptional(azurepagegroup.ListItemInputExtensibleEnum(anotherUnescaped))
		var options *azurepagegroup.PageClientListWithParametersOptions
		if anotherParam != nil {
			options = &azurepagegroup.PageClientListWithParametersOptions{
				Another: anotherParam,
			}
		}
		resp := p.srv.NewListWithParametersPager(body, options)
		newListWithParametersPager = &resp
		p.newListWithParametersPager.add(req, newListWithParametersPager)
		server.PagerResponderInjectNextLinks(newListWithParametersPager, req, func(page *azurepagegroup.PageClientListWithParametersResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWithParametersPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListWithParametersPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWithParametersPager) {
		p.newListWithParametersPager.remove(req)
	}
	return resp, nil
}
