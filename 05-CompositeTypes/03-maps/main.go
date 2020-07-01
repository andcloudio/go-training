package main

import "fmt"

func main() {
	// map[k]v
	ages := make(map[string]int)

	ages["alice"] = 31
	ages["charlie"] = 34

	fmt.Println(ages["alice"])
	fmt.Println(ages["charlie"])
}
