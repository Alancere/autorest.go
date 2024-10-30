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

// SecurityRulesServer is a fake server for instances of the armnetwork.SecurityRulesClient type.
type SecurityRulesServer struct {
	// BeginCreateOrUpdate is the fake for method SecurityRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters armnetwork.SecurityRule, options *armnetwork.SecurityRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.SecurityRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SecurityRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *armnetwork.SecurityRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.SecurityRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SecurityRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *armnetwork.SecurityRulesClientGetOptions) (resp azfake.Responder[armnetwork.SecurityRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SecurityRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkSecurityGroupName string, options *armnetwork.SecurityRulesClientListOptions) (resp azfake.PagerResponder[armnetwork.SecurityRulesClientListResponse])
}

// NewSecurityRulesServerTransport creates a new instance of SecurityRulesServerTransport with the provided implementation.
// The returned SecurityRulesServerTransport instance is connected to an instance of armnetwork.SecurityRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSecurityRulesServerTransport(srv *SecurityRulesServer) *SecurityRulesServerTransport {
	return &SecurityRulesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.SecurityRulesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.SecurityRulesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.SecurityRulesClientListResponse]](),
	}
}

// SecurityRulesServerTransport connects instances of armnetwork.SecurityRulesClient to instances of SecurityRulesServer.
// Don't use this type directly, use NewSecurityRulesServerTransport instead.
type SecurityRulesServerTransport struct {
	srv                 *SecurityRulesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.SecurityRulesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.SecurityRulesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.SecurityRulesClientListResponse]]
}

// Do implements the policy.Transporter interface for SecurityRulesServerTransport.
func (s *SecurityRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SecurityRulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if securityRulesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = securityRulesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SecurityRulesClient.BeginCreateOrUpdate":
				res.resp, res.err = s.dispatchBeginCreateOrUpdate(req)
			case "SecurityRulesClient.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "SecurityRulesClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SecurityRulesClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
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

func (s *SecurityRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityGroups/(?P<networkSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityRules/(?P<securityRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.SecurityRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkSecurityGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityGroupName")])
		if err != nil {
			return nil, err
		}
		securityRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("securityRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkSecurityGroupNameParam, securityRuleNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *SecurityRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityGroups/(?P<networkSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityRules/(?P<securityRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkSecurityGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityGroupName")])
		if err != nil {
			return nil, err
		}
		securityRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("securityRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkSecurityGroupNameParam, securityRuleNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SecurityRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityGroups/(?P<networkSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityRules/(?P<securityRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityGroupName")])
	if err != nil {
		return nil, err
	}
	securityRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("securityRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, networkSecurityGroupNameParam, securityRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecurityRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityGroups/(?P<networkSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkSecurityGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, networkSecurityGroupNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.SecurityRulesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SecurityRulesServerTransport
var securityRulesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
