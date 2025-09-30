package findtriangularsumofanarray

func triangularSum(nums []int) int {
	length := len(nums)
	for length != 1 {
		for i := range length - 1 {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		length--
	}

	return nums[0]
}
