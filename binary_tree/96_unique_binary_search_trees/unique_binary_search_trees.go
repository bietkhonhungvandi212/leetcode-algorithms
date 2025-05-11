package uniquebinarysearchtrees

func numTrees(n int) int {
	backtrack := make([]int, n+1)

	backtrack[0], backtrack[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			backtrack[i] += backtrack[j-1] * backtrack[i-j]
		}
	}

	return backtrack[n]
}
