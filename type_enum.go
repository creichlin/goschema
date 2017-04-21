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

func (g *enumType) docString(prefix, name string) string {
	doc := prefix + name + " // "
	if g.optional {
		doc += "optional, "
	}
	doc += g.description + "\n"
	for _, item := range g.items {
		doc += prefix + "  " + item.key + " // " + item.description + "\n"
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

func (g *enumType) Optional() EnumType {
	g.optional = true
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
