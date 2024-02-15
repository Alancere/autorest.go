//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bytesgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// BytesClient - Test for encode decorator on bytes.
// Don't use this type directly, use a constructor function instead.
type BytesClient struct {
	internal *azcore.Client
}

// NewHeaderClient creates a new instance of [HeaderClient].
func (client *BytesClient) NewHeaderClient() *HeaderClient {
	return &HeaderClient{
		internal: client.internal,
	}
}

// NewPropertyClient creates a new instance of [PropertyClient].
func (client *BytesClient) NewPropertyClient() *PropertyClient {
	return &PropertyClient{
		internal: client.internal,
	}
}

// NewQueryClient creates a new instance of [QueryClient].
func (client *BytesClient) NewQueryClient() *QueryClient {
	return &QueryClient{
		internal: client.internal,
	}
}

// NewRequestBodyClient creates a new instance of [RequestBodyClient].
func (client *BytesClient) NewRequestBodyClient() *RequestBodyClient {
	return &RequestBodyClient{
		internal: client.internal,
	}
}

// NewResponseBodyClient creates a new instance of [ResponseBodyClient].
func (client *BytesClient) NewResponseBodyClient() *ResponseBodyClient {
	return &ResponseBodyClient{
		internal: client.internal,
	}
}
