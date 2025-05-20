package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	mapper := make(map[int]bool)

	dfs(root, &result, &mapper, 0)

	return result
}

func dfs(root *TreeNode, result *[]int, mapper *map[int]bool, level int) {
	if root == nil {
		return
	}

	if _, ok := (*mapper)[level]; ok == false {
		*result = append(*result, root.Val)
		(*mapper)[level] = true
	}

	dfs(root.Right, result, mapper, level+1)
	dfs(root.Left, result, mapper, level+1)
}
