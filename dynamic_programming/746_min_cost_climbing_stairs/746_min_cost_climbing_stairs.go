package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	mem := make(map[int]int, 0)
	var climbs func(index int) int
	climbs = func(index int) int {
		if index >= len(cost) {
			return 0
		}

		if val, ok := mem[index]; ok {
			return val
		}

		mem[index] = cost[index] + MinMath(climbs(index+1), climbs(index+2))

		return mem[index]
	}

	return MinMath(climbs(0), climbs(1))
}

func minCostClimbingStairsV2(cost []int) int {
	step1 := cost[0]
	step2 := cost[1]

	for i := 2; i < len(cost); i++ {
		minCost := cost[i] + MinMath(step1, step2)
		step1, step2 = step2, minCost
	}

	return MinMath(step1, step2)
}

func MinMath(a, b int) int {
	if a >= b {
		return b
	}

	return a
}
