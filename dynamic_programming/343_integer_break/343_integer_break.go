package integerbreak

func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[1] = 1
	for i := 2; i <= n; i++ {
		max := 0
		for j := 1; j < i; j++ {
			product1 := j * (i - j)
			product2 := j * dp[i-j]

			product := product1
			if product2 > product1 {
				product = product2
			}

			if max < product {
				max = product
			}
		}

		dp[i] = max
	}

	return dp[n]
}
