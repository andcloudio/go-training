package main

import "fmt"

func main() {
	x := 100
	p := &x
	fmt.Println(*p)

	*p = 2
	fmt.Println(x)
}
