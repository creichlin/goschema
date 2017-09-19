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

func (g *boolType) docString(field string, docPrefix string) string {
	return docString(field, g.description, docPrefix, "bool")
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
