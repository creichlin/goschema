package goschema

type baseType struct {
	description string
	optional    bool
}

func (g *baseType) isRequired() bool {
	return !g.optional
}
