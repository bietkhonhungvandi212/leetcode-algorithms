package plusone

func plusOne(digits []int) []int {
	n := len(digits) - 1
	for n >= 0 {
		digits[n] += 1
		if digits[n] < 10 {
			return digits
		}

		digits[n] = 0
		n--
	}

	tmp := make([]int, len(digits)+1)
	copy(tmp[1:], digits)
	tmp[0] = 1
	return tmp
}
