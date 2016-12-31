package goschema

type jsWriter interface {
	write() map[string]interface{}
	isRequired() bool
}
