package main

import "fmt"

// Channel Direction

func ping(out chan<- string) {
	out <- "message"
}

func pong(in <-chan string, out chan<- string) {
	msg := <-in
	out <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go ping(ch1)
	go pong(ch1, ch2)

	msg := <-ch2
	fmt.Println(msg)
}
