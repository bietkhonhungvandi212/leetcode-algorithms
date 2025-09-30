package maximumaveragesubarrayi

import "math"

func findMaxAverage(nums []int, k int) float64 {
	maxSum := math.MinInt
	n := len(nums)
	sum := 0

	for left, right := 0, 0; right < n; right++ {
		sum += nums[right]
		if right-left+1 == k {
			maxSum = max(maxSum, sum)
			sum -= nums[left]
			left++
		}
	}

	return float64(maxSum) / float64(k)
}
