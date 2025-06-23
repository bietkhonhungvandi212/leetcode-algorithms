package removeduplicateletters

func removeDuplicateLetters(s string) string {
	lastOccur := make(map[byte]int, 0)

	stack := make([]byte, 0)
	distinctChar := make(map[byte]bool, 0)

	for i := range s {
		lastOccur[s[i]] = i
	}

	for i := range s {
		if distinctChar[s[i]] == true {
			continue
		}

		for len(stack) > 0 && s[i] < stack[len(stack)-1] && i < lastOccur[stack[len(stack)-1]] {
			removed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			distinctChar[removed] = false
		}

		distinctChar[s[i]] = true
		stack = append(stack, s[i])
	}

	return string(stack)
}
