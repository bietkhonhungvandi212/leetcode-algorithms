package makearrayelementsequaltozero

func countValidSelections(nums []int) int {
	n := len(nums)
	count := 0
	if n == 1 {
		return 2
	}
	prefix := make([]int, n+1)

	for i := range nums {
		prefix[i+1] = prefix[i] + nums[i]
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			left := prefix[i+1] - prefix[0]
			right := prefix[n] - prefix[i]

			if right-left == 0 {
				count += 2
			} else if abs(right-left) == 1 {
				count++
			}

		}
	}

	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
