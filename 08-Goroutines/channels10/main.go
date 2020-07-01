package main

import "fmt"

// Non blocking communication

func main() {
	ch := make(chan string)

	select {
	case m := <-ch:
		fmt.Println("received message", m)
	default:
		fmt.Println("no message received")
	}
}
