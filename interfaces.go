package goschema

type Type interface {
	asJSONSchema() map[string]interface{}
	docString(prefix, name string, descPrefix string) string
	isRequired() bool
}

type ObjectType interface {
	Type
	Optional() ObjectType
	String(name string, desc string) StringType
	Int(name string, desc string) IntType
	Bool(name string, desc string) BoolType
	Enum(name string, desc string) EnumType
	Object(name string, desc string, ops func(ObjectType)) ObjectType
	Map(name string, desc string, ops func(MapType)) MapType
	List(name string, ops func(ListType)) ListType
	Any(name string, desc string) AnyType
}

type MapType interface {
	Type
	Optional() MapType
	String(desc string) StringType
	Int(desc string) IntType
	Bool(desc string) BoolType
	SomeOf(desc string, ops func(SomeOf)) SomeOf
	Enum(desc string) EnumType
	Object(desc string, ops func(ObjectType)) ObjectType
}

type ListType interface {
	Type
	Optional() ListType
	String(desc string) StringType
	Int(desc string) IntType
	Bool(desc string) BoolType
	Enum(desc string) EnumType
	Object(desc string, ops func(ObjectType)) ObjectType
}

type StringType interface {
	Type
	Optional() StringType
}

type EnumType interface {
	Type
	Optional() EnumType
	Add(key string, desc string) EnumType
}

type IntType interface {
	Type
	Min(min int) IntType
	Max(min int) IntType
	Optional() IntType
}

type BoolType interface {
	Type
	Optional() BoolType
}

type SomeOf interface {
	Type
	String(desc string) StringType
	Bool(desc string) BoolType
}

type AnyType interface {
	Type
}
