package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout

	// type assertion returns two value
	f, ok := w.(*os.File)
	fmt.Println(f, ok)

	c, ok := w.(*bytes.Buffer)
	fmt.Println(c, ok)
}
