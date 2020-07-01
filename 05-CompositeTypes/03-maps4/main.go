package main

import "fmt"

func main() {
	ages := map[string]int{
		"charlie": 34,
	}

	ages["bob"]++
	fmt.Println(ages)

}
