package main

import "fmt"

// Flags - interface state
type Flags uint

const (
	// FlagUp - interface up
	FlagUp Flags = 1 << iota //0001

	// FlagBroadcast - Broadcast enabled
	FlagBroadcast //0010

	// FlagLoopback - loopback interface
	FlagLoopback //0100
)

// IsUp - check if interface is UP
func IsUp(v Flags) bool {
	return (v&FlagUp == FlagUp)
}
func main() {
	var v = FlagLoopback | FlagUp

	fmt.Printf("%b, %t\n", v, IsUp(v))
}
