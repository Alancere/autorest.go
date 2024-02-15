//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package coreusagegroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// UsageClient - Test for internal decorator.
// Don't use this type directly, use a constructor function instead.
type UsageClient struct {
	internal *azcore.Client
}

// NewModelInOperationClient creates a new instance of [ModelInOperationClient].
func (client *UsageClient) NewModelInOperationClient() *ModelInOperationClient {
	return &ModelInOperationClient{
		internal: client.internal,
	}
}
