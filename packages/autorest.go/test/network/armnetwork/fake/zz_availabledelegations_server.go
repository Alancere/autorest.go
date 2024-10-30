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

// AvailableDelegationsServer is a fake server for instances of the armnetwork.AvailableDelegationsClient type.
type AvailableDelegationsServer struct {
	// NewListPager is the fake for method AvailableDelegationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armnetwork.AvailableDelegationsClientListOptions) (resp azfake.PagerResponder[armnetwork.AvailableDelegationsClientListResponse])
}

// NewAvailableDelegationsServerTransport creates a new instance of AvailableDelegationsServerTransport with the provided implementation.
// The returned AvailableDelegationsServerTransport instance is connected to an instance of armnetwork.AvailableDelegationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAvailableDelegationsServerTransport(srv *AvailableDelegationsServer) *AvailableDelegationsServerTransport {
	return &AvailableDelegationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.AvailableDelegationsClientListResponse]](),
	}
}

// AvailableDelegationsServerTransport connects instances of armnetwork.AvailableDelegationsClient to instances of AvailableDelegationsServer.
// Don't use this type directly, use NewAvailableDelegationsServerTransport instead.
type AvailableDelegationsServerTransport struct {
	srv          *AvailableDelegationsServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.AvailableDelegationsClientListResponse]]
}

// Do implements the policy.Transporter interface for AvailableDelegationsServerTransport.
func (a *AvailableDelegationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AvailableDelegationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if availableDelegationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = availableDelegationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AvailableDelegationsClient.NewListPager":
				res.resp, res.err = a.dispatchNewListPager(req)
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

func (a *AvailableDelegationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/availableDelegations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(locationParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.AvailableDelegationsClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to AvailableDelegationsServerTransport
var availableDelegationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
