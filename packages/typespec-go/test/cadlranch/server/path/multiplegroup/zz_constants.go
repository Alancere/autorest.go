//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package multiplegroup

// Versions - Service versions
type Versions string

const (
	// VersionsV10 - Version 1.0
	VersionsV10 Versions = "v1.0"
)

// PossibleVersionsValues returns the possible values for the Versions const type.
func PossibleVersionsValues() []Versions {
	return []Versions{
		VersionsV10,
	}
}
