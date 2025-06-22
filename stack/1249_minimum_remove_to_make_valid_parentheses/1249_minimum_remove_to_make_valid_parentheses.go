package minimumremovetomakevalidparentheses

import "strings"

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

func minRemoveToMakeValid_Optimization(s string) string {
	bytes := []byte(s)
	openCount := 0

	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '(' {
			openCount++
		} else if bytes[i] == ')' {
			if openCount == 0 {
				bytes[i] = '*' // Mark for removal
			} else {
				openCount--
			}
		}
	}

	for i := len(bytes) - 1; i >= 0; i-- {
		if openCount > 0 && bytes[i] == '(' {
			bytes[i] = '*' // Mark for removal
			openCount--
		}
	}

	var result strings.Builder
	for _, b := range bytes {
		if b != '*' {
			result.WriteByte(b)
		}
	}

	return result.String()
}
