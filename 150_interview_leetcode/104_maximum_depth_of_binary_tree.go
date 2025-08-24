package interviewleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var findMaxDepth func(root *TreeNode) int
	findMaxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftDepth := findMaxDepth(root.Left)
		rightDepth := findMaxDepth(root.Right)

		return max(leftDepth, rightDepth) + 1
	}

	return findMaxDepth(root)
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
