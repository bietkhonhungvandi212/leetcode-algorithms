package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSameMirror(root.Left, root.Right)
}

func isSameMirror(left *TreeNode, right *TreeNode) bool {
	if right == nil && left == nil {
		return true
	}

	if right == nil || left == nil || right.Val != left.Val {
		return false
	}

	return isSameMirror(left.Left, right.Right) && isSameMirror(left.Right, right.Left)
}
