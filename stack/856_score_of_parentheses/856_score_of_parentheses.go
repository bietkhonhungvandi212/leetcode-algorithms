package scoreofparentheses

func scoreOfParentheses(s string) int {
	stack := make([]byte, 0)
	score := 0

	nestedCount := 1
	isNested := false
	for i := range s {
		// nested = 0 && len of stack = 1
		if s[i] == '(' {
			if len(stack) > 0 {
				nestedCount *= 2
				isNested = false
			}
			stack = append(stack, s[i])
		} else {
			if nestedCount == 1 {
				if isNested == true {
					isNested = false
					stack = stack[:len(stack)-1]
					continue
				}
				score += 1
			} else if nestedCount > 1 {
				if isNested != true {
					score += nestedCount
				}
				nestedCount /= 2
				isNested = true
			}
			stack = stack[:len(stack)-1]
		}
	}

	return score
}
