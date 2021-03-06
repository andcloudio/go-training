package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("false") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("Palindrome") {
		t.Error(`IsPalindrome("Palindrome") = true`)
	}
}
