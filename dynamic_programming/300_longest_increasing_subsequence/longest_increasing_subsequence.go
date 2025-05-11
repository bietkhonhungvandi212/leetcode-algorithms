package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	maxLen := 1
	dpArr := make([]int, length)

	dpArr[0] = 1
	for i := 1; i < length; i++ {
		dpArr[i] = 1
		for j := range i {
			if nums[i] > nums[j] && dpArr[i] < dpArr[j]+1 {
				dpArr[i] = 1 + dpArr[j]
			}
		}

		if maxLen < dpArr[i] {
			maxLen = dpArr[i]
		}
	}

	return maxLen
}
