package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	d := make([]string, len(s))
	copy(d, s)
	fmt.Println("cpy", d)
}
