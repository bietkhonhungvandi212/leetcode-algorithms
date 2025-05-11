package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	traversal(root, &result)

	return result

}

func postorderTraversalWithoutRecursion(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	var previousCurr *TreeNode
	current := root

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			if current.Right == nil || current.Right == previousCurr {
				stack = stack[:len(stack)-1]
				result = append(result, current.Val)
				previousCurr = current
				current = nil
			} else {
				current = current.Right
			}
		}
	}

	return result
}

// func postorderTraversalWithoutRecursion(root *TreeNode) []int {
// 	var result []int
// 	var stack []*TreeNode
// 	current := root
// 	var previous *TreeNode

// 	for current != nil || len(stack) > 0 {
// 		if current != nil {
// 			stack = append(stack, current)
// 			current = current.Left
// 		} else {
// 			current = stack[len(stack)-1]
// 			// result = append(result, current.Val)
// 			if current.Right == nil || current.Right == previous {
// 				stack = stack[:len(stack)-1]
// 				result = append(result, current.Val)
// 				previous = current
// 				current = nil
// 			} else {
// 				current = current.Right
// 			}
// 		}
// 	}

// 	return result
// }

func traversal(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, arr)
	traversal(root.Right, arr)
	*arr = append(*arr, root.Val)
}
