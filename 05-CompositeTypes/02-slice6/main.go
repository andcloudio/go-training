package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	l := s[:2]
	fmt.Printf("len=%d, cap=%d\n", len(l), cap(l))
}
