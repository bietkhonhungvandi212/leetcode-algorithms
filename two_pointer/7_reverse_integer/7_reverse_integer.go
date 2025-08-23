package reverseinteger

import "math"

func reverse(x int) int {
	if x == 0 {
		return x
	}

	signedDigit := 1

	if x < 0 {
		signedDigit = -1

	}

	const (
		INT_MAX = math.MaxInt32
		INT_MIN = math.MinInt32
	)

	result := 0
	current := abs(x)
	for current > 0 {
		mod := current % 10
		current /= 10
		if result > INT_MAX/10 || (result == INT_MAX/10 && mod > 7) {
			return 0
		}

		if result < INT_MIN/10 || (result == INT_MIN/10 && mod < -8) {
			return 0
		}
		result = result*10 + mod
	}

	return signedDigit * result
	//1 2 3 / 10 -> 12 3
	// => n * 10 + 3 = 3
	// 12 / 10 -> 1 2
	// => n * 10 + 2 = 30 + 2 = 32
	// 1 / 10 -> 0 1
	// =>
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
