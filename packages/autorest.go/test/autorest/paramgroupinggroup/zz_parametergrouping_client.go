//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paramgroupinggroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ParameterGroupingClient contains the methods for the ParameterGrouping group.
// Don't use this type directly, use a constructor function instead.
type ParameterGroupingClient struct {
	internal *azcore.Client
}

// PostMultiParamGroups - Post parameters from multiple different parameter groups
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - FirstParameterGroup - FirstParameterGroup contains a group of parameters for the ParameterGroupingClient.PostMultiParamGroups
//     method.
//   - ParameterGroupingClientPostMultiParamGroupsSecondParamGroup - ParameterGroupingClientPostMultiParamGroupsSecondParamGroup
//     contains a group of parameters for the ParameterGroupingClient.PostMultiParamGroups method.
//   - options - ParameterGroupingClientPostMultiParamGroupsOptions contains the optional parameters for the ParameterGroupingClient.PostMultiParamGroups
//     method.
func (client *ParameterGroupingClient) PostMultiParamGroups(ctx context.Context, firstParameterGroup *FirstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup *ParameterGroupingClientPostMultiParamGroupsSecondParamGroup, options *ParameterGroupingClientPostMultiParamGroupsOptions) (ParameterGroupingClientPostMultiParamGroupsResponse, error) {
	var err error
	const operationName = "ParameterGroupingClient.PostMultiParamGroups"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postMultiParamGroupsCreateRequest(ctx, firstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup, options)
	if err != nil {
		return ParameterGroupingClientPostMultiParamGroupsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ParameterGroupingClientPostMultiParamGroupsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ParameterGroupingClientPostMultiParamGroupsResponse{}, err
	}
	return ParameterGroupingClientPostMultiParamGroupsResponse{}, nil
}

