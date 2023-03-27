//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package errorsgroup

type Pet struct {
	AniType *string

	// READ-ONLY; Gets the Pet by id.
	Name *string
}

type PetAction struct {
	// action feedback
	ActionResponse *string
}

// PetClientDoSomethingOptions contains the optional parameters for the PetClient.DoSomething method.
type PetClientDoSomethingOptions struct {
	// placeholder for future optional parameters
}

// PetClientGetPetByIDOptions contains the optional parameters for the PetClient.GetPetByID method.
type PetClientGetPetByIDOptions struct {
	// placeholder for future optional parameters
}

// PetClientHasModelsParamOptions contains the optional parameters for the PetClient.HasModelsParam method.
type PetClientHasModelsParamOptions struct {
	// Make sure model deserialization doesn't conflict with this param name, which has input name 'models'. Use client default
	// value in call
	Models *string
}
