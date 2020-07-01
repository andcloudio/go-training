package main

import "fmt"

// Encapsulation

// Counter structure
type Counter struct{ n int }

// N - get value of n
func (c *Counter) N() int { return c.n }

// Increment value of n
func (c *Counter) Increment() { c.n++ }

// Reset value of n
func (c *Counter) Reset() { c.n = 0 }

func main() {
	c := &Counter{}
	fmt.Println(c)

	c.Increment()
	fmt.Println(c.N())
	c.Reset()
	fmt.Println(c.N())
}

// Summary:
// Methods - opertions on named type
// Embedding - methods are promoted
// Hide the internal representation.
