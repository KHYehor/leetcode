package tasks

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"sorted array", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted array", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"unsorted array", []int{3, 1, 4, 5, 2}, []int{1, 2, 3, 4, 5}},
		{"array with duplicates", []int{3, 1, 2, 3, 1}, []int{1, 1, 2, 3, 3}},
		{"empty array", []int{}, []int{}},
		{"single element array", []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.input)
			for i, v := range tt.input {
				if v != tt.expected[i] {
					t.Errorf("got %v, want %v", tt.input, tt.expected)
					break
				}
			}
		})
	}
}

func TestQuickSortLargeArray(t *testing.T) {
	largeArray := make([]int, 10000000)
	for i := range largeArray {
		largeArray[i] = rand.Intn(10000000)
	}

	copyArray := make([]int, len(largeArray))
	copy(copyArray, largeArray)

	t.Run("single-threaded", func(t *testing.T) {
		start := time.Now()
		QuickSort(largeArray)
		t.Logf("Single-threaded sort took %v", time.Since(start))
	})

	t.Run("multi-threaded", func(t *testing.T) {
		start := time.Now()
		QuickSortMT(copyArray)
		t.Logf("Multi-threaded sort took %v", time.Since(start))
	})
}
