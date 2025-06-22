package minimumremovetomakevalidparentheses

func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0)

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, i)
			}
		}
	}

	result := make([]byte, 0)
	stackIndex := 0

	for i := range s {
		if stackIndex < len(stack) && i == stack[stackIndex] {
			stackIndex++
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}
