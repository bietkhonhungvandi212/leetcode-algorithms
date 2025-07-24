package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	right := 0
	mapper := make(map[int]int)

	for right < n {
		// 1
		// 2
		if val, ok := mapper[nums[right]]; ok && right-val <= k {
			return true
		}
		mapper[nums[right]] = right
		right++

	}

	return false
}
