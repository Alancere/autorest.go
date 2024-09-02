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

// NumericServer is a fake server for instances of the numericgroup.NumericClient type.
type NumericServer struct {
	// NumericPropertyServer contains the fakes for client NumericPropertyClient
	NumericPropertyServer NumericPropertyServer
}

// NewNumericServerTransport creates a new instance of NumericServerTransport with the provided implementation.
// The returned NumericServerTransport instance is connected to an instance of numericgroup.NumericClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNumericServerTransport(srv *NumericServer) *NumericServerTransport {
	return &NumericServerTransport{srv: srv}
}

// NumericServerTransport connects instances of numericgroup.NumericClient to instances of NumericServer.
// Don't use this type directly, use NewNumericServerTransport instead.
type NumericServerTransport struct {
	srv                     *NumericServer
	trMu                    sync.Mutex
	trNumericPropertyServer *NumericPropertyServerTransport
}

// Do implements the policy.Transporter interface for NumericServerTransport.
func (n *NumericServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (n *NumericServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "NumericPropertyClient":
		initServer(&n.trMu, &n.trNumericPropertyServer, func() *NumericPropertyServerTransport {
			return NewNumericPropertyServerTransport(&n.srv.NumericPropertyServer)
		})
		resp, err = n.trNumericPropertyServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
