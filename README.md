goschema
========

Define data-structure validation in go code. Can be used to validate data from json/yaml files.
Uses https://github.com/xeipuuv/gojsonschema to do the actual validation.

Its implemented in a fluent interface style:

    personValidator := NewGOSchema("Example person schema").
      Properties(func(p Properties) {
	    p.String("firstName", "Given name")
	    p.String("lastName", "Second name and additional names if available").Optional()
	    p.Int("age", "Age in years").Min(0).Max(5).Optional()
	    p.Enum("gender", "").Add("male", "A male speciemen").Add("female", "A female speciemen")
	  })
	// person must be a map with string keys
	// it can also be loaded from a json file with json.Unmarshal(data, mapPointer)
	personValidator.Validate(person)
	
Todo
----

Creating plain text output and also markdown output
with documentation of the validation format in the sense of:

    Example person schema
    Properties
      firstName // Given name as string
      lastName  // optional Second name and additional names if available as string
      age       // optional Age in years as int from 0 to 5
      gender    // gender as string (male=A malespeciemen, femal=Afemale speciemen)
       
This data can be extracted from the defined valuies above including the help string which is usually
the second parameter in data-type definition calls
