//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azacr"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ContainerRegistryBlobServer is a fake server for instances of the azacr.ContainerRegistryBlobClient type.
type ContainerRegistryBlobServer struct {
	// CancelUpload is the fake for method ContainerRegistryBlobClient.CancelUpload
	// HTTP status codes to indicate success: http.StatusNoContent
	CancelUpload func(ctx context.Context, location string, options *azacr.ContainerRegistryBlobClientCancelUploadOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientCancelUploadResponse], errResp azfake.ErrorResponder)

	// CheckBlobExists is the fake for method ContainerRegistryBlobClient.CheckBlobExists
	// HTTP status codes to indicate success: http.StatusOK
	CheckBlobExists func(ctx context.Context, name string, digest string, options *azacr.ContainerRegistryBlobClientCheckBlobExistsOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientCheckBlobExistsResponse], errResp azfake.ErrorResponder)

	// CheckChunkExists is the fake for method ContainerRegistryBlobClient.CheckChunkExists
	// HTTP status codes to indicate success: http.StatusOK
	CheckChunkExists func(ctx context.Context, name string, digest string, rangeParam string, options *azacr.ContainerRegistryBlobClientCheckChunkExistsOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientCheckChunkExistsResponse], errResp azfake.ErrorResponder)

	// CompleteUpload is the fake for method ContainerRegistryBlobClient.CompleteUpload
	// HTTP status codes to indicate success: http.StatusCreated
	CompleteUpload func(ctx context.Context, digest string, location string, value io.ReadSeekCloser, options *azacr.ContainerRegistryBlobClientCompleteUploadOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientCompleteUploadResponse], errResp azfake.ErrorResponder)

	// DeleteBlob is the fake for method ContainerRegistryBlobClient.DeleteBlob
	// HTTP status codes to indicate success: http.StatusAccepted
	DeleteBlob func(ctx context.Context, name string, digest string, options *azacr.ContainerRegistryBlobClientDeleteBlobOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientDeleteBlobResponse], errResp azfake.ErrorResponder)

	// GetBlob is the fake for method ContainerRegistryBlobClient.GetBlob
	// HTTP status codes to indicate success: http.StatusOK
	GetBlob func(ctx context.Context, name string, digest string, options *azacr.ContainerRegistryBlobClientGetBlobOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientGetBlobResponse], errResp azfake.ErrorResponder)

	// GetChunk is the fake for method ContainerRegistryBlobClient.GetChunk
	// HTTP status codes to indicate success: http.StatusPartialContent
	GetChunk func(ctx context.Context, name string, digest string, rangeParam string, options *azacr.ContainerRegistryBlobClientGetChunkOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientGetChunkResponse], errResp azfake.ErrorResponder)

	// GetUploadStatus is the fake for method ContainerRegistryBlobClient.GetUploadStatus
	// HTTP status codes to indicate success: http.StatusNoContent
	GetUploadStatus func(ctx context.Context, location string, options *azacr.ContainerRegistryBlobClientGetUploadStatusOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientGetUploadStatusResponse], errResp azfake.ErrorResponder)

	// MountBlob is the fake for method ContainerRegistryBlobClient.MountBlob
	// HTTP status codes to indicate success: http.StatusCreated
	MountBlob func(ctx context.Context, name string, from string, mount string, options *azacr.ContainerRegistryBlobClientMountBlobOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientMountBlobResponse], errResp azfake.ErrorResponder)

	// StartUpload is the fake for method ContainerRegistryBlobClient.StartUpload
	// HTTP status codes to indicate success: http.StatusAccepted
	StartUpload func(ctx context.Context, name string, options *azacr.ContainerRegistryBlobClientStartUploadOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientStartUploadResponse], errResp azfake.ErrorResponder)

	// UploadChunk is the fake for method ContainerRegistryBlobClient.UploadChunk
	// HTTP status codes to indicate success: http.StatusAccepted
	UploadChunk func(ctx context.Context, location string, value io.ReadSeekCloser, options *azacr.ContainerRegistryBlobClientUploadChunkOptions) (resp azfake.Responder[azacr.ContainerRegistryBlobClientUploadChunkResponse], errResp azfake.ErrorResponder)
}

