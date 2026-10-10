package tasks

import (
	"testing"
)

func TestFirstUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"package example", textString, 6},
		{"empty string", "", -1},
		{"single char", "a", 0},
		{"first char is unique", "abb", 0},
		{"last char is unique", "aab", 2},
		{"unique in the middle", "aabcc", 2},
		{"all repeated", "aabb", -1},
		{"repeated non-adjacent", "abab", -1},
		{"several uniques, earliest wins", "zzxyw", 2},
		{"char repeated many times", "aaaab", 4},
		{"leetcode", "leetcode", 0},
		{"loveleetcode", "loveleetcode", 2},
		{"case sensitive", "aA", 0},
		{"multibyte returns byte index", "ééx", 4}, // 'é' is 2 bytes
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := firstUnique(tt.input)
			if result != tt.expected {
				t.Errorf("firstUnique(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
