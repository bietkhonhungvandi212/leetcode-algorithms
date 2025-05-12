package threeconsecutiveodds

func threeConsecutiveOdds(arr []int) bool {
	count := 0

	for i := range arr {
		if arr[i]%2 == 0 {
			count = 0
		} else {
			count++
		}

		if count == 3 {
			return true
		}
	}

	return false
}
