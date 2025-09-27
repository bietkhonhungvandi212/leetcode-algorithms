package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	product := 1
	for i := 1; i < n; i++ {
		product *= nums[i-1]
		result[i] = product
	}
	result[n-1] = product
	product = 1

	for i := n - 2; i >= 0; i-- {
		product *= nums[i+1]
		result[i] *= product
	}
	result[0] = product

	return result
}
