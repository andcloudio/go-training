package main

import "fmt"

// Employee structure
type Employee struct {
	ID      int
	Name    string
	Address string
}

func main() {
	var dilbert Employee

	dilbert.ID = 123
	fmt.Printf("%#v\n", dilbert)

	var emp *Employee = &dilbert
	emp.ID = 321
	fmt.Printf("%#v\n", dilbert)
}
