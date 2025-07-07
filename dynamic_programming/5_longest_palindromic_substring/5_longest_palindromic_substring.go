package longestpalindromicsubstring

func longestPalindrome(s string) string {
	left := 0
	right := 0
	var isPalindrome func(s string, j, i int) bool
	isPalindrome = func(s string, j, i int) bool {
		for j < i {
			if s[j] != s[i] {
				return false
			}
			i--
			j++
		}

		return true
	}

	for i := range len(s) {
		j := 0
		for i-j+1 > right-left+1 {
			if isPalindrome(s, j, i) {
				left = j
				right = i
				break
			}
			j++
		}
	}

	return s[left : right+1]
}
