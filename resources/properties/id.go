// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
IDPropertyType - A property used by one or more TAXII resources.
*/
type IDPropertyType struct {
	ID string `json:"id,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - IdPropertyType
//
// ----------------------------------------------------------------------

/*
CreateTAXIIUUID - This method does not take in any parameters. It is used to create
a new ID based on the approved TAXII UUIDv4 format.
*/
func (p *IDPropertyType) CreateTAXIIUUID() (string, error) {
	id := uuid.New()
	return id, nil
}

/*
SetNewID - This method does not take in any parameters. It is used to create
a new ID based on the approved TAXII UUIDv4 format and assigns it to the ID
property.
*/
func (p *IDPropertyType) SetNewID() error {
	p.ID, _ = p.CreateTAXIIUUID()
	return nil
}

/*
SetID - This method takes in a string value representing a TAXII id and
updates the ID property.
*/
func (p *IDPropertyType) SetID(s string) error {
	// TODO add check to validate input value
	p.ID = s
	return nil
}

/*
GetID - This method returns the id.
*/
func (p *IDPropertyType) GetID() string {
	return p.ID
}
