//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package twoopgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// Group1Client contains the methods for the Client.Structure.Service namespace.
// Don't use this type directly, use [TwoOperationGroupClient.NewGroup1Client] instead.
type Group1Client struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// - options - Group1ClientFourOptions contains the optional parameters for the Group1Client.Four method.
func (client *Group1Client) Four(ctx context.Context, options *Group1ClientFourOptions) (Group1ClientFourResponse, error) {
	var err error
	req, err := client.fourCreateRequest(ctx, options)
	if err != nil {
		return Group1ClientFourResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Group1ClientFourResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Group1ClientFourResponse{}, err
	}
	return Group1ClientFourResponse{}, nil
}

// fourCreateRequest creates the Four request.
func (client *Group1Client) fourCreateRequest(ctx context.Context, options *Group1ClientFourOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/four"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - Group1ClientOneOptions contains the optional parameters for the Group1Client.One method.
func (client *Group1Client) One(ctx context.Context, options *Group1ClientOneOptions) (Group1ClientOneResponse, error) {
	var err error
	req, err := client.oneCreateRequest(ctx, options)
	if err != nil {
		return Group1ClientOneResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Group1ClientOneResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Group1ClientOneResponse{}, err
	}
	return Group1ClientOneResponse{}, nil
}

// oneCreateRequest creates the One request.
func (client *Group1Client) oneCreateRequest(ctx context.Context, options *Group1ClientOneOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/one"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - Group1ClientThreeOptions contains the optional parameters for the Group1Client.Three method.
func (client *Group1Client) Three(ctx context.Context, options *Group1ClientThreeOptions) (Group1ClientThreeResponse, error) {
	var err error
	req, err := client.threeCreateRequest(ctx, options)
	if err != nil {
		return Group1ClientThreeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Group1ClientThreeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Group1ClientThreeResponse{}, err
	}
	return Group1ClientThreeResponse{}, nil
}

// threeCreateRequest creates the Three request.
func (client *Group1Client) threeCreateRequest(ctx context.Context, options *Group1ClientThreeOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/three"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
