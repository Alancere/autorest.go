// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import "encoding/json"

type notFoundErrorBase struct {
	wrapped NotFoundErrorBaseClassification
}

func (n *notFoundErrorBase) UnmarshalJSON(data []byte) (err error) {
	n.wrapped, err = unmarshalNotFoundErrorBaseClassification(data)
	return
}

func unmarshalNotFoundErrorBaseClassification(body []byte) (NotFoundErrorBaseClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b NotFoundErrorBaseClassification
	switch m["whatNotFound"] {
	case "AnimalNotFound":
		b = &AnimalNotFound{}
	case "InvalidResourceLink":
		b = &LinkNotFound{}
	default:
		b = &NotFoundErrorBase{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalNotFoundErrorBaseClassificationArray(body []byte) (*[]NotFoundErrorBaseClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]NotFoundErrorBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalNotFoundErrorBaseClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

type petActionError struct {
	wrapped PetActionErrorClassification
}

func (p *petActionError) UnmarshalJSON(data []byte) (err error) {
	p.wrapped, err = unmarshalPetActionErrorClassification(data)
	return
}

func unmarshalPetActionErrorClassification(body []byte) (PetActionErrorClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b PetActionErrorClassification
	switch m["errorType"] {
	case "PetHungryOrThirstyError":
		b = &PetHungryOrThirstyError{}
	case "PetSadError":
		b = &PetSadError{}
	default:
		b = &PetActionError{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalPetActionErrorClassificationArray(body []byte) (*[]PetActionErrorClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]PetActionErrorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalPetActionErrorClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalPetSadErrorClassification(body []byte) (PetSadErrorClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b PetSadErrorClassification
	switch m["errorType"] {
	case "PetHungryOrThirstyError":
		b = &PetHungryOrThirstyError{}
	default:
		b = &PetSadError{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalPetSadErrorClassificationArray(body []byte) (*[]PetSadErrorClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]PetSadErrorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalPetSadErrorClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func strptr(s string) *string {
	return &s
}
