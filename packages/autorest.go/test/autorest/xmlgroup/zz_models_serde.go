// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package xmlgroup

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
	"time"
)

// MarshalXML implements the xml.Marshaller interface for type AccessPolicy.
func (a AccessPolicy) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *dateTimeRFC3339 `xml:"Expiry"`
		Start  *dateTimeRFC3339 `xml:"Start"`
	}{
		alias:  (*alias)(&a),
		Expiry: (*dateTimeRFC3339)(a.Expiry),
		Start:  (*dateTimeRFC3339)(a.Start),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type AccessPolicy.
func (a *AccessPolicy) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *dateTimeRFC3339 `xml:"Expiry"`
		Start  *dateTimeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	if aux.Expiry != nil && !(*time.Time)(aux.Expiry).IsZero() {
		a.Expiry = (*time.Time)(aux.Expiry)
	}
	if aux.Start != nil && !(*time.Time)(aux.Start).IsZero() {
		a.Start = (*time.Time)(aux.Start)
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type AppleBarrel.
func (a AppleBarrel) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias AppleBarrel
	aux := &struct {
		*alias
		BadApples  *[]*string `xml:"BadApples>Apple"`
		GoodApples *[]*string `xml:"GoodApples>Apple"`
	}{
		alias: (*alias)(&a),
	}
	if a.BadApples != nil {
		aux.BadApples = &a.BadApples
	}
	if a.GoodApples != nil {
		aux.GoodApples = &a.GoodApples
	}
	return enc.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type Banana.
func (b Banana) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "banana"
	type alias Banana
	aux := &struct {
		*alias
		Expiration *dateTimeRFC3339 `xml:"expiration"`
	}{
		alias:      (*alias)(&b),
		Expiration: (*dateTimeRFC3339)(b.Expiration),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type Banana.
func (b *Banana) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias Banana
	aux := &struct {
		*alias
		Expiration *dateTimeRFC3339 `xml:"expiration"`
	}{
		alias: (*alias)(b),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	if aux.Expiration != nil && !(*time.Time)(aux.Expiration).IsZero() {
		b.Expiration = (*time.Time)(aux.Expiration)
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type Blob.
func (b Blob) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias Blob
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(&b),
	}
	aux.Metadata = (additionalProperties)(b.Metadata)
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type Blob.
func (b *Blob) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias Blob
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(b),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	b.Metadata = (map[string]*string)(aux.Metadata)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type BlobProperties.
func (b BlobProperties) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias BlobProperties
	aux := &struct {
		*alias
		CopyCompletionTime *dateTimeRFC1123 `xml:"CopyCompletionTime"`
		DeletedTime        *dateTimeRFC1123 `xml:"DeletedTime"`
		LastModified       *dateTimeRFC1123 `xml:"Last-Modified"`
	}{
		alias:              (*alias)(&b),
		CopyCompletionTime: (*dateTimeRFC1123)(b.CopyCompletionTime),
		DeletedTime:        (*dateTimeRFC1123)(b.DeletedTime),
		LastModified:       (*dateTimeRFC1123)(b.LastModified),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type BlobProperties.
func (b *BlobProperties) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias BlobProperties
	aux := &struct {
		*alias
		CopyCompletionTime *dateTimeRFC1123 `xml:"CopyCompletionTime"`
		DeletedTime        *dateTimeRFC1123 `xml:"DeletedTime"`
		LastModified       *dateTimeRFC1123 `xml:"Last-Modified"`
	}{
		alias: (*alias)(b),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	if aux.CopyCompletionTime != nil && !(*time.Time)(aux.CopyCompletionTime).IsZero() {
		b.CopyCompletionTime = (*time.Time)(aux.CopyCompletionTime)
	}
	if aux.DeletedTime != nil && !(*time.Time)(aux.DeletedTime).IsZero() {
		b.DeletedTime = (*time.Time)(aux.DeletedTime)
	}
	if aux.LastModified != nil && !(*time.Time)(aux.LastModified).IsZero() {
		b.LastModified = (*time.Time)(aux.LastModified)
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type Blobs.
func (b Blobs) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias Blobs
	aux := &struct {
		*alias
		Blob       *[]*Blob       `xml:"Blob"`
		BlobPrefix *[]*BlobPrefix `xml:"BlobPrefix"`
	}{
		alias: (*alias)(&b),
	}
	if b.Blob != nil {
		aux.Blob = &b.Blob
	}
	if b.BlobPrefix != nil {
		aux.BlobPrefix = &b.BlobPrefix
	}
	return enc.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type Container.
func (c Container) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias Container
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(&c),
	}
	aux.Metadata = (additionalProperties)(c.Metadata)
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type Container.
func (c *Container) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias Container
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(c),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	c.Metadata = (map[string]*string)(aux.Metadata)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ContainerProperties.
func (c ContainerProperties) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias ContainerProperties
	aux := &struct {
		*alias
		LastModified *dateTimeRFC1123 `xml:"Last-Modified"`
	}{
		alias:        (*alias)(&c),
		LastModified: (*dateTimeRFC1123)(c.LastModified),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type ContainerProperties.
func (c *ContainerProperties) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias ContainerProperties
	aux := &struct {
		*alias
		LastModified *dateTimeRFC1123 `xml:"Last-Modified"`
	}{
		alias: (*alias)(c),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	if aux.LastModified != nil && !(*time.Time)(aux.LastModified).IsZero() {
		c.LastModified = (*time.Time)(aux.LastModified)
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JSONInput.
func (j JSONInput) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", j.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JSONInput.
func (j *JSONInput) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &j.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JSONOutput.
func (j JSONOutput) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", j.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JSONOutput.
func (j *JSONOutput) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &j.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ListContainersResponse.
func (l ListContainersResponse) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias ListContainersResponse
	aux := &struct {
		*alias
		Containers *[]*Container `xml:"Containers>Container"`
	}{
		alias: (*alias)(&l),
	}
	if l.Containers != nil {
		aux.Containers = &l.Containers
	}
	return enc.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type ModelWithByteProperty.
func (m ModelWithByteProperty) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias ModelWithByteProperty
	aux := &struct {
		*alias
		Bytes *string `xml:"Bytes"`
	}{
		alias: (*alias)(&m),
	}
	if m.Bytes != nil {
		encodedBytes := runtime.EncodeByteArray(m.Bytes, runtime.Base64StdFormat)
		aux.Bytes = &encodedBytes
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type ModelWithByteProperty.
func (m *ModelWithByteProperty) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias ModelWithByteProperty
	aux := &struct {
		*alias
		Bytes *string `xml:"Bytes"`
	}{
		alias: (*alias)(m),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	if aux.Bytes != nil {
		if err := runtime.DecodeByteArray(*aux.Bytes, &m.Bytes, runtime.Base64StdFormat); err != nil {
			return err
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type Slide.
func (s Slide) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias Slide
	aux := &struct {
		*alias
		Items *[]*string `xml:"item"`
	}{
		alias: (*alias)(&s),
	}
	if s.Items != nil {
		aux.Items = &s.Items
	}
	return enc.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type Slideshow.
func (s Slideshow) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "slideshow"
	type alias Slideshow
	aux := &struct {
		*alias
		Slides *[]*Slide `xml:"slide"`
	}{
		alias: (*alias)(&s),
	}
	if s.Slides != nil {
		aux.Slides = &s.Slides
	}
	return enc.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type StorageServiceProperties.
func (s StorageServiceProperties) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias StorageServiceProperties
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
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
