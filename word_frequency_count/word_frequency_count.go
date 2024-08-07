package main

import (
	"fmt"
	"strings"
	"unicode"
)

// removePunctuation removes all punctuation marks from a string.
func removePunctuation(s string) string {
	var builder strings.Builder
	for _, char := range s {
		if !unicode.IsPunct(char) {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

// wordFrequencyCount takes a string as input and returns a map containing the frequency of each word in the string.
func wordFrequencyCount(sentence string) map[string]int {
	sentence = removePunctuation(sentence)
	frequencyCount := make(map[string]int)
	words := strings.Fields(sentence)
	for _, word := range words {
		frequencyCount[strings.ToLower(word)]++
	}
	return frequencyCount
}

func main() {
	wordFrequency := wordFrequencyCount("Hello, how are you doing today? This is a test! This test is good.")
	for word, count := range wordFrequency {
		fmt.Printf("word: %s, frequency: %d\n", word, count)
	}
}
