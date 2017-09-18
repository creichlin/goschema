package goschema_test

import (
	"fmt"
	"testing"

	"github.com/creichlin/goschema"
)

func TestReadmeExample(t *testing.T) {
	person = goschema.NewObjectType("Example Schema", func(p goschema.ObjectType) {
		p.String("firstName", "")
		p.String("lastName", "").Optional()
		p.Bool("is-old", "The person is considered old").Optional()
		p.Int("age", "Age in years").Min(0).Max(5).Optional()
		p.Enum("gender", "").Add("male", "A male speciemen").Add("female", "A female speciemen")
		p.List("hobbies", func(p goschema.ListType) {
			p.String("all my hobbies")
		})
		p.List("siblings", func(p goschema.ListType) {
			p.Object("all my siblings", func(p goschema.ObjectType) {
				p.String("firstName", "")
				p.String("lastName", "").Optional()
			})
		})
		p.Map("results", "a map of ints or bools", func(g goschema.MapType) {
			g.SomeOf("bla", func(g goschema.SomeOf) {
				g.String("prosa")
				g.Bool("technical")
			})
		})
	})
	fmt.Println(goschema.Doc(person))
	fmt.Println()
	fmt.Println(goschema.AsJSONSchema(person))
}
