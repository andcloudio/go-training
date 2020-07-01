package main

import "fmt"

func main() {
	var a [3]int
	fmt.Printf("first element = %d\n", a[0])
	fmt.Printf("last element = %d\n", a[len(a)-1])

	var b = [3]int{1, 2, 3}

	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Println()

	var c = []int{4, 5, 6}

	for i, v := range c {
		fmt.Println(i, v)
	}
}
