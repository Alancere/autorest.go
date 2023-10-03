//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armcompute"
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

// RestorePointCollectionsServer is a fake server for instances of the armcompute.RestorePointCollectionsClient type.
type RestorePointCollectionsServer struct {
	// CreateOrUpdate is the fake for method RestorePointCollectionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters armcompute.RestorePointCollection, options *armcompute.RestorePointCollectionsClientCreateOrUpdateOptions) (resp azfake.Responder[armcompute.RestorePointCollectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RestorePointCollectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *armcompute.RestorePointCollectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.RestorePointCollectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RestorePointCollectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *armcompute.RestorePointCollectionsClientGetOptions) (resp azfake.Responder[armcompute.RestorePointCollectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RestorePointCollectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armcompute.RestorePointCollectionsClientListOptions) (resp azfake.PagerResponder[armcompute.RestorePointCollectionsClientListResponse])

	// NewListAllPager is the fake for method RestorePointCollectionsClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armcompute.RestorePointCollectionsClientListAllOptions) (resp azfake.PagerResponder[armcompute.RestorePointCollectionsClientListAllResponse])

	// Update is the fake for method RestorePointCollectionsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters armcompute.RestorePointCollectionUpdate, options *armcompute.RestorePointCollectionsClientUpdateOptions) (resp azfake.Responder[armcompute.RestorePointCollectionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewRestorePointCollectionsServerTransport creates a new instance of RestorePointCollectionsServerTransport with the provided implementation.
// The returned RestorePointCollectionsServerTransport instance is connected to an instance of armcompute.RestorePointCollectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRestorePointCollectionsServerTransport(srv *RestorePointCollectionsServer) *RestorePointCollectionsServerTransport {
	return &RestorePointCollectionsServerTransport{
		srv:             srv,
		beginDelete:     newTracker[azfake.PollerResponder[armcompute.RestorePointCollectionsClientDeleteResponse]](),
		newListPager:    newTracker[azfake.PagerResponder[armcompute.RestorePointCollectionsClientListResponse]](),
		newListAllPager: newTracker[azfake.PagerResponder[armcompute.RestorePointCollectionsClientListAllResponse]](),
	}
}

// RestorePointCollectionsServerTransport connects instances of armcompute.RestorePointCollectionsClient to instances of RestorePointCollectionsServer.
// Don't use this type directly, use NewRestorePointCollectionsServerTransport instead.
type RestorePointCollectionsServerTransport struct {
	srv             *RestorePointCollectionsServer
	beginDelete     *tracker[azfake.PollerResponder[armcompute.RestorePointCollectionsClientDeleteResponse]]
	newListPager    *tracker[azfake.PagerResponder[armcompute.RestorePointCollectionsClientListResponse]]
	newListAllPager *tracker[azfake.PagerResponder[armcompute.RestorePointCollectionsClientListAllResponse]]
}

// Do implements the policy.Transporter interface for RestorePointCollectionsServerTransport.
func (r *RestorePointCollectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RestorePointCollectionsClient.CreateOrUpdate":
		resp, err = r.dispatchCreateOrUpdate(req)
	case "RestorePointCollectionsClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RestorePointCollectionsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RestorePointCollectionsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "RestorePointCollectionsClient.NewListAllPager":
		resp, err = r.dispatchNewListAllPager(req)
	case "RestorePointCollectionsClient.Update":
		resp, err = r.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RestorePointCollectionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.RestorePointCollection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	restorePointCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, restorePointCollectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RestorePointCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RestorePointCollectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		restorePointCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, restorePointCollectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *RestorePointCollectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	restorePointCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(armcompute.RestorePointCollectionExpandOptions(expandUnescaped))
	var options *armcompute.RestorePointCollectionsClientGetOptions
	if expandParam != nil {
		options = &armcompute.RestorePointCollectionsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, restorePointCollectionNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RestorePointCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RestorePointCollectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/restorePointCollections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.RestorePointCollectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

func (r *RestorePointCollectionsServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := r.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/restorePointCollections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := r.srv.NewListAllPager(nil)
		newListAllPager = &resp
		r.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armcompute.RestorePointCollectionsClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		r.newListAllPager.remove(req)
	}
	return resp, nil
}

func (r *RestorePointCollectionsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/restorePointCollections/(?P<restorePointCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.RestorePointCollectionUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	restorePointCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("restorePointCollectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Update(req.Context(), resourceGroupNameParam, restorePointCollectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RestorePointCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
