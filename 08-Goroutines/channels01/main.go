package main

import "fmt"

// Channels
func main() {
	ch := make(chan string)

	go func() { ch <- "hello" }()

	m := <-ch
	fmt.Println(m)
}
