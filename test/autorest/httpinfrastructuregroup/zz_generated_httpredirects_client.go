//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPRedirectsClient contains the methods for the HTTPRedirects group.
// Don't use this type directly, use NewHTTPRedirectsClient() instead.
type HTTPRedirectsClient struct {
	pl runtime.Pipeline
}

// NewHTTPRedirectsClient creates a new instance of HTTPRedirectsClient with the specified values.
// options - pass nil to accept the default values.
func NewHTTPRedirectsClient(options *azcore.ClientOptions) *HTTPRedirectsClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &HTTPRedirectsClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// Delete307 - Delete redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsDelete307Options contains the optional parameters for the HTTPRedirectsClient.Delete307 method.
func (client *HTTPRedirectsClient) Delete307(ctx context.Context, options *HTTPRedirectsDelete307Options) (HTTPRedirectsDelete307Response, error) {
	req, err := client.delete307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsDelete307Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsDelete307Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRedirectsDelete307Response{}, client.delete307HandleError(resp)
	}
	return HTTPRedirectsDelete307Response{RawResponse: resp}, nil
}

// delete307CreateRequest creates the Delete307 request.
func (client *HTTPRedirectsClient) delete307CreateRequest(ctx context.Context, options *HTTPRedirectsDelete307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// delete307HandleError handles the Delete307 error response.
func (client *HTTPRedirectsClient) delete307HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get300 - Return 300 status code and redirect to /http/success/200
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsGet300Options contains the optional parameters for the HTTPRedirectsClient.Get300 method.
func (client *HTTPRedirectsClient) Get300(ctx context.Context, options *HTTPRedirectsGet300Options) (HTTPRedirectsGet300Response, error) {
	req, err := client.get300CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsGet300Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsGet300Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusMultipleChoices) {
		return HTTPRedirectsGet300Response{}, client.get300HandleError(resp)
	}
	return client.get300HandleResponse(resp)
}

