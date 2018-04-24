// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"github.com/dairehoman/libstix2/common/timestamp"
)

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
CreatedPropertyType - A property used by one or more STIX objects that
captures the time that this object was created in STIX timestamp format,
which is an RFC3339 format.
*/
type CreatedPropertyType struct {
	Created string `json:"created,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - CreatedPropertyType
//
// ----------------------------------------------------------------------

/*
SetCreatedToCurrentTime - This methods sets the object created time to the
current time
*/
func (p *CreatedPropertyType) SetCreatedToCurrentTime() error {
	p.Created = timestamp.GetCurrentTime("milli")
	return nil
}

/*
SetCreated - This method takes in a timestamp in either time.Time or string
format and updates the created property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (p *CreatedPropertyType) SetCreated(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	p.Created = ts
	return nil
}

/*
GetCreated - This method will return the created timestamp as a string.
*/
func (p *CreatedPropertyType) GetCreated() string {
	return p.Created
}
