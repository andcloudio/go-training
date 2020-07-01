package main

import (
	"fmt"
)

// Buffered Channels

func naturals() <-chan int {
	ch := make(chan int, 10)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	ch := naturals()

	for v := range ch {
		fmt.Printf("%d ", v)
	}
}
