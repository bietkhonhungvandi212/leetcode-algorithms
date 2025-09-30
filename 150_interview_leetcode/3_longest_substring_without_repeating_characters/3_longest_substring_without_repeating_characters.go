package longestsubstringwithoutrepeatingcharacters

import "strings"

func lengthOfLongestSubstring(s string) int {

	l, r, maxlen := 0, 1, 0
	if len(s) <= 1 {
		return len(s)
	}
	for r < len(s) {
		str := s[l:r]
		newChar := s[r]
		for strings.Contains(str, string(newChar)) {
			l++
			str = s[l:r]
		}
		maxlen = max(maxlen, r-l+1)
		r = r + 1
	}
	return maxlen
}
