package houserobbertopdown

import "math"

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}

	var topDown func(startIndex int) int
	topDown = func(startIndex int) int {
		if startIndex >= n {
			return 0
		}

		if dp[startIndex] != 0 {
			return dp[startIndex]
		}

		max := math.MinInt
		for i := startIndex; i < n; i++ {
			cost := nums[i] + topDown(i+2)
			if max < cost {
				max = cost
			}
		}
		dp[startIndex] = max

		return dp[startIndex]
	}

	return topDown(0)
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
