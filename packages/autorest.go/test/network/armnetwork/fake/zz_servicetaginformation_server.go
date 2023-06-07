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
// The returned ServiceTagInformationServerTransport instance is connected to an instance of armnetwork.ServiceTagInformationClient by way of the
// undefined.Transporter field.
func NewServiceTagInformationServerTransport(srv *ServiceTagInformationServer) *ServiceTagInformationServerTransport {
	return &ServiceTagInformationServerTransport{srv: srv}
}

// ServiceTagInformationServerTransport connects instances of armnetwork.ServiceTagInformationClient to instances of ServiceTagInformationServer.
// Don't use this type directly, use NewServiceTagInformationServerTransport instead.
type ServiceTagInformationServerTransport struct {
	srv          *ServiceTagInformationServer
	newListPager *azfake.PagerResponder[armnetwork.ServiceTagInformationClientListResponse]
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
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if s.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/locations/(?P<location>[a-zA-Z0-9-_]+)/serviceTagDetails"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		noAddressPrefixesParam, err := parseOptional(qp.Get("noAddressPrefixes"), strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		tagNameParam := getOptional(qp.Get("tagName"))
		var options *armnetwork.ServiceTagInformationClientListOptions
		if noAddressPrefixesParam != nil || tagNameParam != nil {
			options = &armnetwork.ServiceTagInformationClientListOptions{
				NoAddressPrefixes: noAddressPrefixesParam,
				TagName:           tagNameParam,
			}
		}
		resp := s.srv.NewListPager(matches[regex.SubexpIndex("location")], options)
		s.newListPager = &resp
		server.PagerResponderInjectNextLinks(s.newListPager, req, func(page *armnetwork.ServiceTagInformationClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newListPager) {
		s.newListPager = nil
	}
	return resp, nil
}