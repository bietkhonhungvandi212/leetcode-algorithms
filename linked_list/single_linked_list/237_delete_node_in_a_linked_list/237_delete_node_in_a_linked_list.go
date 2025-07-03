package deletenodeinalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	curr := node
	var prev *ListNode
	for curr.Next != nil {
		curr.Val = curr.Next.Val

		prev = curr
		curr = curr.Next
	}

	prev.Next = nil
}
