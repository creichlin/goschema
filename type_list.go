package goschema

import (
	"strings"
)

type listType struct {
	baseType
	subtype Type
}

func NewListType(description string, subType func(m ListType)) ListType {
	gol := &listType{
		baseType: baseType{
			description: description,
		},
	}
	subType(gol)
	return gol
}

func (g *listType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{}
	data["title"] = g.description
	data["type"] = "array"
	data["additionalItems"] = false
	if g.subtype != nil {
		data["items"] = g.subtype.asJSONSchema()
	}

	return data
}

func (g *listType) docString(prefix string, name string) string {
	result := prefix
	if name != "" { // we are not on root level
		result += name + " // " + g.description + "\n"
	} else {
		result += g.description + "\n"
		result += strings.Repeat("-", len(g.description)) + "\n"
	}

	return result
}

func (g *listType) Enum(desc string) EnumType {
	t := NewEnumType(desc)
	g.subtype = t
	return t
}

func (g *listType) Optional() ListType {
	g.optional = true
	return g
}

func (g *listType) Int(desc string) IntType {
	t := NewIntType(desc)
	g.subtype = t
	return t
}

func (g *listType) String(desc string) StringType {
	t := NewStringType(desc)
	g.subtype = t
	return t
}

func (g *listType) Object(desc string, ops func(ObjectType)) ObjectType {
	t := NewObjectType(desc, ops)
	g.subtype = t
	return t
}
