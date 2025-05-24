package findwordscontainingcharacter

func findWordsContaining(words []string, x byte) []int {
	result := make([]int, 0)

	for i := range words {
		for j := range words[i] {
			if words[i][j] == x {
				result = append(result, i)
				break
			}
		}
	}

	return result
}
