package tasks

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			"one array empty",
			[]int{1, 3},
			[]int{},
			2.0,
		},
		{
			"even total length",
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
		{
			"odd total length",
			[]int{1, 3},
			[]int{2},
			2.0,
		},
		{
			"different lengths",
			[]int{1, 2, 3},
			[]int{4, 5},
			3.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMedianSortedArrays(tt.nums1, tt.nums2)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
