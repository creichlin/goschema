package goschema

type goIntProperty struct {
	parent      *goProperties
	description string
	optional    bool
	min         *int
	max         *int
}

func (g *goIntProperty) Optional() IntProperty {
	g.optional = true
	return g
}

func (g *goIntProperty) Min(min int) IntProperty {
	g.min = &min
	return g
}

func (g *goIntProperty) Max(max int) IntProperty {
	g.max = &max
	return g
}

func (g *goIntProperty) write() map[string]interface{} {
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

func (g *goIntProperty) isRequired() bool {
	return !g.optional
}
