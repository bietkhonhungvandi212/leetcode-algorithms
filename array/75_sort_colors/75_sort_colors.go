package sortcolors

func sortColors(nums []int) {
	freq := make([]int, 3)

	for i := range nums {
		freq[nums[i]]++
	}

	index := 0
	for i := range freq {
		for freq[i] != 0 {
			nums[index] = i
			index++
			freq[i]--
		}
	}
}
