package recursive

import "testing"

func TestRecursiveSum(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"single element", []int{7}, 7},
		{"two elements", []int{1, 2}, 3},
		{"several positive elements", []int{1, 2, 3, 4, 5}, 15},
		{"all negative elements", []int{-3, -2, -1}, -6},
		{"mixed positive and negative", []int{5, -3, 2}, 4},
		{"elements summing to zero", []int{4, -4}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.arr)
			if got != tt.want {
				t.Errorf("RecursiveSum(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}

// Documents current behavior: the only base case is len(arr) == 1, so an
// empty (or nil) slice falls through to "arr[0] + ...", which panics with
// an index-out-of-range error instead of returning 0. Unlike the
// zero/negative case in Factorial, this is a regular panic and can
// safely be recovered in a test.
func TestRecursiveSumPanicsOnEmptySlice(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"empty slice", []int{}},
		{"nil slice", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Fatal("expected panic (index out of range) for a slice with no elements, got none")
				}
			}()

			Sum(tt.arr)
		})
	}
}
