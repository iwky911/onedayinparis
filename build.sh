#!/bin/sh
export GOPATH=$PWD
go build src/matrix_computation/*.go
go build src/tsp/*.go