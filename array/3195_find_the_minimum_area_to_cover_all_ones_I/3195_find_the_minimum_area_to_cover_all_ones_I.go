package findtheminimumareatocoverallonesi

func minimumArea(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	rightMost := -1
	leftMost := cols
	topMost := -1
	bottomMost := -1

	for j := range cols {
		hasOne := false
		for i := range rows {
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

	for i := range rows {
		hasOne := false
		for j := range cols {
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
