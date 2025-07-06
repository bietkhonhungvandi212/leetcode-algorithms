package rangesumofbst

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0

	var dfs func(root *TreeNode, low int, high int)
	dfs = func(root *TreeNode, low, high int) {
		if root == nil {
			return
		}

		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}

		if root.Val > low {
			dfs(root.Left, low, high)
		}

		if root.Val < high {
			dfs(root.Right, low, high)
		}
	}

	dfs(root, low, high)

	return sum
}
