// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// AccessServer is a fake server for instances of the accessgroup.AccessClient type.
type AccessServer struct {
	// AccessPublicOperationServer contains the fakes for client AccessPublicOperationClient
	AccessPublicOperationServer AccessPublicOperationServer

	// AccessSharedModelInOperationServer contains the fakes for client AccessSharedModelInOperationClient
	AccessSharedModelInOperationServer AccessSharedModelInOperationServer
}

// NewAccessServerTransport creates a new instance of AccessServerTransport with the provided implementation.
// The returned AccessServerTransport instance is connected to an instance of accessgroup.AccessClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessServerTransport(srv *AccessServer) *AccessServerTransport {
	return &AccessServerTransport{srv: srv}
}

// AccessServerTransport connects instances of accessgroup.AccessClient to instances of AccessServer.
// Don't use this type directly, use NewAccessServerTransport instead.
type AccessServerTransport struct {
	srv                                  *AccessServer
	trMu                                 sync.Mutex
	trAccessPublicOperationServer        *AccessPublicOperationServerTransport
	trAccessSharedModelInOperationServer *AccessSharedModelInOperationServerTransport
}

// Do implements the policy.Transporter interface for AccessServerTransport.
func (a *AccessServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (a *AccessServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "AccessPublicOperationClient":
		initServer(&a.trMu, &a.trAccessPublicOperationServer, func() *AccessPublicOperationServerTransport {
			return NewAccessPublicOperationServerTransport(&a.srv.AccessPublicOperationServer)
		})
		resp, err = a.trAccessPublicOperationServer.Do(req)
	case "AccessSharedModelInOperationClient":
		initServer(&a.trMu, &a.trAccessSharedModelInOperationServer, func() *AccessSharedModelInOperationServerTransport {
			return NewAccessSharedModelInOperationServerTransport(&a.srv.AccessSharedModelInOperationServer)
		})
		resp, err = a.trAccessSharedModelInOperationServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

// set this to conditionally intercept incoming requests to AccessServerTransport
var accessServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
