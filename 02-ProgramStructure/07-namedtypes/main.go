package main

import "fmt"

type celcius float64
type fahreneit float64

func main() {
	var c celcius
	var f fahreneit

	fmt.Println(c == f)
}
