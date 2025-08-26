package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		elements := make([]int, 0, size)

		for range size {
			first := queue[0]
			queue = queue[1:]

			elements = append(elements, first.Val)

			if first.Left != nil {
				queue = append(queue, first.Left)
			}

			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}

		result = append(result, elements)
	}

	return result
}
