package main

import "fmt"

func main() {
	var s string
	var s1 string = "hello"
	var s2 = "world"

	var i, j, k int
	var b, f, s3 = true, 2.3, "hello"

	fmt.Printf("s=%s, s1=%s s2=%s\n", s, s1, s2)
	fmt.Printf("i=%d, j=%d, k=%d\n", i, j, k)
	fmt.Printf("b=%t, f=%f, s3=%s\n", b, f, s3)

	m := 10
	fmt.Printf("m=%d\n", m)

}
