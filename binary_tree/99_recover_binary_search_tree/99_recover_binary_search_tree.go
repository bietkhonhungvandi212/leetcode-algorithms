package recoverbinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var pred, first, second *TreeNode

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if pred != nil && pred.Val > root.Val {
			if first == nil {
				first = pred
				second = root
			} else {
				second = root
			}

		}

		pred = root
		dfs(root.Right)
	}

	if first != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
