// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package migroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

// MultipleInheritanceServiceClient contains the methods for the MultipleInheritanceServiceClient group.
// Don't use this type directly, use NewMultipleInheritanceServiceClient() instead.
type MultipleInheritanceServiceClient struct {
	con *Connection
}

// NewMultipleInheritanceServiceClient creates a new instance of MultipleInheritanceServiceClient with the specified values.
func NewMultipleInheritanceServiceClient(con *Connection) MultipleInheritanceServiceClient {
	return MultipleInheritanceServiceClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client MultipleInheritanceServiceClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetCat - Get a cat with name 'Whiskers' where likesMilk, meows, and hisses is true
func (client MultipleInheritanceServiceClient) GetCat(ctx context.Context, options *MultipleInheritanceServiceClientGetCatOptions) (CatResponse, error) {
	req, err := client.getCatCreateRequest(ctx, options)
	if err != nil {
		return CatResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return CatResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return CatResponse{}, client.getCatHandleError(resp)
	}
	result, err := client.getCatHandleResponse(resp)
	if err != nil {
		return CatResponse{}, err
	}
	return result, nil
}

// getCatCreateRequest creates the GetCat request.
func (client MultipleInheritanceServiceClient) getCatCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetCatOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/cat"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getCatHandleResponse handles the GetCat response.
func (client MultipleInheritanceServiceClient) getCatHandleResponse(resp *azcore.Response) (CatResponse, error) {
	result := CatResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Cat)
	return result, err
}

// getCatHandleError handles the GetCat error response.
func (client MultipleInheritanceServiceClient) getCatHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetFeline - Get a feline where meows and hisses are true
func (client MultipleInheritanceServiceClient) GetFeline(ctx context.Context, options *MultipleInheritanceServiceClientGetFelineOptions) (FelineResponse, error) {
	req, err := client.getFelineCreateRequest(ctx, options)
	if err != nil {
		return FelineResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return FelineResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FelineResponse{}, client.getFelineHandleError(resp)
	}
	result, err := client.getFelineHandleResponse(resp)
	if err != nil {
		return FelineResponse{}, err
	}
	return result, nil
}

// getFelineCreateRequest creates the GetFeline request.
func (client MultipleInheritanceServiceClient) getFelineCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetFelineOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/feline"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getFelineHandleResponse handles the GetFeline response.
func (client MultipleInheritanceServiceClient) getFelineHandleResponse(resp *azcore.Response) (FelineResponse, error) {
	result := FelineResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Feline)
	return result, err
}

// getFelineHandleError handles the GetFeline error response.
func (client MultipleInheritanceServiceClient) getFelineHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetHorse - Get a horse with name 'Fred' and isAShowHorse true
func (client MultipleInheritanceServiceClient) GetHorse(ctx context.Context, options *MultipleInheritanceServiceClientGetHorseOptions) (HorseResponse, error) {
	req, err := client.getHorseCreateRequest(ctx, options)
	if err != nil {
		return HorseResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return HorseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HorseResponse{}, client.getHorseHandleError(resp)
	}
	result, err := client.getHorseHandleResponse(resp)
	if err != nil {
		return HorseResponse{}, err
	}
	return result, nil
}

// getHorseCreateRequest creates the GetHorse request.
func (client MultipleInheritanceServiceClient) getHorseCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetHorseOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/horse"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHorseHandleResponse handles the GetHorse response.
func (client MultipleInheritanceServiceClient) getHorseHandleResponse(resp *azcore.Response) (HorseResponse, error) {
	result := HorseResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Horse)
	return result, err
}

// getHorseHandleError handles the GetHorse error response.
func (client MultipleInheritanceServiceClient) getHorseHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetKitten - Get a kitten with name 'Gatito' where likesMilk and meows is true, and hisses and eatsMiceYet is false
func (client MultipleInheritanceServiceClient) GetKitten(ctx context.Context, options *MultipleInheritanceServiceClientGetKittenOptions) (KittenResponse, error) {
	req, err := client.getKittenCreateRequest(ctx, options)
	if err != nil {
		return KittenResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return KittenResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return KittenResponse{}, client.getKittenHandleError(resp)
	}
	result, err := client.getKittenHandleResponse(resp)
	if err != nil {
		return KittenResponse{}, err
	}
	return result, nil
}

// getKittenCreateRequest creates the GetKitten request.
func (client MultipleInheritanceServiceClient) getKittenCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetKittenOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/kitten"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getKittenHandleResponse handles the GetKitten response.
func (client MultipleInheritanceServiceClient) getKittenHandleResponse(resp *azcore.Response) (KittenResponse, error) {
	result := KittenResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Kitten)
	return result, err
}

// getKittenHandleError handles the GetKitten error response.
func (client MultipleInheritanceServiceClient) getKittenHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPet - Get a pet with name 'Peanut'
func (client MultipleInheritanceServiceClient) GetPet(ctx context.Context, options *MultipleInheritanceServiceClientGetPetOptions) (PetResponse, error) {
	req, err := client.getPetCreateRequest(ctx, options)
	if err != nil {
		return PetResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return PetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PetResponse{}, client.getPetHandleError(resp)
	}
	result, err := client.getPetHandleResponse(resp)
	if err != nil {
		return PetResponse{}, err
	}
	return result, nil
}

