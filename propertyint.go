package goschema

import "fmt"

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

func (g *goIntProperty) docString(prefix, name string) string {
	doc := prefix + name + " // "
	if g.optional {
		doc += " optional, "
	}

	if g.description == "" {
		doc += name + " "
	} else {
		doc += g.description + " "
	}

	doc += "as int "

	if g.min != nil && g.max != nil {
		doc += fmt.Sprintf("from %v to %v", *g.min, *g.max)
	} else if g.min != nil {
		doc += fmt.Sprintf("%v or more", *g.min)
	} else if g.max != nil {
		doc += fmt.Sprintf("%v or less", *g.max)
	}
	return doc + "\n"
}

func (g *goIntProperty) writeJSONSchema() map[string]interface{} {
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
