#! /bin/sh

go test ./... -race -covermode=atomic -coverprofile=coverage.out

codecov -f coverage.out

