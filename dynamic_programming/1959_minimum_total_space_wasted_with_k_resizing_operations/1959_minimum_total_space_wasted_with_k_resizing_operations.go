package minimumtotalspacewastedwithkresizingoperations

import "math"

func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var minWasted func(i, k int) int
	minWasted = func(i, k int) int {
		if i == n {
			return 0
		}
		if k < 0 {
			return math.MaxInt
		}

		if dp[i][k] != -1 {
			return dp[i][k]
		}

		totalSum := 0
		maxNum := nums[i]
		minCost := math.MaxInt

		for j := i; j < n; j++ {
			maxNum = max(maxNum, nums[j])
			totalSum += nums[j]
			wasted := maxNum*(j-i+1) - totalSum

			if j == n-1 {
				minCost = min(minCost, wasted)
			} else {
				if k > 0 {
					minCost = min(minCost, minWasted(j+1, k-1)+wasted)
				}
			}
		}

		dp[i][k] = minCost
		return minCost
	}

	return minWasted(0, k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
