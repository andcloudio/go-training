package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []string{"Alice", "Bob", "Dave"}

	// function literal
	less := func(i, j int) bool {
		return len(people[i]) < len(people[j])
	}
	sort.Slice(people, less)
	fmt.Println(people)
}
