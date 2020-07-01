package main

import "fmt"

// Struct Embedding

// Person structure
type Person struct {
	Name string
	Age  int
}

// PrintName of person
func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}

// Musician structure
type Musician struct {
	Person
	works []string
}

func main() {
	m := Musician{Person: Person{Name: "ARRahman", Age: 52}}
	m.PrintName()
}
