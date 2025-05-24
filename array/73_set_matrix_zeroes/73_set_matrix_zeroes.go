package setmatrixzeroes

func setZeroes(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	column := len(matrix[0])
	if column == 0 {
		return
	}
	colMap := make(map[int]bool)
	rowMap := make(map[int]bool)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				colMap[j] = true
				rowMap[i] = true

			}
		}
	}

	for i := range row {
		for j := range column {
			if _, col := colMap[j]; col == true {
				matrix[i][j] = 0
			} else if _, row := rowMap[i]; row == true {
				matrix[i][j] = 0

			}
		}
	}
}
