package minimumtimetomakeropecolorful

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	time := 0
	maxTime := neededTime[0]
	for i := 1; i < n; i++ {
		if colors[i] != colors[i-1] {
			maxTime = 0
		}
		time += min(maxTime, neededTime[i])
		maxTime = max(maxTime, neededTime[i])
	}

	return time
}
