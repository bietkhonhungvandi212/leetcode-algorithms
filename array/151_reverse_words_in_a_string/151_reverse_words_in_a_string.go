package reversewordsinastring

import (
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	var result strings.Builder
	stack := make([]byte, 0)

	j := 0
	for j < n && s[j] == ' ' {
		j++
	}

	for i := n - 1; i >= j; i-- {
		if len(stack) == 0 && s[i] == ' ' {
			continue
		}

		if s[i] != ' ' {
			stack = append(stack, s[i])
		} else {
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				result.WriteByte(top)
				stack = stack[:len(stack)-1]
			}
			result.WriteByte(' ')
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		result.WriteByte(top)
		stack = stack[:len(stack)-1]
	}

	return result.String()
}
