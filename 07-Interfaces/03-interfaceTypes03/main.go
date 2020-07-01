package main

import (
	"io"
	"os"
)

func main() {
	os.Stdout.Write([]byte("hello"))
	os.Stdout.Close()

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello"))

	w.Close() // compilation error
	// only methods defined by interface are revealed
}
