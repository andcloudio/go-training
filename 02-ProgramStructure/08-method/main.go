package main

import "fmt"

// Celcius type declaration
type Celcius float64

func (c Celcius) String() string {
	return fmt.Sprintf("%gC", c)
}

func main() {
	var c Celcius
	c = 100
	fmt.Println(c)
}
