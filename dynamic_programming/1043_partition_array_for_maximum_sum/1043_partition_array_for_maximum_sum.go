package partitionarrayformaximumsum

import "math"

func maxSumAfterPartitioning_TopDown(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}

	var maxSumPartition func(index int) int
	maxSumPartition = func(index int) int {
		if index < 0 {
			return 0
		}

		if dp[index] != -1 {
			return dp[index]
		}

		maxSum := math.MinInt
		currMax := arr[index]
		for i := range min(index+1, k) {
			currMax = max(arr[index-i], currMax)
			cost := currMax*(i+1) + maxSumPartition(index-i-1)
			if maxSum < cost {
				maxSum = cost
			}
		}

		dp[index] = maxSum
		return dp[index]
	}

	return maxSumPartition(n - 1)
}

func maxSumAfterPartitioning_BottomUp(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxSum := math.MinInt
		currMax := math.MinInt
		for j := 1; j <= min(i, k); j++ {
			currMax = max(currMax, arr[i-j])
			cost := currMax*j + dp[i-j]
			if cost > maxSum {
				maxSum = cost
			}
		}
		dp[i] = maxSum

	}

	return dp[n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}
