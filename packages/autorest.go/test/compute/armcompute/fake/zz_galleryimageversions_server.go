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
	"regexp"
)

// GalleryImageVersionsServer is a fake server for instances of the armcompute.GalleryImageVersionsClient type.
type GalleryImageVersionsServer struct {
	// BeginCreateOrUpdate is the fake for method GalleryImageVersionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion armcompute.GalleryImageVersion, options *armcompute.GalleryImageVersionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleryImageVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GalleryImageVersionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *armcompute.GalleryImageVersionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.GalleryImageVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GalleryImageVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *armcompute.GalleryImageVersionsClientGetOptions) (resp azfake.Responder[armcompute.GalleryImageVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByGalleryImagePager is the fake for method GalleryImageVersionsClient.NewListByGalleryImagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByGalleryImagePager func(resourceGroupName string, galleryName string, galleryImageName string, options *armcompute.GalleryImageVersionsClientListByGalleryImageOptions) (resp azfake.PagerResponder[armcompute.GalleryImageVersionsClientListByGalleryImageResponse])

	// BeginUpdate is the fake for method GalleryImageVersionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion armcompute.GalleryImageVersionUpdate, options *armcompute.GalleryImageVersionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleryImageVersionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGalleryImageVersionsServerTransport creates a new instance of GalleryImageVersionsServerTransport with the provided implementation.
// The returned GalleryImageVersionsServerTransport instance is connected to an instance of armcompute.GalleryImageVersionsClient by way of the
// undefined.Transporter field.
func NewGalleryImageVersionsServerTransport(srv *GalleryImageVersionsServer) *GalleryImageVersionsServerTransport {
	return &GalleryImageVersionsServerTransport{srv: srv}
}

// GalleryImageVersionsServerTransport connects instances of armcompute.GalleryImageVersionsClient to instances of GalleryImageVersionsServer.
// Don't use this type directly, use NewGalleryImageVersionsServerTransport instead.
type GalleryImageVersionsServerTransport struct {
	srv                        *GalleryImageVersionsServer
	beginCreateOrUpdate        *azfake.PollerResponder[armcompute.GalleryImageVersionsClientCreateOrUpdateResponse]
	beginDelete                *azfake.PollerResponder[armcompute.GalleryImageVersionsClientDeleteResponse]
	newListByGalleryImagePager *azfake.PagerResponder[armcompute.GalleryImageVersionsClientListByGalleryImageResponse]
	beginUpdate                *azfake.PollerResponder[armcompute.GalleryImageVersionsClientUpdateResponse]
}

// Do implements the policy.Transporter interface for GalleryImageVersionsServerTransport.
func (g *GalleryImageVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GalleryImageVersionsClient.BeginCreateOrUpdate":
		resp, err = g.dispatchBeginCreateOrUpdate(req)
	case "GalleryImageVersionsClient.BeginDelete":
		resp, err = g.dispatchBeginDelete(req)
	case "GalleryImageVersionsClient.Get":
		resp, err = g.dispatchGet(req)
	case "GalleryImageVersionsClient.NewListByGalleryImagePager":
		resp, err = g.dispatchNewListByGalleryImagePager(req)
	case "GalleryImageVersionsClient.BeginUpdate":
		resp, err = g.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if g.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)/images/(?P<galleryImageName>[a-zA-Z0-9-_]+)/versions/(?P<galleryImageVersionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryImageVersion](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], matches[regex.SubexpIndex("galleryImageName")], matches[regex.SubexpIndex("galleryImageVersionName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		g.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(g.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(g.beginCreateOrUpdate) {
		g.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if g.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)/images/(?P<galleryImageName>[a-zA-Z0-9-_]+)/versions/(?P<galleryImageVersionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], matches[regex.SubexpIndex("galleryImageName")], matches[regex.SubexpIndex("galleryImageVersionName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		g.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(g.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(g.beginDelete) {
		g.beginDelete = nil
	}

	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)/images/(?P<galleryImageName>[a-zA-Z0-9-_]+)/versions/(?P<galleryImageVersionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(armcompute.ReplicationStatusTypes(qp.Get("$expand")))
	var options *armcompute.GalleryImageVersionsClientGetOptions
	if expandParam != nil {
		options = &armcompute.GalleryImageVersionsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := g.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], matches[regex.SubexpIndex("galleryImageName")], matches[regex.SubexpIndex("galleryImageVersionName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GalleryImageVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchNewListByGalleryImagePager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByGalleryImagePager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByGalleryImagePager not implemented")}
	}
	if g.newListByGalleryImagePager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)/images/(?P<galleryImageName>[a-zA-Z0-9-_]+)/versions"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := g.srv.NewListByGalleryImagePager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], matches[regex.SubexpIndex("galleryImageName")], nil)
		g.newListByGalleryImagePager = &resp
		server.PagerResponderInjectNextLinks(g.newListByGalleryImagePager, req, func(page *armcompute.GalleryImageVersionsClientListByGalleryImageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(g.newListByGalleryImagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(g.newListByGalleryImagePager) {
		g.newListByGalleryImagePager = nil
	}
	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if g.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/galleries/(?P<galleryName>[a-zA-Z0-9-_]+)/images/(?P<galleryImageName>[a-zA-Z0-9-_]+)/versions/(?P<galleryImageVersionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryImageVersionUpdate](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("galleryName")], matches[regex.SubexpIndex("galleryImageName")], matches[regex.SubexpIndex("galleryImageVersionName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		g.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(g.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(g.beginUpdate) {
		g.beginUpdate = nil
	}

	return resp, nil
}