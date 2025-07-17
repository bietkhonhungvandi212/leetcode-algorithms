package houserobbertopdown

func rob_top_down(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}

	var topDown func(startIndex int) int
	topDown = func(startIndex int) int {
		if startIndex >= n {
			return 0
		}

		if dp[startIndex] != -1 {
			return dp[startIndex]
		}

		// Two choices for each house:
		// 1. Rob current house: nums[startIndex] + topDown(startIndex + 2)
		// 2. Skip current house: topDown(startIndex + 1)
		rob := nums[startIndex] + topDown(startIndex+2)
		skip := topDown(startIndex + 1)

		dp[startIndex] = max(rob, skip)
		return dp[startIndex]
	}

	return topDown(0)
}

func rob_bottom_up(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]               // Base case: only house 0
	dp[1] = max(nums[0], nums[1]) // Base case: better of house 0 or 1

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	return dp[n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
