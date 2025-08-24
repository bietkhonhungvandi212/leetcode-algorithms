package maxconsecutiveonesiii

func longestOnes(nums []int, k int) int {
	left := 0
	zeros := 0
	maxLength := 0

	for right, val := range nums {
		if val == 0 {
			zeros++
		}

		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}

			left++
		}
		length := right - left + 1

		if maxLength < length {
			maxLength = length
		}
	}

	return maxLength
}
