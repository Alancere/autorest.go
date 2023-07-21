//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ServiceTagInformationServer is a fake server for instances of the armnetwork.ServiceTagInformationClient type.
type ServiceTagInformationServer struct {
	// NewListPager is the fake for method ServiceTagInformationClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armnetwork.ServiceTagInformationClientListOptions) (resp azfake.PagerResponder[armnetwork.ServiceTagInformationClientListResponse])
}

// NewServiceTagInformationServerTransport creates a new instance of ServiceTagInformationServerTransport with the provided implementation.
// The returned ServiceTagInformationServerTransport instance is connected to an instance of armnetwork.ServiceTagInformationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceTagInformationServerTransport(srv *ServiceTagInformationServer) *ServiceTagInformationServerTransport {
	return &ServiceTagInformationServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.ServiceTagInformationClientListResponse]](),
	}
}

// ServiceTagInformationServerTransport connects instances of armnetwork.ServiceTagInformationClient to instances of ServiceTagInformationServer.
// Don't use this type directly, use NewServiceTagInformationServerTransport instead.
type ServiceTagInformationServerTransport struct {
	srv          *ServiceTagInformationServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.ServiceTagInformationClientListResponse]]
}

// Do implements the policy.Transporter interface for ServiceTagInformationServerTransport.
func (s *ServiceTagInformationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServiceTagInformationClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServiceTagInformationServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceTagDetails`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		noAddressPrefixesUnescaped, err := url.QueryUnescape(qp.Get("noAddressPrefixes"))
		if err != nil {
			return nil, err
		}
		noAddressPrefixesParam, err := parseOptional(noAddressPrefixesUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		tagNameUnescaped, err := url.QueryUnescape(qp.Get("tagName"))
		if err != nil {
			return nil, err
		}
		tagNameParam := getOptional(tagNameUnescaped)
		var options *armnetwork.ServiceTagInformationClientListOptions
		if noAddressPrefixesParam != nil || tagNameParam != nil {
			options = &armnetwork.ServiceTagInformationClientListOptions{
				NoAddressPrefixes: noAddressPrefixesParam,
				TagName:           tagNameParam,
			}
		}
		resp := s.srv.NewListPager(locationUnescaped, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ServiceTagInformationClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}
