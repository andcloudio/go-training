package main

import "fmt"

func main() {
	a := 0
Loop:
	for i := 0; i < 20; i++ {
		for j := 0; j < 10; i++ {
			if a == 15 {
				break Loop
			}
			a++
		}
	}
	fmt.Println(a)
}
