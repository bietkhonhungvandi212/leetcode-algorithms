package findkclosestelements

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr)

	for left < right {
		mid := (left + right) / 2
		if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	start := left - 1
	end := left

	for k > 0 {
		if start < 0 {
			end++
		} else if end >= len(arr) {
			start--
		} else {
			if abs(arr[start], x) <= abs(arr[end], x) {
				start--
			} else {
				end++
			}
		}
		k--
	}

	return arr[start+1 : end]
}

func abs(a, b int) int {
	sub := a - b
	if sub > 0 {
		return sub
	}

	return (-1) * sub
}
