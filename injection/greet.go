package injection

import (
	"fmt"
	"io"
)

// Greet greets a person
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
