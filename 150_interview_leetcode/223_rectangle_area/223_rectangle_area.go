package rectanglearea

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	width := min(ax2, bx2) - max(ax1, bx1)
	height := min(ay2, by2) - max(ay1, by1)

	commonArea := width * height

	if bx1 >= ax2 || bx2 <= ax1 || by2 <= ay1 || by1 >= ay2 {
		commonArea = 0
	}

	totalArea1 := (ax2 - ax1) * (ay2 - ay1)
	totalArea2 := (bx2 - bx1) * (by2 - by1)

	return totalArea1 + totalArea2 - commonArea
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
