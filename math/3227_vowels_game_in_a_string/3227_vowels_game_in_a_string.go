package vowelsgameinastring

func doesAliceWin(s string) bool {
	for i := range s {
		if isVowel(s[i]) {
			return true
		}
	}

	return false
}

func isVowel(b byte) bool {
	if b == 'u' || b == 'a' || b == 'e' || b == 'o' || b == 'i' {
		return true
	}

	return false
}
