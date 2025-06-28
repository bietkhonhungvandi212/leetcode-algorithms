package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	count := 0
	stack := make([]*ListNode, 0)
	cur := head

	for cur != nil {
		count++
		stack = append(stack, cur)
		cur = cur.Next
	}

	top := stack[len(stack)-1]
	top.Next = head
	k = k % count

	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	head = stack[len(stack)-1]
	result := head.Next
	head.Next = nil

	return result
}
