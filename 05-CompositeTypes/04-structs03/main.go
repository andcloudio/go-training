package main

import "fmt"

// Point structure
type Point struct {
	X, Y int
}

func main() {
	p := Point{X: 1}

	fmt.Printf("%#v\n", p)
}
