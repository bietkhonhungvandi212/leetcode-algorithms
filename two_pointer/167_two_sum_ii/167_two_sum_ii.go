package twosumii

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	if n == 2 {
		return []int{1, 2}
	}

	left := 0
	right := n - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}

		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
