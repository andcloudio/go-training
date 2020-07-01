package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p := Point{1, 2}
	q := Point{1, 2}
	r := Point{3, 4}
	fmt.Println(p == q, q == r)
}
