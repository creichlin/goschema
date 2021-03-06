package goschema

import (
	"os"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func MustValidateJSONSchema(t *testing.T, json string) {
	errs := validateJSONSchema(t, json)
	if len(errs) != 0 {
		t.Errorf("invalid json schema %v", errs)
	}
}

func validateJSONSchema(t *testing.T, json string) []string {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("Could not get working directory. really?")
	}

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + wd + "/testdata/jsonschema.json")
	documentLoader := gojsonschema.NewStringLoader(json)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		t.Errorf("Could not load validator, %v", err)
	}

	if result.Valid() {
		return []string{}
	}

	errs := []string{}
	for _, desc := range result.Errors() {
		errs = append(errs, desc.Description())
	}
	return errs
}
