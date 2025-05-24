package zeroarraytransformationi

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)

	for i := range queries {
		diff[queries[i][0]]++
		diff[queries[i][1]+1]--
	}

	for i := 1; i < n+1; i++ {
		diff[i] += diff[i-1]
	}

	for i := range nums {
		if diff[i] < nums[i] {
			return false
		}
	}

	return true
}
