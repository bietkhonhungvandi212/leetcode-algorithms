package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	current := head

	var tmp1 *ListNode
	for current != nil {
		tmp2 := current.Next
		if current == head {
			current.Next = nil
		} else {
			current.Next = tmp1
		}

		tmp1 = current
		current = tmp2
	}

	return tmp1
}
