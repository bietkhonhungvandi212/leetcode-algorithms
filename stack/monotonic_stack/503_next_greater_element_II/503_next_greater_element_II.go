package nextgreaterelementii

func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = -1
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, nums[i])
	}

	stack = stack[:0]
	for i := range nums {
		if result[i] != -1 {
			continue
		}

		for j := i; j >= 0; j-- {
			for len(stack) > 0 && nums[j] >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				result[j] = stack[len(stack)-1]
			}

			stack = append(stack, nums[j])
		}
	}

	return result
}

func nextGreaterElements_OptimizedVersion(nums []int) []int {
	n := len(nums)
	stack := make([]int, 0, 2*n)
	res := make([]int, n)

	for i := range res {
		res[i] = -1
	}

	for i := 2*n - 1; i >= 0; i-- {
		currElement := nums[i%n]

		for len(stack) != 0 && stack[len(stack)-1] <= currElement {
			stack = stack[:len(stack)-1]
		}

		if i < n && len(stack) != 0 {
			res[i] = stack[len(stack)-1]
		}

		stack = append(stack, currElement)
	}

	return res

}
