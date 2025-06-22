package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	mapIndex := make(map[int]int, n1)
	result := make([]int, n1)
	stack := make([]int, 0)

	for i := range nums1 {
		mapIndex[nums1[i]] = i
		result[i] = -1
	}

	for i := n2 - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			if idx, ok := mapIndex[nums2[i]]; ok {
				result[idx] = stack[len(stack)-1]
			}
		}

		stack = append(stack, nums2[i])
	}

	return result
}
