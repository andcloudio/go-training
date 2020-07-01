package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	f("direct")

	go f("goroutine 1")

	go func(msg string) {
		f(msg)
	}("goroutine 2")

	time.Sleep(time.Second * 2)
	fmt.Println("done")
}
