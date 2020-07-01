package main

import (
	"fmt"
	"log"
	"os"
)

func readFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	// Deferred function call
	defer f.Close()

	b := make([]byte, 10)

	_, err = f.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	b, err := readFile("abc")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(string(b))
}