// postMultiParamGroupsCreateRequest creates the PostMultiParamGroups request.
func (client *ParameterGroupingClient) postMultiParamGroupsCreateRequest(ctx context.Context, firstParameterGroup *FirstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup *ParameterGroupingClientPostMultiParamGroupsSecondParamGroup, options *ParameterGroupingClientPostMultiParamGroupsOptions) (*policy.Request, error) {
	urlPath := "/parameterGrouping/postMultipleParameterGroups"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if firstParameterGroup != nil && firstParameterGroup.QueryOne != nil {
		reqQP.Set("query-one", strconv.FormatInt(int64(*firstParameterGroup.QueryOne), 10))
	}
	if parameterGroupingClientPostMultiParamGroupsSecondParamGroup != nil && parameterGroupingClientPostMultiParamGroupsSecondParamGroup.QueryTwo != nil {
		reqQP.Set("query-two", strconv.FormatInt(int64(*parameterGroupingClientPostMultiParamGroupsSecondParamGroup.QueryTwo), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if firstParameterGroup != nil && firstParameterGroup.HeaderOne != nil {
		req.Raw().Header["header-one"] = []string{*firstParameterGroup.HeaderOne}
	}
	if parameterGroupingClientPostMultiParamGroupsSecondParamGroup != nil && parameterGroupingClientPostMultiParamGroupsSecondParamGroup.HeaderTwo != nil {
		req.Raw().Header["header-two"] = []string{*parameterGroupingClientPostMultiParamGroupsSecondParamGroup.HeaderTwo}
	}
	return req, nil
}

// PostOptional - Post a bunch of optional parameters grouped
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - ParameterGroupingClientPostOptionalParameters - ParameterGroupingClientPostOptionalParameters contains a group of parameters
//     for the ParameterGroupingClient.PostOptional method.
//   - options - ParameterGroupingClientPostOptionalOptions contains the optional parameters for the ParameterGroupingClient.PostOptional
//     method.
func (client *ParameterGroupingClient) PostOptional(ctx context.Context, parameterGroupingClientPostOptionalParameters *ParameterGroupingClientPostOptionalParameters, options *ParameterGroupingClientPostOptionalOptions) (ParameterGroupingClientPostOptionalResponse, error) {
	var err error
	const operationName = "ParameterGroupingClient.PostOptional"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postOptionalCreateRequest(ctx, parameterGroupingClientPostOptionalParameters, options)
	if err != nil {
		return ParameterGroupingClientPostOptionalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ParameterGroupingClientPostOptionalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ParameterGroupingClientPostOptionalResponse{}, err
	}
	return ParameterGroupingClientPostOptionalResponse{}, nil
}

// postOptionalCreateRequest creates the PostOptional request.
func (client *ParameterGroupingClient) postOptionalCreateRequest(ctx context.Context, parameterGroupingClientPostOptionalParameters *ParameterGroupingClientPostOptionalParameters, options *ParameterGroupingClientPostOptionalOptions) (*policy.Request, error) {
	urlPath := "/parameterGrouping/postOptional"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if parameterGroupingClientPostOptionalParameters != nil && parameterGroupingClientPostOptionalParameters.Query != nil {
		reqQP.Set("query", strconv.FormatInt(int64(*parameterGroupingClientPostOptionalParameters.Query), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if parameterGroupingClientPostOptionalParameters != nil && parameterGroupingClientPostOptionalParameters.CustomHeader != nil {
		req.Raw().Header["customHeader"] = []string{*parameterGroupingClientPostOptionalParameters.CustomHeader}
	}
	return req, nil
}

// PostRequired - Post a bunch of required parameters grouped
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - ParameterGroupingClientPostRequiredParameters - ParameterGroupingClientPostRequiredParameters contains a group of parameters
//     for the ParameterGroupingClient.PostRequired method.
//   - options - ParameterGroupingClientPostRequiredOptions contains the optional parameters for the ParameterGroupingClient.PostRequired
//     method.
func (client *ParameterGroupingClient) PostRequired(ctx context.Context, parameterGroupingClientPostRequiredParameters ParameterGroupingClientPostRequiredParameters, options *ParameterGroupingClientPostRequiredOptions) (ParameterGroupingClientPostRequiredResponse, error) {
	var err error
	const operationName = "ParameterGroupingClient.PostRequired"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postRequiredCreateRequest(ctx, parameterGroupingClientPostRequiredParameters, options)
	if err != nil {
		return ParameterGroupingClientPostRequiredResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ParameterGroupingClientPostRequiredResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ParameterGroupingClientPostRequiredResponse{}, err
	}
	return ParameterGroupingClientPostRequiredResponse{}, nil
}

// postRequiredCreateRequest creates the PostRequired request.
func (client *ParameterGroupingClient) postRequiredCreateRequest(ctx context.Context, parameterGroupingClientPostRequiredParameters ParameterGroupingClientPostRequiredParameters, options *ParameterGroupingClientPostRequiredOptions) (*policy.Request, error) {
	urlPath := "/parameterGrouping/postRequired/{path}"
	if parameterGroupingClientPostRequiredParameters.Path == "" {
		return nil, errors.New("parameter parameterGroupingClientPostRequiredParameters.Path cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{path}", url.PathEscape(parameterGroupingClientPostRequiredParameters.Path))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if parameterGroupingClientPostRequiredParameters.Query != nil {
		reqQP.Set("query", strconv.FormatInt(int64(*parameterGroupingClientPostRequiredParameters.Query), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if parameterGroupingClientPostRequiredParameters.CustomHeader != nil {
		req.Raw().Header["customHeader"] = []string{*parameterGroupingClientPostRequiredParameters.CustomHeader}
	}
	if err := runtime.MarshalAsJSON(req, parameterGroupingClientPostRequiredParameters.Body); err != nil {
		return nil, err
	}
	return req, nil
}

// PostReservedWords - Post a grouped parameters with reserved words
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - ParameterGroupingClientPostReservedWordsParameters - ParameterGroupingClientPostReservedWordsParameters contains a group
//     of parameters for the ParameterGroupingClient.PostReservedWords method.
//   - options - ParameterGroupingClientPostReservedWordsOptions contains the optional parameters for the ParameterGroupingClient.PostReservedWords
//     method.
func (client *ParameterGroupingClient) PostReservedWords(ctx context.Context, parameterGroupingClientPostReservedWordsParameters *ParameterGroupingClientPostReservedWordsParameters, options *ParameterGroupingClientPostReservedWordsOptions) (ParameterGroupingClientPostReservedWordsResponse, error) {
	var err error
	const operationName = "ParameterGroupingClient.PostReservedWords"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postReservedWordsCreateRequest(ctx, parameterGroupingClientPostReservedWordsParameters, options)
	if err != nil {
		return ParameterGroupingClientPostReservedWordsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ParameterGroupingClientPostReservedWordsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ParameterGroupingClientPostReservedWordsResponse{}, err
	}
	return ParameterGroupingClientPostReservedWordsResponse{}, nil
}

// postReservedWordsCreateRequest creates the PostReservedWords request.
func (client *ParameterGroupingClient) postReservedWordsCreateRequest(ctx context.Context, parameterGroupingClientPostReservedWordsParameters *ParameterGroupingClientPostReservedWordsParameters, options *ParameterGroupingClientPostReservedWordsOptions) (*policy.Request, error) {
	urlPath := "/parameterGrouping/postReservedWords"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if parameterGroupingClientPostReservedWordsParameters != nil && parameterGroupingClientPostReservedWordsParameters.Accept != nil {
		reqQP.Set("accept", *parameterGroupingClientPostReservedWordsParameters.Accept)
	}
	if parameterGroupingClientPostReservedWordsParameters != nil && parameterGroupingClientPostReservedWordsParameters.From != nil {
		reqQP.Set("from", *parameterGroupingClientPostReservedWordsParameters.From)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - FirstParameterGroup - FirstParameterGroup contains a group of parameters for the ParameterGroupingClient.PostMultiParamGroups
//     method.
//   - options - ParameterGroupingClientPostSharedParameterGroupObjectOptions contains the optional parameters for the ParameterGroupingClient.PostSharedParameterGroupObject
//     method.
func (client *ParameterGroupingClient) PostSharedParameterGroupObject(ctx context.Context, firstParameterGroup *FirstParameterGroup, options *ParameterGroupingClientPostSharedParameterGroupObjectOptions) (ParameterGroupingClientPostSharedParameterGroupObjectResponse, error) {
	var err error
	const operationName = "ParameterGroupingClient.PostSharedParameterGroupObject"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postSharedParameterGroupObjectCreateRequest(ctx, firstParameterGroup, options)
	if err != nil {
		return ParameterGroupingClientPostSharedParameterGroupObjectResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ParameterGroupingClientPostSharedParameterGroupObjectResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ParameterGroupingClientPostSharedParameterGroupObjectResponse{}, err
	}
	return ParameterGroupingClientPostSharedParameterGroupObjectResponse{}, nil
}

// postSharedParameterGroupObjectCreateRequest creates the PostSharedParameterGroupObject request.
func (client *ParameterGroupingClient) postSharedParameterGroupObjectCreateRequest(ctx context.Context, firstParameterGroup *FirstParameterGroup, options *ParameterGroupingClientPostSharedParameterGroupObjectOptions) (*policy.Request, error) {
	urlPath := "/parameterGrouping/sharedParameterGroupObject"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if firstParameterGroup != nil && firstParameterGroup.QueryOne != nil {
		reqQP.Set("query-one", strconv.FormatInt(int64(*firstParameterGroup.QueryOne), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if firstParameterGroup != nil && firstParameterGroup.HeaderOne != nil {
		req.Raw().Header["header-one"] = []string{*firstParameterGroup.HeaderOne}
	}
	return req, nil
}
