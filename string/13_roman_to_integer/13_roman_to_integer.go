package romantointeger

func romanToInt(s string) int {
	romanToInt := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	i := 0
	for i = 0; i < len(s)-1; i++ {
		tmp := romanToInt[s[i+1]] - romanToInt[s[i]]
		division := tmp / romanToInt[s[i]]

		if division == 9 || division == 4 {
			sum += tmp
			i = i + 1
			continue
		}

		sum += romanToInt[s[i]]
	}

	if i >= len(s) {
		return sum
	}

	return sum + romanToInt[s[len(s)-1]]
}
