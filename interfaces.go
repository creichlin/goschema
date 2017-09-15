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
	Enum(name string, desc string) EnumType
	Object(name string, desc string, ops func(ObjectType)) ObjectType
	Map(name string, desc string, ops func(MapType)) MapType
	List(name string, ops func(ListType)) ListType
}

type MapType interface {
	Type
	Optional() MapType
	String(desc string) StringType
	Int(desc string) IntType
	Enum(desc string) EnumType
	Object(desc string, ops func(ObjectType)) ObjectType
}

type ListType interface {
	Type
	Optional() ListType
	String(desc string) StringType
	Int(desc string) IntType
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
