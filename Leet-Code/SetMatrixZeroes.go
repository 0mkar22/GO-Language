package leetcode

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rowzero := make([]bool, m)
	colzero := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowzero[i] = true
				colzero[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowzero[i] || colzero[j] {
				matrix[i][j] = 0
			}
		}
	}

}
