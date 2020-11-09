// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimerfc1123group

import (
	"strings"
	"time"
)

const (
	rfc1123JSON = `"` + time.RFC1123 + `"`
)

type timeRFC1123 time.Time

func (t timeRFC1123) MarshalJSON() ([]byte, error) {
	b := []byte(time.Time(t).Format(rfc1123JSON))
	return b, nil
}

func (t timeRFC1123) MarshalText() ([]byte, error) {
	b := []byte(time.Time(t).Format(time.RFC1123))
	return b, nil
}

func (t *timeRFC1123) UnmarshalJSON(data []byte) error {
	p, err := time.Parse(rfc1123JSON, strings.ToUpper(string(data)))
	*t = timeRFC1123(p)
	return err
}

func (t *timeRFC1123) UnmarshalText(data []byte) error {
	p, err := time.Parse(time.RFC1123, string(data))
	*t = timeRFC1123(p)
	return err
}
