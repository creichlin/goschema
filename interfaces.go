package goschema

type Type interface {
	asJSONSchema() map[string]interface{}
	docString(prefix, name string, descPrefix string) string
	isRequired() bool
}

type TypeDefs interface {
	String(desc string) StringType
	Int(desc string) IntType
	Bool(desc string) BoolType
	SomeOf(desc string, ops func(SomeOf)) SomeOf
	Enum(desc string) EnumType
	Null(desc string) NullType
	Any(desc string) AnyType
	Object(desc string, ops func(ObjectType)) ObjectType
	Map(ops func(MapType)) MapType
	List(ops func(ListType)) ListType
}

type ObjectType interface {
	Type
	Attribute(name string) ObjectAttribute
	Optional(name string) ObjectAttribute
}

type ObjectAttribute interface {
	TypeDefs
}

type MapType interface {
	Type
	TypeDefs
}

type ListType interface {
	Type
	TypeDefs
}

type StringType interface {
	Type
}

type EnumType interface {
	Type
	Add(key string, desc string) EnumType
}

type IntType interface {
	Type
	Min(min int) IntType
	Max(min int) IntType
}

type BoolType interface {
	Type
}

type SomeOf interface {
	Type
	TypeDefs
}

type AnyType interface {
	Type
}

type NullType interface {
	Type
}
