package goschema

import "fmt"

type intType struct {
	baseType
	min *int
	max *int
}

func NewIntType(description string) IntType {
	return &intType{
		baseType: baseType{
			description: description,
		},
	}
}

func (g *intType) Min(min int) IntType {
	g.min = &min
	return g
}

func (g *intType) Max(max int) IntType {
	g.max = &max
	return g
}

func (g *intType) docString(prefix, name string, docPrefix string) string {
	rnge := ""
	if g.min != nil && g.max != nil {
		rnge += fmt.Sprintf("from %v to %v", *g.min, *g.max)
	} else if g.min != nil {
		rnge += fmt.Sprintf("%v or more", *g.min)
	} else if g.max != nil {
		rnge += fmt.Sprintf("%v or less", *g.max)
	}
	return docString(prefix, name, docPrefix, g.description, "as int", rnge)
}

func (g *intType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{
		"type": "integer",
	}
	if g.description != "" {
		data["description"] = g.description
	}
	if g.min != nil {
		data["minimum"] = g.min
	}
	if g.max != nil {
		data["maximum"] = g.max
	}
	return data
}
