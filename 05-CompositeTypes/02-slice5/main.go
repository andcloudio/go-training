package main

import "fmt"

func main() {
	s := make([]int, 3)
	fmt.Println(s[:20]) // panic
}
