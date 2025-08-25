package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currentSum := 0
	minLen := len(nums) + 1

	for left, right := 0, 0; right < len(nums); right++ {
		currentSum += nums[right]

		for currentSum >= target {
			length := right - left + 1
			if length < minLen {
				minLen = length
			}

			currentSum -= nums[left]
			left++
		}

	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
