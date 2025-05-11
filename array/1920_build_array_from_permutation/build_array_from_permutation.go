package buildarrayfrompermutation

func buildArray(nums []int) []int {
	for i := range nums {
		nums[i] += 1000 * (nums[nums[i]] % 1000)
	}

	for i := range nums {
		nums[i] = nums[i] / 1000
	}

	return nums
}
