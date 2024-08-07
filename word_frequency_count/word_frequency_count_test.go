package main

import (
	"reflect"
	"testing"
)

func TestWordFrequencyCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"Hello, hello!", map[string]int{"hello": 2}},
		{"How are you? Are you okay?", map[string]int{"how": 1, "are": 2, "you": 2, "okay": 1}},
		{"This is a test. This test is good.", map[string]int{"this": 2, "is": 2, "a": 1, "test": 2, "good": 1}},
	}

	for _, tt := range tests {
		result := wordFrequencyCount(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("wordFrequencyCount(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
