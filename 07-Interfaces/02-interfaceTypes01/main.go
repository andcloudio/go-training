package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	var w io.Writer
	// dynamic type = nil, dynamic value = nil

	var b bytes.Buffer

	w = &b
	// type = bytes.Buffer value = &b

	w.Write([]byte("hello world"))

	// type assertion
	c := w.(*bytes.Buffer)

	fmt.Println(c)
}
