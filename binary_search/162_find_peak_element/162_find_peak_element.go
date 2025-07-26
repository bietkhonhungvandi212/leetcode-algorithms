package findpeakelement

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	// left = 0
	// right = 1
	for left <= right {

		// mid = 0
		mid := (left + right) / 2

		if mid+1 < n && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else if mid-1 >= 0 && nums[mid] < nums[mid-1] {
			right = mid - 1
		} else {
			return mid
		}

	}

	return left
}
