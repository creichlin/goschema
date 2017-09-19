package goschema

import (
	"sort"
	"strings"
)

type objectType struct {
	baseType
	props    map[string]Type
	optional map[string]bool
}

type objectAttribute struct {
	object *objectType
	name   string
}

func NewObjectType(description string, each func(ObjectType)) ObjectType {
	gop := &objectType{
		baseType: baseType{
			description: description,
		},
		props:    map[string]Type{},
		optional: map[string]bool{},
	}
	each(gop)
	return gop
}

func (g *objectType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{}

	data["title"] = g.description
	data["type"] = "object"

	g.addProperties(data)

	return data
}

func (g *objectType) addProperties(data map[string]interface{}) {
	props := map[string]interface{}{}
	required := []string{}

	for name, value := range g.props {
		props[name] = value.asJSONSchema()
		if !g.optional[name] {
			required = append(required, name)
		}
	}

	data["properties"] = props
	data["additionalProperties"] = false
	if len(required) > 0 {
		data["required"] = required
	}
}

func (g *objectType) docString(prefix string, name string, docPrefix string) string {
	result := ""
	if name != "" { // we are not on root level
		result += docString(prefix, name, docPrefix, g.description)
		prefix += "  "
	} else {
		result += g.description + "\n"
		result += strings.Repeat("-", len(g.description)) + "\n"
	}

	keys := []string{}
	for key := range g.props {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		result += g.props[key].docString(prefix, key, "")
	}
	return result
}

func (g *objectType) Attribute(name string) ObjectAttribute {
	return &objectAttribute{
		object: g,
		name:   name,
	}
}

func (g *objectType) Optional(name string) ObjectAttribute {
	g.optional[name] = true
	return &objectAttribute{
		object: g,
		name:   name,
	}
}

func (g *objectAttribute) Enum(desc string) EnumType {
	t := NewEnumType(desc)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) Int(desc string) IntType {
	t := NewIntType(desc)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) Bool(desc string) BoolType {
	t := NewBoolType(desc)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) String(desc string) StringType {
	t := NewStringType(desc)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) Object(desc string, ops func(ObjectType)) ObjectType {
	t := NewObjectType(desc, ops)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) Map(ops func(MapType)) MapType {
	t := NewMapType("", ops)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) List(ops func(ListType)) ListType {
	t := NewListType(ops)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) Any(description string) AnyType {
	t := NewAnyType(description)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) Null(description string) NullType {
	t := NewNullType(description)
	g.object.props[g.name] = t
	return t
}

func (g *objectAttribute) SomeOf(desc string, ops func(SomeOf)) SomeOf {
	t := NewSomeOf(desc, ops)
	g.object.props[g.name] = t
	return t
}
