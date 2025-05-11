package partitionequalsubsetsum

func canPartition(nums []int) bool {
	sum := 0

	for i := range nums {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dpArr := make([]int, target+1)
	dpArr[0] = 1
	for _, num := range nums {
		for currSum := target; currSum >= num; currSum-- {
			if dpArr[currSum] == 1 {
				continue
			}

			if dpArr[currSum-num] == 1 {
				dpArr[currSum] = 1
			}

			if dpArr[target] == 1 {
				return true
			}
		}
	}

	return false
}
