package convertintegertothesumoftwonozerointegers

func getNoZeroIntegers(n int) []int {
	num := n / 2
	var a, b int

	for i := 1; i <= num; i++ {
		if isNoZero(i) && isNoZero(n-i) {
			a = i
			b = n - i
			break
		}
	}

	return []int{a, b}
}

func isNoZero(n int) bool {
	tmp := n
	for tmp > 0 {
		mod := tmp % 10
		if mod == 0 {
			return false
		}
		tmp /= 10
	}

	return true
}
