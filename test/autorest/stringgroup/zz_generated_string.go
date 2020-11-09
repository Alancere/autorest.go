// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// StringOperations contains the methods for the String group.
type StringOperations interface {
	// GetBase64Encoded - Get value that is base64 encoded
	GetBase64Encoded(ctx context.Context, options *StringGetBase64EncodedOptions) (*ByteArrayResponse, error)
	// GetBase64URLEncoded - Get value that is base64url encoded
	GetBase64URLEncoded(ctx context.Context, options *StringGetBase64URLEncodedOptions) (*ByteArrayResponse, error)
	// GetEmpty - Get empty string value value ''
	GetEmpty(ctx context.Context, options *StringGetEmptyOptions) (*StringResponse, error)
	// GetMBCS - Get mbcs string value '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
	GetMBCS(ctx context.Context, options *StringGetMBCSOptions) (*StringResponse, error)
	// GetNotProvided - Get String value when no string value is sent in response payload
	GetNotProvided(ctx context.Context, options *StringGetNotProvidedOptions) (*StringResponse, error)
	// GetNull - Get null string value value
	GetNull(ctx context.Context, options *StringGetNullOptions) (*StringResponse, error)
	// GetNullBase64URLEncoded - Get null value that is expected to be base64url encoded
	GetNullBase64URLEncoded(ctx context.Context, options *StringGetNullBase64URLEncodedOptions) (*ByteArrayResponse, error)
	// GetWhitespace - Get string value with leading and trailing whitespace 'Now is the time for all good men to come to the aid of their country'
	GetWhitespace(ctx context.Context, options *StringGetWhitespaceOptions) (*StringResponse, error)
	// PutBase64URLEncoded - Put value that is base64url encoded
	PutBase64URLEncoded(ctx context.Context, stringBody []byte, options *StringPutBase64URLEncodedOptions) (*http.Response, error)
	// PutEmpty - Set string value empty ''
	PutEmpty(ctx context.Context, options *StringPutEmptyOptions) (*http.Response, error)
	// PutMBCS - Set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
	PutMBCS(ctx context.Context, options *StringPutMBCSOptions) (*http.Response, error)
	// PutNull - Set string value null
	PutNull(ctx context.Context, options *StringPutNullOptions) (*http.Response, error)
	// PutWhitespace - Set String value with leading and trailing whitespace 'Now is the time for all good men to come to the aid of their country'
	PutWhitespace(ctx context.Context, options *StringPutWhitespaceOptions) (*http.Response, error)
}

// StringClient implements the StringOperations interface.
// Don't use this type directly, use NewStringClient() instead.
type StringClient struct {
	con *Connection
}

// NewStringClient creates a new instance of StringClient with the specified values.
func NewStringClient(con *Connection) StringOperations {
	return &StringClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *StringClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetBase64Encoded - Get value that is base64 encoded
func (client *StringClient) GetBase64Encoded(ctx context.Context, options *StringGetBase64EncodedOptions) (*ByteArrayResponse, error) {
	req, err := client.GetBase64EncodedCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetBase64EncodedHandleError(resp)
	}
	result, err := client.GetBase64EncodedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBase64EncodedCreateRequest creates the GetBase64Encoded request.
func (client *StringClient) GetBase64EncodedCreateRequest(ctx context.Context, options *StringGetBase64EncodedOptions) (*azcore.Request, error) {
	urlPath := "/string/base64Encoding"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetBase64EncodedHandleResponse handles the GetBase64Encoded response.
func (client *StringClient) GetBase64EncodedHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64StdFormat)
}

// GetBase64EncodedHandleError handles the GetBase64Encoded error response.
func (client *StringClient) GetBase64EncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetBase64URLEncoded - Get value that is base64url encoded
func (client *StringClient) GetBase64URLEncoded(ctx context.Context, options *StringGetBase64URLEncodedOptions) (*ByteArrayResponse, error) {
	req, err := client.GetBase64URLEncodedCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetBase64URLEncodedHandleError(resp)
	}
	result, err := client.GetBase64URLEncodedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBase64URLEncodedCreateRequest creates the GetBase64URLEncoded request.
