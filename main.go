package main

import (
	"fmt"
	"os"

	"github.com/prabhatsharma/go-codecov/pkg/greet"
)

func main() {
	fmt.Println(greet.Hello("Someone", "english"))
	greet.HelloPrint(os.Stdout, "Someone", "english")
}
