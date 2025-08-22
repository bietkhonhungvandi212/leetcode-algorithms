package findtheminimumareatocoverallonesi

func minimumArea(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	rightMost := -1
	leftMost := cols
	topMost := -1
	bottomMost := -1

	for j := 0; j < cols; j++ {
		hasOne := false
		for i := 0; i < rows; i++ {
			if grid[i][j] == 1 {
				hasOne = true
				break
			}
		}
		if hasOne {
			if leftMost == cols {
				leftMost = j
			}
			rightMost = j
		}
	}

	for i := 0; i < rows; i++ {
		hasOne := false
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				hasOne = true
				break
			}
		}
		if hasOne {
			if topMost == -1 {
				topMost = i
			}
			bottomMost = i
		}
	}

	if rightMost == -1 || topMost == -1 {
		return 0
	}

	return (rightMost - leftMost + 1) * (bottomMost - topMost + 1)
}