func (client *StringClient) GetBase64URLEncodedCreateRequest(ctx context.Context, options *StringGetBase64URLEncodedOptions) (*azcore.Request, error) {
	urlPath := "/string/base64UrlEncoding"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetBase64URLEncodedHandleResponse handles the GetBase64URLEncoded response.
func (client *StringClient) GetBase64URLEncodedHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64URLFormat)
}

// GetBase64URLEncodedHandleError handles the GetBase64URLEncoded error response.
func (client *StringClient) GetBase64URLEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetEmpty - Get empty string value value ''
func (client *StringClient) GetEmpty(ctx context.Context, options *StringGetEmptyOptions) (*StringResponse, error) {
	req, err := client.GetEmptyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetEmptyHandleError(resp)
	}
	result, err := client.GetEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEmptyCreateRequest creates the GetEmpty request.
func (client *StringClient) GetEmptyCreateRequest(ctx context.Context, options *StringGetEmptyOptions) (*azcore.Request, error) {
	urlPath := "/string/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (client *StringClient) GetEmptyHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetEmptyHandleError handles the GetEmpty error response.
func (client *StringClient) GetEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetMBCS - Get mbcs string value '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
func (client *StringClient) GetMBCS(ctx context.Context, options *StringGetMBCSOptions) (*StringResponse, error) {
	req, err := client.GetMBCSCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetMBCSHandleError(resp)
	}
	result, err := client.GetMBCSHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMBCSCreateRequest creates the GetMBCS request.
func (client *StringClient) GetMBCSCreateRequest(ctx context.Context, options *StringGetMBCSOptions) (*azcore.Request, error) {
	urlPath := "/string/mbcs"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetMBCSHandleResponse handles the GetMBCS response.
func (client *StringClient) GetMBCSHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetMBCSHandleError handles the GetMBCS error response.
func (client *StringClient) GetMBCSHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotProvided - Get String value when no string value is sent in response payload
func (client *StringClient) GetNotProvided(ctx context.Context, options *StringGetNotProvidedOptions) (*StringResponse, error) {
	req, err := client.GetNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNotProvidedHandleError(resp)
	}
	result, err := client.GetNotProvidedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNotProvidedCreateRequest creates the GetNotProvided request.
func (client *StringClient) GetNotProvidedCreateRequest(ctx context.Context, options *StringGetNotProvidedOptions) (*azcore.Request, error) {
	urlPath := "/string/notProvided"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNotProvidedHandleResponse handles the GetNotProvided response.
func (client *StringClient) GetNotProvidedHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNotProvidedHandleError handles the GetNotProvided error response.
func (client *StringClient) GetNotProvidedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get null string value value
func (client *StringClient) GetNull(ctx context.Context, options *StringGetNullOptions) (*StringResponse, error) {
	req, err := client.GetNullCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNullHandleError(resp)
	}
	result, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNullCreateRequest creates the GetNull request.
func (client *StringClient) GetNullCreateRequest(ctx context.Context, options *StringGetNullOptions) (*azcore.Request, error) {
	urlPath := "/string/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNullHandleResponse handles the GetNull response.
func (client *StringClient) GetNullHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNullHandleError handles the GetNull error response.
func (client *StringClient) GetNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNullBase64URLEncoded - Get null value that is expected to be base64url encoded
func (client *StringClient) GetNullBase64URLEncoded(ctx context.Context, options *StringGetNullBase64URLEncodedOptions) (*ByteArrayResponse, error) {
	req, err := client.GetNullBase64URLEncodedCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNullBase64URLEncodedHandleError(resp)
	}
	result, err := client.GetNullBase64URLEncodedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNullBase64URLEncodedCreateRequest creates the GetNullBase64URLEncoded request.
func (client *StringClient) GetNullBase64URLEncodedCreateRequest(ctx context.Context, options *StringGetNullBase64URLEncodedOptions) (*azcore.Request, error) {
	urlPath := "/string/nullBase64UrlEncoding"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNullBase64URLEncodedHandleResponse handles the GetNullBase64URLEncoded response.
func (client *StringClient) GetNullBase64URLEncodedHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64URLFormat)
}

// GetNullBase64URLEncodedHandleError handles the GetNullBase64URLEncoded error response.
func (client *StringClient) GetNullBase64URLEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetWhitespace - Get string value with leading and trailing whitespace 'Now is the time for all good men to come to the aid of their country'
func (client *StringClient) GetWhitespace(ctx context.Context, options *StringGetWhitespaceOptions) (*StringResponse, error) {
	req, err := client.GetWhitespaceCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetWhitespaceHandleError(resp)
	}
	result, err := client.GetWhitespaceHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWhitespaceCreateRequest creates the GetWhitespace request.
func (client *StringClient) GetWhitespaceCreateRequest(ctx context.Context, options *StringGetWhitespaceOptions) (*azcore.Request, error) {
	urlPath := "/string/whitespace"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetWhitespaceHandleResponse handles the GetWhitespace response.
func (client *StringClient) GetWhitespaceHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetWhitespaceHandleError handles the GetWhitespace error response.
func (client *StringClient) GetWhitespaceHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutBase64URLEncoded - Put value that is base64url encoded
func (client *StringClient) PutBase64URLEncoded(ctx context.Context, stringBody []byte, options *StringPutBase64URLEncodedOptions) (*http.Response, error) {
	req, err := client.PutBase64URLEncodedCreateRequest(ctx, stringBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutBase64URLEncodedHandleError(resp)
	}
	return resp.Response, nil
}

// PutBase64URLEncodedCreateRequest creates the PutBase64URLEncoded request.
func (client *StringClient) PutBase64URLEncodedCreateRequest(ctx context.Context, stringBody []byte, options *StringPutBase64URLEncodedOptions) (*azcore.Request, error) {
	urlPath := "/string/base64UrlEncoding"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsByteArray(stringBody, azcore.Base64URLFormat)
}

// PutBase64URLEncodedHandleError handles the PutBase64URLEncoded error response.
func (client *StringClient) PutBase64URLEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutEmpty - Set string value empty ''
func (client *StringClient) PutEmpty(ctx context.Context, options *StringPutEmptyOptions) (*http.Response, error) {
	req, err := client.PutEmptyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// PutEmptyCreateRequest creates the PutEmpty request.
func (client *StringClient) PutEmptyCreateRequest(ctx context.Context, options *StringPutEmptyOptions) (*azcore.Request, error) {
	urlPath := "/string/empty"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON("")
}

// PutEmptyHandleError handles the PutEmpty error response.
func (client *StringClient) PutEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutMBCS - Set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
func (client *StringClient) PutMBCS(ctx context.Context, options *StringPutMBCSOptions) (*http.Response, error) {
	req, err := client.PutMBCSCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutMBCSHandleError(resp)
	}
	return resp.Response, nil
}

// PutMBCSCreateRequest creates the PutMBCS request.
func (client *StringClient) PutMBCSCreateRequest(ctx context.Context, options *StringPutMBCSOptions) (*azcore.Request, error) {
	urlPath := "/string/mbcs"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON("啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€")
}

// PutMBCSHandleError handles the PutMBCS error response.
func (client *StringClient) PutMBCSHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutNull - Set string value null
func (client *StringClient) PutNull(ctx context.Context, options *StringPutNullOptions) (*http.Response, error) {
	req, err := client.PutNullCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutNullHandleError(resp)
	}
	return resp.Response, nil
}

// PutNullCreateRequest creates the PutNull request.
func (client *StringClient) PutNullCreateRequest(ctx context.Context, options *StringPutNullOptions) (*azcore.Request, error) {
	urlPath := "/string/null"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.StringBody)
	}
	return req, nil
}

// PutNullHandleError handles the PutNull error response.
func (client *StringClient) PutNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutWhitespace - Set String value with leading and trailing whitespace 'Now is the time for all good men to come to the aid of their country'
func (client *StringClient) PutWhitespace(ctx context.Context, options *StringPutWhitespaceOptions) (*http.Response, error) {
	req, err := client.PutWhitespaceCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutWhitespaceHandleError(resp)
	}
	return resp.Response, nil
}

// PutWhitespaceCreateRequest creates the PutWhitespace request.
func (client *StringClient) PutWhitespaceCreateRequest(ctx context.Context, options *StringPutWhitespaceOptions) (*azcore.Request, error) {
	urlPath := "/string/whitespace"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON("    Now is the time for all good men to come to the aid of their country    ")
}

// PutWhitespaceHandleError handles the PutWhitespace error response.
func (client *StringClient) PutWhitespaceHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
