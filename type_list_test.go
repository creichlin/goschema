package goschema_test

import (
	"github.com/creichlin/goschema"
	"github.com/creichlin/gutil"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	ints := goschema.NewListType("A List with int values", func(m goschema.ListType) {
		m.Int("Int values")
	})

	floatList := []float64{
		5,
		9,
	}
	if goschema.ValidateGO(ints, floatList).Has() {
		t.Errorf("list with floats should be valid")
		t.Log(goschema.ValidateGO(ints, floatList))
		gutil.PrintAsYAML(goschema.AsJSONSchemaTree(ints))
	}

	stringList := []string{
		"aa",
		"77",
	}
	if !reflect.DeepEqual(goschema.ValidateGO(ints, stringList).StringList(),
		[]string{"0: Invalid type. Expected: integer, given: string", "1: Invalid type. Expected: integer, given: string"}) {
		t.Errorf("list with strings should not be valid")
		t.Log(goschema.ValidateGO(ints, stringList).StringList())
		gutil.PrintAsYAML(goschema.AsJSONSchemaTree(ints))
	}
}
