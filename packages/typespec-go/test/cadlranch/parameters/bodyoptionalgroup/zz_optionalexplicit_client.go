//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bodyoptionalgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OptionalExplicitClient contains the methods for the Parameters.BodyOptionality namespace.
// Don't use this type directly, use [BodyOptionalityClient.NewOptionalExplicitClient] instead.
type OptionalExplicitClient struct {
	internal *azcore.Client
}

// - options - OptionalExplicitClientOmitOptions contains the optional parameters for the OptionalExplicitClient.Omit method.
func (client *OptionalExplicitClient) Omit(ctx context.Context, options *OptionalExplicitClientOmitOptions) (OptionalExplicitClientOmitResponse, error) {
	var err error
	req, err := client.omitCreateRequest(ctx, options)
	if err != nil {
		return OptionalExplicitClientOmitResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalExplicitClientOmitResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalExplicitClientOmitResponse{}, err
	}
	return OptionalExplicitClientOmitResponse{}, nil
}

// omitCreateRequest creates the Omit request.
func (client *OptionalExplicitClient) omitCreateRequest(ctx context.Context, options *OptionalExplicitClientOmitOptions) (*policy.Request, error) {
	urlPath := "/parameters/body-optionality/optional-explicit/omit"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// - options - OptionalExplicitClientSetOptions contains the optional parameters for the OptionalExplicitClient.Set method.
func (client *OptionalExplicitClient) Set(ctx context.Context, options *OptionalExplicitClientSetOptions) (OptionalExplicitClientSetResponse, error) {
	var err error
	req, err := client.setCreateRequest(ctx, options)
	if err != nil {
		return OptionalExplicitClientSetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalExplicitClientSetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalExplicitClientSetResponse{}, err
	}
	return OptionalExplicitClientSetResponse{}, nil
}

// setCreateRequest creates the Set request.
func (client *OptionalExplicitClient) setCreateRequest(ctx context.Context, options *OptionalExplicitClientSetOptions) (*policy.Request, error) {
	urlPath := "/parameters/body-optionality/optional-explicit/set"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}
