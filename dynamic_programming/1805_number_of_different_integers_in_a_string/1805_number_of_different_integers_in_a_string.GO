package numberofdifferentintegersinastring

func numDifferentIntegers(word string) int {
	frequency := make(map[string]bool)

	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			continue
		}

		start := i
		for i < len(word) && (word[i] < 'a' || word[i] > 'z') {
			i++
		}

		digits := word[start:i]

		for len(digits) > 1 && digits[0] == '0' {
			digits = digits[1:]
		}
		frequency[digits] = true
	}

	return len(frequency)
}
