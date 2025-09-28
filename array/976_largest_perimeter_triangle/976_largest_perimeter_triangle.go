package largestperimetertriangle

import "sort"

func largestPerimeter(nums []int) int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < n-2; i++ {
		if nums[i] < nums[i+1]+nums[i+2] {
			return nums[i+1] + nums[i+2] + nums[i]
		}
	}

	return 0
}
