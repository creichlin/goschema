package goschema

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/creichlin/gutil"
	"github.com/creichlin/gutil/format"
	"github.com/xeipuuv/gojsonschema"
)

type baseType struct {
	description string
}

// ValidateGO will validate the provided interface which can be a composite
// of maps, slices and scalars, restricted to constructs that are allowed
// in json (map keys must be strings, only float64, string, bool as scalars)
func ValidateGO(t Type, document interface{}) *gutil.ErrorCollector {
	errs := gutil.NewErrorCollector()
	schema := t.asJSONSchema()

	schemaLoader := gojsonschema.NewGoLoader(schema)
	documentLoader := gojsonschema.NewGoLoader(document)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		errs.Add(err)
		return errs
	}

	if !result.Valid() {
		for _, desc := range result.Errors() {
			errs.Add(errors.New(desc.Field() + ": " + desc.Description()))
		}
	}
	return errs
}

// Doc returns a string containing a very coarse documentation built from the
// validator definition
func Doc(t Type) string {
	docstr := t.docString("", "")
	return format.Align(docstr, "//")
}

// AsJSONSchema returns the validator as a json schema string
func AsJSONSchema(t Type) (string, error) {
	json, err := json.Marshal(t.asJSONSchema())
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// AsGOJSONSchema returns the validator as a map, slice, scalar composite defining the json schema
func AsGOJSONSchema(t Type) interface{} {
	return t.asJSONSchema()
}

func docString(field string, doc string, docs ...string) string {
	parts := []string{}
	for _, d := range docs {
		if d != "" {
			parts = append(parts, d)
		}
	}

	name := extractName(field)
	nameDoc := mergeNameDoc(name, doc)

	return cleanUpField(field) + " // " + nameDoc + strings.Join(parts, " ") + "\n"
}

func mergeNameDoc(name, doc string) string {
	if name == "" && doc == "" {
		return ""
	}
	if name == "" {
		return doc + ", "
	}
	if doc == "" {
		return name + ", "
	}
	return name + " " + doc + ", "
}

func cleanUpField(field string) string {
	if field == "" {
		return "."
	}
	if len(field) > 1 && field[0] == '.' {
		return field[1:]
	}
	return field
}

func extractName(field string) string {
	ln := strings.Split(field, ".")
	name := ln[len(ln)-1]

	if strings.HasSuffix(name, "]") || // if field is a list
		strings.HasSuffix(name, ")") || // or a someOf
		name == "*" { // or a map, there is no name
		return ""
	}
	return name
}
