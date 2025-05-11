package numberofequivalentdominopairs

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	mapper := make(map[string]int)

	for i := range dominoes {
		if dominoes[i][0] < dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}

		key := fmt.Sprintf("%d%d", dominoes[i][0], dominoes[i][1])

		_, ok := mapper[key]

		if ok {
			mapper[key]++
		} else {
			mapper[key] = 1
		}
	}

	var result int
	for _, value := range mapper {
		if value > 1 {
			result += value * (value - 1) / 2
		}
	}

	return result
}
