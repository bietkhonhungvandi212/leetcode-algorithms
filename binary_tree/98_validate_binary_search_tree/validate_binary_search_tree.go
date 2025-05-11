package validatebinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var arr []int

	InOrderTraversal(root, &arr)

	for i := range arr {
		if i == 0 {
			continue
		}
		if arr[i] <= arr[i-1] {
			return false
		}
	}

	return true
}

func InOrderTraversal(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	InOrderTraversal(root.Left, arr)
	*arr = append(*arr, root.Val)
	InOrderTraversal(root.Right, arr)
}
