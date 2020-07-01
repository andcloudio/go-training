package main

import "fmt"

func incr(p *int) {
	*p++
}
func main() {
	x := 100
	incr(&x)
	fmt.Println(x)
}
