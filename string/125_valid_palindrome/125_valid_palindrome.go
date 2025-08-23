package validpalindrome

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if !isAlphaNumeric(s[right]) {
			right--
			continue
		}

		if !isAlphaNumeric(s[left]) {
			left++
			continue
		}

		if toLowerCase(s[right]) != toLowerCase(s[left]) {
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