// NewContainerRegistryBlobServerTransport creates a new instance of ContainerRegistryBlobServerTransport with the provided implementation.
// The returned ContainerRegistryBlobServerTransport instance is connected to an instance of azacr.ContainerRegistryBlobClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerRegistryBlobServerTransport(srv *ContainerRegistryBlobServer) *ContainerRegistryBlobServerTransport {
	return &ContainerRegistryBlobServerTransport{srv: srv}
}

// ContainerRegistryBlobServerTransport connects instances of azacr.ContainerRegistryBlobClient to instances of ContainerRegistryBlobServer.
// Don't use this type directly, use NewContainerRegistryBlobServerTransport instead.
type ContainerRegistryBlobServerTransport struct {
	srv *ContainerRegistryBlobServer
}

// Do implements the policy.Transporter interface for ContainerRegistryBlobServerTransport.
func (c *ContainerRegistryBlobServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerRegistryBlobClient.CancelUpload":
		resp, err = c.dispatchCancelUpload(req)
	case "ContainerRegistryBlobClient.CheckBlobExists":
		resp, err = c.dispatchCheckBlobExists(req)
	case "ContainerRegistryBlobClient.CheckChunkExists":
		resp, err = c.dispatchCheckChunkExists(req)
	case "ContainerRegistryBlobClient.CompleteUpload":
		resp, err = c.dispatchCompleteUpload(req)
	case "ContainerRegistryBlobClient.DeleteBlob":
		resp, err = c.dispatchDeleteBlob(req)
	case "ContainerRegistryBlobClient.GetBlob":
		resp, err = c.dispatchGetBlob(req)
	case "ContainerRegistryBlobClient.GetChunk":
		resp, err = c.dispatchGetChunk(req)
	case "ContainerRegistryBlobClient.GetUploadStatus":
		resp, err = c.dispatchGetUploadStatus(req)
	case "ContainerRegistryBlobClient.MountBlob":
		resp, err = c.dispatchMountBlob(req)
	case "ContainerRegistryBlobClient.StartUpload":
		resp, err = c.dispatchStartUpload(req)
	case "ContainerRegistryBlobClient.UploadChunk":
		resp, err = c.dispatchUploadChunk(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchCancelUpload(req *http.Request) (*http.Response, error) {
	if c.srv.CancelUpload == nil {
		return nil, &nonRetriableError{errors.New("fake for method CancelUpload not implemented")}
	}
	const regexStr = `/(?P<nextBlobUuidLink>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("nextBlobUuidLink")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CancelUpload(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchCheckBlobExists(req *http.Request) (*http.Response, error) {
	if c.srv.CheckBlobExists == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckBlobExists not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobs/(?P<digest>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	digestParam, err := url.PathUnescape(matches[regex.SubexpIndex("digest")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CheckBlobExists(req.Context(), nameParam, digestParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentLength; val != nil {
		resp.Header.Set("Content-Length", strconv.FormatInt(*val, 10))
	}
	if val := server.GetResponse(respr).DockerContentDigest; val != nil {
		resp.Header.Set("Docker-Content-Digest", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchCheckChunkExists(req *http.Request) (*http.Response, error) {
	if c.srv.CheckChunkExists == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckChunkExists not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobs/(?P<digest>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	digestParam, err := url.PathUnescape(matches[regex.SubexpIndex("digest")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CheckChunkExists(req.Context(), nameParam, digestParam, getHeaderValue(req.Header, "Range"), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentLength; val != nil {
		resp.Header.Set("Content-Length", strconv.FormatInt(*val, 10))
	}
	if val := server.GetResponse(respr).ContentRange; val != nil {
		resp.Header.Set("Content-Range", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchCompleteUpload(req *http.Request) (*http.Response, error) {
	if c.srv.CompleteUpload == nil {
		return nil, &nonRetriableError{errors.New("fake for method CompleteUpload not implemented")}
	}
	const regexStr = `/(?P<nextBlobUuidLink>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	digestParam, err := url.QueryUnescape(qp.Get("digest"))
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("nextBlobUuidLink")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CompleteUpload(req.Context(), digestParam, locationParam, req.Body.(io.ReadSeekCloser), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).DockerContentDigest; val != nil {
		resp.Header.Set("Docker-Content-Digest", *val)
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).Range; val != nil {
		resp.Header.Set("Range", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchDeleteBlob(req *http.Request) (*http.Response, error) {
	if c.srv.DeleteBlob == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteBlob not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobs/(?P<digest>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	digestParam, err := url.PathUnescape(matches[regex.SubexpIndex("digest")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.DeleteBlob(req.Context(), nameParam, digestParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).DockerContentDigest; val != nil {
		resp.Header.Set("Docker-Content-Digest", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchGetBlob(req *http.Request) (*http.Response, error) {
	if c.srv.GetBlob == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBlob not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobs/(?P<digest>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	digestParam, err := url.PathUnescape(matches[regex.SubexpIndex("digest")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetBlob(req.Context(), nameParam, digestParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentLength; val != nil {
		resp.Header.Set("Content-Length", strconv.FormatInt(*val, 10))
	}
	if val := server.GetResponse(respr).DockerContentDigest; val != nil {
		resp.Header.Set("Docker-Content-Digest", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchGetChunk(req *http.Request) (*http.Response, error) {
	if c.srv.GetChunk == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetChunk not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobs/(?P<digest>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	digestParam, err := url.PathUnescape(matches[regex.SubexpIndex("digest")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetChunk(req.Context(), nameParam, digestParam, getHeaderValue(req.Header, "Range"), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusPartialContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusPartialContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ContentLength; val != nil {
		resp.Header.Set("Content-Length", strconv.FormatInt(*val, 10))
	}
	if val := server.GetResponse(respr).ContentRange; val != nil {
		resp.Header.Set("Content-Range", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchGetUploadStatus(req *http.Request) (*http.Response, error) {
	if c.srv.GetUploadStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUploadStatus not implemented")}
	}
	const regexStr = `/(?P<nextBlobUuidLink>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("nextBlobUuidLink")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetUploadStatus(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).DockerUploadUUID; val != nil {
		resp.Header.Set("Docker-Upload-UUID", *val)
	}
	if val := server.GetResponse(respr).Range; val != nil {
		resp.Header.Set("Range", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchMountBlob(req *http.Request) (*http.Response, error) {
	if c.srv.MountBlob == nil {
		return nil, &nonRetriableError{errors.New("fake for method MountBlob not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobs/uploads/`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	fromParam, err := url.QueryUnescape(qp.Get("from"))
	if err != nil {
		return nil, err
	}
	mountParam, err := url.QueryUnescape(qp.Get("mount"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.MountBlob(req.Context(), nameParam, fromParam, mountParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).DockerContentDigest; val != nil {
		resp.Header.Set("Docker-Content-Digest", *val)
	}
	if val := server.GetResponse(respr).DockerUploadUUID; val != nil {
		resp.Header.Set("Docker-Upload-UUID", *val)
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchStartUpload(req *http.Request) (*http.Response, error) {
	if c.srv.StartUpload == nil {
		return nil, &nonRetriableError{errors.New("fake for method StartUpload not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/blobs/uploads/`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.StartUpload(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).DockerUploadUUID; val != nil {
		resp.Header.Set("Docker-Upload-UUID", *val)
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).Range; val != nil {
		resp.Header.Set("Range", *val)
	}
	return resp, nil
}

func (c *ContainerRegistryBlobServerTransport) dispatchUploadChunk(req *http.Request) (*http.Response, error) {
	if c.srv.UploadChunk == nil {
		return nil, &nonRetriableError{errors.New("fake for method UploadChunk not implemented")}
	}
	const regexStr = `/(?P<nextBlobUuidLink>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("nextBlobUuidLink")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.UploadChunk(req.Context(), locationParam, req.Body.(io.ReadSeekCloser), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).DockerUploadUUID; val != nil {
		resp.Header.Set("Docker-Upload-UUID", *val)
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).Range; val != nil {
		resp.Header.Set("Range", *val)
	}
	return resp, nil
}
