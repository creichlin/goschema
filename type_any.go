package goschema

type anyType struct {
	baseType
}

func NewAnyType(description string) AnyType {
	return &anyType{
		baseType: baseType{
			description: description,
		},
	}
}

func (g *anyType) docString(prefix, name string, docPrefix string) string {
	return docString(prefix, name, docPrefix, g.description, "can be anything")
}

func (g *anyType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{}
	if g.description != "" {
		data["description"] = g.description
	}
	return data
}
