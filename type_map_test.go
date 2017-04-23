package goschema_test

import (
	"github.com/creichlin/goschema"
	"github.com/creichlin/gutil"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	ints := goschema.NewMapType("A map with int values", func(m goschema.MapType) {
		m.Int("Int values")
	})

	floatMap := map[string]float64{
		"a": 5,
		"7": 9,
	}
	if goschema.ValidateGO(ints, floatMap).Has() {
		t.Errorf("map with floats should be valid")
		t.Log(goschema.ValidateGO(ints, floatMap))
		gutil.PrintAsYAML(goschema.AsJSONSchemaTree(ints))
	}

	stringMap := map[string]string{
		"a": "aa",
		"7": "77",
	}
	if !reflect.DeepEqual(goschema.ValidateGO(ints, stringMap).StringList(),
		[]string{"(root): Invalid type. Expected: integer, given: string", "(root): Invalid type. Expected: integer, given: string"}) {
		t.Errorf("map with strings should not be valid")
		t.Log(goschema.ValidateGO(ints, stringMap).StringList())
		gutil.PrintAsYAML(goschema.AsJSONSchemaTree(ints))
	}
}
