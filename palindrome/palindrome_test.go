package main 

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	} {
		{"hello", false},
        {"A man, a plan, a canal, Panama!", true},
        {"racecar", true},
        {"Was it a car or a cat I saw?", true},
        {"No 'x' in Nixon", true},
        {"not a palindrome", false},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)

		if result != test.expected { 
			t.Errorf("isPalindrome(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}