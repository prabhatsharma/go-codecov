#! /bin/sh

go test ./... -race -covermode=atomic -coverprofile=coverage.out

codecov -f coverage.out
# or 
# bash <(curl -s https://codecov.io/bash) 

