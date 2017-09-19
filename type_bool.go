package goschema

type boolType struct {
	baseType
}

func NewBoolType(description string) BoolType {
	return &boolType{
		baseType: baseType{
			description: description,
		},
	}
}

func (g *boolType) Optional() BoolType {
	g.optional = true
	return g
}

func (g *boolType) docString(prefix, name string, docPrefix string) string {
	return docString(prefix, name, docPrefix, optionalMap[g.optional], g.description, "as bool")
}

func (g *boolType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{
		"type": "boolean",
	}
	if g.description != "" {
		data["description"] = g.description
	}
	return data
}