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
)

// AvailablePrivateEndpointTypesServer is a fake server for instances of the armnetwork.AvailablePrivateEndpointTypesClient type.
type AvailablePrivateEndpointTypesServer struct {
	// NewListPager is the fake for method AvailablePrivateEndpointTypesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armnetwork.AvailablePrivateEndpointTypesClientListOptions) (resp azfake.PagerResponder[armnetwork.AvailablePrivateEndpointTypesClientListResponse])

	// NewListByResourceGroupPager is the fake for method AvailablePrivateEndpointTypesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(location string, resourceGroupName string, options *armnetwork.AvailablePrivateEndpointTypesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.AvailablePrivateEndpointTypesClientListByResourceGroupResponse])
}

// NewAvailablePrivateEndpointTypesServerTransport creates a new instance of AvailablePrivateEndpointTypesServerTransport with the provided implementation.
// The returned AvailablePrivateEndpointTypesServerTransport instance is connected to an instance of armnetwork.AvailablePrivateEndpointTypesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAvailablePrivateEndpointTypesServerTransport(srv *AvailablePrivateEndpointTypesServer) *AvailablePrivateEndpointTypesServerTransport {
	return &AvailablePrivateEndpointTypesServerTransport{
		srv:                         srv,
		newListPager:                newTracker[azfake.PagerResponder[armnetwork.AvailablePrivateEndpointTypesClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetwork.AvailablePrivateEndpointTypesClientListByResourceGroupResponse]](),
	}
}

// AvailablePrivateEndpointTypesServerTransport connects instances of armnetwork.AvailablePrivateEndpointTypesClient to instances of AvailablePrivateEndpointTypesServer.
// Don't use this type directly, use NewAvailablePrivateEndpointTypesServerTransport instead.
type AvailablePrivateEndpointTypesServerTransport struct {
	srv                         *AvailablePrivateEndpointTypesServer
	newListPager                *tracker[azfake.PagerResponder[armnetwork.AvailablePrivateEndpointTypesClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetwork.AvailablePrivateEndpointTypesClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for AvailablePrivateEndpointTypesServerTransport.
func (a *AvailablePrivateEndpointTypesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AvailablePrivateEndpointTypesClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AvailablePrivateEndpointTypesClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AvailablePrivateEndpointTypesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/availablePrivateEndpointTypes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(locationUnescaped, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.AvailablePrivateEndpointTypesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *AvailablePrivateEndpointTypesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/availablePrivateEndpointTypes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(locationUnescaped, resourceGroupNameUnescaped, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetwork.AvailablePrivateEndpointTypesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}
