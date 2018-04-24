package objects

import (
	"github.com/dairehoman/libstix2/objects/properties"
)

//MarkingDefintionType -
type MarkingDefintionType struct {
	properties.CommonMarkingDefinitionPropertiesType
	DefinitionType string `json:"definition_type"`
	Definition     struct {
		Tlp string `json:"tlp"`
	} `json:"definition"`
}

//NewMarkingDefinition -
func NewMarkingDefinition(ver string, colour string) *MarkingDefintionType {
	var obj MarkingDefintionType
	obj.InitMarkingDefinitionObjectProperties("marking_definition", ver)
	obj.addType("tlp")
	obj.Definition.Tlp = colour
	return &obj
}

func (o *MarkingDefintionType) addType(s string) error {
	o.DefinitionType = s
	return nil
}
