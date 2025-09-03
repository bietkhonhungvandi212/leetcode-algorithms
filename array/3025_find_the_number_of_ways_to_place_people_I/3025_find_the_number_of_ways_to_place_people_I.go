package findthenumberofwaystoplacepeoplei

import (
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	n := len(points)

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	count := 0
	// Iterate over pairs (A, B)
	for i := range n {
		upperY := points[i][1]
		minLow := math.MinInt
		for j := i + 1; j < n; j++ {
			currentY := points[j][1]
			if currentY <= upperY && currentY > minLow {
				count++
				minLow = currentY

				if currentY == upperY {
					break
				}
			}
		}
	}

	return count
}
