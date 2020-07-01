package main

import "fmt"

type rect struct {
	width, height int
}

// Methods
func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 5, height: 4}

	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())

}
