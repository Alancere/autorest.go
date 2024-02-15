//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package specialwordsgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// SpecialWordsClient - Scenarios to verify that reserved words can be used in service and generators will handle it appropriately.
// Current list of special words
// ```txt
// and
// as
// assert
// async
// await
// break
// class
// constructor
// continue
// def
// del
// elif
// else
// except
// exec
// finally
// for
// from
// global
// if
// import
// in
// is
// lambda
// not
// or
// pass
// raise
// return
// try
// while
// with
// yield
// ```
// Don't use this type directly, use a constructor function instead.
type SpecialWordsClient struct {
	internal *azcore.Client
}

// NewModelPropertiesClient creates a new instance of [ModelPropertiesClient].
func (client *SpecialWordsClient) NewModelPropertiesClient() *ModelPropertiesClient {
	return &ModelPropertiesClient{
		internal: client.internal,
	}
}

// NewModelsClient creates a new instance of [ModelsClient].
func (client *SpecialWordsClient) NewModelsClient() *ModelsClient {
	return &ModelsClient{
		internal: client.internal,
	}
}

// NewOperationsClient creates a new instance of [OperationsClient].
func (client *SpecialWordsClient) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: client.internal,
	}
}

// NewParametersClient creates a new instance of [ParametersClient].
func (client *SpecialWordsClient) NewParametersClient() *ParametersClient {
	return &ParametersClient{
		internal: client.internal,
	}
}
