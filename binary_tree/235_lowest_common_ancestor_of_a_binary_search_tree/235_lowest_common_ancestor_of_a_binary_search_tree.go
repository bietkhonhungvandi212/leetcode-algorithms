package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var min func(num1, num2 int) int
	var max func(num1, num2 int) int
	min = func(num1, num2 int) int {
		if num1 > num2 {
			return num2
		}

		return num1
	}

	max = func(num1, num2 int) int {
		if num1 < num2 {
			return num2
		}

		return num1
	}

	if root.Val < min(p.Val, q.Val) {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if root.Val > max(p.Val, q.Val) {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}
