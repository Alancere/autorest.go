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
)

// BgpServiceCommunitiesServer is a fake server for instances of the armnetwork.BgpServiceCommunitiesClient type.
type BgpServiceCommunitiesServer struct {
	// NewListPager is the fake for method BgpServiceCommunitiesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.BgpServiceCommunitiesClientListOptions) (resp azfake.PagerResponder[armnetwork.BgpServiceCommunitiesClientListResponse])
}

// NewBgpServiceCommunitiesServerTransport creates a new instance of BgpServiceCommunitiesServerTransport with the provided implementation.
// The returned BgpServiceCommunitiesServerTransport instance is connected to an instance of armnetwork.BgpServiceCommunitiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBgpServiceCommunitiesServerTransport(srv *BgpServiceCommunitiesServer) *BgpServiceCommunitiesServerTransport {
	return &BgpServiceCommunitiesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.BgpServiceCommunitiesClientListResponse]](),
	}
}

// BgpServiceCommunitiesServerTransport connects instances of armnetwork.BgpServiceCommunitiesClient to instances of BgpServiceCommunitiesServer.
// Don't use this type directly, use NewBgpServiceCommunitiesServerTransport instead.
type BgpServiceCommunitiesServerTransport struct {
	srv          *BgpServiceCommunitiesServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.BgpServiceCommunitiesClientListResponse]]
}

// Do implements the policy.Transporter interface for BgpServiceCommunitiesServerTransport.
func (b *BgpServiceCommunitiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BgpServiceCommunitiesClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BgpServiceCommunitiesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/bgpServiceCommunities`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := b.srv.NewListPager(nil)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.BgpServiceCommunitiesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}
