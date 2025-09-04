package findclosestperson

func findClosest(x int, y int, z int) int {
	distance1 := abs(x - z)
	distance2 := abs(y - z)

	if distance1 == distance2 {
		return 0
	}

	if distance1 < distance2 {
		return 1
	}

	return 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
