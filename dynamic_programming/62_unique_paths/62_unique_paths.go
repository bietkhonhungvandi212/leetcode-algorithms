package uniquepaths

func uniquePaths_TopDown(m int, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	var countPath func(down, right int) int
	countPath = func(down, right int) int {
		if down == m-1 && right == n-1 {
			return 1
		}

		if down >= m || right >= n {
			return 0
		}

		if dp[down][right] != 0 {
			return dp[down][right]
		}

		dp[down][right] = countPath(down+1, right) + countPath(down, right+1)

		return dp[down][right]
	}

	return countPath(0, 0)
}

func uniquePaths_Bottomup(m int, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range m {
		dp[i][0] = 1
	}

	for i := range n {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
