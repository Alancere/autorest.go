//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

// BoolGetFalseOptions contains the optional parameters for the BoolClient.GetFalse method.
type BoolGetFalseOptions struct {
	// placeholder for future optional parameters
}

// BoolGetInvalidOptions contains the optional parameters for the BoolClient.GetInvalid method.
type BoolGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// BoolGetNullOptions contains the optional parameters for the BoolClient.GetNull method.
type BoolGetNullOptions struct {
	// placeholder for future optional parameters
}

// BoolGetTrueOptions contains the optional parameters for the BoolClient.GetTrue method.
type BoolGetTrueOptions struct {
	// placeholder for future optional parameters
}

// BoolPutFalseOptions contains the optional parameters for the BoolClient.PutFalse method.
type BoolPutFalseOptions struct {
	// placeholder for future optional parameters
}

// BoolPutTrueOptions contains the optional parameters for the BoolClient.PutTrue method.
type BoolPutTrueOptions struct {
	// placeholder for future optional parameters
}

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}
