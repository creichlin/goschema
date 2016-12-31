package goschema

import (
	"encoding/json"
	"errors"
	"github.com/creichlin/gutil"
	"github.com/xeipuuv/gojsonschema"
)

type goSchema struct {
	description string
	properties  *goProperties
}

func NewGOSchema(description string) GOSchema {
	gs := &goSchema{
		description: description,
	}
	gs.properties = &goProperties{
		parent: gs,
		props:  map[string]jsWriter{},
	}
	return gs
}

func (g *goSchema) Properties(pf func(Properties)) GOSchema {
	pf(g.properties)
	return g
}

func (g *goSchema) AsJSONSchema() (string, error) {
	root := map[string]interface{}{}

	root["title"] = g.description

	g.write(root)

	data, err := json.Marshal(root)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (g *goSchema) Validate(document map[string]interface{}) *gutil.ErrorCollector {
	errs := gutil.NewErrorCollector()
	schema, err := g.AsJSONSchema()
	if err != nil {
		errs.Add(err)
		return errs
	}
	schemaLoader := gojsonschema.NewStringLoader(schema)
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

func (g *goSchema) write(data map[string]interface{}) {
	data["type"] = "object"

	props := map[string]interface{}{}
	required := []string{}

	for name, value := range g.properties.props {
		props[name] = value.write()
		if value.isRequired() {
			required = append(required, name)
		}
	}

	data["properties"] = props
	data["additionalProperties"] = false
	if len(required) > 0 {
		data["required"] = required
	}
}
