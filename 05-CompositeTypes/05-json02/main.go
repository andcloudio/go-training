package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movies structure
type Movies struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movies{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json marshalling failed: %v", err)
	}

	fmt.Printf("%s\n", data)

	var titles []struct {
		Title string
	}

	err = json.Unmarshal(data, &titles)
	if err != nil {
		log.Fatalf("Json unmarshalling failed: %v", err)
	}

	fmt.Printf("%v\n", titles)
}
