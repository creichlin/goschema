package goschema

type someOf struct {
	baseType
	subtypes []Type
}

func NewSomeOf(description string, subTypes func(m SomeOf)) SomeOf {
	gom := &someOf{
		baseType: baseType{
			description: description,
		},
	}
	subTypes(gom)
	return gom
}

func (g *someOf) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{}

	data["title"] = g.description
	sts := []interface{}{}

	for _, st := range g.subtypes {
		sts = append(sts, st.asJSONSchema())
	}

	data["anyOf"] = sts

	return data
}

func (g *someOf) docString(prefix string, name string, docPrefix string) string {
	dstr := docString(prefix, name, "is one of")
	for _, st := range g.subtypes {
		dstr += st.docString(prefix+"  ", "- ", "")
	}
	return dstr
}

func (g *someOf) Optional() SomeOf {
	g.optional = true
	return g
}

func (g *someOf) Bool(desc string) BoolType {
	t := NewBoolType(desc)
	g.subtypes = append(g.subtypes, t)
	return t
}

func (g *someOf) String(desc string) StringType {
	t := NewStringType(desc)
	g.subtypes = append(g.subtypes, t)
	return t
}

func (g *someOf) Null(desc string) NullType {
	t := NewNullType(desc)
	g.subtypes = append(g.subtypes, t)
	return t
}

func (g *someOf) Object(desc string, ops func(ObjectType)) ObjectType {
	t := NewObjectType(desc, ops)
	g.subtypes = append(g.subtypes, t)
	return t
}
