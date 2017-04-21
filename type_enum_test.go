package goschema_test

import (
	"github.com/creichlin/goschema"
	"testing"
)

func TestEnum(t *testing.T) {
	bi := goschema.NewEnumType("Gender").Add("male", "Male").Add("female", "Female")

	if goschema.ValidateGO(bi, "male").Has() {
		t.Errorf("male enum value should be valid")
	}
	if goschema.ValidateGO(bi, "female").Has() {
		t.Errorf("female enum value should be valid")
	}
	if !goschema.ValidateGO(bi, "mal").Has() {
		t.Errorf("mal should return an error")
	}
}
