package greet

import (
	"fmt"
	"io"
)

func HelloPrint(writer io.Writer, name string, language string) {
	if name == "" {
		name = "World"
	}

	greeting := Hello(name, language)

	fmt.Fprintf(writer, "%s", greeting)
}
