package removeelement

func removeElement(nums []int, val int) int {
	// Not equal val
	p := 0
	for i := range nums {
		if nums[i] != val {
			nums[p] = nums[i]
			p++
		}
	}

	return p
}
