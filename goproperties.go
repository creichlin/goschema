package goschema

type goProperties struct {
	parent *goSchema
	props  map[string]jsWriter
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
