package tasks

import "slices"

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	if m == 0 {
		return
	}

	// Transpose
	for i := 0; i < n; i++ {
		for j := i + 1; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < n; i++ {
		slices.Reverse(matrix[i])
	}
}
