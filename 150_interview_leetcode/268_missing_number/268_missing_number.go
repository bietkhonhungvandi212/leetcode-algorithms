package missingnumber

func missingNumber(nums []int) int {
	length := len(nums)
	sum1 := (length + 1) * length / 2
	sum2 := 0
	for _, val := range nums {
		sum2 += val
	}

	return sum1 - sum2
}
