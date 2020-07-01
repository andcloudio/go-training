package main

import (
	"fmt"
	"time"
)

// Timeout

func main() {
	c1 := make(chan string, 1)

	select {
	case r := <-c1:
		fmt.Println(r)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
