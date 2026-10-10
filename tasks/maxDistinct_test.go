package tasks

import (
	"testing"
)

func TestMaxDistinct(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"one unique", "aaaa", 1},
		{"two unique", "abab", 2},
		{"three unique", "abab", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxDistinct(tt.input)
			if result != tt.expected {
				t.Errorf("firstUnique(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
