package goschema

import (
	"sort"
	"strings"
)

type objectType struct {
	baseType
	props map[string]Type
}

func NewObjectType(description string, each func(ObjectType)) ObjectType {
	gop := &objectType{
		baseType: baseType{
			description: description,
		},
		props: map[string]Type{},
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
		if value.isRequired() {
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

func (g *objectType) Optional() ObjectType {
	g.optional = true
	return g
}

func (g *objectType) Enum(name string, desc string) EnumType {
	t := NewEnumType(desc)
	g.props[name] = t
	return t
}

func (g *objectType) Int(name string, desc string) IntType {
	prop := NewIntType(desc)
	g.props[name] = prop
	return prop
}

func (g *objectType) Bool(name string, desc string) BoolType {
	prop := NewBoolType(desc)
	g.props[name] = prop
	return prop
}

func (g *objectType) String(name string, desc string) StringType {
	prop := NewStringType(desc)
	g.props[name] = prop
	return prop
}

func (g *objectType) Object(name string, desc string, ops func(ObjectType)) ObjectType {
	prop := NewObjectType(desc, ops)
	g.props[name] = prop
	return prop
}

func (g *objectType) Map(name string, desc string, ops func(MapType)) MapType {
	prop := NewMapType(desc, ops)
	g.props[name] = prop
	return prop
}

func (g *objectType) List(name string, ops func(ListType)) ListType {
	prop := NewListType(ops)
	g.props[name] = prop
	return prop
}

func (g *objectType) Any(name string, description string) AnyType {
	prop := NewAnyType(description)
	g.props[name] = prop
	return prop
}
