package recursive

import "testing"

func TestRecursiveFactorial(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"base case", 1, 1},
		{"two", 2, 2},
		{"three", 3, 6},
		{"five", 5, 120},
		{"ten", 10, 3628800},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Factorial(tt.num)
			if got != tt.want {
				t.Errorf("RecursiveFactorial(%d) = %d, want %d", tt.num, got, tt.want)
			}
		})
	}
}

// Documents a current limitation rather than exercising it: the only base
// case is num == 1, so num <= 0 never hits it and recurses forever through
// negative numbers, ending in an unrecoverable stack overflow rather than a
// panic. That crashes the whole test binary, so it can't be asserted with
// recover() and is intentionally not run here - RecursiveFactorial should
// only be called with num >= 1 as it stands.
