package goschema

type nullType struct {
	baseType
}

func NewNullType(description string) NullType {
	return &nullType{
		baseType: baseType{
			description: description,
		},
	}
}

func (g *nullType) docString(field string, docPrefix string) string {
	return docString(field, g.description, docPrefix, "must be nothing (null)")
}

func (g *nullType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{
		"type": "null",
	}
	if g.description != "" {
		data["description"] = g.description
	}
	return data
}
