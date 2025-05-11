package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	differ := maxDepth(root.Left) - maxDepth(root.Right)

	if abs(differ) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + MaxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func MaxInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
