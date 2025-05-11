package largestdivisiblesubset

import "sort"

func largestDivisibleSubset(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}

	sort.Ints(nums)

	dp_arr, prev_arr := make([]int, length), make([]int, length)
	maxLength, maxIndex := 1, 0

	dp_arr[0] = 1
	prev_arr[0] = -1
	for i := 1; i < length; i++ {
		for j := range i {
			dp_arr[i] = 1
			prev_arr[i] = -1

			if nums[i]%nums[j] == 0 && dp_arr[i] < dp_arr[j]+1 {
				dp_arr[i] = 1 + dp_arr[j]
				prev_arr[i] = j
			}
		}

		if maxLength < dp_arr[i] {
			maxLength = dp_arr[i]
			maxIndex = i
		}
	}

	ans := make([]int, 0)
	for maxIndex >= 0 {
		ans = append(ans, nums[maxIndex])
		maxIndex = prev_arr[maxIndex]
	}

	return ans
}
