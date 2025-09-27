package findminimuminrotatedsortedarray

/*
Constraint O(logn) -> This is a binary search problem
It is monotoic problem that
	- If meet the condition then move to right
	- else move to left
*/
func findMin(nums []int) int {
	n := len(nums)

	left := 0
	right := n - 1

	for left < right {

		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}

	}

	return nums[left]
}
