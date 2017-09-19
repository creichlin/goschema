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

func (g *stringType) docString(field string, docPrefix string) string {
	return docString(field, g.description, docPrefix, "string")
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
