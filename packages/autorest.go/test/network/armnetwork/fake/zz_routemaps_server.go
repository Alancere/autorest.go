//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"context"
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

// RouteMapsServer is a fake server for instances of the armnetwork.RouteMapsClient type.
type RouteMapsServer struct {
	// BeginCreateOrUpdate is the fake for method RouteMapsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, routeMapParameters armnetwork.RouteMap, options *armnetwork.RouteMapsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.RouteMapsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RouteMapsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *armnetwork.RouteMapsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.RouteMapsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RouteMapsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *armnetwork.RouteMapsClientGetOptions) (resp azfake.Responder[armnetwork.RouteMapsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RouteMapsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualHubName string, options *armnetwork.RouteMapsClientListOptions) (resp azfake.PagerResponder[armnetwork.RouteMapsClientListResponse])
}

// NewRouteMapsServerTransport creates a new instance of RouteMapsServerTransport with the provided implementation.
// The returned RouteMapsServerTransport instance is connected to an instance of armnetwork.RouteMapsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRouteMapsServerTransport(srv *RouteMapsServer) *RouteMapsServerTransport {
	return &RouteMapsServerTransport{srv: srv}
}

// RouteMapsServerTransport connects instances of armnetwork.RouteMapsClient to instances of RouteMapsServer.
// Don't use this type directly, use NewRouteMapsServerTransport instead.
type RouteMapsServerTransport struct {
	srv                 *RouteMapsServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.RouteMapsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.RouteMapsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.RouteMapsClientListResponse]
}

// Do implements the policy.Transporter interface for RouteMapsServerTransport.
func (r *RouteMapsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RouteMapsClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RouteMapsClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RouteMapsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RouteMapsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RouteMapsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if r.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routeMaps/(?P<routeMapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.RouteMap](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		routeMapNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeMapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, virtualHubNameUnescaped, routeMapNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		r.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(r.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(r.beginCreateOrUpdate) {
		r.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (r *RouteMapsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if r.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routeMaps/(?P<routeMapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		routeMapNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeMapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, virtualHubNameUnescaped, routeMapNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		r.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(r.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(r.beginDelete) {
		r.beginDelete = nil
	}

	return resp, nil
}

func (r *RouteMapsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routeMaps/(?P<routeMapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualHubNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
	if err != nil {
		return nil, err
	}
	routeMapNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeMapName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameUnescaped, virtualHubNameUnescaped, routeMapNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RouteMap, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RouteMapsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if r.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routeMaps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(resourceGroupNameUnescaped, virtualHubNameUnescaped, nil)
		r.newListPager = &resp
		server.PagerResponderInjectNextLinks(r.newListPager, req, func(page *armnetwork.RouteMapsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(r.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(r.newListPager) {
		r.newListPager = nil
	}
	return resp, nil
}
