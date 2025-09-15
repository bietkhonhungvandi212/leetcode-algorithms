package maximumnumberofwordsyoucantype

func canBeTypedWords(text string, brokenLetters string) int {
	brokens := make([]byte, 26)

	for i := range brokenLetters {
		brokens[brokenLetters[i]-'a']++
	}

	count := 0
	isTyped := true
	for i := range text {
		if text[i] == ' ' {
			if isTyped {
				count++
			}
			isTyped = true

			continue
		}

		if brokens[text[i]-'a'] != 0 {
			isTyped = false
		}
	}

	if isTyped {
		count++
	}
	return count

}
