package validatebinarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, min, max int) bool
	dfs = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}

		if root.Val >= max || root.Val <= min {
			return false
		}

		if root.Left != nil && root.Left.Val >= root.Val {
			return false
		}

		if root.Right != nil && root.Right.Val <= root.Val {
			return false
		}

		return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}
