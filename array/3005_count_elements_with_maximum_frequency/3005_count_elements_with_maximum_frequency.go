package countelementswithmaximumfrequency

func maxFrequencyElements(nums []int) int {
	freq := make(map[int]int)
	result, max := 0, 0

	for _, num := range nums {
		freq[num]++
		if freq[num] > max {
			max = freq[num]
			result = 1
		} else if freq[num] == max {
			result++
		}
	}

	return result * max
}
