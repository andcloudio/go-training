package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	buf.WriteByte('[')
	buf.WriteString("abc")
	buf.WriteByte(']')
	fmt.Println(buf.String())
}
