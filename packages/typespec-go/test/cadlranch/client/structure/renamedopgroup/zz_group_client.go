//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package renamedopgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// GroupClient contains the methods for the Client.Structure.Service namespace.
// Don't use this type directly, use [RenamedOperationClient.NewGroupClient] instead.
type GroupClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// - options - GroupClientRenamedFourOptions contains the optional parameters for the GroupClient.RenamedFour method.
func (client *GroupClient) RenamedFour(ctx context.Context, options *GroupClientRenamedFourOptions) (GroupClientRenamedFourResponse, error) {
	var err error
	req, err := client.renamedFourCreateRequest(ctx, options)
	if err != nil {
		return GroupClientRenamedFourResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GroupClientRenamedFourResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return GroupClientRenamedFourResponse{}, err
	}
	return GroupClientRenamedFourResponse{}, nil
}

// renamedFourCreateRequest creates the RenamedFour request.
func (client *GroupClient) renamedFourCreateRequest(ctx context.Context, options *GroupClientRenamedFourOptions) (*policy.Request, error) {
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

// - options - GroupClientRenamedSixOptions contains the optional parameters for the GroupClient.RenamedSix method.
func (client *GroupClient) RenamedSix(ctx context.Context, options *GroupClientRenamedSixOptions) (GroupClientRenamedSixResponse, error) {
	var err error
	req, err := client.renamedSixCreateRequest(ctx, options)
	if err != nil {
		return GroupClientRenamedSixResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GroupClientRenamedSixResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return GroupClientRenamedSixResponse{}, err
	}
	return GroupClientRenamedSixResponse{}, nil
}

// renamedSixCreateRequest creates the RenamedSix request.
func (client *GroupClient) renamedSixCreateRequest(ctx context.Context, options *GroupClientRenamedSixOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/six"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - GroupClientRenamedTwoOptions contains the optional parameters for the GroupClient.RenamedTwo method.
func (client *GroupClient) RenamedTwo(ctx context.Context, options *GroupClientRenamedTwoOptions) (GroupClientRenamedTwoResponse, error) {
	var err error
	req, err := client.renamedTwoCreateRequest(ctx, options)
	if err != nil {
		return GroupClientRenamedTwoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GroupClientRenamedTwoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return GroupClientRenamedTwoResponse{}, err
	}
	return GroupClientRenamedTwoResponse{}, nil
}

// renamedTwoCreateRequest creates the RenamedTwo request.
func (client *GroupClient) renamedTwoCreateRequest(ctx context.Context, options *GroupClientRenamedTwoOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/two"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
