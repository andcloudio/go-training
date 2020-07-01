package main

import "fmt"

func coinflip(x int) string {
	if x%2 == 0 {
		return "heads"
	} else {
		return "tails"
	}
}

func main() {
	var heads, tails int
	for i := 0; i < 10; i++ {
		switch coinflip(i) {
		case "heads":
			heads++
		case "tails":
			tails++
		default:
		}
	}
	fmt.Println(heads, tails)
}
