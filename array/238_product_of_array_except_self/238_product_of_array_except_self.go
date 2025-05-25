package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	product := 1
	zeroCnt := 0
	result := make([]int, len(nums))

	for i := range nums {
		if nums[i] == 0 {
			zeroCnt++
		}

		if zeroCnt < 2 && nums[i] != 0 {
			product *= nums[i]
		}

		if zeroCnt >= 2 {
			product = 0
			break
		}
	}

	for i := range nums {
		if nums[i] == 0 && zeroCnt < 2 {
			result[i] = product
		} else if zeroCnt >= 2 || (zeroCnt > 0 && nums[i] != 0) {
			result[i] = 0
		} else {
			result[i] = product / nums[i]
		}
	}

	return result
}

func productExceptSelf2(nums []int) []int {
	output := make([]int, len(nums))

	previous := 1
	for i := range nums {
		if i == 0 {
			continue
		}
		output[i] = previous * nums[i-1]
		previous *= nums[i-1]
	}

	current := 1
	for i := len(nums) - 2; i >= 0; i-- {
		output[i] = current * nums[i+1]
		current = output[i]

	}

	return output
}
