// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"io"
	"mime"
	"mime/multipart"
	"multipartgroup"
	"net/http"
)

// MultiPartFormDataServer is a fake server for instances of the multipartgroup.MultiPartFormDataClient type.
type MultiPartFormDataServer struct {
	// AnonymousModel is the fake for method MultiPartFormDataClient.AnonymousModel
	// HTTP status codes to indicate success: http.StatusNoContent
	AnonymousModel func(ctx context.Context, anonymousModelRequest multipartgroup.AnonymousModelRequest, options *multipartgroup.MultiPartFormDataClientAnonymousModelOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientAnonymousModelResponse], errResp azfake.ErrorResponder)

	// Basic is the fake for method MultiPartFormDataClient.Basic
	// HTTP status codes to indicate success: http.StatusNoContent
	Basic func(ctx context.Context, body multipartgroup.MultiPartRequest, options *multipartgroup.MultiPartFormDataClientBasicOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientBasicResponse], errResp azfake.ErrorResponder)

	// BinaryArrayParts is the fake for method MultiPartFormDataClient.BinaryArrayParts
	// HTTP status codes to indicate success: http.StatusNoContent
	BinaryArrayParts func(ctx context.Context, body multipartgroup.BinaryArrayPartsRequest, options *multipartgroup.MultiPartFormDataClientBinaryArrayPartsOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientBinaryArrayPartsResponse], errResp azfake.ErrorResponder)

	// CheckFileNameAndContentType is the fake for method MultiPartFormDataClient.CheckFileNameAndContentType
	// HTTP status codes to indicate success: http.StatusNoContent
	CheckFileNameAndContentType func(ctx context.Context, body multipartgroup.MultiPartRequest, options *multipartgroup.MultiPartFormDataClientCheckFileNameAndContentTypeOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientCheckFileNameAndContentTypeResponse], errResp azfake.ErrorResponder)

	// Complex is the fake for method MultiPartFormDataClient.Complex
	// HTTP status codes to indicate success: http.StatusNoContent
	Complex func(ctx context.Context, body multipartgroup.ComplexPartsRequest, options *multipartgroup.MultiPartFormDataClientComplexOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientComplexResponse], errResp azfake.ErrorResponder)

	// JSONArrayParts is the fake for method MultiPartFormDataClient.JSONArrayParts
	// HTTP status codes to indicate success: http.StatusNoContent
	JSONArrayParts func(ctx context.Context, body multipartgroup.JSONArrayPartsRequest, options *multipartgroup.MultiPartFormDataClientJSONArrayPartsOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientJSONArrayPartsResponse], errResp azfake.ErrorResponder)

	// JSONPart is the fake for method MultiPartFormDataClient.JSONPart
	// HTTP status codes to indicate success: http.StatusNoContent
	JSONPart func(ctx context.Context, body multipartgroup.JSONPartRequest, options *multipartgroup.MultiPartFormDataClientJSONPartOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientJSONPartResponse], errResp azfake.ErrorResponder)

	// MultiBinaryParts is the fake for method MultiPartFormDataClient.MultiBinaryParts
	// HTTP status codes to indicate success: http.StatusNoContent
	MultiBinaryParts func(ctx context.Context, body multipartgroup.MultiBinaryPartsRequest, options *multipartgroup.MultiPartFormDataClientMultiBinaryPartsOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientMultiBinaryPartsResponse], errResp azfake.ErrorResponder)
}

// NewMultiPartFormDataServerTransport creates a new instance of MultiPartFormDataServerTransport with the provided implementation.
// The returned MultiPartFormDataServerTransport instance is connected to an instance of multipartgroup.MultiPartFormDataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMultiPartFormDataServerTransport(srv *MultiPartFormDataServer) *MultiPartFormDataServerTransport {
	return &MultiPartFormDataServerTransport{srv: srv}
}

// MultiPartFormDataServerTransport connects instances of multipartgroup.MultiPartFormDataClient to instances of MultiPartFormDataServer.
// Don't use this type directly, use NewMultiPartFormDataServerTransport instead.
type MultiPartFormDataServerTransport struct {
	srv *MultiPartFormDataServer
}

