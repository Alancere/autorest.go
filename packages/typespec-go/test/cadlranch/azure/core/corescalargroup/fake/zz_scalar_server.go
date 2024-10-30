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

// ScalarServer is a fake server for instances of the corescalargroup.ScalarClient type.
type ScalarServer struct {
	// ScalarAzureLocationScalarServer contains the fakes for client ScalarAzureLocationScalarClient
	ScalarAzureLocationScalarServer ScalarAzureLocationScalarServer
}

// NewScalarServerTransport creates a new instance of ScalarServerTransport with the provided implementation.
// The returned ScalarServerTransport instance is connected to an instance of corescalargroup.ScalarClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScalarServerTransport(srv *ScalarServer) *ScalarServerTransport {
	return &ScalarServerTransport{srv: srv}
}

// ScalarServerTransport connects instances of corescalargroup.ScalarClient to instances of ScalarServer.
// Don't use this type directly, use NewScalarServerTransport instead.
type ScalarServerTransport struct {
	srv                               *ScalarServer
	trMu                              sync.Mutex
	trScalarAzureLocationScalarServer *ScalarAzureLocationScalarServerTransport
}

// Do implements the policy.Transporter interface for ScalarServerTransport.
func (s *ScalarServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (s *ScalarServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "ScalarAzureLocationScalarClient":
		initServer(&s.trMu, &s.trScalarAzureLocationScalarServer, func() *ScalarAzureLocationScalarServerTransport {
			return NewScalarAzureLocationScalarServerTransport(&s.srv.ScalarAzureLocationScalarServer)
		})
		resp, err = s.trScalarAzureLocationScalarServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

// set this to conditionally intercept incoming requests to ScalarServerTransport
var scalarServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
