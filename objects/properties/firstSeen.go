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
FirstSeenPropertyType - A property used by one or more STIX objects that
captures the time that this object was first seen in STIX timestamp format,
which is an RFC3339 format.
*/
type FirstSeenPropertyType struct {
	FirstSeen string `json:"first_seen,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - FirstSeenPropertyType
//
// ----------------------------------------------------------------------

/*
SetFirstSeenToCurrentTime - This methods sets the first seen time to the
current time
*/
func (p *FirstSeenPropertyType) SetFirstSeenToCurrentTime() error {
	p.FirstSeen = timestamp.GetCurrentTime("micro")
	return nil
}

/*
SetFirstSeen -  This method takes in a timestamp in either time.Time or string
format and updates the first seen property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (p *FirstSeenPropertyType) SetFirstSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	p.FirstSeen = ts
	return nil
}
