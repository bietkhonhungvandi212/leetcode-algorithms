package diagonaltraverse

func findDiagonalOrder(mat [][]int) []int {
	rows := len(mat)
	cols := len(mat[0])
	rowD, colD := 0, 0
	isUp := true
	result := make([]int, 0, rows*cols)

	for len(result) < rows*cols {
		result = append(result, mat[rowD][colD])

		if isUp {
			if colD == cols-1 {
				rowD++
				isUp = false
			} else if rowD == 0 {
				colD++
				isUp = false
			} else {
				rowD--
				colD++
			}
		} else {
			if rowD == rows-1 {
				colD++
				isUp = true
			} else if colD == 0 {
				rowD++
				isUp = true
			} else {
				rowD++
				colD--
			}
		}
	}

	return result
}
