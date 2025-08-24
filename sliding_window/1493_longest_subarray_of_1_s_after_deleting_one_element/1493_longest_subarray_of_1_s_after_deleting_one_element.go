package longestsubarrayof1safterdeletingoneelement

func longestSubarray(nums []int) int {
	//  1  1 0 1
	n := len(nums)
	countZero := 0
	max := 0
	left := 0
	right := 0

	for right < n {
		if nums[right] == 0 {
			countZero++
		}

		for countZero > 1 {
			if nums[left] == 0 {
				countZero--
			}
			left++
		}

		length := right - left

		if max < length {
			max = length
		}
		right++
	}

	if countZero == 0 {
		return n - 1
	}

	return max
}
