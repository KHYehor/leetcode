package two_pointers

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"no duplicates", []int{1, 2, 3, 4, 5}, 5},
		{"all duplicates", []int{1, 1, 1, 1, 1}, 1},
		{"some duplicates", []int{1, 1, 2, 2, 3}, 3},
		{"duplicates at start", []int{1, 1, 1, 2, 3}, 3},
		{"duplicates at end", []int{1, 2, 3, 3, 3}, 3},
		{"example from problem", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicates(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
