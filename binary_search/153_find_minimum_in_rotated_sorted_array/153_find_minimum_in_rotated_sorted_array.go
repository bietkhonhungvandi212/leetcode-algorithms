package findminimuminrotatedsortedarray

func findMin(nums []int) int {
	n := len(nums)

	left := 0
	right := n - 1

	for left < right {

		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}

	}

	return nums[left]
}
