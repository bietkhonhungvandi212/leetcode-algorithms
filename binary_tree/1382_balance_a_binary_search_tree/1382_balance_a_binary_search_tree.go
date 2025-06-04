package balanceabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	sortedArr := preOrderTravesal(root)

	return sortedArrayToBST(0, len(sortedArr)-1, sortedArr)

}

func sortedArrayToBST(start, end int, sortedArray []int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	current := TreeNode{Val: sortedArray[mid]}
	current.Left = sortedArrayToBST(start, mid, sortedArray)
	current.Right = sortedArrayToBST(mid+1, end, sortedArray)
	return &current
}

func preOrderTravesal(root *TreeNode) []int {
	stack := make([]int, 0)
	result := make([]int, 0)

	current := root
	for current != nil || len(stack) > 0 {
		if current != nil {
			if current.Right != nil {
				stack = append(stack, current.Right.Val)
			}
			stack = append(stack, current.Left.Val)
			current = current.Left
		} else {
			currentValue := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, currentValue)
		}
	}

	return result
}
