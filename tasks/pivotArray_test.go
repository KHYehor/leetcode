package tasks

import (
	"slices"
	"testing"
)

func TestPivotArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		pivot    int
		expected []int
	}{
		{"leetcode example 1", []int{9, 12, 5, 10, 14, 3, 10}, 10, []int{9, 5, 3, 10, 10, 12, 14}},
		{"leetcode example 2", []int{-3, 4, 3, 2}, 2, []int{-3, 2, 4, 3}},
		{"single element equal to pivot", []int{5}, 5, []int{5}},
		{"single element less than pivot", []int{1}, 5, []int{1}},
		{"single element greater than pivot", []int{9}, 5, []int{9}},
		{"all equal to pivot", []int{7, 7, 7}, 7, []int{7, 7, 7}},
		{"all less than pivot", []int{3, 1, 2}, 10, []int{3, 1, 2}},
		{"all greater than pivot", []int{30, 10, 20}, 5, []int{30, 10, 20}},
		{"greater elements keep relative order", []int{5, 8, 6, 7}, 5, []int{5, 8, 6, 7}},
		{"less elements keep relative order", []int{4, 1, 3, 2, 5}, 5, []int{4, 1, 3, 2, 5}},
		{"pivot at the end", []int{6, 2, 8, 1, 4}, 4, []int{2, 1, 4, 6, 8}},
		{"pivot at the start", []int{4, 6, 2, 8, 1}, 4, []int{2, 1, 4, 6, 8}},
		{"multiple pivots scattered", []int{3, 0, 3, 5, 3, 1}, 3, []int{0, 1, 3, 3, 3, 5}},
		{"negative numbers", []int{-1, -5, 0, -3, 2}, -3, []int{-5, -3, -1, 0, 2}},
		{"no greater elements", []int{1, 2, 3, 3}, 3, []int{1, 2, 3, 3}},
		{"no less elements", []int{3, 3, 4, 5}, 3, []int{3, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.nums)
			result := pivotArray(input, tt.pivot)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("pivotArray(%v, %d) = %v, want %v", tt.nums, tt.pivot, result, tt.expected)
			}
		})
	}
}
