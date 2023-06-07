//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aztables

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// MarshalXML implements the xml.Marshaller interface for type AccessPolicy.
func (a AccessPolicy) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias:  (*alias)(&a),
		Expiry: (*timeRFC3339)(a.Expiry),
		Start:  (*timeRFC3339)(a.Start),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type AccessPolicy.
func (a *AccessPolicy) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	a.Expiry = (*time.Time)(aux.Expiry)
	a.Start = (*time.Time)(aux.Start)
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EntityQueryResponse.
func (e EntityQueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "odata.metadata", e.ODataMetadata)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EntityQueryResponse.
func (e *EntityQueryResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "odata.metadata":
			err = unpopulate(val, "ODataMetadata", &e.ODataMetadata)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &e.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type GeoReplication.
func (g GeoReplication) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias:        (*alias)(&g),
		LastSyncTime: (*timeRFC1123)(g.LastSyncTime),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type GeoReplication.
func (g *GeoReplication) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias: (*alias)(g),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	g.LastSyncTime = (*time.Time)(aux.LastSyncTime)
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "TableName", p.TableName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "TableName":
			err = unpopulate(val, "TableName", &p.TableName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QueryResponse.
func (q QueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "odata.metadata", q.ODataMetadata)
	populate(objectMap, "value", q.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryResponse.
func (q *QueryResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", q, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "odata.metadata":
			err = unpopulate(val, "ODataMetadata", &q.ODataMetadata)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &q.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", q, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Response.
func (r Response) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "odata.editLink", r.ODataEditLink)
	populate(objectMap, "odata.id", r.ODataID)
	populate(objectMap, "odata.metadata", r.ODataMetadata)
	populate(objectMap, "odata.type", r.ODataType)
	populate(objectMap, "TableName", r.TableName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Response.
func (r *Response) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "odata.editLink":
			err = unpopulate(val, "ODataEditLink", &r.ODataEditLink)
			delete(rawMsg, key)
		case "odata.id":
			err = unpopulate(val, "ODataID", &r.ODataID)
			delete(rawMsg, key)
		case "odata.metadata":
			err = unpopulate(val, "ODataMetadata", &r.ODataMetadata)
			delete(rawMsg, key)
		case "odata.type":
			err = unpopulate(val, "ODataType", &r.ODataType)
			delete(rawMsg, key)
		case "TableName":
			err = unpopulate(val, "TableName", &r.TableName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResponseProperties.
func (r ResponseProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "odata.editLink", r.ODataEditLink)
	populate(objectMap, "odata.id", r.ODataID)
	populate(objectMap, "odata.type", r.ODataType)
	populate(objectMap, "TableName", r.TableName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResponseProperties.
func (r *ResponseProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "odata.editLink":
			err = unpopulate(val, "ODataEditLink", &r.ODataEditLink)
			delete(rawMsg, key)
		case "odata.id":
			err = unpopulate(val, "ODataID", &r.ODataID)
			delete(rawMsg, key)
		case "odata.type":
			err = unpopulate(val, "ODataType", &r.ODataType)
			delete(rawMsg, key)
		case "TableName":
			err = unpopulate(val, "TableName", &r.TableName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceError.
func (s ServiceError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "Message", s.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceError.
func (s *ServiceError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "Message":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ServiceProperties.
func (s ServiceProperties) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "StorageServiceProperties"
	type alias ServiceProperties
	aux := &struct {
		*alias
		Cors *[]*CorsRule `xml:"Cors>CorsRule"`
	}{
		alias: (*alias)(&s),
	}
	if s.Cors != nil {
		aux.Cors = &s.Cors
	}
	return enc.EncodeElement(aux, start)
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}