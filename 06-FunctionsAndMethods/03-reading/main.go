package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("abc")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()
	b := make([]byte, 10)

	for i := 0; i < 10; i++ {
		_, err = f.Read(b)

		// Distinguished Errors
		if err == io.EOF {
			log.Printf("end of file")
			break
		}
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}
	}

	fmt.Println(string(b))
}
