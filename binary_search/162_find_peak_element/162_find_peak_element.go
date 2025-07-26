package findpeakelement

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	// Binary search for peak element
	for left < right {
		mid := (left + right) / 2

		// If mid element is smaller than its right neighbor,
		// then there is always a peak in right half
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			// If mid element is greater than or equal to its right neighbor,
			// then there is always a peak in left half (including mid)
			right = mid
		}
	}

	// When left == right, we found the peak
	return left
}
