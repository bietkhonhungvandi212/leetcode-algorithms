package usingarobottoprintthelexicographicallysmalleststring

func robotWithString(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	minSuffix := make([]byte, n)
	minSuffix[n-1] = s[n-1]
	for i := n - 2; i >= 0; i-- {
		minSuffix[i] = min(s[i], minSuffix[i+1])
	}

	stack := make([]byte, 0)
	result := make([]byte, 0)

	for i := range n {
		stack = append(stack, s[i])
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if i == n-1 || top <= minSuffix[i+1] {
				result = append(result, top)
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

	}

	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return string(result)
}
