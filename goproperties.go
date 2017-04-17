package goschema

import "sort"

type goProperties struct {
	parent *goSchema
	props  map[string]jsWriter
}

func (g *goProperties) docString(prefix string) string {
	result := ""

	keys := []string{}
	for key := range g.props {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		result += g.props[key].docString(prefix, key)
	}
	return result
}

func (g *goProperties) Enum(name string, desc string) EnumProperty {
	prop := &goEnumProperty{
		parent:      g,
		description: desc,
	}
	g.props[name] = prop
	return prop
}

func (g *goProperties) Int(name string, desc string) IntProperty {
	prop := &goIntProperty{
		parent:      g,
		description: desc,
	}
	g.props[name] = prop
	return prop
}

func (g *goProperties) String(name string, desc string) StringProperty {
	prop := &goStringProperty{
		parent:      g,
		description: desc,
	}
	g.props[name] = prop
	return prop
}
