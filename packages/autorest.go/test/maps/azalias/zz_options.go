//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

// ClientCreateOptions contains the optional parameters for the Client.Create method.
type ClientCreateOptions struct {
	GroupBy []SomethingCount

	// The unique id that references the assigned data item to be aliased.
	AssignedID *float32

	// The unique id that references a creator data item to be aliased.
	CreatorID *int32
}

// ClientGetScriptOptions contains the optional parameters for the Client.GetScript method.
type ClientGetScriptOptions struct {
	// placeholder for future optional parameters
}

// ClientListOptions contains the optional parameters for the Client.NewListPager method.
type ClientListOptions struct {
	GroupBy []LogMetricsGroupBy
}

// ClientPolicyAssignmentOptions contains the optional parameters for the Client.PolicyAssignment method.
type ClientPolicyAssignmentOptions struct {
	Interval *string
}
