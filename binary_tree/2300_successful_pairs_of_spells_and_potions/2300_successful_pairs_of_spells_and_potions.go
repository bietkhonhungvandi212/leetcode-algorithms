package successfulpairsofspellsandpotions

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	n1 := len(spells)
	n2 := len(potions)
	pairs := make([]int, n1)
	sort.Ints(potions)

	for i := range n1 {
		left := 0
		right := n2 - 1
		for left <= right {
			mid := (left + right) / 2

			if int64(spells[i]*potions[mid]) >= success {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		pairs[i] = n2 - left
	}

	return pairs
}
