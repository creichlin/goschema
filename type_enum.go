package goschema

type enumItem struct {
	key         string
	description string
}

type enumType struct {
	baseType
	items []enumItem
}

func NewEnumType(description string) EnumType {
	return &enumType{
		baseType: baseType{
			description: description,
		},
	}
}

func (g *enumType) docString(field string, docPrefix string) string {
	doc := docString(field, g.description, docPrefix, "must be one of:")

	for _, item := range g.items {
		doc += " //   - \"" + item.key + "\" " + item.description + "\n"
	}
	return doc
}

func (g *enumType) Add(key string, desc string) EnumType {
	g.items = append(g.items, enumItem{
		key:         key,
		description: desc,
	})
	return g
}

func (g *enumType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{
		"type": "string",
	}
	values := []string{}
	for _, item := range g.items {
		values = append(values, item.key)
	}
	data["enum"] = values
	if g.description != "" {
		data["description"] = g.description
	}
	return data
}
