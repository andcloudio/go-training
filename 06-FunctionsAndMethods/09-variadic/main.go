package main

import "fmt"

// Variadic functions
func sum(vs ...int) int {
	var sum int
	for _, v := range vs {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3))
}
