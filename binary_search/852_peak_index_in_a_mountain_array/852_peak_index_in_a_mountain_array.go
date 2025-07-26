package peakindexinamountainarray

func peakIndexInMountainArray(arr []int) int {
	//0
	left := 0
	//n-1
	right := len(arr) - 1

	//3 4 5 1
	for left < right {
		mid := (left + right) / 2

		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
