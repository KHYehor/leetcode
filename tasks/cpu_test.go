package tasks

import "testing"

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []byte
		n        int
		expected int
	}{
		{"empty slice", []byte{}, 1, 0},
		{"empty slice", []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{"empty slice", []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K'}, 7, 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leastInterval(tt.tasks, tt.n)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
