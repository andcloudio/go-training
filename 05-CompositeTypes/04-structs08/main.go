package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{
		Circle: Circle{
			Point:  Point{X: 1, Y: 2},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
}

// Summary:
// structs are aggregate type
// dot notation to access fields
// struct literal to initilize values to fields
// struct can be embedding in another struct
