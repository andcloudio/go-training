package main

import "os"

// Panic
func main() {
	panic("a problem occurred")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// Summary:
// func name(param-list) (result-list) {}
// multiple return values
// function values
// function literal & anonymous functions
// closures
// variadic functions
// defer
// panic
