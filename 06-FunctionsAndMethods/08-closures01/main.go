package main

import "fmt"

// Closures
func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	f := increment()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
