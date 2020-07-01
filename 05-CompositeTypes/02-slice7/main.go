package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	s := []int{1, 2, 3}
	reverse(s[:])
	fmt.Println(s)
}

// Note:
// zero value of slice is nil, ie no underlying array.
// []int{} - empty slice with len=cap=0

// slices are not comparable.
// For comparing []bytes - bytes.Equal()
// write own function for other types.

// Summary:
// slice - []T - variable length array
// make([]T, len, cap) - create slice
// append()
// slice operator - s[l:h] - refer to elements within slice
