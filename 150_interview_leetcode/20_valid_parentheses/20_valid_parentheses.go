package validparentheses

func isValid(s string) bool {
	stack := make([]byte, 0)

	for i := range s {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			top := stack[len(stack)-1]

			if (top == '(' && s[i] == ')') || (top == '{' && s[i] == '}') || (top == '[' && s[i] == ']') {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}

		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
