package countsquaresubmatriceswithallones

func countSquares(matrix [][]int) int {
	var res int

	row, col := len(matrix), len(matrix[0])

	for i := range row {
		for j := range col {
			if matrix[i][j] == 1 && i > 0 && j > 0 {
				matrix[i][j] = 1 + minOf3(matrix[i][j-1], matrix[i-1][j-1], matrix[i-1][j])
			}
			res += matrix[i][j]
		}
	}

	return res
}

func minOf3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
