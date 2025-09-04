package longestpalindromicsubstring

func longestPalindrome_dp(s string) string {
	n := len(s)
	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
	}

	max := 0
	start := 0
	end := 0

	for i := range n {
		dp[i][i] = true
		for j := range i {
			if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1] == true) {
				dp[j][i] = true
				length := i - j + 1
				if length > max {
					max = length
					start = j
					end = i
				}
			}
		}
	}

	return s[start : end+1]
}

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
		}
	}

	return s[left : right+1]
}
