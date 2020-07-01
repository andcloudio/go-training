package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len", len(s))
	fmt.Println("cap", cap(s))
}
