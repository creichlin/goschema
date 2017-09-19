package goschema

type listType struct {
	baseType
	subtype Type
}

func NewListType(subType func(m ListType)) ListType {
	gol := &listType{
		baseType: baseType{},
	}
	subType(gol)
	return gol
}

func (g *listType) asJSONSchema() map[string]interface{} {
	data := map[string]interface{}{}
	data["title"] = g.description
	data["type"] = "array"
	data["additionalItems"] = false
	if g.subtype != nil {
		data["items"] = g.subtype.asJSONSchema()
	}

	return data
}

func (g *listType) docString(field string, docPrefix string) string {
	return g.subtype.docString(field+"[]", "list of")
}

func (g *listType) Enum(desc string) EnumType {
	t := NewEnumType(desc)
	g.subtype = t
	return t
}

func (g *listType) Int(desc string) IntType {
	t := NewIntType(desc)
	g.subtype = t
	return t
}

func (g *listType) Bool(desc string) BoolType {
	t := NewBoolType(desc)
	g.subtype = t
	return t
}

func (g *listType) String(desc string) StringType {
	t := NewStringType(desc)
	g.subtype = t
	return t
}

func (g *listType) Object(desc string, ops func(ObjectType)) ObjectType {
	t := NewObjectType(desc, ops)
	g.subtype = t
	return t
}

func (g *listType) Null(desc string) NullType {
	t := NewNullType(desc)
	g.subtype = t
	return t
}

func (g *listType) Any(desc string) AnyType {
	t := NewAnyType(desc)
	g.subtype = t
	return t
}

func (g *listType) SomeOf(ops func(SomeOf)) SomeOf {
	t := NewSomeOf(ops)
	g.subtype = t
	return t
}

func (g *listType) List(ops func(ListType)) ListType {
	t := NewListType(ops)
	g.subtype = t
	return t
}

func (g *listType) Map(ops func(MapType)) MapType {
	t := NewMapType("", ops)
	g.subtype = t
	return t
}
