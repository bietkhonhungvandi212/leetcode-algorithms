package jumpgame

func canJump(nums []int) bool {
	memo := make([]int, len(nums))
	return checkJumpToLastIndex(nums, 0, memo)
}

func checkJumpToLastIndex(nums []int, currIndex int, memo []int) bool {
	if currIndex == len(nums)-1 {
		return true
	}

	if currIndex >= len(nums) || nums[currIndex] == 0 {
		return false
	}

	if memo[currIndex] == 1 {
		return false
	}

	for i := nums[currIndex]; i >= 1; i-- {
		if checkJumpToLastIndex(nums, currIndex+i, memo) {
			return true
		}
	}

	memo[currIndex] = 1
	return false
}
