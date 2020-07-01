package main

import "fmt"

type Address struct {
	Hostname string
	Port     int
}

func main() {
	hits := make(map[Address]int)
	hits[Address{"golang.org", 443}]++
	fmt.Println(hits)
}
