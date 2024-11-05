// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"datetimegroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"time"
)

// DatetimeHeaderServer is a fake server for instances of the datetimegroup.DatetimeHeaderClient type.
type DatetimeHeaderServer struct {
	// Default is the fake for method DatetimeHeaderClient.Default
	// HTTP status codes to indicate success: http.StatusNoContent
	Default func(ctx context.Context, value time.Time, options *datetimegroup.DatetimeHeaderClientDefaultOptions) (resp azfake.Responder[datetimegroup.DatetimeHeaderClientDefaultResponse], errResp azfake.ErrorResponder)

	// RFC3339 is the fake for method DatetimeHeaderClient.RFC3339
	// HTTP status codes to indicate success: http.StatusNoContent
	RFC3339 func(ctx context.Context, value time.Time, options *datetimegroup.DatetimeHeaderClientRFC3339Options) (resp azfake.Responder[datetimegroup.DatetimeHeaderClientRFC3339Response], errResp azfake.ErrorResponder)

	// RFC7231 is the fake for method DatetimeHeaderClient.RFC7231
	// HTTP status codes to indicate success: http.StatusNoContent
	RFC7231 func(ctx context.Context, value time.Time, options *datetimegroup.DatetimeHeaderClientRFC7231Options) (resp azfake.Responder[datetimegroup.DatetimeHeaderClientRFC7231Response], errResp azfake.ErrorResponder)

	// UnixTimestamp is the fake for method DatetimeHeaderClient.UnixTimestamp
	// HTTP status codes to indicate success: http.StatusNoContent
	UnixTimestamp func(ctx context.Context, value time.Time, options *datetimegroup.DatetimeHeaderClientUnixTimestampOptions) (resp azfake.Responder[datetimegroup.DatetimeHeaderClientUnixTimestampResponse], errResp azfake.ErrorResponder)

	// UnixTimestampArray is the fake for method DatetimeHeaderClient.UnixTimestampArray
	// HTTP status codes to indicate success: http.StatusNoContent
	UnixTimestampArray func(ctx context.Context, value []time.Time, options *datetimegroup.DatetimeHeaderClientUnixTimestampArrayOptions) (resp azfake.Responder[datetimegroup.DatetimeHeaderClientUnixTimestampArrayResponse], errResp azfake.ErrorResponder)
}

// NewDatetimeHeaderServerTransport creates a new instance of DatetimeHeaderServerTransport with the provided implementation.
// The returned DatetimeHeaderServerTransport instance is connected to an instance of datetimegroup.DatetimeHeaderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatetimeHeaderServerTransport(srv *DatetimeHeaderServer) *DatetimeHeaderServerTransport {
	return &DatetimeHeaderServerTransport{srv: srv}
}

// DatetimeHeaderServerTransport connects instances of datetimegroup.DatetimeHeaderClient to instances of DatetimeHeaderServer.
// Don't use this type directly, use NewDatetimeHeaderServerTransport instead.
type DatetimeHeaderServerTransport struct {
	srv *DatetimeHeaderServer
}

// Do implements the policy.Transporter interface for DatetimeHeaderServerTransport.
func (d *DatetimeHeaderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DatetimeHeaderServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if datetimeHeaderServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = datetimeHeaderServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DatetimeHeaderClient.Default":
				res.resp, res.err = d.dispatchDefault(req)
			case "DatetimeHeaderClient.RFC3339":
				res.resp, res.err = d.dispatchRFC3339(req)
			case "DatetimeHeaderClient.RFC7231":
				res.resp, res.err = d.dispatchRFC7231(req)
			case "DatetimeHeaderClient.UnixTimestamp":
				res.resp, res.err = d.dispatchUnixTimestamp(req)
			case "DatetimeHeaderClient.UnixTimestampArray":
				res.resp, res.err = d.dispatchUnixTimestampArray(req)
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

func (d *DatetimeHeaderServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if d.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	valueParam, err := time.Parse(time.RFC1123, getHeaderValue(req.Header, "value"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Default(req.Context(), valueParam, nil)
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

func (d *DatetimeHeaderServerTransport) dispatchRFC3339(req *http.Request) (*http.Response, error) {
	if d.srv.RFC3339 == nil {
		return nil, &nonRetriableError{errors.New("fake for method RFC3339 not implemented")}
	}
	valueParam, err := time.Parse(time.RFC3339Nano, getHeaderValue(req.Header, "value"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.RFC3339(req.Context(), valueParam, nil)
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

func (d *DatetimeHeaderServerTransport) dispatchRFC7231(req *http.Request) (*http.Response, error) {
	if d.srv.RFC7231 == nil {
		return nil, &nonRetriableError{errors.New("fake for method RFC7231 not implemented")}
	}
	valueParam, err := time.Parse(time.RFC1123, getHeaderValue(req.Header, "value"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.RFC7231(req.Context(), valueParam, nil)
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

func (d *DatetimeHeaderServerTransport) dispatchUnixTimestamp(req *http.Request) (*http.Response, error) {
	if d.srv.UnixTimestamp == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnixTimestamp not implemented")}
	}
	valueParam, err := parseWithCast(getHeaderValue(req.Header, "value"), func(v string) (time.Time, error) {
		p, parseErr := strconv.ParseInt(v, 10, 64)
		if parseErr != nil {
			return time.Time{}, parseErr
		}
		return time.Unix(p, 0).UTC(), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.UnixTimestamp(req.Context(), valueParam, nil)
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

func (d *DatetimeHeaderServerTransport) dispatchUnixTimestampArray(req *http.Request) (*http.Response, error) {
	if d.srv.UnixTimestampArray == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnixTimestampArray not implemented")}
	}
	valueElements := splitHelper(getHeaderValue(req.Header, "value"), ",")
	valueParam := make([]time.Time, len(valueElements))
	for i := 0; i < len(valueElements); i++ {
		p, parseErr := strconv.ParseInt(valueElements[i], 10, 64)
		if parseErr != nil {
			return nil, parseErr
		}
		parsedTimeUnix := time.Unix(p, 0).UTC()
		valueParam[i] = time.Time(parsedTimeUnix)
	}
	respr, errRespr := d.srv.UnixTimestampArray(req.Context(), valueParam, nil)
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

// set this to conditionally intercept incoming requests to DatetimeHeaderServerTransport
var datetimeHeaderServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
