package besttimetobuyandsellstockii

import "math"

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))

	for i := range dp {
		dp[i] = math.MinInt
	}

	return greedy(0, prices, &dp)
}

func greedy(buyIndex int, prices []int, dp *[]int) int {
	n := len(prices)
	if buyIndex == n {
		return 0
	}

	if (*dp)[buyIndex] != math.MinInt {
		return (*dp)[buyIndex]
	}

	max := 0
	for i := buyIndex; i < n; i++ {
		if prices[i] >= prices[buyIndex] {
			revenue := prices[i] - prices[buyIndex] + greedy(i+1, prices, dp)
			if max < revenue {
				max = revenue
			}
		}
	}
	(*dp)[buyIndex] = max

	return max
}

func maxProfit_Optimized(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		dayProfit := prices[i] - prices[i-1]
		if dayProfit > 0 {
			profit += dayProfit
		}
	}
	return profit

}
