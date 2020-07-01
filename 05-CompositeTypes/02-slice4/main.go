package main

import "fmt"

func main() {
	s := make([]string, 6)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"
	s[5] = "f"

	// s[low:high]
	// 0 <= low <= high <= cap(s)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	l = s[:]
	fmt.Println("sl4", l)
}
