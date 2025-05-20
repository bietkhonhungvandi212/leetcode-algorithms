package subtreeofanothertree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	var helper func(p, q *TreeNode) bool
	helper = func(p, q *TreeNode) bool {
		if p == nil {
			return false
		}
		if isSameTree(p, q) {
			return true
		}
		return helper(p.Left, q) || helper(p.Right, q)
	}

	return helper(root, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
