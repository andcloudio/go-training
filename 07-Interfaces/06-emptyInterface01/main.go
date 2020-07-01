package main

import "fmt"

func main() {
	// Empty Interface
	var any interface{}
	any = true
	any = 1.2
	any = "hello"
	fmt.Printf("%s\n", any.(string))
}
