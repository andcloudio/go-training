package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	pp := &Point{1, 2}
	fmt.Printf("%#v\n", pp)
}
