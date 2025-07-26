package intervalarray

import "sort"

func merge(intervals [][]int) [][]int {
	stack := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	stack = append(stack, intervals[0])

	for i := range intervals {
		// [1,3] - [2,6]
		head := stack[len(stack)-1]
		if head[1] >= intervals[i][0] {
			if head[1] < intervals[i][1] {
				head[1] = intervals[i][1]
			}
		} else {
			stack = append(stack, intervals[i])
		}

	}

	return stack
}
