package removekdigits

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	if len(num) == 1 {
		return "0"
	}

	for i := range num {
		for len(stack) > 0 && num[i] < stack[len(stack)-1] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}

		if num[i] != '0' || len(stack) != 0 {
			stack = append(stack, num[i])
		}
	}

	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
