package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	n := maxDepth(root)

	arr2D := make([][]int, n)
	for i := range arr2D {
		arr2D[i] = []int{}
	}

	dfs(root, &arr2D, 0)

	return arr2D
}

func dfs(root *TreeNode, arr2D *[][]int, level int) {
	if root == nil {
		return
	}

	(*arr2D)[level] = append((*arr2D)[level], root.Val)

	dfs(root.Left, arr2D, level+1)
	dfs(root.Right, arr2D, level+1)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + MaxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
