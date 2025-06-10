package maximumdifferencebetweenevenandoddfrequencyi

import "math"

func maxDifference(s string) int {
	characters := make([]int, 26)
	maxOdd := 0
	minEven := math.MaxInt

	for i := range s {
		characters[s[i]-'a']++
	}

	for i := range characters {
		if characters[i] == 0 {
			continue
		}

		if characters[i]%2 == 0 && minEven > characters[i] {
			minEven = characters[i]
		}

		if characters[i]%2 != 0 && maxOdd < characters[i] {
			maxOdd = characters[i]
		}
	}

	return maxOdd - minEven
}
