package minimumscoretriangulationofpolygon

import "math"

func minScoreTriangulationV1(values []int) int {
	n := len(values)
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length <= n; length++ { // 2
		for i := 0; i < n-length; i++ { // 0 < (n-2) = 1 -> true
			endVertex := i + length
			dp[i][endVertex] = math.MaxInt       // dp[0][2]
			for j := i + 1; j < endVertex; j++ { // j = 1; j <= 2
				cost := values[i]*values[j]*values[endVertex] + dp[i][j] + dp[j][endVertex]
				//1 *  2 * 3 = 6
				if cost < dp[i][endVertex] {
					dp[i][endVertex] = cost
				}
			}
		}
	}

	return dp[0][n-1]

}

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
