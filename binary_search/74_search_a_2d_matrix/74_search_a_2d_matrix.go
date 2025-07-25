package searcha2dmatrix

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	length1 := len(matrix) - 1
	left := 0
	right := length1

	for left <= right {
		mid := (left + right) / 2
		if target > matrix[mid][0] {
			left = mid + 1
		} else if target < matrix[mid][0] {
			right = mid - 1
		} else {
			return true
		}
	}

	if left == 0 {
		return false
	}

	findArray := matrix[left-1]
	left = 0
	right = len(findArray) - 1

	//4
	for left <= right {
		fmt.Println(left)
		fmt.Println(right)
		//2
		//1
		mid := (left + right) / 2

		// findarray[2] = 5
		//findarray[1] = 3
		if target > findArray[mid] {
			left = mid + 1
		} else if target < findArray[mid] {
			//left = 0, right = 1
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}