// get300CreateRequest creates the Get300 request.
func (client *HTTPRedirectsClient) get300CreateRequest(ctx context.Context, options *HTTPRedirectsGet300Options) (*policy.Request, error) {
	urlPath := "/http/redirect/300"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// get300HandleResponse handles the Get300 response.
func (client *HTTPRedirectsClient) get300HandleResponse(resp *http.Response) (HTTPRedirectsGet300Response, error) {
	result := HTTPRedirectsGet300Response{RawResponse: resp}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StringArray); err != nil {
		return HTTPRedirectsGet300Response{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// get300HandleError handles the Get300 error response.
func (client *HTTPRedirectsClient) get300HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get301 - Return 301 status code and redirect to /http/success/200
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsGet301Options contains the optional parameters for the HTTPRedirectsClient.Get301 method.
func (client *HTTPRedirectsClient) Get301(ctx context.Context, options *HTTPRedirectsGet301Options) (HTTPRedirectsGet301Response, error) {
	req, err := client.get301CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsGet301Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsGet301Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRedirectsGet301Response{}, client.get301HandleError(resp)
	}
	return HTTPRedirectsGet301Response{RawResponse: resp}, nil
}

// get301CreateRequest creates the Get301 request.
func (client *HTTPRedirectsClient) get301CreateRequest(ctx context.Context, options *HTTPRedirectsGet301Options) (*policy.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// get301HandleError handles the Get301 error response.
func (client *HTTPRedirectsClient) get301HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get302 - Return 302 status code and redirect to /http/success/200
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsGet302Options contains the optional parameters for the HTTPRedirectsClient.Get302 method.
func (client *HTTPRedirectsClient) Get302(ctx context.Context, options *HTTPRedirectsGet302Options) (HTTPRedirectsGet302Response, error) {
	req, err := client.get302CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsGet302Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsGet302Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRedirectsGet302Response{}, client.get302HandleError(resp)
	}
	return HTTPRedirectsGet302Response{RawResponse: resp}, nil
}

// get302CreateRequest creates the Get302 request.
func (client *HTTPRedirectsClient) get302CreateRequest(ctx context.Context, options *HTTPRedirectsGet302Options) (*policy.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// get302HandleError handles the Get302 error response.
func (client *HTTPRedirectsClient) get302HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get307 - Redirect get with 307, resulting in a 200 success
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsGet307Options contains the optional parameters for the HTTPRedirectsClient.Get307 method.
func (client *HTTPRedirectsClient) Get307(ctx context.Context, options *HTTPRedirectsGet307Options) (HTTPRedirectsGet307Response, error) {
	req, err := client.get307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsGet307Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsGet307Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRedirectsGet307Response{}, client.get307HandleError(resp)
	}
	return HTTPRedirectsGet307Response{RawResponse: resp}, nil
}

// get307CreateRequest creates the Get307 request.
func (client *HTTPRedirectsClient) get307CreateRequest(ctx context.Context, options *HTTPRedirectsGet307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// get307HandleError handles the Get307 error response.
func (client *HTTPRedirectsClient) get307HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Head300 - Return 300 status code and redirect to /http/success/200
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsHead300Options contains the optional parameters for the HTTPRedirectsClient.Head300 method.
func (client *HTTPRedirectsClient) Head300(ctx context.Context, options *HTTPRedirectsHead300Options) (HTTPRedirectsHead300Response, error) {
	req, err := client.head300CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsHead300Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsHead300Response{}, err
	}
	return client.head300HandleResponse(resp)
}

// head300CreateRequest creates the Head300 request.
func (client *HTTPRedirectsClient) head300CreateRequest(ctx context.Context, options *HTTPRedirectsHead300Options) (*policy.Request, error) {
	urlPath := "/http/redirect/300"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// head300HandleResponse handles the Head300 response.
func (client *HTTPRedirectsClient) head300HandleResponse(resp *http.Response) (HTTPRedirectsHead300Response, error) {
	result := HTTPRedirectsHead300Response{RawResponse: resp}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// Head301 - Return 301 status code and redirect to /http/success/200
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsHead301Options contains the optional parameters for the HTTPRedirectsClient.Head301 method.
func (client *HTTPRedirectsClient) Head301(ctx context.Context, options *HTTPRedirectsHead301Options) (HTTPRedirectsHead301Response, error) {
	req, err := client.head301CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsHead301Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsHead301Response{}, err
	}
	result := HTTPRedirectsHead301Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head301CreateRequest creates the Head301 request.
func (client *HTTPRedirectsClient) head301CreateRequest(ctx context.Context, options *HTTPRedirectsHead301Options) (*policy.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Head302 - Return 302 status code and redirect to /http/success/200
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsHead302Options contains the optional parameters for the HTTPRedirectsClient.Head302 method.
func (client *HTTPRedirectsClient) Head302(ctx context.Context, options *HTTPRedirectsHead302Options) (HTTPRedirectsHead302Response, error) {
	req, err := client.head302CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsHead302Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsHead302Response{}, err
	}
	result := HTTPRedirectsHead302Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head302CreateRequest creates the Head302 request.
func (client *HTTPRedirectsClient) head302CreateRequest(ctx context.Context, options *HTTPRedirectsHead302Options) (*policy.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Head307 - Redirect with 307, resulting in a 200 success
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsHead307Options contains the optional parameters for the HTTPRedirectsClient.Head307 method.
func (client *HTTPRedirectsClient) Head307(ctx context.Context, options *HTTPRedirectsHead307Options) (HTTPRedirectsHead307Response, error) {
	req, err := client.head307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsHead307Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsHead307Response{}, err
	}
	result := HTTPRedirectsHead307Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head307CreateRequest creates the Head307 request.
func (client *HTTPRedirectsClient) head307CreateRequest(ctx context.Context, options *HTTPRedirectsHead307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Options307 - options redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsOptions307Options contains the optional parameters for the HTTPRedirectsClient.Options307 method.
func (client *HTTPRedirectsClient) Options307(ctx context.Context, options *HTTPRedirectsOptions307Options) (HTTPRedirectsOptions307Response, error) {
	req, err := client.options307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsOptions307Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsOptions307Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRedirectsOptions307Response{}, client.options307HandleError(resp)
	}
	return HTTPRedirectsOptions307Response{RawResponse: resp}, nil
}

// options307CreateRequest creates the Options307 request.
func (client *HTTPRedirectsClient) options307CreateRequest(ctx context.Context, options *HTTPRedirectsOptions307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// options307HandleError handles the Options307 error response.
func (client *HTTPRedirectsClient) options307HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Patch302 - Patch true Boolean value in request returns 302. This request should not be automatically redirected, but should
// return the received 302 to the caller for evaluation
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsPatch302Options contains the optional parameters for the HTTPRedirectsClient.Patch302 method.
func (client *HTTPRedirectsClient) Patch302(ctx context.Context, options *HTTPRedirectsPatch302Options) (HTTPRedirectsPatch302Response, error) {
	req, err := client.patch302CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsPatch302Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsPatch302Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusFound) {
		return HTTPRedirectsPatch302Response{}, client.patch302HandleError(resp)
	}
	return client.patch302HandleResponse(resp)
}

// patch302CreateRequest creates the Patch302 request.
func (client *HTTPRedirectsClient) patch302CreateRequest(ctx context.Context, options *HTTPRedirectsPatch302Options) (*policy.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// patch302HandleResponse handles the Patch302 response.
func (client *HTTPRedirectsClient) patch302HandleResponse(resp *http.Response) (HTTPRedirectsPatch302Response, error) {
	result := HTTPRedirectsPatch302Response{RawResponse: resp}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// patch302HandleError handles the Patch302 error response.
func (client *HTTPRedirectsClient) patch302HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Patch307 - Patch redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsPatch307Options contains the optional parameters for the HTTPRedirectsClient.Patch307 method.
func (client *HTTPRedirectsClient) Patch307(ctx context.Context, options *HTTPRedirectsPatch307Options) (HTTPRedirectsPatch307Response, error) {
	req, err := client.patch307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsPatch307Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsPatch307Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRedirectsPatch307Response{}, client.patch307HandleError(resp)
	}
	return HTTPRedirectsPatch307Response{RawResponse: resp}, nil
}

// patch307CreateRequest creates the Patch307 request.
func (client *HTTPRedirectsClient) patch307CreateRequest(ctx context.Context, options *HTTPRedirectsPatch307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// patch307HandleError handles the Patch307 error response.
func (client *HTTPRedirectsClient) patch307HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Post303 - Post true Boolean value in request returns 303. This request should be automatically redirected usign a get,
// ultimately returning a 200 status code
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsPost303Options contains the optional parameters for the HTTPRedirectsClient.Post303 method.
func (client *HTTPRedirectsClient) Post303(ctx context.Context, options *HTTPRedirectsPost303Options) (HTTPRedirectsPost303Response, error) {
	req, err := client.post303CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsPost303Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsPost303Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusSeeOther) {
		return HTTPRedirectsPost303Response{}, client.post303HandleError(resp)
	}
	return client.post303HandleResponse(resp)
}

// post303CreateRequest creates the Post303 request.
func (client *HTTPRedirectsClient) post303CreateRequest(ctx context.Context, options *HTTPRedirectsPost303Options) (*policy.Request, error) {
	urlPath := "/http/redirect/303"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// post303HandleResponse handles the Post303 response.
func (client *HTTPRedirectsClient) post303HandleResponse(resp *http.Response) (HTTPRedirectsPost303Response, error) {
	result := HTTPRedirectsPost303Response{RawResponse: resp}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// post303HandleError handles the Post303 error response.
func (client *HTTPRedirectsClient) post303HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Post307 - Post redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsPost307Options contains the optional parameters for the HTTPRedirectsClient.Post307 method.
func (client *HTTPRedirectsClient) Post307(ctx context.Context, options *HTTPRedirectsPost307Options) (HTTPRedirectsPost307Response, error) {
	req, err := client.post307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsPost307Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsPost307Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRedirectsPost307Response{}, client.post307HandleError(resp)
	}
	return HTTPRedirectsPost307Response{RawResponse: resp}, nil
}

// post307CreateRequest creates the Post307 request.
func (client *HTTPRedirectsClient) post307CreateRequest(ctx context.Context, options *HTTPRedirectsPost307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// post307HandleError handles the Post307 error response.
func (client *HTTPRedirectsClient) post307HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Put301 - Put true Boolean value in request returns 301. This request should not be automatically redirected, but should
// return the received 301 to the caller for evaluation
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsPut301Options contains the optional parameters for the HTTPRedirectsClient.Put301 method.
func (client *HTTPRedirectsClient) Put301(ctx context.Context, options *HTTPRedirectsPut301Options) (HTTPRedirectsPut301Response, error) {
	req, err := client.put301CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsPut301Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsPut301Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusMovedPermanently) {
		return HTTPRedirectsPut301Response{}, client.put301HandleError(resp)
	}
	return client.put301HandleResponse(resp)
}

// put301CreateRequest creates the Put301 request.
func (client *HTTPRedirectsClient) put301CreateRequest(ctx context.Context, options *HTTPRedirectsPut301Options) (*policy.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// put301HandleResponse handles the Put301 response.
func (client *HTTPRedirectsClient) put301HandleResponse(resp *http.Response) (HTTPRedirectsPut301Response, error) {
	result := HTTPRedirectsPut301Response{RawResponse: resp}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// put301HandleError handles the Put301 error response.
func (client *HTTPRedirectsClient) put301HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Put307 - Put redirected with 307, resulting in a 200 after redirect
// If the operation fails it returns the *Error error type.
// options - HTTPRedirectsPut307Options contains the optional parameters for the HTTPRedirectsClient.Put307 method.
func (client *HTTPRedirectsClient) Put307(ctx context.Context, options *HTTPRedirectsPut307Options) (HTTPRedirectsPut307Response, error) {
	req, err := client.put307CreateRequest(ctx, options)
	if err != nil {
		return HTTPRedirectsPut307Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPRedirectsPut307Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPRedirectsPut307Response{}, client.put307HandleError(resp)
	}
	return HTTPRedirectsPut307Response{RawResponse: resp}, nil
}

// put307CreateRequest creates the Put307 request.
func (client *HTTPRedirectsClient) put307CreateRequest(ctx context.Context, options *HTTPRedirectsPut307Options) (*policy.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// put307HandleError handles the Put307 error response.
func (client *HTTPRedirectsClient) put307HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
