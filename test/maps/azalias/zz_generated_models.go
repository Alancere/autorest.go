//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// AliasCreateOptions contains the optional parameters for the aliasClient.Create method.
type AliasCreateOptions struct {
	// The unique id that references a creator data item to be aliased.
	CreatorDataItemID *string
	GroupBy           []SomethingCount
}

// AliasListItem - Detailed information for the alias.
type AliasListItem struct {
	// READ-ONLY; The id for the alias.
	AliasID *string `json:"aliasId,omitempty" azure:"ro"`

	// READ-ONLY; The created timestamp for the alias.
	CreatedTimestamp *string `json:"createdTimestamp,omitempty" azure:"ro"`

	// READ-ONLY; The id for the creator data item that this alias references (could be null if the alias has not been assigned).
	CreatorDataItemID *string `json:"creatorDataItemId,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of the last time the alias was assigned.
	LastUpdatedTimestamp *string `json:"lastUpdatedTimestamp,omitempty" azure:"ro"`
}

// AliasListOptions contains the optional parameters for the aliasClient.List method.
type AliasListOptions struct {
	GroupBy []LogMetricsGroupBy
}

// AliasListResponse - The response model for the List API. Returns a list of all the previously created aliases.
type AliasListResponse struct {
	// READ-ONLY; A list of all the previously created aliases.
	Aliases []*AliasListItem `json:"aliases,omitempty" azure:"ro"`

	// READ-ONLY; If present, the location of the next page of data.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type AliasListResponse.
func (a AliasListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aliases", a.Aliases)
	populate(objectMap, "nextLink", a.NextLink)
	return json.Marshal(objectMap)
}

// AliasesCreateResponse - The response model for the Alias Create API for the case when the alias was successfully created.
type AliasesCreateResponse struct {
	// READ-ONLY; The id for the alias.
	AliasID *string `json:"aliasId,omitempty" azure:"ro"`

	// READ-ONLY; The created timestamp for the alias.
	CreatedTimestamp *string `json:"createdTimestamp,omitempty" azure:"ro"`

	// READ-ONLY; The id for the creator data item that this alias references (could be null if the alias has not been assigned).
	CreatorDataItemID *string `json:"creatorDataItemId,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of the last time the alias was assigned.
	LastUpdatedTimestamp *string `json:"lastUpdatedTimestamp,omitempty" azure:"ro"`
}

// ErrorResponse - An error happened.
type ErrorResponse struct {
	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// GeoJSONFeature - A valid GeoJSON Feature object type. Please refer to RFC 7946 [https://tools.ietf.org/html/rfc7946#section-3.2]
// for details.
type GeoJSONFeature struct {
	GeoJSONFeatureData
	GeoJSONObject
}

// MarshalJSON implements the json.Marshaller interface for type GeoJSONFeature.
func (g GeoJSONFeature) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	g.GeoJSONObject.marshalInternal(objectMap, GeoJSONObjectTypeGeoJSONFeature)
	g.GeoJSONFeatureData.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

type GeoJSONFeatureData struct {
	// The type of the feature. The value depends on the data model the current feature is part of. Some data models may have
	// an empty value.
	FeatureType *string `json:"featureType,omitempty"`

	// Identifier for the feature.
	ID *string `json:"id,omitempty"`

	// Properties can contain any additional metadata about the Feature. Value can be any JSON object or a JSON null value
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type GeoJSONFeatureData.
func (g GeoJSONFeatureData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	g.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GeoJSONFeatureData.
func (g *GeoJSONFeatureData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return g.unmarshalInternal(rawMsg)
}

func (g GeoJSONFeatureData) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "featureType", g.FeatureType)
	populate(objectMap, "id", g.ID)
	populate(objectMap, "properties", g.Properties)
}

func (g *GeoJSONFeatureData) unmarshalInternal(rawMsg map[string]json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "featureType":
			err = unpopulate(val, &g.FeatureType)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, &g.ID)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &g.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GeoJSONObjectClassification provides polymorphic access to related types.
// Call the interface's GetGeoJSONObject() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *GeoJSONObject, *GeoJsonFeature
type GeoJSONObjectClassification interface {
	// GetGeoJSONObject returns the GeoJSONObject content of the underlying type.
	GetGeoJSONObject() *GeoJSONObject
}

// GeoJSONObject - A valid GeoJSON object. Please refer to RFC 7946 [https://tools.ietf.org/html/rfc7946#section-3] for details.
type GeoJSONObject struct {
	// REQUIRED; Specifies the GeoJSON type. Must be one of the nine valid GeoJSON object types - Point, MultiPoint, LineString,
	// MultiLineString, Polygon, MultiPolygon, GeometryCollection, Feature and
	// FeatureCollection.
	Type *GeoJSONObjectType `json:"type,omitempty"`
}

// GetGeoJSONObject implements the GeoJSONObjectClassification interface for type GeoJSONObject.
func (g *GeoJSONObject) GetGeoJSONObject() *GeoJSONObject { return g }

// UnmarshalJSON implements the json.Unmarshaller interface for type GeoJSONObject.
func (g *GeoJSONObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return g.unmarshalInternal(rawMsg)
}

func (g GeoJSONObject) marshalInternal(objectMap map[string]interface{}, discValue GeoJSONObjectType) {
	g.Type = &discValue
	objectMap["type"] = g.Type
}

func (g *GeoJSONObject) unmarshalInternal(rawMsg map[string]json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "type":
			err = unpopulate(val, &g.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GeoJSONObjectNamedCollection - A named collection of GeoJSON object
type GeoJSONObjectNamedCollection struct {
	// Name of the collection
	CollectionName *string `json:"collectionName,omitempty"`

	// Dictionary of
	Objects map[string]GeoJSONObjectClassification `json:"objects,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type GeoJSONObjectNamedCollection.
func (g GeoJSONObjectNamedCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "collectionName", g.CollectionName)
	populate(objectMap, "objects", g.Objects)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GeoJSONObjectNamedCollection.
func (g *GeoJSONObjectNamedCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "collectionName":
			err = unpopulate(val, &g.CollectionName)
			delete(rawMsg, key)
		case "objects":
			g.Objects, err = unmarshalGeoJSONObjectClassificationMap(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ScheduleCreateOrUpdateProperties - The parameters supplied to the create or update schedule operation.
type ScheduleCreateOrUpdateProperties struct {
	// A list of all the previously created aliases.
	Aliases []*string `json:"aliases,omitempty"`

	// Gets or sets the description of the schedule.
	Description *string `json:"description,omitempty"`

	// Gets or sets the interval of the schedule.
	Interval interface{} `json:"interval,omitempty"`

	// Gets or sets the start time of the schedule.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ScheduleCreateOrUpdateProperties.
func (s ScheduleCreateOrUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aliases", s.Aliases)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "interval", &s.Interval)
	populateTimeRFC3339(objectMap, "startTime", s.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ScheduleCreateOrUpdateProperties.
func (s *ScheduleCreateOrUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aliases":
			err = unpopulate(val, &s.Aliases)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &s.Description)
			delete(rawMsg, key)
		case "interval":
			err = unpopulate(val, &s.Interval)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &s.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
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

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
