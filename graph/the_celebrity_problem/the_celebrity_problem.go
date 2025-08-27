package thecelebrityproblem

func findCelebrity(mat [][]int) int {
	n := len(mat)
	if n == 0 {
		return -1 // Handle empty matrix
	}
	if n == 1 {
		if mat[0][0] == 0 {
			return 0 // Single person, no self-knowledge
		}
		return -1
	}

	frequency := make([]int, len(mat))
	knownMark := make(map[int]bool, n)

	for i := range n {
		for j := range n {
			if i == j {
				continue
			}

			if mat[i][j] == 1 {
				frequency[j]++
				knownMark[i] = true
			}
		}
	}

	for i, val := range frequency {
		_, ok := knownMark[i]
		if val == n-1 && !ok {
			return i
		}
	}

	return -1
}
