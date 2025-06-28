package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	stack := make([]*ListNode, 0)
	current := head

	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	for len(stack) > 0 && n > 0 {
		stack = stack[:len(stack)-1]
		n--
	}

	if len(stack) == 0 {
		return head.Next
	}
	top := stack[len(stack)-1]
	top.Next = top.Next.Next

	return head
}
