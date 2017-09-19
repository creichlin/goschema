package goschema_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/creichlin/goschema"
	"github.com/creichlin/gutil"
)

type testCase struct {
	Name     string
	Schema   goschema.Type
	Document interface{}
	Errors   []string
}

var (
	person = goschema.NewObjectType("Example Schema", func(p goschema.ObjectType) {
		p.String("firstName", "")
		p.String("lastName", "").Optional()
		p.Int("age", "Age in years").Min(0).Max(5).Optional()
		p.Enum("gender", "").Add("male", "A male speciemen").Add("female", "A female speciemen")
	})

	strings = goschema.NewObjectType("Example String Schema", func(p goschema.ObjectType) {
		p.String("foo", "")
		p.String("bar", "").Optional()
	})

	integers = goschema.NewObjectType("Example integer schema", func(p goschema.ObjectType) {
		p.Int("foo", "")
		p.Int("min", "").Min(5).Optional()
		p.Int("max", "").Max(6).Optional()
		p.Int("minmax1", "").Min(3).Max(6).Optional()
		p.Int("minmax2", "").Min(3).Max(6).Optional()
	})

	enums = goschema.NewObjectType("Example enum schema", func(p goschema.ObjectType) {
		p.Enum("foo", "").Add("A", "").Add("B", "")
		p.Enum("bar", "").Add("X", "").Add("Y", "").Optional()
	})

	maps = goschema.NewObjectType("Example map schema", func(p goschema.ObjectType) {
		p.Map("map1", "mdesc", func(g goschema.MapType) {
			g.String("foooo")
		})
	})

	nested = goschema.NewObjectType("Nested object example", func(p goschema.ObjectType) {
		p.Object("nested", "Nested object", func(p goschema.ObjectType) {
			p.String("foo", "Foo")
			p.String("bar", "Bar")
		})
	})

	some = goschema.NewMapType("Some example", func(p goschema.MapType) {
		p.SomeOf("someof", func(p goschema.SomeOf) {
			p.String("foo")
			p.Bool("f")
		})
	})

	any = goschema.NewObjectType("Some example", func(p goschema.ObjectType) {
		p.Any("any", "this can be just anything")
	})

	nullSchema = goschema.NewSomeOf("null or string", func(p goschema.SomeOf) {
		p.String("either a tring")
		p.Null("or just null")
	})

	testCases = []testCase{
		{
			"Minimal valid person",
			person,
			map[string]interface{}{
				"firstName": "lala",
				"gender":    "male",
			},
			[]string{},
		}, {
			"required string",
			strings,
			map[string]interface{}{},
			[]string{"foo: foo is required"},
		}, {
			"extra string",
			strings,
			map[string]interface{}{
				"foo":  "x",
				"bloo": "y",
			},
			[]string{"bloo: Additional property bloo is not allowed"},
		}, {
			"required integer",
			integers,
			map[string]interface{}{},
			[]string{"foo: foo is required"},
		}, {
			"extra integer",
			integers,
			map[string]interface{}{
				"foo":  5,
				"bloo": 9,
			},
			[]string{"bloo: Additional property bloo is not allowed"},
		}, {
			"int out of range",
			integers,
			map[string]interface{}{
				"foo":     5,
				"min":     4,
				"max":     7,
				"minmax1": 2,
				"minmax2": 7,
			},
			[]string{
				"max: Must be less than or equal to 6",
				"min: Must be greater than or equal to 5",
				"minmax1: Must be greater than or equal to 3",
				"minmax2: Must be less than or equal to 6",
			},
		}, {
			"int in range",
			integers,
			map[string]interface{}{
				"foo":     5,
				"min":     5,
				"max":     6,
				"minmax1": 4,
				"minmax2": 5,
			},
			[]string{},
		}, {
			"required enum",
			enums,
			map[string]interface{}{},
			[]string{"foo: foo is required"},
		}, {
			"extra enum",
			enums,
			map[string]interface{}{
				"foo":  "A",
				"bloo": "Q",
			},
			[]string{"bloo: Additional property bloo is not allowed"},
		}, {
			"enum invalid value",
			enums,
			map[string]interface{}{
				"foo": "X",
			},
			[]string{`foo: foo must be one of the following: "A", "B"`},
		}, {
			"nested object",
			nested,
			map[string]interface{}{
				"nested": map[string]interface{}{
					"foo": "foo-value",
					"bar": "bar-value",
				},
			},
			[]string{},
		}, {
			"nested object with error",
			nested,
			map[string]interface{}{
				"nested": map[string]interface{}{
					"foo": "foo-value",
				},
			},
			[]string{"bar: bar is required"},
		}, {
			"map with error",
			maps,
			map[string]interface{}{
				"map1": map[string]interface{}{
					"foo": 6,
				},
			},
			[]string{"map1: Invalid type. Expected: string, given: integer"},
		}, {
			"some with string or bool",
			some,
			map[string]interface{}{
				"a1": "foo",
				"a2": false,
			},
			[]string{},
		}, {
			"some with string or bool AND int",
			some,
			map[string]interface{}{
				"a1": "foo",
				"a2": false,
				"a3": 5,
			},
			[]string{"(root): Invalid type. Expected: string, given: integer",
				"(root): Must validate at least one schema (anyOf)"},
		}, {
			"any as string",
			any,
			map[string]interface{}{
				"any": "foo",
			},
			[]string{},
		}, {
			"any as list",
			any,
			map[string]interface{}{
				"any": []string{"a", "b"},
			},
			[]string{},
		}, {
			"no any at all",
			any,
			map[string]interface{}{},
			[]string{"any: any is required"},
		}, {
			"null doc",
			nullSchema,
			nil,
			[]string{},
		}, {
			"string doc",
			nullSchema,
			"hey",
			[]string{},
		}, {
			"int doc",
			nullSchema,
			6,
			[]string{"(root): Invalid type. Expected: string, given: integer",
				"(root): Must validate at least one schema (anyOf)"},
		},
	}
)

func TestBasicExample(t *testing.T) {

	fmt.Println(goschema.Doc(person))
	fmt.Println(goschema.Doc(integers))

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			errs := goschema.ValidateGO(testCase.Schema, testCase.Document)
			errsList := errs.StringList()
			sort.Strings(errsList)
			if !reflect.DeepEqual(errsList, testCase.Errors) {
				t.Errorf("Errors don't match:\nexpected: '%v'\nactual: '%v'", testCase.Errors, errsList)
				gutil.PrintAsYAML(goschema.AsJSONSchemaTree(testCase.Schema))
			}
			js, _ := goschema.AsJSONSchema(testCase.Schema)
			goschema.MustValidateJSONSchema(t, string(js))
		})
	}
}
