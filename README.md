goschema
========

Define data-structure validation in go code. Can be used to validate data from json/yaml files.
Uses https://github.com/xeipuuv/gojsonschema to do the actual validation.

usage
-----

It's implemented in a fluent interface style:

    personSchema := NewObjectType("Example person schema", func(p ObjectType) {
		p.String("firstName", "")
		p.String("lastName", "").Optional()
		p.Int("age", "Age in years").Min(0).Max(5).Optional()
		p.Enum("gender", "").Add("male", "A male specimen").Add("female", "A female specimen")
	})
	// person must be a map with string keys
	// it can also be loaded from a json file with json.Unmarshal(data, mapPointer)
	goschema.Validate(personSchema, person)


doc generation
--------------

Creating plain text output with documentation of the validation format
in the sense of:

    Example person schema
    ---------------------
    age       //  optional, Age in years as int from 0 to 5
    firstName // firstName as string
    gender    // 
      male    // A male speciemen
      female  // A female speciemen
    lastName  //  optional, lastName as string

can be done with

    str := goschema.Doc(personSchema)

This data can be extracted from the defined values above including the help string which is usually
the second parameter in data-type definition calls
