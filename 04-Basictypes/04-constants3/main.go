package main

import "fmt"

type Weekdays int

const (
	Sunday Weekdays = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d\n",
		Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
