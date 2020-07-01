package main

import "sync"

var mu sync.Mutex
var m map[string]int

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func main() {
	m = map[string]int{
		"alice": 31,
		"dave":  24,
	}
	lookup("alice")
}
