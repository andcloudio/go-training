package bank

import "sync"

var (
	mu      sync.RWMutex
	balance int
)

// Deposit amount
func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

// Balance read
func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	b := balance
	return b
}
