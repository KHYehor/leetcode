package tasks

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "two most frequent values",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "single value",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "negative values",
			nums: []int{-1, -1, -1, -2, -2, -3},
			k:    2,
			want: []int{-1, -2},
		},
		{
			name: "all unique values",
			nums: []int{4, 3, 2, 1},
			k:    4,
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)

			// The problem allows the result in any order.
			slices.Sort(got)
			slices.Sort(tt.want)

			if !slices.Equal(got, tt.want) {
				t.Fatalf("topKFrequent(%v, %d) = %v; want %v (in any order)", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
