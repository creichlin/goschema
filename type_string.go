package goschema

type stringType struct {
	baseType
}

func NewStringType(description string) StringType {
	return &stringType{
		baseType: baseType{
			description: description,
		},
	}
}

func (g *stringType) Optional() StringType {
	g.optional = true
	return g
}

func (g *stringType) docString(prefix, name string) string {
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

func (g *stringType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{
		"type": "string",
	}
	if g.description != "" {
		data["description"] = g.description
	}
	return data
}
