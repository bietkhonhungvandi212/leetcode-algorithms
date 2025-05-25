package numberofsubarrayswithboundedmaximum

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	return count(nums, right) - count(nums, left-1)

}

func count(nums []int, bound int) int {
	cnt := 0
	answer := 0
	for i := range nums {
		if nums[i] <= bound {
			cnt += 1
		} else {
			cnt = 0
		}

		answer += cnt
	}

	return answer
}
