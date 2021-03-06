// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
ObjectTypePropertyType - A property used by one or more STIX objects that
captures the STIX object type in string format.
*/
type ObjectTypePropertyType struct {
	ObjectType string `json:"type,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - ObjectTypePropertyType
//
// ----------------------------------------------------------------------

/*
SetObjectType - This method takes in a string value representing a STIX object
type and updates the type property.
*/
func (p *ObjectTypePropertyType) SetObjectType(s string) error {
	p.ObjectType = s
	return nil
}

/*
GetObjectType - This method returns the object type.
*/
func (p *ObjectTypePropertyType) GetObjectType() string {
	return p.ObjectType
}
