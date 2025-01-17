//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package xmlgroup

// XMLClientGetACLsResponse contains the response from method XMLClient.GetACLs.
type XMLClientGetACLsResponse struct {
	// a collection of signed identifiers
	SignedIdentifiers []*SignedIdentifier `xml:"SignedIdentifier"`
}

// XMLClientGetBytesResponse contains the response from method XMLClient.GetBytes.
type XMLClientGetBytesResponse struct {
	ModelWithByteProperty
}

// XMLClientGetComplexTypeRefNoMetaResponse contains the response from method XMLClient.GetComplexTypeRefNoMeta.
type XMLClientGetComplexTypeRefNoMetaResponse struct {
	// I am root, and I ref a model with no meta
	RootWithRefAndNoMeta
}

// XMLClientGetComplexTypeRefWithMetaResponse contains the response from method XMLClient.GetComplexTypeRefWithMeta.
type XMLClientGetComplexTypeRefWithMetaResponse struct {
	// I am root, and I ref a model WITH meta
	RootWithRefAndMeta
}

// XMLClientGetEmptyChildElementResponse contains the response from method XMLClient.GetEmptyChildElement.
type XMLClientGetEmptyChildElementResponse struct {
	// A banana.
	Banana
}

// XMLClientGetEmptyListResponse contains the response from method XMLClient.GetEmptyList.
type XMLClientGetEmptyListResponse struct {
	// Data about a slideshow
	Slideshow
}

// XMLClientGetEmptyRootListResponse contains the response from method XMLClient.GetEmptyRootList.
type XMLClientGetEmptyRootListResponse struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`
}

// XMLClientGetEmptyWrappedListsResponse contains the response from method XMLClient.GetEmptyWrappedLists.
type XMLClientGetEmptyWrappedListsResponse struct {
	// A barrel of apples.
	AppleBarrel
}

// XMLClientGetHeadersResponse contains the response from method XMLClient.GetHeaders.
type XMLClientGetHeadersResponse struct {
	// CustomHeader contains the information returned from the Custom-Header header response.
	CustomHeader *string
}

// XMLClientGetRootListResponse contains the response from method XMLClient.GetRootList.
type XMLClientGetRootListResponse struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`
}

// XMLClientGetRootListSingleItemResponse contains the response from method XMLClient.GetRootListSingleItem.
type XMLClientGetRootListSingleItemResponse struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`
}

// XMLClientGetServicePropertiesResponse contains the response from method XMLClient.GetServiceProperties.
type XMLClientGetServicePropertiesResponse struct {
	// Storage Service Properties.
	StorageServiceProperties
}

// XMLClientGetSimpleResponse contains the response from method XMLClient.GetSimple.
type XMLClientGetSimpleResponse struct {
	// Data about a slideshow
	Slideshow
}

// XMLClientGetURIResponse contains the response from method XMLClient.GetURI.
type XMLClientGetURIResponse struct {
	ModelWithURLProperty
}

// XMLClientGetWrappedListsResponse contains the response from method XMLClient.GetWrappedLists.
type XMLClientGetWrappedListsResponse struct {
	// A barrel of apples.
	AppleBarrel
}

// XMLClientGetXMsTextResponse contains the response from method XMLClient.GetXMsText.
type XMLClientGetXMsTextResponse struct {
	// Contans property
	ObjectWithXMsTextProperty
}

// XMLClientJSONInputResponse contains the response from method XMLClient.JSONInput.
type XMLClientJSONInputResponse struct {
	// placeholder for future response values
}

// XMLClientJSONOutputResponse contains the response from method XMLClient.JSONOutput.
type XMLClientJSONOutputResponse struct {
	JSONOutput
}

// XMLClientListBlobsResponse contains the response from method XMLClient.ListBlobs.
type XMLClientListBlobsResponse struct {
	// An enumeration of blobs
	ListBlobsResponse
}

// XMLClientListContainersResponse contains the response from method XMLClient.ListContainers.
type XMLClientListContainersResponse struct {
	// An enumeration of containers
	ListContainersResponse
}

// XMLClientPutACLsResponse contains the response from method XMLClient.PutACLs.
type XMLClientPutACLsResponse struct {
	// placeholder for future response values
}

// XMLClientPutBinaryResponse contains the response from method XMLClient.PutBinary.
type XMLClientPutBinaryResponse struct {
	// placeholder for future response values
}

// XMLClientPutComplexTypeRefNoMetaResponse contains the response from method XMLClient.PutComplexTypeRefNoMeta.
type XMLClientPutComplexTypeRefNoMetaResponse struct {
	// placeholder for future response values
}

// XMLClientPutComplexTypeRefWithMetaResponse contains the response from method XMLClient.PutComplexTypeRefWithMeta.
type XMLClientPutComplexTypeRefWithMetaResponse struct {
	// placeholder for future response values
}

// XMLClientPutEmptyChildElementResponse contains the response from method XMLClient.PutEmptyChildElement.
type XMLClientPutEmptyChildElementResponse struct {
	// placeholder for future response values
}

// XMLClientPutEmptyListResponse contains the response from method XMLClient.PutEmptyList.
type XMLClientPutEmptyListResponse struct {
	// placeholder for future response values
}

// XMLClientPutEmptyRootListResponse contains the response from method XMLClient.PutEmptyRootList.
type XMLClientPutEmptyRootListResponse struct {
	// placeholder for future response values
}

// XMLClientPutEmptyWrappedListsResponse contains the response from method XMLClient.PutEmptyWrappedLists.
type XMLClientPutEmptyWrappedListsResponse struct {
	// placeholder for future response values
}

// XMLClientPutRootListResponse contains the response from method XMLClient.PutRootList.
type XMLClientPutRootListResponse struct {
	// placeholder for future response values
}

// XMLClientPutRootListSingleItemResponse contains the response from method XMLClient.PutRootListSingleItem.
type XMLClientPutRootListSingleItemResponse struct {
	// placeholder for future response values
}

// XMLClientPutServicePropertiesResponse contains the response from method XMLClient.PutServiceProperties.
type XMLClientPutServicePropertiesResponse struct {
	// placeholder for future response values
}

// XMLClientPutSimpleResponse contains the response from method XMLClient.PutSimple.
type XMLClientPutSimpleResponse struct {
	// placeholder for future response values
}

// XMLClientPutURIResponse contains the response from method XMLClient.PutURI.
type XMLClientPutURIResponse struct {
	// placeholder for future response values
}

// XMLClientPutWrappedListsResponse contains the response from method XMLClient.PutWrappedLists.
type XMLClientPutWrappedListsResponse struct {
	// placeholder for future response values
}
