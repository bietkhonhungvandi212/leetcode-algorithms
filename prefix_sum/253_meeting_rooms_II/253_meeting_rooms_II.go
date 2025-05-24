package meetingroomsii

func minMeetingRooms(intervals [][]int) int {
	max := 0
	for i := range intervals {
		if max < intervals[i][1] {
			max = intervals[i][1]
		}
	}

	trackingOverlap := make([]int, max+1)
	for i := range intervals {
		start := intervals[i][0]
		end := intervals[i][1]

		trackingOverlap[start]++
		trackingOverlap[end]--
	}

	res := trackingOverlap[0]
	for i := 1; i <= max; i++ {
		trackingOverlap[i] += trackingOverlap[i-1]

		if res < trackingOverlap[i] {
			res = trackingOverlap[i]
		}
	}

	return res
}
