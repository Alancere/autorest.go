// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"nullablegroup"
)

// NullableBytesServer is a fake server for instances of the nullablegroup.NullableBytesClient type.
type NullableBytesServer struct {
	// GetNonNull is the fake for method NullableBytesClient.GetNonNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNonNull func(ctx context.Context, options *nullablegroup.NullableBytesClientGetNonNullOptions) (resp azfake.Responder[nullablegroup.NullableBytesClientGetNonNullResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method NullableBytesClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *nullablegroup.NullableBytesClientGetNullOptions) (resp azfake.Responder[nullablegroup.NullableBytesClientGetNullResponse], errResp azfake.ErrorResponder)

	// PatchNonNull is the fake for method NullableBytesClient.PatchNonNull
	// HTTP status codes to indicate success: http.StatusNoContent
	PatchNonNull func(ctx context.Context, body nullablegroup.BytesProperty, options *nullablegroup.NullableBytesClientPatchNonNullOptions) (resp azfake.Responder[nullablegroup.NullableBytesClientPatchNonNullResponse], errResp azfake.ErrorResponder)

	// PatchNull is the fake for method NullableBytesClient.PatchNull
	// HTTP status codes to indicate success: http.StatusNoContent
	PatchNull func(ctx context.Context, body nullablegroup.BytesProperty, options *nullablegroup.NullableBytesClientPatchNullOptions) (resp azfake.Responder[nullablegroup.NullableBytesClientPatchNullResponse], errResp azfake.ErrorResponder)
}

// NewNullableBytesServerTransport creates a new instance of NullableBytesServerTransport with the provided implementation.
// The returned NullableBytesServerTransport instance is connected to an instance of nullablegroup.NullableBytesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNullableBytesServerTransport(srv *NullableBytesServer) *NullableBytesServerTransport {
	return &NullableBytesServerTransport{srv: srv}
}

// NullableBytesServerTransport connects instances of nullablegroup.NullableBytesClient to instances of NullableBytesServer.
// Don't use this type directly, use NewNullableBytesServerTransport instead.
type NullableBytesServerTransport struct {
	srv *NullableBytesServer
}

// Do implements the policy.Transporter interface for NullableBytesServerTransport.
func (n *NullableBytesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NullableBytesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if nullableBytesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = nullableBytesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "NullableBytesClient.GetNonNull":
				res.resp, res.err = n.dispatchGetNonNull(req)
			case "NullableBytesClient.GetNull":
				res.resp, res.err = n.dispatchGetNull(req)
			case "NullableBytesClient.PatchNonNull":
				res.resp, res.err = n.dispatchPatchNonNull(req)
			case "NullableBytesClient.PatchNull":
				res.resp, res.err = n.dispatchPatchNull(req)
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

func (n *NullableBytesServerTransport) dispatchGetNonNull(req *http.Request) (*http.Response, error) {
	if n.srv.GetNonNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNonNull not implemented")}
	}
	respr, errRespr := n.srv.GetNonNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BytesProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NullableBytesServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if n.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := n.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BytesProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NullableBytesServerTransport) dispatchPatchNonNull(req *http.Request) (*http.Response, error) {
	if n.srv.PatchNonNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PatchNonNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nullablegroup.BytesProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PatchNonNull(req.Context(), body, nil)
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

func (n *NullableBytesServerTransport) dispatchPatchNull(req *http.Request) (*http.Response, error) {
	if n.srv.PatchNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PatchNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nullablegroup.BytesProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.PatchNull(req.Context(), body, nil)
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

// set this to conditionally intercept incoming requests to NullableBytesServerTransport
var nullableBytesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
