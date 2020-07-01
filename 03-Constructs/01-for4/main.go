package main

import "fmt"

func main() {
	for i, v := range []string{"simon", "jessy"} {
		fmt.Println(i, v)
	}

	for _, v := range []string{"peter", "mark"} {
		fmt.Println(v)
	}
}
