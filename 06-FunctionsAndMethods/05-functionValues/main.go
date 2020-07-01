// Function Values

package main

import "fmt"

func square(n int) int { return n * n }

func main() {
	// function values
	f := square
	n := f(3)
	fmt.Println(n)
}
