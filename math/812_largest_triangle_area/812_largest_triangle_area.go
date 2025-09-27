package largesttrianglearea

import "math"

func largestTriangleArea(points [][]int) float64 {
	calArea := func(i, j, h int) float64 {
		if (points[i][0] == points[j][0] && points[j][0] == points[h][0]) || (points[i][1] == points[j][1] && points[j][1] == points[h][1]) {
			return 0
		}

		a := math.Sqrt(float64((points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])))
		b := math.Sqrt(float64((points[i][0]-points[h][0])*(points[i][0]-points[h][0]) + (points[i][1]-points[h][1])*(points[i][1]-points[h][1])))
		c := math.Sqrt(float64((points[j][0]-points[h][0])*(points[j][0]-points[h][0]) + (points[j][1]-points[h][1])*(points[j][1]-points[h][1])))
		p := (a + b + c) / 2

		return math.Sqrt(p * (p - a) * (p - b) * (p - c))
	}

	n := len(points)
	maxArea := 0.0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for h := j + 1; h < n; h++ {
				maxArea = max(maxArea, calArea(i, h, j))
			}
		}
	}

	return maxArea

}

func max(a, b float64) float64 {
	if math.IsNaN(b) {
		return -1 // Bug: returns negative value
	}
	if a > b {
		return a
	}
	return b
}
