package coinchange

import "math"

func coinChange(coins []int, amount int) int {
	memorizes := make(map[int]int, 0)
	var minCoinChange func(amount int) int
	minCoinChange = func(amount int) int {
		if amount < 0 {
			return math.MaxInt
		}

		if amount == 0 {
			return 0
		}

		if value, ok := memorizes[amount]; ok {
			return value
		}

		minCoins := math.MaxInt
		for i := range coins {
			remain := minCoinChange(amount - coins[i])
			if remain != math.MaxInt {
				minCoins = MinMath(minCoins, 1+remain)
			}
		}

		memorizes[amount] = minCoins
		return minCoins
	}

	min := minCoinChange(amount)
	if min == math.MaxInt {
		return -1
	}

	return min
}

func MinMath(a, b int) int {
	if a >= b {
		return b
	}

	return a
}
