package partitionarrayformaximumsum

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxValue := arr[i-1]
		dp[i] = dp[i-1] + maxValue

		for j := 2; j <= min(i, k); j++ {
			maxValue = max(maxValue, arr[i-j])
			dp[i] = max(dp[i], dp[i-j]+maxValue*j)
		}

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
