package nthtribonaccinumber

func tribonacci(n int) int {
	mem := make(map[int]int, 0)
	var calculateTribonacci func(n int) int
	calculateTribonacci = func(n int) int {
		if n == 0 {
			return 0
		}

		if n == 1 || n == 2 {
			return 1
		}

		if val, ok := mem[n]; ok {
			return val
		}

		mem[n] = calculateTribonacci(n-3) + calculateTribonacci(n-2) + calculateTribonacci(n-1)

		return mem[n]
	}

	return calculateTribonacci(n)

}
