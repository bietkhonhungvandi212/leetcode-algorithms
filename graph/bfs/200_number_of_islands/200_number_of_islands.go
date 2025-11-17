package numberofislands

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	dir := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	cnt := 0

	bfs := func(row, col int) {
		queue := [][2]int{}
		queue = append(queue, [2]int{row, col})
		for len(queue) > 0 {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			for i := range dir {
				nr, nc := r+dir[i][0], c+dir[i][1]

				if nr >= 0 && nc >= 0 && nr < rows && nc < cols && grid[nr][nc] == '1' {
					grid[nr][nc] = '0'
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}

	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				bfs(i, j)
				cnt++
			}
		}
	}

	return cnt
}
