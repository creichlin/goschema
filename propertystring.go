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

func (g *goStringProperty) write() map[string]interface{} {
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
