// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"regexp"
	"strings"
	"time"
)

// Azure reports time in UTC but it doesn't include the 'Z' time zone suffix in some cases.
var tzOffsetRegex = regexp.MustCompile(`(?:Z|z|\+|-)(?:\d+:\d+)*"*$`)

const (
	utcDateTime        = "2006-01-02T15:04:05.999999999"
	utcDateTimeJSON    = `"` + utcDateTime + `"`
	utcDateTimeNoT     = "2006-01-02 15:04:05.999999999"
	utcDateTimeJSONNoT = `"` + utcDateTimeNoT + `"`
	dateTimeNoT        = `2006-01-02 15:04:05.999999999Z07:00`
	dateTimeJSON       = `"` + time.RFC3339Nano + `"`
	dateTimeJSONNoT    = `"` + dateTimeNoT + `"`
)

type dateTimeRFC3339 time.Time

func (t dateTimeRFC3339) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	return tt.MarshalJSON()
}

func (t dateTimeRFC3339) MarshalText() ([]byte, error) {
	tt := time.Time(t)
	return tt.MarshalText()
}

func (t *dateTimeRFC3339) UnmarshalJSON(data []byte) error {
	tzOffset := tzOffsetRegex.Match(data)
	hasT := strings.Contains(string(data), "T") || strings.Contains(string(data), "t")
	var layout string
	if tzOffset && hasT {
		layout = dateTimeJSON
	} else if tzOffset {
		layout = dateTimeJSONNoT
	} else if hasT {
		layout = utcDateTimeJSON
	} else {
		layout = utcDateTimeJSONNoT
	}
	return t.Parse(layout, string(data))
}

func (t *dateTimeRFC3339) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	tzOffset := tzOffsetRegex.Match(data)
	hasT := strings.Contains(string(data), "T") || strings.Contains(string(data), "t")
	var layout string
	if tzOffset && hasT {
		layout = time.RFC3339Nano
	} else if tzOffset {
		layout = dateTimeNoT
	} else if hasT {
		layout = utcDateTime
	} else {
		layout = utcDateTimeNoT
	}
	return t.Parse(layout, string(data))
}

func (t *dateTimeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = dateTimeRFC3339(p)
	return err
}

func (t dateTimeRFC3339) String() string {
	return time.Time(t).Format(time.RFC3339Nano)
}