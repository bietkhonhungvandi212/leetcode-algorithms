package fractiontorecurringdecimal

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return "0"
	}

	var str []byte
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		str = append(str, '-')
	}

	numerator = abs(numerator)
	denominator = abs(denominator)

	integer := numerator / denominator
	str = append(str, []byte(strconv.Itoa(integer))...)

	mod := numerator % denominator
	if mod == 0 {
		return string(str)
	}

	str = append(str, '.')
	freq := map[int]int{}
	repeatIdx := len(str)

	for mod != 0 {
		if idx, ok := freq[mod]; ok {
			return string(str[:idx-1]) + "(" + string(str[idx-1:]) + ")"
		}
		freq[mod] = repeatIdx

		mod *= 10
		digit := mod / denominator
		str = append(str, []byte(strconv.Itoa(digit))...)
		mod %= denominator
		repeatIdx++
	}

	return string(str)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
