package minimumdistancebetweenbstnodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	min := math.MaxInt
	var pred *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if pred != nil {
			min = MathMin(min, root.Val-pred.Val)

		}
		pred = root
		dfs(root.Right)
	}
	dfs(root)

	return min
}

func MathMin(a, b int) int {
	if a >= b {
		return b
	}

	return a
}
