package findmodeinbinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	max := 0
	mapper := make(map[int]int, 0)
	result := make([]int, 0)
	findModeDFS(root, &mapper, &max)

	for key := range mapper {
		if mapper[key] == max {
			result = append(result, key)
		}
	}

	return result
}

func findModeDFS(root *TreeNode, mapper *map[int]int, max *int) {
	if root == nil {
		return
	}

	findModeDFS(root.Left, mapper, max)
	(*mapper)[root.Val]++

	if (*mapper)[root.Val] > *max {
		*max = (*mapper)[root.Val]

	}

	findModeDFS(root.Right, mapper, max)
}
