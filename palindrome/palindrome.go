package main 

import (
	"fmt"
	"strings"
	"unicode"
)

// isPalindrome checks whether a given string is a palindrome, ignoring spaces, punctuation, and capitalization.
func isPalindrome(s string) bool {
	var builder strings.Builder
	s = strings.ToLower(s)

	// Normalize the input string:  convert to lowercase, remove spaces, punctuation and capitalization
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			builder.WriteRune(char)
		}
	}
	
	normalized_string := builder.String()

	// check if the normalized string is  a palindrome
	length := len(normalized_string)

	for i := 0; i < length / 2; i ++ {
		if normalized_string[i] != normalized_string[length - i - 1] {
			return false
		}
	}
	return true
}

func main() {
	testcases := []string {
		"hello",
		"A man, a plan, a canal, Panama!",
		"racecar",
		"Was it a car or a cat I saw?",
		"No 'x' in Nixon",
		"not a palindrome",
	}

	for _, testcase := range testcases {
		fmt.Printf("isPalindrome(%q) = %v\n", testcase, isPalindrome(testcase))
	}
}