// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"github.com/dairehoman/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Message Type
//
// ----------------------------------------------------------------------

/*
InfrastructureType - This type implements the STIX 2 Infrastructure SDO and defines
all of the properties methods needed to create and work with the STIX Infrastructure
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.
*/
type InfrastructureType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewInfrastructure - This function will create a new STIX Infrastructure object
and return it as a pointer.
*/
func NewInfrastructure(ver string) *InfrastructureType {
	var obj InfrastructureType
	obj.InitObjectProperties("infrastructure", ver)
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - InfrastructureType
//
// ----------------------------------------------------------------------
