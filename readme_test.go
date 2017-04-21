package goschema_test

import (
	"fmt"
	"github.com/creichlin/goschema"
	"testing"
)

func TestReadmeExample(t *testing.T) {
	person = goschema.NewObjectType("Example Schema", func(p goschema.ObjectType) {
		p.String("firstName", "")
		p.String("lastName", "").Optional()
		p.Int("age", "Age in years").Min(0).Max(5).Optional()
		p.Enum("gender", "").Add("male", "A male speciemen").Add("female", "A female speciemen")
	})
	fmt.Println(goschema.Doc(person))
	fmt.Println()
	fmt.Println(goschema.AsJSONSchema(person))
}
