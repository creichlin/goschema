package goschema

import (
	"encoding/json"
	"errors"
	"github.com/creichlin/gutil"
	"github.com/creichlin/gutil/format"
	"github.com/xeipuuv/gojsonschema"
)

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

func Doc(t Type) string {
	docstr := t.docString("", "")
	return format.Align(docstr, "//")
}

func AsJSONSchema(t Type) (string, error) {
	json, err := json.Marshal(t.asJSONSchema())
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func AsJSONSchemaTree(t Type) interface{} {
	return t.asJSONSchema()
}
