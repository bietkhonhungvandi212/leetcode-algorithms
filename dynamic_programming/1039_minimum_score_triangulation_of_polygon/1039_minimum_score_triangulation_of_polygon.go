package minimumscoretriangulationofpolygon

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for len := 3; len <= n; len++ {
		for i := 0; i < n-len+1; i++ {
			j := i + len - 1
			min := math.MaxInt32
			for k := i + 1; k < j; k++ {
				cost := values[i]*values[j]*values[k] + dp[i][k] + dp[k][j]

				if cost < min {
					min = cost
				}
			}

			dp[i][j] = min
		}
	}

	return dp[0][n-1]
}
