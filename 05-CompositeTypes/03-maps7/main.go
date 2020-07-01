package main

import "fmt"

func main() {
	ages := map[string]int{
		"charlie": 34,
	}

	if age, ok := ages["peter"]; !ok {
		fmt.Println("key peter is not present")
	} else {
		fmt.Printf("peter's age is %d\n", age)
	}

	if age, ok := ages["charlie"]; !ok {
		fmt.Println("key charlie is not present")
	} else {
		fmt.Printf("charlie's age is %d\n", age)
	}
}

// maps cannot be compared to each other, only legal comparison is with nil.

// Summary:
// map[k]v are unordered collection of key/value pair.
// make() to create map
// map literal to create and initialize
// map subscript notation m[k] to access values
// m[k] return two values- v, ok
