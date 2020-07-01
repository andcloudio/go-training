package main

import "fmt"

// func name1(param-list)(result-list) {
// }

// func name2(param-list) {
// }

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
}
