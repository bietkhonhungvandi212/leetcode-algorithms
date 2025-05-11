package besttimetobuyandsellstockii

func maxProfit(prices []int) int {
	n := len(prices)
	left := prices[0]
	max := 0
	maxRange := 0

	for i := 1; i < n; i++ {
		if left < prices[i] {
			maxRange = MaxInt(maxRange, prices[i]-left)
		}
		left = prices[i]
		max = max + maxRange
		maxRange = 0
	}

	return max

}

func maxProfitDP(prices []int) int {
	n := len(prices)
	dpArr := make([]int, n)
	dpArr[n-1] = 0

	max := 0
	maximumPrice := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		if maximumPrice-prices[i] > 0 {
			dpArr[i] = maximumPrice - prices[i] + max

			if max < dpArr[i] {
				max = dpArr[i]
			}
		}

		maximumPrice = prices[i]
		dpArr[i] = dpArr[i+1]
	}

	return max
}

func MaxInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
