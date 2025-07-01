package oddevenlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	oddDummy := &ListNode{}
	evenDummy := &ListNode{}

	current := head
	prevOdd := oddDummy
	prevEven := evenDummy

	i := 0
	for current != nil {
		if i%2 == 0 {
			prevOdd.Next = current
			prevOdd = prevOdd.Next
		} else {
			prevEven.Next = current
			prevEven = prevEven.Next
		}
		i++
		current = current.Next
	}

	prevOdd.Next = evenDummy.Next
	prevEven.Next = nil

	return oddDummy.Next
}
