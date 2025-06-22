package scoreofparentheses

import "math"

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

func scoreOfParentheses_StackOfScores(s string) int {
	// stack will hold the scores of nested parenthetical levels.
	stack := []int{0} // Start with a base score of 0 for the outermost level.

	for _, char := range s {
		if char == '(' {
			// Entering a new, deeper level. Start its score at 0.
			stack = append(stack, 0)
		} else { // char == ')'
			// Finishing a level.
			// 1. Pop the score of the completed level.
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 2. Get the score of the parent level (now at the top).
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 3. Calculate the value of the completed level and add it to the parent.
			// If v was 0, it was a "()", worth 1.
			// If v > 0, it was an "(A)", worth 2 * v.
			// max(2*v, 1) is a clever way to combine these two cases.
			scoreToAdd := int(math.Max(float64(2*v), 1.0))

			// 4. Push the updated parent score back onto the stack.
			stack = append(stack, p+scoreToAdd)
		}
	}

	return stack[0]
}
