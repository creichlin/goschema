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

func (g *stringType) docString(prefix, name string, docPrefix string) string {
	return docString(prefix, name, docPrefix, g.description, "as string")
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
