package stonegame

import "math"

func stoneGame_TopDown(piles []int) bool {
	n := len(piles)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = math.MinInt
		}
	}

	var miniMax func(left, right int) int
	miniMax = func(left, right int) int {
		if left > right {
			return 0
		}

		if memo[left][right] != math.MinInt {
			return memo[left][right]
		}

		takeLeft := piles[left] - miniMax(left+1, right)
		takeRight := piles[right] - miniMax(left, right-1)

		memo[left][right] = max(takeLeft, takeRight)
		return memo[left][right]
	}

	scores := miniMax(0, len(piles)-1)

	return scores > 0
}

func stoneGame_BottomUp(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			takeLeft := piles[i] - dp[i+1][j]
			takeRight := piles[j] - dp[i][j-1]

			dp[i][j] = max(takeLeft, takeRight)
		}
	}

	return dp[0][n-1] > 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
