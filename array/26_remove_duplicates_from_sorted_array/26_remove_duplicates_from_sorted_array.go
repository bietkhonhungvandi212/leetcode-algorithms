package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	p := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[p] {
			p++
			nums[p] = nums[i]
		}
	}

	return p + 1
}
