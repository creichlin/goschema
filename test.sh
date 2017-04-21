#!/bin/bash

export GOPATH=$(dirname $(dirname $(dirname $(dirname $(pwd)))))

go test -coverprofile=.coverprofile github.com/creichlin/goschema


