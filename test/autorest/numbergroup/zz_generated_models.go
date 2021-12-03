//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package numbergroup

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

// NumberGetBigDecimalNegativeDecimalOptions contains the optional parameters for the NumberClient.GetBigDecimalNegativeDecimal
// method.
type NumberGetBigDecimalNegativeDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberGetBigDecimalOptions contains the optional parameters for the NumberClient.GetBigDecimal method.
type NumberGetBigDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberGetBigDecimalPositiveDecimalOptions contains the optional parameters for the NumberClient.GetBigDecimalPositiveDecimal
// method.
type NumberGetBigDecimalPositiveDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberGetBigDoubleNegativeDecimalOptions contains the optional parameters for the NumberClient.GetBigDoubleNegativeDecimal
// method.
type NumberGetBigDoubleNegativeDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberGetBigDoubleOptions contains the optional parameters for the NumberClient.GetBigDouble method.
type NumberGetBigDoubleOptions struct {
	// placeholder for future optional parameters
}

// NumberGetBigDoublePositiveDecimalOptions contains the optional parameters for the NumberClient.GetBigDoublePositiveDecimal
// method.
type NumberGetBigDoublePositiveDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberGetBigFloatOptions contains the optional parameters for the NumberClient.GetBigFloat method.
type NumberGetBigFloatOptions struct {
	// placeholder for future optional parameters
}

// NumberGetInvalidDecimalOptions contains the optional parameters for the NumberClient.GetInvalidDecimal method.
type NumberGetInvalidDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberGetInvalidDoubleOptions contains the optional parameters for the NumberClient.GetInvalidDouble method.
type NumberGetInvalidDoubleOptions struct {
	// placeholder for future optional parameters
}

// NumberGetInvalidFloatOptions contains the optional parameters for the NumberClient.GetInvalidFloat method.
type NumberGetInvalidFloatOptions struct {
	// placeholder for future optional parameters
}

// NumberGetNullOptions contains the optional parameters for the NumberClient.GetNull method.
type NumberGetNullOptions struct {
	// placeholder for future optional parameters
}

// NumberGetSmallDecimalOptions contains the optional parameters for the NumberClient.GetSmallDecimal method.
type NumberGetSmallDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberGetSmallDoubleOptions contains the optional parameters for the NumberClient.GetSmallDouble method.
type NumberGetSmallDoubleOptions struct {
	// placeholder for future optional parameters
}

// NumberGetSmallFloatOptions contains the optional parameters for the NumberClient.GetSmallFloat method.
type NumberGetSmallFloatOptions struct {
	// placeholder for future optional parameters
}

// NumberPutBigDecimalNegativeDecimalOptions contains the optional parameters for the NumberClient.PutBigDecimalNegativeDecimal
// method.
type NumberPutBigDecimalNegativeDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberPutBigDecimalOptions contains the optional parameters for the NumberClient.PutBigDecimal method.
type NumberPutBigDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberPutBigDecimalPositiveDecimalOptions contains the optional parameters for the NumberClient.PutBigDecimalPositiveDecimal
// method.
type NumberPutBigDecimalPositiveDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberPutBigDoubleNegativeDecimalOptions contains the optional parameters for the NumberClient.PutBigDoubleNegativeDecimal
// method.
type NumberPutBigDoubleNegativeDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberPutBigDoubleOptions contains the optional parameters for the NumberClient.PutBigDouble method.
type NumberPutBigDoubleOptions struct {
	// placeholder for future optional parameters
}

// NumberPutBigDoublePositiveDecimalOptions contains the optional parameters for the NumberClient.PutBigDoublePositiveDecimal
// method.
type NumberPutBigDoublePositiveDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberPutBigFloatOptions contains the optional parameters for the NumberClient.PutBigFloat method.
type NumberPutBigFloatOptions struct {
	// placeholder for future optional parameters
}

// NumberPutSmallDecimalOptions contains the optional parameters for the NumberClient.PutSmallDecimal method.
type NumberPutSmallDecimalOptions struct {
	// placeholder for future optional parameters
}

// NumberPutSmallDoubleOptions contains the optional parameters for the NumberClient.PutSmallDouble method.
type NumberPutSmallDoubleOptions struct {
	// placeholder for future optional parameters
}

// NumberPutSmallFloatOptions contains the optional parameters for the NumberClient.PutSmallFloat method.
type NumberPutSmallFloatOptions struct {
	// placeholder for future optional parameters
}
