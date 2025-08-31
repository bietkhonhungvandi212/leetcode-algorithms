package sudokusolver

func solveSudoku(board [][]byte) {
	rows := make([][10]bool, 9)
	cols := make([][10]bool, 9)
	boxes := make([][10]bool, 9)

	for i := range 9 {
		for j := range 9 {
			if board[i][j] != '.' {
				digit := board[i][j] - '0'
				boxIndex := (i/3)*3 + j/3
				rows[i][digit] = true
				cols[j][digit] = true
				boxes[boxIndex][digit] = true
			}
		}
	}

	var backtracking func(i, j int) bool
	backtracking = func(i, j int) bool {
		if i == 9 {
			return true
		}

		if j == 9 {
			return backtracking(i+1, 0)
		}

		if board[i][j] != '.' {
			return backtracking(i, j+1)
		}

		for digit := byte(1); digit <= 9; digit++ {
			boxIndex := (i/3)*3 + j/3
			if !rows[i][digit] && !cols[j][digit] && !boxes[boxIndex][digit] {
				board[i][j] = digit + '0'
				rows[i][digit] = true
				cols[j][digit] = true
				boxes[boxIndex][digit] = true

				if backtracking(i, j+1) {
					return true
				}

				board[i][j] = '.'
				rows[i][digit] = false
				cols[j][digit] = false
				boxes[boxIndex][digit] = false
			}
		}

		return false
	}

	backtracking(0, 0)
}
