package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var compute = func() {
	defer wg.Done()

	if cancelled() {
		return
	}

	time.Sleep(1 * time.Second)
	fmt.Println("compute1 done")

	if cancelled() {
		return
	}

	time.Sleep(3 * time.Second)
	fmt.Println("compute2 done")
}

// Cancel goroutines

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	wg.Add(2)
	go compute()
	go compute()
	close(done)
	fmt.Println("all goroutines cancelled")
	wg.Wait()
	fmt.Println("all goroutines exited")
}
