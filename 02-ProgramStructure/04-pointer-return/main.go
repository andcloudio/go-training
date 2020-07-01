package main

import "fmt"

func f() *int {
	v := 100
	return &v
}

func main() {
	var p = f()
	fmt.Println(*p)
}
