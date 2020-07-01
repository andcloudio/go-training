package main

import "fmt"

func main() {

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	delete(ages, "alice")

	fmt.Println(ages)

}
