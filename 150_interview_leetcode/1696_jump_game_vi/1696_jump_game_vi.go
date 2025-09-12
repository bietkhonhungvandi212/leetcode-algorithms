package jumpgamevi

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	deque := make([]int, 0)
	dp[0] = nums[0]
	deque = append(deque, 0)

	for i := 1; i < n; i++ {
		if len(deque) > 0 && deque[0] < i-k {
			deque = deque[1:]
		}

		dp[i] = nums[i] + dp[deque[0]]
		for len(deque) > 0 && dp[i] > dp[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
	}

	return dp[n-1]
}
