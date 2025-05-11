package lowestcommonancestorofabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Set struct {
	elements map[string]struct{}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var stack []*TreeNode
	parentMapper := make(map[*TreeNode]*TreeNode)

	parentMapper[root] = nil
	stack = append(stack, root)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.Left != nil {
			parentMapper[curr.Left] = curr
			stack = append(stack, curr.Left)
		}

		if curr.Right != nil {
			parentMapper[curr.Right] = curr
			stack = append(stack, curr.Right)
		}
	}

	set := make(map[*TreeNode]bool)
	tmpNode1 := p
	for tmpNode1 != nil {
		set[tmpNode1] = true

		node, ok := parentMapper[tmpNode1]

		if ok {
			tmpNode1 = node
		} else {
			tmpNode1 = nil
		}
	}

	tmpNode2 := q
	for tmpNode2 != nil {
		_, ok := set[tmpNode2]

		if ok {
			return tmpNode2
		} else {
			node, ok := parentMapper[tmpNode2]

			if ok {
				tmpNode2 = node
			} else {
				tmpNode2 = nil
			}

		}
	}

	return nil
}
