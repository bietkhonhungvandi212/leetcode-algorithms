package longestvalidparentheses

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	maxLen := 0
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				top := stack[len(stack)-1]
				maxLen = Max(maxLen, i-top)
			}
		}
	}

	return maxLen

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
