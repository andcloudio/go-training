package main

var global *int

func f() {
	// heap allocated
	var x int
	x = 100
	global = &x
}

func g() {
	// stack allocated
	y := new(int)
	*y = 1
}

func main() {
	f()
	g()
}
