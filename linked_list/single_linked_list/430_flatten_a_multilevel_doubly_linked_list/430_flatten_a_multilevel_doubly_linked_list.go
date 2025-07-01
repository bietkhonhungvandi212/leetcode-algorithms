package flattenamultileveldoublylinkedlist

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	stack := make([]*Node, 0)
	current := root

	for current != nil || len(stack) > 0 {
		if current.Child != nil {
			if current.Next != nil {
				stack = append(stack, current.Next)
			}
			current.Next = current.Child
			current.Child.Prev = current
			tmp := current
			current = current.Child
			tmp.Child = nil

		} else {
			if current.Next == nil && len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				top.Prev = current
				current.Next = top

			}

			current = current.Next
		}

	}

	return root
}

func flatten_v2(root *Node) *Node {
	if root == nil {
		return root
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)

	var prev *Node
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil {
			prev.Next = curr
			curr.Prev = prev
		}

		if curr.Next != nil {
			stack = append(stack, curr.Next)
		}

		if curr.Child != nil {
			stack = append(stack, curr.Child)
			curr.Child = nil
		}

		prev = curr
	}

	return root
}
