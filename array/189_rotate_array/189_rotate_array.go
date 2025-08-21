package rotatearray

func rotate(nums []int, k int) {
	n := len(nums)
	times := k % n

	if times == 0 {
		return
	}

	tmp := make([]int, times)

	copy(tmp, nums[n-times:])
	copy(nums[times:], nums[:n-times])
	copy(nums[:times], tmp)
}
