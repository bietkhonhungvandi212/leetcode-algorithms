package binarytreepreordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalV2(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	stack = append(stack, root)
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			result = append(result, current.Val)
			if current.Right != nil {
				stack = append(stack, current.Right)
			}
			current = current.Left
		}

		if len(stack) > 0 {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return result
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	traversal(root, &result)

	return result
}

func traversal(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	if root.Left != nil {
		traversal(root.Left, arr)
	}

	if root.Right != nil {
		traversal(root.Right, arr)
	}
}
