package main

import "fmt"

func main() {
	var apples int32 = 1
	var oranges int16 = 2
	var compose int = apples + oranges // error
	fmt.Println(compose)
}
