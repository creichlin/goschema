package goschema

type goStringProperty struct {
	parent      *goProperties
	description string
	optional    bool
}

func (g *goStringProperty) Optional() StringProperty {
	g.optional = true
	return g
}

func (g *goStringProperty) docString(prefix, name string) string {
	doc := prefix + name + " // "
	if g.optional {
		doc += " optional, "
	}

	if g.description == "" {
		doc += name + " "
	} else {
		doc += g.description + " "
	}

	return doc + "as string\n"
}

func (g *goStringProperty) writeJSONSchema() map[string]interface{} {
	data := map[string]interface{}{
		"type": "string",
	}
	if g.description != "" {
		data["description"] = g.description
	}
	return data
}

func (g *goStringProperty) isRequired() bool {
	return !g.optional
}
