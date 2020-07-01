package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.ReadWriteCloser
	w = os.Stdout
	w = new(bytes.Buffer)
}
