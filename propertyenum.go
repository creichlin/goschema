package goschema

type goEnumItem struct {
	key         string
	description string
}

type goEnumProperty struct {
	parent      *goProperties
	description string
	items       []goEnumItem
	optional    bool
}

func (g *goEnumProperty) Add(key string, desc string) EnumProperty {
	g.items = append(g.items, goEnumItem{
		key:         key,
		description: desc,
	})
	return g
}

func (g *goEnumProperty) Optional() EnumProperty {
	g.optional = true
	return g
}

func (g *goEnumProperty) write() map[string]interface{} {
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

func (g *goEnumProperty) isRequired() bool {
	return !g.optional
}
