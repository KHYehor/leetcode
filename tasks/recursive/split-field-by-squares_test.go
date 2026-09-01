package recursive

import "testing"

// SplitFieldBySquares is meant to be called with size=0 on the first
// invocation; it then recursively narrows size the way the Euclidean
// algorithm narrows its remainder, until it lands on the greatest
// common divisor of width and height (the side of the largest square
// that tiles the width x height field evenly).
func TestSplitFieldBySquares(t *testing.T) {
	tests := []struct {
		name          string
		width, height int
		want          int
	}{
		{"already a perfect square", 6, 6, 6},
		{"simple rectangle", 8, 12, 4},
		{"height is a multiple of width", 3, 9, 3},
		{"width is a multiple of height", 9, 3, 3},
		{"coprime dimensions", 5, 17, 1},
		{"coprime dimensions, width larger", 17, 5, 1},
		{"unit width", 1, 7, 1},
		{"unit height", 7, 1, 1},
		{"1x1 field", 1, 1, 1},
		{"classic euclidean example", 1071, 462, 21},
		{"classic euclidean example, swapped", 462, 1071, 21},
		{"zero width", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitFieldBySquares(tt.width, tt.height, 0)
			if got != tt.want {
				t.Errorf("SplitFieldBySquares(%d, %d, 0) = %d, want %d", tt.width, tt.height, got, tt.want)
			}
		})
	}
}

// When size is already non-zero, the function skips the "size = height"
// initialization and treats the caller-supplied size as the current
// candidate square. This is how the function drives its own recursion,
// but it means calling it directly with an arbitrary non-zero size can
// return a divisor smaller than the true GCD of width and height.
func TestSplitFieldBySquaresNonZeroStartingSize(t *testing.T) {
	tests := []struct {
		name                string
		width, height, size int
		want                int
	}{
		{"starting size equal to height matches default init", 8, 12, 12, 4},
		{"starting size already a valid divisor", 8, 12, 8, 4},
		{"arbitrary starting size can undershoot the true GCD", 8, 12, 6, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitFieldBySquares(tt.width, tt.height, tt.size)
			if got != tt.want {
				t.Errorf("SplitFieldBySquares(%d, %d, %d) = %d, want %d", tt.width, tt.height, tt.size, got, tt.want)
			}
		})
	}
}

// Documents current behavior: with size=0 and height=0, the "size = height"
// initialization leaves size at 0, and the subsequent "width % size" panics
// with a division by zero instead of returning a sentinel value.
func TestSplitFieldBySquaresPanicsOnZeroHeightWithZeroSize(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic (division by zero) when height=0 and size=0, got none")
		}
	}()

	SplitFieldBySquares(5, 0, 0)
}
