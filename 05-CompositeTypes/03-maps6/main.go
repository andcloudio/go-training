package main

import "fmt"

func main() {
	var ages = make(map[string]int)
	ages["carol"] = 35 // panic
	fmt.Println(ages)
}
