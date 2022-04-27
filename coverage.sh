#! /bin/sh

go test ./... -race -covermode=atomic -coverprofile=coverage.out

# make sure to set CODECOV_TOKEN env variable before doing this
codecov -f coverage.out
# or 
# bash <(curl -s https://codecov.io/bash) 

