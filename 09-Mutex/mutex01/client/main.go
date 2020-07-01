package main

import (
	"fmt"
	"sync"

	"go-training/09-Mutex/mutex01/bank"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
	}()
	go func() {
		defer wg.Done()
		bank.Deposit(100)
	}()

	wg.Wait()
	fmt.Println("=", bank.Balance())
}
