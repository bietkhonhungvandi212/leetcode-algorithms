package evaluatereversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		n := len(stack)
		isOperator := false
		var head1, head2, result int
		if len(stack) >= 2 {
			head1 = stack[n-1]
			head2 = stack[n-2]
		}

		if token == "+" {
			isOperator = true
			result = head1 + head2
		}

		if token == "-" {
			isOperator = true
			result = head2 - head1
		}

		if token == "*" {
			isOperator = true
			result = head1 * head2
		}

		if token == "/" {
			isOperator = true
			result = head2 / head1
		}

		if isOperator {
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
