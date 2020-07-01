package palindrome

import (
	"fmt"
	"testing"
)

// Unit testing
func TestPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.in); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", test.in, got, test.want)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// false
	// false
}
