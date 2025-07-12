package minimumcosttocutastick

import (
	"math"
	"sort"
)

func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)

	lengthCuts := len(cuts)

	mem := make([][]int, lengthCuts)
	for i := range mem {
		mem[i] = make([]int, lengthCuts)
	}

	for pairLength := 2; pairLength < lengthCuts; pairLength++ {
		for i := 0; i < lengthCuts-pairLength; i++ {
			j := i + pairLength
			min := math.MaxInt
			for k := i + 1; k < j; k++ {
				cost := cuts[j] - cuts[i] + mem[i][k] + mem[k][j]
				if cost < min {
					min = cost
				}
			}
			mem[i][j] = min
		}
	}

	return mem[0][lengthCuts-1]
}

func minCost_TopDown(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)

	lengthCuts := len(cuts)

	mem := make([][]int, lengthCuts)
	for i := range mem {
		mem[i] = make([]int, lengthCuts)
	}

	var cutStick func(start, end int) int
	cutStick = func(start, end int) int {
		if mem[start][end] != 0 {
			return mem[start][end]
		}

		if end-start == 1 {
			return 0
		}

		min := math.MaxInt
		for mid := start + 1; mid < end; mid++ {
			cost := cutStick(start, mid) + cutStick(mid, end) + cuts[end] - cuts[start]
			if min > cost {
				min = cost
			}
		}

		mem[start][end] = min

		return mem[start][end]
	}

	return cutStick(0, lengthCuts-1)
}
