package goschema

type Type interface {
	asJSONSchema() map[string]interface{}
	docString(prefix, name string) string
	isRequired() bool
}

type ObjectType interface {
	Type
	String(name string, desc string) StringType
	Int(name string, desc string) IntType
	Enum(name string, desc string) EnumType
	Object(name string, desc string, ops func(ObjectType)) ObjectType
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
