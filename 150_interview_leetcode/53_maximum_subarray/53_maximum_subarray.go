package maximumsubarray

import "math"

/*
1. see that at each loop the parent problem depending on the subproblems.
2. And the subproblems have attribute that overlap

-> dynamic programing

We should think aloud this problem there are 2 choice that either end at index `i` or start at index `i`.

- If it ends at index `i`, it should including nums[i] from the sum at dp[i-1]: `dp[i] = nums[i] + dp[i-1]`
- if it start at index `i`, the dp[i] = nums[i]
*/

func maxSubArray_bottom_up(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1] + nums[i])

		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}

func maxSubArray(nums []int) int {
	n := len(nums)
	mem := make([]int, n)
	for i := range mem {
		mem[i] = math.MinInt
	}
	var dp func(index int) int
	dp = func(index int) int {
		if mem[index] != math.MinInt {
			return mem[index]
		}
		if index == 0 {
			mem[index] = nums[0]
			return mem[index]
		}
		mem[index] = max(nums[index], nums[index]+dp(index-1))
		return mem[index]
	}

	maxSum := math.MinInt
	for i := 0; i < n; i++ {
		maxSum = max(maxSum, dp(i))
	}
	return maxSum
}
