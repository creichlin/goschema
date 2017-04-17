package goschema

type jsWriter interface {
	writeJSONSchema() map[string]interface{}
	docString(prefix, name string) string
	isRequired() bool
}
