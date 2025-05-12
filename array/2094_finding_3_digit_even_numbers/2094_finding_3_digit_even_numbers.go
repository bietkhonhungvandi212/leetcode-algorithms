package finding3digitevennumbers

func findEvenNumbers(digits []int) []int {
	res := make([]int, 0)
	frequency := make(map[int]int)

	for i := range digits {
		frequency[digits[i]]++
	}

	for i := 1; i <= 10; i++ {
		if frequency[i] == 0 {
			continue
		}
		frequency[i]--

		for j := 0; j <= 10; j++ {
			if frequency[j] == 0 {
				continue
			}
			frequency[j]--
			for k := 0; k <= 10; k++ {
				if k%2 != 0 || frequency[k] == 0 {
					continue
				}

				res = append(res, i*100+j*10+k)
			}
			frequency[j]++
		}
		frequency[i]++
	}

	return res
}
