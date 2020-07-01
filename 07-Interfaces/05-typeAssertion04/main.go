package main

import (
	"io"
	"os"
)

// StringWriter interface adds additional methods
type StringWriter interface {
	io.Writer
	WriteString(p string) (int, error)
}

func main() {
	var w io.Writer
	w = os.Stdout

	// Querying behaviors with interface type assertions
	if sw, ok := w.(StringWriter); ok {
		sw.WriteString("hello")
	}
}
