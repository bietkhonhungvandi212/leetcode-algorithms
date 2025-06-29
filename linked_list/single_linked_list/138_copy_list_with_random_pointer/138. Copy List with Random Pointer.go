package copylistwithrandompointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	current := head
	mapper := make(map[*Node]*Node)
	mapper[nil] = nil

	var getCloneNode func(*Node) *Node
	getCloneNode = func(n *Node) *Node {
		if n == nil {
			return n
		}

		if clone, existed := mapper[n]; existed {
			return clone
		}

		cloned := &Node{Val: n.Val}
		mapper[n] = cloned
		return cloned
	}

	for current != nil {
		clonedNode := getCloneNode(current)
		clonedNode.Next = getCloneNode(current.Next)
		clonedNode.Random = getCloneNode(current.Random)

		current = current.Next
	}

	return mapper[head]
}
