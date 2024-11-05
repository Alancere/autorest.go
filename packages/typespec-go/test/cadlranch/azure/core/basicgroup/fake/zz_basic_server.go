// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"basicgroup"
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
	"strconv"
)

// BasicServer is a fake server for instances of the basicgroup.BasicClient type.
type BasicServer struct {
	// CreateOrReplace is the fake for method BasicClient.CreateOrReplace
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrReplace func(ctx context.Context, id int32, resource basicgroup.User, options *basicgroup.BasicClientCreateOrReplaceOptions) (resp azfake.Responder[basicgroup.BasicClientCreateOrReplaceResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method BasicClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, id int32, resource basicgroup.User, options *basicgroup.BasicClientCreateOrUpdateOptions) (resp azfake.Responder[basicgroup.BasicClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method BasicClient.Delete
	// HTTP status codes to indicate success: http.StatusNoContent
	Delete func(ctx context.Context, id int32, options *basicgroup.BasicClientDeleteOptions) (resp azfake.Responder[basicgroup.BasicClientDeleteResponse], errResp azfake.ErrorResponder)

	// Export is the fake for method BasicClient.Export
	// HTTP status codes to indicate success: http.StatusOK
	Export func(ctx context.Context, id int32, formatParam string, options *basicgroup.BasicClientExportOptions) (resp azfake.Responder[basicgroup.BasicClientExportResponse], errResp azfake.ErrorResponder)

	// ExportAllUsers is the fake for method BasicClient.ExportAllUsers
	// HTTP status codes to indicate success: http.StatusOK
	ExportAllUsers func(ctx context.Context, formatParam string, options *basicgroup.BasicClientExportAllUsersOptions) (resp azfake.Responder[basicgroup.BasicClientExportAllUsersResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BasicClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, id int32, options *basicgroup.BasicClientGetOptions) (resp azfake.Responder[basicgroup.BasicClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BasicClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *basicgroup.BasicClientListOptions) (resp azfake.PagerResponder[basicgroup.BasicClientListResponse])
}

// NewBasicServerTransport creates a new instance of BasicServerTransport with the provided implementation.
// The returned BasicServerTransport instance is connected to an instance of basicgroup.BasicClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBasicServerTransport(srv *BasicServer) *BasicServerTransport {
	return &BasicServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[basicgroup.BasicClientListResponse]](),
	}
}

// BasicServerTransport connects instances of basicgroup.BasicClient to instances of BasicServer.
// Don't use this type directly, use NewBasicServerTransport instead.
type BasicServerTransport struct {
	srv          *BasicServer
	newListPager *tracker[azfake.PagerResponder[basicgroup.BasicClientListResponse]]
}

// Do implements the policy.Transporter interface for BasicServerTransport.
func (b *BasicServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BasicServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if basicServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = basicServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BasicClient.CreateOrReplace":
				res.resp, res.err = b.dispatchCreateOrReplace(req)
			case "BasicClient.CreateOrUpdate":
				res.resp, res.err = b.dispatchCreateOrUpdate(req)
			case "BasicClient.Delete":
				res.resp, res.err = b.dispatchDelete(req)
			case "BasicClient.Export":
				res.resp, res.err = b.dispatchExport(req)
			case "BasicClient.ExportAllUsers":
				res.resp, res.err = b.dispatchExportAllUsers(req)
			case "BasicClient.Get":
				res.resp, res.err = b.dispatchGet(req)
			case "BasicClient.NewListPager":
				res.resp, res.err = b.dispatchNewListPager(req)
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

func (b *BasicServerTransport) dispatchCreateOrReplace(req *http.Request) (*http.Response, error) {
	if b.srv.CreateOrReplace == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrReplace not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[basicgroup.User](req)
	if err != nil {
		return nil, err
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.CreateOrReplace(req.Context(), idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[basicgroup.User](req)
	if err != nil {
		return nil, err
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.CreateOrUpdate(req.Context(), idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if b.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Delete(req.Context(), idParam, nil)
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

func (b *BasicServerTransport) dispatchExport(req *http.Request) (*http.Response, error) {
	if b.srv.Export == nil {
		return nil, &nonRetriableError{errors.New("fake for method Export not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+):export`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	formatParamParam, err := url.QueryUnescape(qp.Get("format"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Export(req.Context(), idParam, formatParamParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchExportAllUsers(req *http.Request) (*http.Response, error) {
	if b.srv.ExportAllUsers == nil {
		return nil, &nonRetriableError{errors.New("fake for method ExportAllUsers not implemented")}
	}
	qp := req.URL.Query()
	formatParamParam, err := url.QueryUnescape(qp.Get("format"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.ExportAllUsers(req.Context(), formatParamParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UserList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		qp := req.URL.Query()
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		orderbyEscaped := qp["orderby"]
		orderbyParam := make([]string, len(orderbyEscaped))
		for i, v := range orderbyEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			orderbyParam[i] = u
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		selectParamEscaped := qp["select"]
		selectParamParam := make([]string, len(selectParamEscaped))
		for i, v := range selectParamEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			selectParamParam[i] = u
		}
		expandEscaped := qp["expand"]
		expandParam := make([]string, len(expandEscaped))
		for i, v := range expandEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			expandParam[i] = u
		}
		var options *basicgroup.BasicClientListOptions
		if topParam != nil || skipParam != nil || maxpagesizeParam != nil || len(orderbyParam) > 0 || filterParam != nil || len(selectParamParam) > 0 || len(expandParam) > 0 {
			options = &basicgroup.BasicClientListOptions{
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
				Orderby:     orderbyParam,
				Filter:      filterParam,
				SelectParam: selectParamParam,
				Expand:      expandParam,
			}
		}
		resp := b.srv.NewListPager(options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *basicgroup.BasicClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to BasicServerTransport
var basicServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
