package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	// i must maintain a stack of elements that wait for warmer day
	// -> monotoic stack maintains the decreasing order of element
	// if meeting the a higher/ warmer temparature, we must calculate the indice at that element with indices stored in stack
	n := len(temperatures)
	stack := []int{0}
	result := make([]int, n)
	for i := 1; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			head := stack[len(stack)-1]
			result[head] = i - head
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}
