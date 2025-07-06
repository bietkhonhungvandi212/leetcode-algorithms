package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	traverseTree(root, &result)
	return result
}

func inorderTraversalV2(root *TreeNode) []int {
	result := []int{}
	var stack []*TreeNode
	current := root

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, current.Val)
			if current.Right != nil {
				current = current.Right
			} else {
				current = nil
			}
		}
		// for current != nil {
		// 	stack = append(stack, current)
		// 	current = current.Left
		// }

		// if len(stack) > 0 {
		// 	current = stack[len(stack)-1]
		// 	stack = stack[:len(stack)-1]
		// 	result = append(result, current.Val)
		// 	if current.Right != nil {
		// 		current = current.Right
		// 	} else {
		// 		current = nil
		// 	}
		// }
	}

	return result
}

func traverseTree(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	traverseTree(root.Left, result)
	*result = append(*result, root.Val)
	traverseTree(root.Right, result)
}
