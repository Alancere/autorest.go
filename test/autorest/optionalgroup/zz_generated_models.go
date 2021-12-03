//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io"
	"reflect"
)

type ArrayOptionalWrapper struct {
	Value []*string `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ArrayOptionalWrapper.
func (a ArrayOptionalWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

type ArrayWrapper struct {
	// REQUIRED
	Value []*string `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ArrayWrapper.
func (a ArrayWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

type ClassOptionalWrapper struct {
	Value *Product `json:"value,omitempty"`
}

type ClassWrapper struct {
	// REQUIRED
	Value *Product `json:"value,omitempty"`
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

// ExplicitPostOptionalArrayHeaderOptions contains the optional parameters for the ExplicitClient.PostOptionalArrayHeader
// method.
type ExplicitPostOptionalArrayHeaderOptions struct {
	HeaderParameter []string
}

// ExplicitPostOptionalArrayParameterOptions contains the optional parameters for the ExplicitClient.PostOptionalArrayParameter
// method.
type ExplicitPostOptionalArrayParameterOptions struct {
	BodyParameter []*string
}

// ExplicitPostOptionalArrayPropertyOptions contains the optional parameters for the ExplicitClient.PostOptionalArrayProperty
// method.
type ExplicitPostOptionalArrayPropertyOptions struct {
	BodyParameter *ArrayOptionalWrapper
}

// ExplicitPostOptionalClassParameterOptions contains the optional parameters for the ExplicitClient.PostOptionalClassParameter
// method.
type ExplicitPostOptionalClassParameterOptions struct {
	BodyParameter *Product
}

// ExplicitPostOptionalClassPropertyOptions contains the optional parameters for the ExplicitClient.PostOptionalClassProperty
// method.
type ExplicitPostOptionalClassPropertyOptions struct {
	BodyParameter *ClassOptionalWrapper
}

// ExplicitPostOptionalIntegerHeaderOptions contains the optional parameters for the ExplicitClient.PostOptionalIntegerHeader
// method.
type ExplicitPostOptionalIntegerHeaderOptions struct {
	HeaderParameter *int32
}

// ExplicitPostOptionalIntegerParameterOptions contains the optional parameters for the ExplicitClient.PostOptionalIntegerParameter
// method.
type ExplicitPostOptionalIntegerParameterOptions struct {
	BodyParameter *int32
}

// ExplicitPostOptionalIntegerPropertyOptions contains the optional parameters for the ExplicitClient.PostOptionalIntegerProperty
// method.
type ExplicitPostOptionalIntegerPropertyOptions struct {
	BodyParameter *IntOptionalWrapper
}

// ExplicitPostOptionalStringHeaderOptions contains the optional parameters for the ExplicitClient.PostOptionalStringHeader
// method.
type ExplicitPostOptionalStringHeaderOptions struct {
	BodyParameter *string
}

// ExplicitPostOptionalStringParameterOptions contains the optional parameters for the ExplicitClient.PostOptionalStringParameter
// method.
type ExplicitPostOptionalStringParameterOptions struct {
	BodyParameter *string
}

// ExplicitPostOptionalStringPropertyOptions contains the optional parameters for the ExplicitClient.PostOptionalStringProperty
// method.
type ExplicitPostOptionalStringPropertyOptions struct {
	BodyParameter *StringOptionalWrapper
}

// ExplicitPostRequiredArrayHeaderOptions contains the optional parameters for the ExplicitClient.PostRequiredArrayHeader
// method.
type ExplicitPostRequiredArrayHeaderOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredArrayParameterOptions contains the optional parameters for the ExplicitClient.PostRequiredArrayParameter
// method.
type ExplicitPostRequiredArrayParameterOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredArrayPropertyOptions contains the optional parameters for the ExplicitClient.PostRequiredArrayProperty
// method.
type ExplicitPostRequiredArrayPropertyOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredClassParameterOptions contains the optional parameters for the ExplicitClient.PostRequiredClassParameter
// method.
type ExplicitPostRequiredClassParameterOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredClassPropertyOptions contains the optional parameters for the ExplicitClient.PostRequiredClassProperty
// method.
type ExplicitPostRequiredClassPropertyOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredIntegerHeaderOptions contains the optional parameters for the ExplicitClient.PostRequiredIntegerHeader
// method.
type ExplicitPostRequiredIntegerHeaderOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredIntegerParameterOptions contains the optional parameters for the ExplicitClient.PostRequiredIntegerParameter
// method.
type ExplicitPostRequiredIntegerParameterOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredIntegerPropertyOptions contains the optional parameters for the ExplicitClient.PostRequiredIntegerProperty
// method.
type ExplicitPostRequiredIntegerPropertyOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredStringHeaderOptions contains the optional parameters for the ExplicitClient.PostRequiredStringHeader
// method.
type ExplicitPostRequiredStringHeaderOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredStringParameterOptions contains the optional parameters for the ExplicitClient.PostRequiredStringParameter
// method.
type ExplicitPostRequiredStringParameterOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPostRequiredStringPropertyOptions contains the optional parameters for the ExplicitClient.PostRequiredStringProperty
// method.
type ExplicitPostRequiredStringPropertyOptions struct {
	// placeholder for future optional parameters
}

// ExplicitPutOptionalBinaryBodyOptions contains the optional parameters for the ExplicitClient.PutOptionalBinaryBody method.
type ExplicitPutOptionalBinaryBodyOptions struct {
	BodyParameter io.ReadSeekCloser
}

// ExplicitPutRequiredBinaryBodyOptions contains the optional parameters for the ExplicitClient.PutRequiredBinaryBody method.
type ExplicitPutRequiredBinaryBodyOptions struct {
	// placeholder for future optional parameters
}

// ImplicitGetOptionalGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetOptionalGlobalQuery method.
type ImplicitGetOptionalGlobalQueryOptions struct {
}

// ImplicitGetRequiredGlobalPathOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalPath method.
type ImplicitGetRequiredGlobalPathOptions struct {
	// placeholder for future optional parameters
}

// ImplicitGetRequiredGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalQuery method.
type ImplicitGetRequiredGlobalQueryOptions struct {
	// placeholder for future optional parameters
}

// ImplicitGetRequiredPathOptions contains the optional parameters for the ImplicitClient.GetRequiredPath method.
type ImplicitGetRequiredPathOptions struct {
	// placeholder for future optional parameters
}

// ImplicitPutOptionalBinaryBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBinaryBody method.
type ImplicitPutOptionalBinaryBodyOptions struct {
	BodyParameter io.ReadSeekCloser
}

// ImplicitPutOptionalBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBody method.
type ImplicitPutOptionalBodyOptions struct {
	BodyParameter *string
}

// ImplicitPutOptionalHeaderOptions contains the optional parameters for the ImplicitClient.PutOptionalHeader method.
type ImplicitPutOptionalHeaderOptions struct {
	QueryParameter *string
}

// ImplicitPutOptionalQueryOptions contains the optional parameters for the ImplicitClient.PutOptionalQuery method.
type ImplicitPutOptionalQueryOptions struct {
	QueryParameter *string
}

type IntOptionalWrapper struct {
	Value *int32 `json:"value,omitempty"`
}

type IntWrapper struct {
	// REQUIRED
	Value *int32 `json:"value,omitempty"`
}

type Product struct {
	// REQUIRED
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type StringOptionalWrapper struct {
	Value *string `json:"value,omitempty"`
}

type StringWrapper struct {
	// REQUIRED
	Value *string `json:"value,omitempty"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
