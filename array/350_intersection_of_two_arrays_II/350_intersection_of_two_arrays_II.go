package intersectionoftwoarraysii

func intersect(nums1 []int, nums2 []int) []int {
	frequency := make(map[int]int, len(nums1))
	result := make([]int, 0)
	for _, val := range nums1 {
		frequency[val]++
	}

	for _, val := range nums2 {
		if count, ok := frequency[val]; ok && count != 0 {
			result = append(result, val)
			frequency[val]--
		}
	}

	return result
}
