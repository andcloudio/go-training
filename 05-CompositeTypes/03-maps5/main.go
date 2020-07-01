package main

import "fmt"

func main() {
	ages := map[string]int{
		"charlie": 34,
		"alice":   32,
	}

	for k, v := range ages {
		fmt.Println(k, v)
	}
}

// Note:
// zero value of map is nil. ie no hash table.
// empty map - map[string]int {} - length map will be 0
