package main

import "fmt"

func maxProfit(prices []int) int {
	left := 0
	right := 0
	max := 0

	for right < len(prices) {
		profit := prices[right] - prices[left]

		if profit < 0 {
			left++
		} else {
			max = MaxInt(max, profit)
			right++
		}
	}

	return max
}

func MaxInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	// Test cases
	testCases := []struct {
		prices   []int
		expected int
		scenario string
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5, "Standard case with profit"},
		{[]int{7, 6, 5, 4, 3, 1}, 0, "Decreasing prices - no profit"},
		{[]int{1, 2}, 1, "Simple case with profit"},
		{[]int{2, 1}, 0, "Simple case without profit"},
		{[]int{1}, 0, "Single price - no transactions possible"},
		{[]int{}, 0, "Empty array"},
		{[]int{3, 2, 6, 5, 0, 3, 10}, 10, "Multiple peaks and valleys"},
	}

	for _, tc := range testCases {
		profit := maxProfit(tc.prices)
		fmt.Printf("Scenario: %s\n", tc.scenario)
		fmt.Printf("Prices: %v\n", tc.prices)
		fmt.Printf("Expected profit: %d, Got: %d\n", tc.expected, profit)
		if profit == tc.expected {
			fmt.Println("✓ PASS")
		} else {
			fmt.Println("✗ FAIL")
		}
		fmt.Println("---------------------")
	}
}
