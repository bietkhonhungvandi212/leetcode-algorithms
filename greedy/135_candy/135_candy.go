package candy

func candy(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}
	results := make([]int, n)

	switchIndex := 0
	results[0] = 1
	for i := 1; i < n; i++ {
		results[i] = 1
		if ratings[i] > ratings[i-1] {
			switchIndex = i
			results[i] = results[i-1] + 1
		} else if ratings[i] < ratings[i-1] {
			for j := i - 1; j >= switchIndex; j-- {
				if results[j] <= results[j+1] {
					results[j]++
				}
			}
		} else {
			switchIndex = i
		}
	}

	min := 0
	for i := range results {
		min += results[i]
	}

	return min

}
