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

func reverseList_V2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	current := dummy.Next

	for current.Next != nil {
		next := current.Next
		current.Next = next.Next
		next.Next = dummy.Next
		dummy.Next = next
	}

	return dummy.Next
}
