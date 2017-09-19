package goschema_test

import (
	"fmt"
	"testing"

	"github.com/creichlin/goschema"
)

func TestReadmeExample(t *testing.T) {
	person = goschema.NewObjectType("example schema", func(p goschema.ObjectType) {
		p.Attribute("firstName").String("")
		p.Optional("lastName").String("")
		p.Optional("is-old").Bool("is set when the person is considered to be old")
		p.Optional("age").Int("in years").Min(0).Max(5)
		p.Attribute("gender").Enum("describes the persons sex").
			Add("male", "specimen").
			Add("female", "specimen")
		p.Attribute("hobbies").List(func(p goschema.ListType) {
			p.String("all my hobbies")
		})
		p.Attribute("siblings").List(func(p goschema.ListType) {
			p.Object("all my siblings", func(p goschema.ObjectType) {
				p.Attribute("firstName").String("")
				p.Optional("lastName").String("")
			})
		})
		p.Attribute("results").Map(func(g goschema.MapType) {
			g.SomeOf(func(g goschema.SomeOf) {
				g.String("prosa")
				g.Bool("technical")
			})
		})
	})
	fmt.Println(goschema.Doc(person))
	fmt.Println()
	fmt.Println(goschema.AsJSONSchema(person))
}
