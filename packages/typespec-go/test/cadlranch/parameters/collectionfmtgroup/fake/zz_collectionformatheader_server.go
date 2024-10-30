// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"collectionfmtgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// CollectionFormatHeaderServer is a fake server for instances of the collectionfmtgroup.CollectionFormatHeaderClient type.
type CollectionFormatHeaderServer struct {
	// CSV is the fake for method CollectionFormatHeaderClient.CSV
	// HTTP status codes to indicate success: http.StatusNoContent
	CSV func(ctx context.Context, colors []string, options *collectionfmtgroup.CollectionFormatHeaderClientCSVOptions) (resp azfake.Responder[collectionfmtgroup.CollectionFormatHeaderClientCSVResponse], errResp azfake.ErrorResponder)
}

// NewCollectionFormatHeaderServerTransport creates a new instance of CollectionFormatHeaderServerTransport with the provided implementation.
// The returned CollectionFormatHeaderServerTransport instance is connected to an instance of collectionfmtgroup.CollectionFormatHeaderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCollectionFormatHeaderServerTransport(srv *CollectionFormatHeaderServer) *CollectionFormatHeaderServerTransport {
	return &CollectionFormatHeaderServerTransport{srv: srv}
}

// CollectionFormatHeaderServerTransport connects instances of collectionfmtgroup.CollectionFormatHeaderClient to instances of CollectionFormatHeaderServer.
// Don't use this type directly, use NewCollectionFormatHeaderServerTransport instead.
type CollectionFormatHeaderServerTransport struct {
	srv *CollectionFormatHeaderServer
}

// Do implements the policy.Transporter interface for CollectionFormatHeaderServerTransport.
func (c *CollectionFormatHeaderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CollectionFormatHeaderServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if collectionFormatHeaderServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = collectionFormatHeaderServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CollectionFormatHeaderClient.CSV":
				res.resp, res.err = c.dispatchCSV(req)
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

func (c *CollectionFormatHeaderServerTransport) dispatchCSV(req *http.Request) (*http.Response, error) {
	if c.srv.CSV == nil {
		return nil, &nonRetriableError{errors.New("fake for method CSV not implemented")}
	}
	respr, errRespr := c.srv.CSV(req.Context(), splitHelper(getHeaderValue(req.Header, "colors"), ","), nil)
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

// set this to conditionally intercept incoming requests to CollectionFormatHeaderServerTransport
var collectionFormatHeaderServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
