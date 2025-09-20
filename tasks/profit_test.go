package tasks

import "testing"

func TestBestProfit(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect [2]int
	}{
		{
			name:   "nil slice",
			input:  nil,
			expect: [2]int{-1, -1},
		},
		{
			name:   "single price",
			input:  []int{5},
			expect: [2]int{-1, -1},
		},
		{
			name:   "descending prices - no profit",
			input:  []int{9, 8, 7, 6},
			expect: [2]int{-1, -1},
		},
		{
			name:   "ascending prices - buy first sell last",
			input:  []int{1, 2, 3, 4},
			expect: [2]int{0, 3},
		},
		{
			name:   "fluctuating prices - classic example",
			input:  []int{7, 1, 5, 3, 6, 4},
			expect: [2]int{1, 4}, // buy at 1 (price=1), sell at 4 (price=6)
		},
		{
			name:   "tie profits keep earliest pair",
			input:  []int{3, 1, 2, 1, 2},
			expect: [2]int{1, 2}, // profit=1 appears twice; function keeps the first best
		},
		{
			name:   "minimal price late, large jump",
			input:  []int{5, 4, 3, 2, 10},
			expect: [2]int{3, 4},
		},
		{
			name:   "all equal prices - no profit",
			input:  []int{2, 2, 2},
			expect: [2]int{-1, -1},
		},
		{
			name:   "same minimal price occurs later but earlier index kept",
			input:  []int{5, 2, 4, 2, 6},
			expect: [2]int{1, 4}, // minimal price repeats at index 3, but earliest minimal kept
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := BestProfit(tc.input)
			if got != tc.expect {
				t.Fatalf("BestProfit(%v) = %v; want %v", tc.input, got, tc.expect)
			}
		})
	}
}