// Do implements the policy.Transporter interface for MultiPartFormDataServerTransport.
func (m *MultiPartFormDataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MultiPartFormDataServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "MultiPartFormDataClient.AnonymousModel":
		resp, err = m.dispatchAnonymousModel(req)
	case "MultiPartFormDataClient.Basic":
		resp, err = m.dispatchBasic(req)
	case "MultiPartFormDataClient.BinaryArrayParts":
		resp, err = m.dispatchBinaryArrayParts(req)
	case "MultiPartFormDataClient.CheckFileNameAndContentType":
		resp, err = m.dispatchCheckFileNameAndContentType(req)
	case "MultiPartFormDataClient.Complex":
		resp, err = m.dispatchComplex(req)
	case "MultiPartFormDataClient.JSONArrayParts":
		resp, err = m.dispatchJSONArrayParts(req)
	case "MultiPartFormDataClient.JSONPart":
		resp, err = m.dispatchJSONPart(req)
	case "MultiPartFormDataClient.MultiBinaryParts":
		resp, err = m.dispatchMultiBinaryParts(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (m *MultiPartFormDataServerTransport) dispatchAnonymousModel(req *http.Request) (*http.Response, error) {
	if m.srv.AnonymousModel == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnonymousModel not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var anonymousModelRequest multipartgroup.AnonymousModelRequest
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "profileImage":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			anonymousModelRequest.ProfileImage.Body = streaming.NopCloser(bytes.NewReader(content))
			anonymousModelRequest.ProfileImage.ContentType = part.Header.Get("Content-Type")
			anonymousModelRequest.ProfileImage.Filename = part.FileName()
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := m.srv.AnonymousModel(req.Context(), anonymousModelRequest, nil)
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

func (m *MultiPartFormDataServerTransport) dispatchBasic(req *http.Request) (*http.Response, error) {
	if m.srv.Basic == nil {
		return nil, &nonRetriableError{errors.New("fake for method Basic not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var body multipartgroup.MultiPartRequest
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "id":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ID = string(content)
		case "profileImage":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ProfileImage.Body = streaming.NopCloser(bytes.NewReader(content))
			body.ProfileImage.ContentType = part.Header.Get("Content-Type")
			body.ProfileImage.Filename = part.FileName()
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := m.srv.Basic(req.Context(), body, nil)
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

func (m *MultiPartFormDataServerTransport) dispatchBinaryArrayParts(req *http.Request) (*http.Response, error) {
	if m.srv.BinaryArrayParts == nil {
		return nil, &nonRetriableError{errors.New("fake for method BinaryArrayParts not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var body multipartgroup.BinaryArrayPartsRequest
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "id":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ID = string(content)
		case "pictures":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.Pictures = append(body.Pictures, streaming.MultipartContent{
				Body:        streaming.NopCloser(bytes.NewReader(content)),
				ContentType: part.Header.Get("Content-Type"),
				Filename:    part.FileName(),
			})
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := m.srv.BinaryArrayParts(req.Context(), body, nil)
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

func (m *MultiPartFormDataServerTransport) dispatchCheckFileNameAndContentType(req *http.Request) (*http.Response, error) {
	if m.srv.CheckFileNameAndContentType == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckFileNameAndContentType not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var body multipartgroup.MultiPartRequest
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "id":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ID = string(content)
		case "profileImage":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ProfileImage.Body = streaming.NopCloser(bytes.NewReader(content))
			body.ProfileImage.ContentType = part.Header.Get("Content-Type")
			body.ProfileImage.Filename = part.FileName()
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := m.srv.CheckFileNameAndContentType(req.Context(), body, nil)
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

func (m *MultiPartFormDataServerTransport) dispatchComplex(req *http.Request) (*http.Response, error) {
	if m.srv.Complex == nil {
		return nil, &nonRetriableError{errors.New("fake for method Complex not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var body multipartgroup.ComplexPartsRequest
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "address":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			if err = json.Unmarshal(content, &body.Address); err != nil {
				return nil, err
			}
		case "id":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ID = string(content)
		case "pictures":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.Pictures = append(body.Pictures, streaming.MultipartContent{
				Body:        streaming.NopCloser(bytes.NewReader(content)),
				ContentType: part.Header.Get("Content-Type"),
				Filename:    part.FileName(),
			})
		case "previousAddresses":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			if err = json.Unmarshal(content, &body.PreviousAddresses); err != nil {
				return nil, err
			}
		case "profileImage":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ProfileImage.Body = streaming.NopCloser(bytes.NewReader(content))
			body.ProfileImage.ContentType = part.Header.Get("Content-Type")
			body.ProfileImage.Filename = part.FileName()
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := m.srv.Complex(req.Context(), body, nil)
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

func (m *MultiPartFormDataServerTransport) dispatchJSONArrayParts(req *http.Request) (*http.Response, error) {
	if m.srv.JSONArrayParts == nil {
		return nil, &nonRetriableError{errors.New("fake for method JSONArrayParts not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var body multipartgroup.JSONArrayPartsRequest
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "previousAddresses":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			if err = json.Unmarshal(content, &body.PreviousAddresses); err != nil {
				return nil, err
			}
		case "profileImage":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ProfileImage.Body = streaming.NopCloser(bytes.NewReader(content))
			body.ProfileImage.ContentType = part.Header.Get("Content-Type")
			body.ProfileImage.Filename = part.FileName()
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := m.srv.JSONArrayParts(req.Context(), body, nil)
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

func (m *MultiPartFormDataServerTransport) dispatchJSONPart(req *http.Request) (*http.Response, error) {
	if m.srv.JSONPart == nil {
		return nil, &nonRetriableError{errors.New("fake for method JSONPart not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var body multipartgroup.JSONPartRequest
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "address":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			if err = json.Unmarshal(content, &body.Address); err != nil {
				return nil, err
			}
		case "profileImage":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ProfileImage.Body = streaming.NopCloser(bytes.NewReader(content))
			body.ProfileImage.ContentType = part.Header.Get("Content-Type")
			body.ProfileImage.Filename = part.FileName()
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := m.srv.JSONPart(req.Context(), body, nil)
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

func (m *MultiPartFormDataServerTransport) dispatchMultiBinaryParts(req *http.Request) (*http.Response, error) {
	if m.srv.MultiBinaryParts == nil {
		return nil, &nonRetriableError{errors.New("fake for method MultiBinaryParts not implemented")}
	}
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	reader := multipart.NewReader(req.Body, params["boundary"])
	var body multipartgroup.MultiBinaryPartsRequest
	for {
		var part *multipart.Part
		part, err = reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		var content []byte
		switch fn := part.FormName(); fn {
		case "profileImage":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.ProfileImage.Body = streaming.NopCloser(bytes.NewReader(content))
			body.ProfileImage.ContentType = part.Header.Get("Content-Type")
			body.ProfileImage.Filename = part.FileName()
		case "picture":
			content, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			body.Picture.Body = streaming.NopCloser(bytes.NewReader(content))
			body.Picture.ContentType = part.Header.Get("Content-Type")
			body.Picture.Filename = part.FileName()
		default:
			return nil, fmt.Errorf("unexpected part %s", fn)
		}
	}
	respr, errRespr := m.srv.MultiBinaryParts(req.Context(), body, nil)
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
