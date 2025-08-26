package maximumareaoflongestdiagonalrectangle

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDiagonal := 0
	maxArea := 0

	for _, val := range dimensions {
		currDiagonal, currArea := val[0]*val[0]+val[1]*val[1], val[0]*val[1]
		if currDiagonal > maxDiagonal || (currDiagonal == maxDiagonal && currArea > maxArea) {
			maxDiagonal = currDiagonal
			maxArea = val[0] * val[1]
		}
	}

	return maxArea
}
