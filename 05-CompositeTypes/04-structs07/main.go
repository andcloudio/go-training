// STRUCT EMBEDDING

package main

import "fmt"

// Point structure
type Point struct {
	X, Y int
}

// Circle structure
type Circle struct {
	Point
	Radius int
}

// Wheel structure
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 1
	w.Y = 2
	w.Radius = 5
	w.Spokes = 10
	fmt.Printf("%#v\n", w)
}
