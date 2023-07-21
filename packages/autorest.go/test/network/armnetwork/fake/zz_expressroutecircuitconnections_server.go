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

// ExpressRouteCircuitConnectionsServer is a fake server for instances of the armnetwork.ExpressRouteCircuitConnectionsClient type.
type ExpressRouteCircuitConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method ExpressRouteCircuitConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, expressRouteCircuitConnectionParameters armnetwork.ExpressRouteCircuitConnection, options *armnetwork.ExpressRouteCircuitConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExpressRouteCircuitConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *armnetwork.ExpressRouteCircuitConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExpressRouteCircuitConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *armnetwork.ExpressRouteCircuitConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRouteCircuitConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExpressRouteCircuitConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, circuitName string, peeringName string, options *armnetwork.ExpressRouteCircuitConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRouteCircuitConnectionsClientListResponse])
}

// NewExpressRouteCircuitConnectionsServerTransport creates a new instance of ExpressRouteCircuitConnectionsServerTransport with the provided implementation.
// The returned ExpressRouteCircuitConnectionsServerTransport instance is connected to an instance of armnetwork.ExpressRouteCircuitConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRouteCircuitConnectionsServerTransport(srv *ExpressRouteCircuitConnectionsServer) *ExpressRouteCircuitConnectionsServerTransport {
	return &ExpressRouteCircuitConnectionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.ExpressRouteCircuitConnectionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.ExpressRouteCircuitConnectionsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.ExpressRouteCircuitConnectionsClientListResponse]](),
	}
}

// ExpressRouteCircuitConnectionsServerTransport connects instances of armnetwork.ExpressRouteCircuitConnectionsClient to instances of ExpressRouteCircuitConnectionsServer.
// Don't use this type directly, use NewExpressRouteCircuitConnectionsServerTransport instead.
type ExpressRouteCircuitConnectionsServerTransport struct {
	srv                 *ExpressRouteCircuitConnectionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.ExpressRouteCircuitConnectionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.ExpressRouteCircuitConnectionsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.ExpressRouteCircuitConnectionsClientListResponse]]
}

// Do implements the policy.Transporter interface for ExpressRouteCircuitConnectionsServerTransport.
func (e *ExpressRouteCircuitConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRouteCircuitConnectionsClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "ExpressRouteCircuitConnectionsClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "ExpressRouteCircuitConnectionsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRouteCircuitConnectionsClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRouteCircuitConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ExpressRouteCircuitConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, peeringNameUnescaped, connectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *ExpressRouteCircuitConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, peeringNameUnescaped, connectionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *ExpressRouteCircuitConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
	if err != nil {
		return nil, err
	}
	peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
	if err != nil {
		return nil, err
	}
	connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, peeringNameUnescaped, connectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteCircuitConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteCircuitConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListPager(resourceGroupNameUnescaped, circuitNameUnescaped, peeringNameUnescaped, nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ExpressRouteCircuitConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}
