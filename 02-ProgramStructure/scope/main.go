package main

import "fmt"

func f() int {
	return 100
}

func g() int {
	return 200
}

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}

	// declaration error
	fmt.Println(x, y)
}
