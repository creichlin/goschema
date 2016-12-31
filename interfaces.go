package goschema

import "github.com/creichlin/gutil"

type GOSchema interface {
	Properties(func(Properties)) GOSchema
	AsJSONSchema() (string, error)
	Validate(map[string]interface{}) *gutil.ErrorCollector
}

type Properties interface {
	String(name string, desc string) StringProperty
	Int(name string, desc string) IntProperty
	Enum(name string, desc string) EnumProperty
}

type StringProperty interface {
	Optional() StringProperty
}

type EnumProperty interface {
	Optional() EnumProperty
	Add(key string, desc string) EnumProperty
}

type IntProperty interface {
	Min(min int) IntProperty
	Max(min int) IntProperty
	Optional() IntProperty
}
