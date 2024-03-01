// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aztables

import (
	"strings"
	"time"
)

const (
	dateTimeRFC1123JSON = `"` + time.RFC1123 + `"`
)

type dateTimeRFC1123 time.Time

func (t dateTimeRFC1123) MarshalJSON() ([]byte, error) {
	b := []byte(time.Time(t).Format(dateTimeRFC1123JSON))
	return b, nil
}

func (t dateTimeRFC1123) MarshalText() ([]byte, error) {
	b := []byte(time.Time(t).Format(time.RFC1123))
	return b, nil
}

func (t *dateTimeRFC1123) UnmarshalJSON(data []byte) error {
	p, err := time.Parse(dateTimeRFC1123JSON, strings.ToUpper(string(data)))
	*t = dateTimeRFC1123(p)
	return err
}

func (t *dateTimeRFC1123) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	p, err := time.Parse(time.RFC1123, string(data))
	*t = dateTimeRFC1123(p)
	return err
}

func (t dateTimeRFC1123) String() string {
	return time.Time(t).Format(time.RFC1123)
}
