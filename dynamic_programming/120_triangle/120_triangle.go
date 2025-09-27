package triangle

import "math"

func minimumTotalBottomUp(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return triangle[0][0]
	}

	for i := n - 2; i >= 0; i-- {
		for j := len(triangle[i]); j >= 0; j-- {
			triangle[i][j] = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]

}

//This is top down approach but not complete
// I first think that sum[currLevel][index] = triangle[level][index] + min(sum[currentLevel + 1][index], sum[currentLevel + 1][index + 1])
// it means that parent problem (higher height of triangle) can depend on child problems (lower height of triangle)
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return triangle[0][0]
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	var backtrack func(prevIdx, currentLevel int) int
	backtrack = func(prevIdx, currentLevel int) int {
		if currentLevel == n {
			return 0
		}

		if dp[currentLevel][prevIdx] != math.MaxInt32 {
			return dp[currentLevel][prevIdx]
		}

		minSum := math.MaxInt32
		for i := prevIdx; i <= min(prevIdx+1, n-1); i++ {
			nextSum := backtrack(i, currentLevel+1)
			minSum = min(minSum, triangle[currentLevel][i]+nextSum)
		}

		dp[currentLevel][prevIdx] = minSum
		return minSum
	}

	return triangle[0][0] + backtrack(0, 1)
}
