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
	"lrorpcgroup"
	"net/http"
)

// RPCServer is a fake server for instances of the lrorpcgroup.RPCClient type.
type RPCServer struct {
	// BeginLongRunningRPC is the fake for method RPCClient.BeginLongRunningRPC
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginLongRunningRPC func(ctx context.Context, generationOptions lrorpcgroup.GenerationOptions, options *lrorpcgroup.RPCClientBeginLongRunningRPCOptions) (resp azfake.PollerResponder[lrorpcgroup.RPCClientLongRunningRPCResponse], errResp azfake.ErrorResponder)
}

// NewRPCServerTransport creates a new instance of RPCServerTransport with the provided implementation.
// The returned RPCServerTransport instance is connected to an instance of lrorpcgroup.RPCClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRPCServerTransport(srv *RPCServer) *RPCServerTransport {
	return &RPCServerTransport{
		srv:                 srv,
		beginLongRunningRPC: newTracker[azfake.PollerResponder[lrorpcgroup.RPCClientLongRunningRPCResponse]](),
	}
}

// RPCServerTransport connects instances of lrorpcgroup.RPCClient to instances of RPCServer.
// Don't use this type directly, use NewRPCServerTransport instead.
type RPCServerTransport struct {
	srv                 *RPCServer
	beginLongRunningRPC *tracker[azfake.PollerResponder[lrorpcgroup.RPCClientLongRunningRPCResponse]]
}

// Do implements the policy.Transporter interface for RPCServerTransport.
func (r *RPCServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RPCServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "RPCClient.BeginLongRunningRPC":
		resp, err = r.dispatchBeginLongRunningRPC(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (r *RPCServerTransport) dispatchBeginLongRunningRPC(req *http.Request) (*http.Response, error) {
	if r.srv.BeginLongRunningRPC == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginLongRunningRPC not implemented")}
	}
	beginLongRunningRPC := r.beginLongRunningRPC.get(req)
	if beginLongRunningRPC == nil {
		body, err := server.UnmarshalRequestAsJSON[lrorpcgroup.GenerationOptions](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginLongRunningRPC(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginLongRunningRPC = &respr
		r.beginLongRunningRPC.add(req, beginLongRunningRPC)
	}

	resp, err := server.PollerResponderNext(beginLongRunningRPC, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		r.beginLongRunningRPC.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginLongRunningRPC) {
		r.beginLongRunningRPC.remove(req)
	}

	return resp, nil
}