// getPetCreateRequest creates the GetPet request.
func (client MultipleInheritanceServiceClient) getPetCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetPetOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/pet"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPetHandleResponse handles the GetPet response.
func (client MultipleInheritanceServiceClient) getPetHandleResponse(resp *azcore.Response) (PetResponse, error) {
	result := PetResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Pet)
	return result, err
}

// getPetHandleError handles the GetPet error response.
func (client MultipleInheritanceServiceClient) getPetHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutCat - Put a cat with name 'Boots' where likesMilk and hisses is false, meows is true
func (client MultipleInheritanceServiceClient) PutCat(ctx context.Context, cat Cat, options *MultipleInheritanceServiceClientPutCatOptions) (StringResponse, error) {
	req, err := client.putCatCreateRequest(ctx, cat, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.putCatHandleError(resp)
	}
	result, err := client.putCatHandleResponse(resp)
	if err != nil {
		return StringResponse{}, err
	}
	return result, nil
}

// putCatCreateRequest creates the PutCat request.
func (client MultipleInheritanceServiceClient) putCatCreateRequest(ctx context.Context, cat Cat, options *MultipleInheritanceServiceClientPutCatOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/cat"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(cat)
}

// putCatHandleResponse handles the PutCat response.
func (client MultipleInheritanceServiceClient) putCatHandleResponse(resp *azcore.Response) (StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// putCatHandleError handles the PutCat error response.
func (client MultipleInheritanceServiceClient) putCatHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// PutFeline - Put a feline who hisses and doesn't meow
func (client MultipleInheritanceServiceClient) PutFeline(ctx context.Context, feline Feline, options *MultipleInheritanceServiceClientPutFelineOptions) (StringResponse, error) {
	req, err := client.putFelineCreateRequest(ctx, feline, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.putFelineHandleError(resp)
	}
	result, err := client.putFelineHandleResponse(resp)
	if err != nil {
		return StringResponse{}, err
	}
	return result, nil
}

// putFelineCreateRequest creates the PutFeline request.
func (client MultipleInheritanceServiceClient) putFelineCreateRequest(ctx context.Context, feline Feline, options *MultipleInheritanceServiceClientPutFelineOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/feline"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(feline)
}

// putFelineHandleResponse handles the PutFeline response.
func (client MultipleInheritanceServiceClient) putFelineHandleResponse(resp *azcore.Response) (StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// putFelineHandleError handles the PutFeline error response.
func (client MultipleInheritanceServiceClient) putFelineHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// PutHorse - Put a horse with name 'General' and isAShowHorse false
func (client MultipleInheritanceServiceClient) PutHorse(ctx context.Context, horse Horse, options *MultipleInheritanceServiceClientPutHorseOptions) (StringResponse, error) {
	req, err := client.putHorseCreateRequest(ctx, horse, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.putHorseHandleError(resp)
	}
	result, err := client.putHorseHandleResponse(resp)
	if err != nil {
		return StringResponse{}, err
	}
	return result, nil
}

// putHorseCreateRequest creates the PutHorse request.
func (client MultipleInheritanceServiceClient) putHorseCreateRequest(ctx context.Context, horse Horse, options *MultipleInheritanceServiceClientPutHorseOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/horse"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(horse)
}

// putHorseHandleResponse handles the PutHorse response.
func (client MultipleInheritanceServiceClient) putHorseHandleResponse(resp *azcore.Response) (StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// putHorseHandleError handles the PutHorse error response.
func (client MultipleInheritanceServiceClient) putHorseHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// PutKitten - Put a kitten with name 'Kitty' where likesMilk and hisses is false, meows and eatsMiceYet is true
func (client MultipleInheritanceServiceClient) PutKitten(ctx context.Context, kitten Kitten, options *MultipleInheritanceServiceClientPutKittenOptions) (StringResponse, error) {
	req, err := client.putKittenCreateRequest(ctx, kitten, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.putKittenHandleError(resp)
	}
	result, err := client.putKittenHandleResponse(resp)
	if err != nil {
		return StringResponse{}, err
	}
	return result, nil
}

// putKittenCreateRequest creates the PutKitten request.
func (client MultipleInheritanceServiceClient) putKittenCreateRequest(ctx context.Context, kitten Kitten, options *MultipleInheritanceServiceClientPutKittenOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/kitten"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(kitten)
}

// putKittenHandleResponse handles the PutKitten response.
func (client MultipleInheritanceServiceClient) putKittenHandleResponse(resp *azcore.Response) (StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// putKittenHandleError handles the PutKitten error response.
func (client MultipleInheritanceServiceClient) putKittenHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// PutPet - Put a pet with name 'Butter'
func (client MultipleInheritanceServiceClient) PutPet(ctx context.Context, pet Pet, options *MultipleInheritanceServiceClientPutPetOptions) (StringResponse, error) {
	req, err := client.putPetCreateRequest(ctx, pet, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.putPetHandleError(resp)
	}
	result, err := client.putPetHandleResponse(resp)
	if err != nil {
		return StringResponse{}, err
	}
	return result, nil
}

// putPetCreateRequest creates the PutPet request.
func (client MultipleInheritanceServiceClient) putPetCreateRequest(ctx context.Context, pet Pet, options *MultipleInheritanceServiceClientPutPetOptions) (*azcore.Request, error) {
	urlPath := "/multipleInheritance/pet"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(pet)
}

// putPetHandleResponse handles the PutPet response.
func (client MultipleInheritanceServiceClient) putPetHandleResponse(resp *azcore.Response) (StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Value)
	return result, err
}

// putPetHandleError handles the PutPet error response.
func (client MultipleInheritanceServiceClient) putPetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
