package corporateflightbookings

func corpFlightBookings(bookings [][]int, n int) []int {
	preservation := make([]int, n+1)
	result := make([]int, n)

	for i := range bookings {
		first := bookings[i][0]
		last := bookings[i][1]
		seat := bookings[i][2]
		preservation[first] += seat
		if last < n {
			preservation[last+1] -= seat

		}
	}

	for i := 1; i <= n; i++ {
		preservation[i] += preservation[i-1]
		result[i-1] = preservation[i]
	}

	return result
}
