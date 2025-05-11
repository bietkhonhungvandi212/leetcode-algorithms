package climbingstairs

func climbStairs_BottomUp(n int) int {
	mem := make([]int, n+1)
	if n <= 2 {
		return n
	}
	mem[1] = 1
	mem[2] = 2

	for i := 3; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}

	return mem[n]
}

func climbStairs(n int) int {
	mem := make([]int, n+1)
	mem[0] = 1
	mem[1] = 1
	for i := range n {
		mem[i+1] = -1
	}
	return countClimbStairs(n, mem)
}

func countClimbStairs(n int, mem []int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if mem[n] != -1 {
		return mem[n]
	}

	mem[n] = countClimbStairs(n-1, mem) + countClimbStairs(n-2, mem)

	return mem[n]
}
