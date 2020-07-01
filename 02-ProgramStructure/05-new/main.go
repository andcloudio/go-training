package main

import "fmt"

func main() {
	p := new(int)
	*p = 2
	fmt.Println(*p)
}
