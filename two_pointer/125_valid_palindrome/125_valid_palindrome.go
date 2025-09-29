package validpalindrome

func isPalindrome(s string) bool {
	n := len(s)
	left := 0
	right := n - 1

	for left < right {
		if left < n && !isAlphaNumeric(s[left]) {
			left++
			continue
		}

		if right > 0 && !isAlphaNumeric(s[right]) {
			right--
			continue
		}

		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLowerCase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}
