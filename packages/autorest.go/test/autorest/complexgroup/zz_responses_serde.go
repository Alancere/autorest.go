// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

// UnmarshalJSON implements the json.Unmarshaller interface for type FlattencomplexClientGetValidResponse.
func (f *FlattencomplexClientGetValidResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalMyBaseTypeClassification(data)
	if err != nil {
		return err
	}
	f.MyBaseTypeClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolymorphicrecursiveClientGetValidResponse.
func (p *PolymorphicrecursiveClientGetValidResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalFishClassification(data)
	if err != nil {
		return err
	}
	p.FishClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolymorphismClientGetComplicatedResponse.
func (p *PolymorphismClientGetComplicatedResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSalmonClassification(data)
	if err != nil {
		return err
	}
	p.SalmonClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolymorphismClientGetDotSyntaxResponse.
func (p *PolymorphismClientGetDotSyntaxResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDotFishClassification(data)
	if err != nil {
		return err
	}
	p.DotFishClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolymorphismClientGetValidResponse.
func (p *PolymorphismClientGetValidResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalFishClassification(data)
	if err != nil {
		return err
	}
	p.FishClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolymorphismClientPutMissingDiscriminatorResponse.
func (p *PolymorphismClientPutMissingDiscriminatorResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSalmonClassification(data)
	if err != nil {
		return err
	}
	p.SalmonClassification = res
	return nil
}