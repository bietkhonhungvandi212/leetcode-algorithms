package containerwithmostwater

/*
- Start with the maximum width: Initialize two pointers, left at the start (index 0) and
 right at the end (index len(height)-1).
- Shrink based on height: Compare the heights at left and right.
 Move the pointer from the side with the smaller height inward (increment left if height[left] < height[right], or decrement right otherwise).
- Reason for moving smaller height: By moving the pointer at the smaller height,
 you aim to find a potentially taller line that could increase the area, as the area is limited by the shorter height, and reducing width by moving the taller height is less likely to yield a larger area.
- Track maximum area: Calculate the area at each step (min(height[left], height[right]) * (right - left))
 and update the maximum area if the current area is larger.
*/
func maxArea(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	maxArea := min(height[left], height[right]) * (right - left)

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if maxArea < area {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return maxArea

}
