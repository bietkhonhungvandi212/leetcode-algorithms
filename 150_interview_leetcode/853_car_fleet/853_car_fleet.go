package carfleet

func carFleet(target int, position []int, speed []int) int {
	hours := make([]float64, target)
	stack := make([]int, 0)

	for i, pos := range position {
		distance := target - pos
		hours[pos] = float64(distance) / float64(speed[i])
	}

	for pos, hour := range hours {
		if hour == 0 {
			continue
		}

		for len(stack) > 0 && hour >= hours[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, pos)
	}

	return len(stack)
}
