package tasks

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive number", 123, 321},
		{"negative number", -123, -321},
		{"number with trailing zeros", 1200, 21},
		{"negative number with trailing zeros", -1200, -21},
		{"zero", 0, 0},
		{"large positive number", 2147483647, 0},  // overflow case
		{"large negative number", -2147483647, 0}, // overflow case
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverse(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
