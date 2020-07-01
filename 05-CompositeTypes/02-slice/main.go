package main

import "fmt"

func main() {
	// []T
	// make([]T, len, cap)
	s := make([]string, 3, 5)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))
	fmt.Println("cap", cap(s))
}
