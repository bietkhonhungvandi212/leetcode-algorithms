package validanagram

func isAnagram(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	frequency := make([]int, 26)

	for _, char := range s {
		frequency[char-'a']++
	}

	for _, char := range t {
		if frequency[char-'a'] == 0 {
			return false
		}
		frequency[char-'a']--
	}

	return true
}
