//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/errorsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
)

// PetServer is a fake server for instances of the errorsgroup.PetClient type.
type PetServer struct {
	// DoSomething is the fake for method PetClient.DoSomething
	// HTTP status codes to indicate success: http.StatusOK
	DoSomething func(ctx context.Context, whatAction string, options *errorsgroup.PetClientDoSomethingOptions) (resp azfake.Responder[errorsgroup.PetClientDoSomethingResponse], errResp azfake.ErrorResponder)

	// GetPetByID is the fake for method PetClient.GetPetByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	GetPetByID func(ctx context.Context, petID string, options *errorsgroup.PetClientGetPetByIDOptions) (resp azfake.Responder[errorsgroup.PetClientGetPetByIDResponse], errResp azfake.ErrorResponder)

	// HasModelsParam is the fake for method PetClient.HasModelsParam
	// HTTP status codes to indicate success: http.StatusOK
	HasModelsParam func(ctx context.Context, options *errorsgroup.PetClientHasModelsParamOptions) (resp azfake.Responder[errorsgroup.PetClientHasModelsParamResponse], errResp azfake.ErrorResponder)
}

// NewPetServerTransport creates a new instance of PetServerTransport with the provided implementation.
// The returned PetServerTransport instance is connected to an instance of errorsgroup.PetClient by way of the
// undefined.Transporter field.
func NewPetServerTransport(srv *PetServer) *PetServerTransport {
	return &PetServerTransport{srv: srv}
}

// PetServerTransport connects instances of errorsgroup.PetClient to instances of PetServer.
// Don't use this type directly, use NewPetServerTransport instead.
type PetServerTransport struct {
	srv *PetServer
}

// Do implements the policy.Transporter interface for PetServerTransport.
func (p *PetServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PetClient.DoSomething":
		resp, err = p.dispatchDoSomething(req)
	case "PetClient.GetPetByID":
		resp, err = p.dispatchGetPetByID(req)
	case "PetClient.HasModelsParam":
		resp, err = p.dispatchHasModelsParam(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PetServerTransport) dispatchDoSomething(req *http.Request) (*http.Response, error) {
	if p.srv.DoSomething == nil {
		return nil, &nonRetriableError{errors.New("method DoSomething not implemented")}
	}
	const regexStr = "/errorStatusCodes/Pets/doSomething/(?P<whatAction>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.DoSomething(req.Context(), matches[regex.SubexpIndex("whatAction")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PetAction, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PetServerTransport) dispatchGetPetByID(req *http.Request) (*http.Response, error) {
	if p.srv.GetPetByID == nil {
		return nil, &nonRetriableError{errors.New("method GetPetByID not implemented")}
	}
	const regexStr = "/errorStatusCodes/Pets/(?P<petId>[a-zA-Z0-9-_]+)/GetPet"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.GetPetByID(req.Context(), matches[regex.SubexpIndex("petId")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PetServerTransport) dispatchHasModelsParam(req *http.Request) (*http.Response, error) {
	if p.srv.HasModelsParam == nil {
		return nil, &nonRetriableError{errors.New("method HasModelsParam not implemented")}
	}
	qp := req.URL.Query()
	modelsParam := getOptional(qp.Get("models"))
	var options *errorsgroup.PetClientHasModelsParamOptions
	if modelsParam != nil {
		options = &errorsgroup.PetClientHasModelsParamOptions{
			Models: modelsParam,
		}
	}
	respr, errRespr := p.srv.HasModelsParam(req.Context(), options)
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
	return resp, nil
}