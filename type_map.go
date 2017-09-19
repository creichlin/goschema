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
		return docString(prefix, name, docPrefix, g.description) + g.subtype.docString(prefix+"  ", "  ", "")
	}
	result := g.description + "\n"
	result += strings.Repeat("-", len(g.description)) + "\n"
	return result
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

func (g *mapType) Bool(desc string) BoolType {
	t := NewBoolType(desc)
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

func (g *mapType) SomeOf(ops func(SomeOf)) SomeOf {
	t := NewSomeOf(ops)
	g.subtype = t
	return t
}

func (g *mapType) Null(desc string) NullType {
	t := NewNullType(desc)
	g.subtype = t
	return t
}

func (g *mapType) Any(desc string) AnyType {
	t := NewAnyType(desc)
	g.subtype = t
	return t
}

func (g *mapType) List(ops func(ListType)) ListType {
	t := NewListType(ops)
	g.subtype = t
	return t
}

func (g *mapType) Map(ops func(MapType)) MapType {
	t := NewMapType("", ops)
	g.subtype = t
	return t
}
