package main

import "fmt"

func main() {
	x := 100
	y := 200

	fmt.Println(&x == &x, &x == &y, &x == nil)
}
