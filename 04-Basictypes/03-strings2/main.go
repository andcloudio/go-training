package main

import "fmt"

func main() {
	b := []byte("abc")
	b[0] = 'A'
	fmt.Println(string(b))
}
