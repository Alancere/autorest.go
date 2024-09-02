// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package multipartgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"

type Address struct {
	// REQUIRED
	City *string
}

type BinaryArrayPartsRequest struct {
	// REQUIRED
	ID string

	// REQUIRED
	Pictures []streaming.MultipartContent
}

type ComplexHTTPPartsModelRequest struct {
	// REQUIRED
	Address Address

	// REQUIRED
	ID string

	// REQUIRED
	Pictures []streaming.MultipartContent

	// REQUIRED
	PreviousAddresses []Address

	// REQUIRED
	ProfileImage streaming.MultipartContent
}

type ComplexPartsRequest struct {
	// REQUIRED
	Address Address

	// REQUIRED
	ID string

	// REQUIRED
	Pictures []streaming.MultipartContent

	// REQUIRED
	ProfileImage streaming.MultipartContent
}

type FileOptionalContentType struct {
	// REQUIRED
	Contents []byte

	// REQUIRED
	Filename    *string
	ContentType *string
}

type FileRequiredMetaData struct {
	// REQUIRED
	ContentType *string

	// REQUIRED
	Contents []byte

	// REQUIRED
	Filename *string
}

type FileSpecificContentType struct {
	// CONSTANT; undefinedField has constant value "image/jpg", any specified value is ignored.
	ContentType *string

	// REQUIRED
	Contents []byte

	// REQUIRED
	Filename *string
}

type FileWithHTTPPartOptionalContentTypeRequest struct {
	// REQUIRED
	ProfileImage streaming.MultipartContent
}

type FileWithHTTPPartRequiredContentTypeRequest struct {
	// REQUIRED
	ProfileImage streaming.MultipartContent
}

type FileWithHTTPPartSpecificContentTypeRequest struct {
	// REQUIRED
	ProfileImage streaming.MultipartContent
}

type JSONPartRequest struct {
	// REQUIRED
	Address Address

	// REQUIRED
	ProfileImage streaming.MultipartContent
}

type MultiBinaryPartsRequest struct {
	// REQUIRED
	ProfileImage streaming.MultipartContent
	Picture      *streaming.MultipartContent
}

type MultiPartRequest struct {
	// REQUIRED
	ID string

	// REQUIRED
	ProfileImage streaming.MultipartContent
}
