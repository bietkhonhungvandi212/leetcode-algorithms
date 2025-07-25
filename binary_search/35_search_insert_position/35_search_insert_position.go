package searchinsertposition

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid
		} else {
			return mid
		}

	}

	return left
}
