package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("abc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f.Stat())
}
