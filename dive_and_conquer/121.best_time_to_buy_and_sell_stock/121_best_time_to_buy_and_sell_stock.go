package besttimetobuyandsellstock

import (
	"fmt"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}

	changes := make([]int, n)
	// create the array to store the sum of adjacent elements
	for i := 1; i < n; i++ {
		changes[i] = prices[i] - prices[i-1]
	}

	left, right, max := findMax(changes, 1, n-1)
	fmt.Printf("Buy at day %d (price: %d)\n", left, prices[left])
	fmt.Printf("Sell at day %d (price: %d)\n", right, prices[right])
	fmt.Printf("Maximum profit: %d\n", max)

	if max < 0 {
		return 0
	}

	return max
}

func findMax(changes []int, left, right int) (int, int, int) {
	if left >= right {
		return left, right, changes[left]
	}
	mid := (left + right) / 2

	// Find the left-side max
	leftLow, leftHigh, sumLeft := findMax(changes, left, mid)
	// Find the right-side max
	rightLow, rightHigh, sumRight := findMax(changes, mid+1, right)
	//Find the accross-mid max
	acrossLow, acrosshigh, sumAcross := findAcrossMax(changes, left, mid, right)

	if sumLeft >= sumRight && sumLeft >= sumAcross {
		return leftLow, leftHigh, sumLeft
	} else if sumRight > sumLeft && sumRight > sumAcross {
		return rightLow, rightHigh, sumRight
	}

	return acrossLow, acrosshigh, sumAcross
}

func findAcrossMax(changes []int, left, mid, right int) (int, int, int) {
	sum := 0
	leftsum := changes[mid]
	leftIndex := mid
	for i := mid; i >= left; i-- {
		sum += changes[i]
		if leftsum < sum {
			leftsum = sum
			leftIndex = i
		}
	}

	sum = 0
	rightSum := changes[mid+1]
	rightIndex := mid + 1
	for i := mid + 1; i <= right; i++ {
		sum += changes[i]
		if rightSum < sum {
			rightSum = sum
			rightIndex = i
		}

	}

	return leftIndex, rightIndex, leftsum + rightSum

}
