package goschema

import (
	"strings"
)

type mapType struct {
	baseType
	subtype Type
}

func NewMapType(description string, subType func(m MapType)) MapType {
	gom := &mapType{
		baseType: baseType{
			description: description,
		},
	}
	subType(gom)
	return gom
}

func (g *mapType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{}

	data["title"] = g.description
	data["type"] = "object"
	if g.subtype != nil {
		data["additionalProperties"] = g.subtype.asJSONSchema()
	} else {
		data["additionalProperties"] = false
	}

	return data
}

func (g *mapType) docString(prefix string, name string, docPrefix string) string {
	if name != "" { // we are not on root level
		return docString(prefix, name, docPrefix, g.description)
	} else {
		result := g.description + "\n"
		result += strings.Repeat("-", len(g.description)) + "\n"
		return result
	}
}

func (g *mapType) Optional() MapType {
	g.optional = true
	return g
}

func (g *mapType) Enum(desc string) EnumType {
	t := NewEnumType(desc)
	g.subtype = t
	return t
}

func (g *mapType) Int(desc string) IntType {
	t := NewIntType(desc)
	g.subtype = t
	return t
}

func (g *mapType) String(desc string) StringType {
	t := NewStringType(desc)
	g.subtype = t
	return t
}

func (g *mapType) Object(desc string, ops func(ObjectType)) ObjectType {
	t := NewObjectType(desc, ops)
	g.subtype = t
	return t
}
