package goschema

// Type is an interface that is implemented by all nodes of a validation tree
type Type interface {
	asJSONSchema() map[string]interface{}
	docString(prefix, name string, descPrefix string) string
}

// TypeDefs provides methods to define a validator for a child of
// collection validators
type TypeDefs interface {
	String(desc string) StringType
	Int(desc string) IntType
	Bool(desc string) BoolType
	SomeOf(ops func(SomeOf)) SomeOf
	Enum(desc string) EnumType
	Null(desc string) NullType
	Any(desc string) AnyType
	Object(desc string, ops func(ObjectType)) ObjectType
	Map(ops func(MapType)) MapType
	List(ops func(ListType)) ListType
}

// ObjectType validates that the element is an object with named attributes
type ObjectType interface {
	Type
	Attribute(name string) ObjectAttribute
	Optional(name string) ObjectAttribute
}

// ObjectAttribute validates one single attribute of an object
type ObjectAttribute interface {
	TypeDefs
}

// MapType validates that the element is a map/hash (key value pairs)
// key is always string, value can be defined
type MapType interface {
	Type
	TypeDefs
}

// ListType validates that the value is a list values
type ListType interface {
	Type
	TypeDefs
}

// StringType validates that the value is a string
type StringType interface {
	Type
}

// EnumType validates that a given string is one of the provided values
type EnumType interface {
	Type
	Add(key string, desc string) EnumType
}

// IntType validates
type IntType interface {
	Type
	Min(min int) IntType
	Max(min int) IntType
}

// BoolType validates that the value is of type bool
type BoolType interface {
	Type
}

// SomeOf validates that the value is of at least one of the defined types
type SomeOf interface {
	Type
	TypeDefs
}

// AnyType can be just anything, map, list, null, scalar...
type AnyType interface {
	Type
}

// NullType validates that the given value is nil/null
type NullType interface {
	Type
}
