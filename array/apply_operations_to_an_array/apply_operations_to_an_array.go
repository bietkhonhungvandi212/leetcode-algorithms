package main

import "fmt"

func applyOperations(nums []int) []int {
	arrLen := len(nums)
	storeNums := make([]int, arrLen)
	index := 0

	for i := range arrLen - 1 {
		if nums[i] == 0 {
			continue
		}

		if nums[i] == nums[i+1] {
			storeNums[index] = nums[i] * 2
			nums[i+1] = 0
		} else if nums[i] != 0 {
			storeNums[index] = nums[i]
		}

		index++
	}

	if nums[arrLen-1] != 0 {
		storeNums[index] = nums[arrLen-1]
	}

	return storeNums
}

func main() {
	testCases := [][]int{
		{0, 1}, // Adjacent duplicates
	}

	for _, test := range testCases {
		fmt.Printf("Input: %v -> Output: %v\n", test, applyOperations(test))
	}
}
