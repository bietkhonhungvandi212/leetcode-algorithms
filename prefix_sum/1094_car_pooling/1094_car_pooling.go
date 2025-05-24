package carpooling

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	max := 0
	for i := range trips {
		if max < trips[i][2] {
			max = trips[i][2]
		}
	}

	pickNumber := make([]int, max+1)

	for i := range trips {
		numPassengers := trips[i][0]
		from := trips[i][1]
		to := trips[i][2]

		pickNumber[from] += numPassengers

		if to < max {
			pickNumber[to+1] -= numPassengers
			fmt.Println(pickNumber[to+1])
		}
	}

	if pickNumber[0] > capacity {
		return false
	}

	for i := 1; i <= max; i++ {
		pickNumber[i] = pickNumber[i] + pickNumber[i-1]

		if pickNumber[i] > capacity {
			return false
		}
	}

	return true
}
