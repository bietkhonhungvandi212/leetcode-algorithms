package numberofpeopleawareofasecret

const mod = 1000000007

func peopleAwareOfSecret(n int, delay int, forget int) int {
	// number newlearners of day xth is dp[x] = sum dp[ith], ith in range [x - forgot + 1, x - delays]
	dp := make([]int, n)
	knows := make([]int, n+1)

	dp[0], knows[1] = 1, 1

	//delay = 2, forget = 4

	// dp[4] ->  dp[1] + dp[2]
	// dp[0] -> forgot; dp[3] not meet delay days

	for d := 2; d <= n; d++ {
		start := max(d-forget+1, 1)

		end := d - delay

		if start <= end {
			dp[d-1] = (knows[end] - knows[start-1] + mod) % mod
		}

		knows[d] = knows[d-1] + dp[d-1]
	}

	start := max(1, n-forget+1)
	result := (knows[n] - knows[start-1] + mod) % mod
	return result
}
