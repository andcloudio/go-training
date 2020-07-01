package main

func f() int {
	return 100
}
func main() {
	switch x := f(); x {
	case 100:
		x = x + 50
	case 200:
	default:
	}
}
