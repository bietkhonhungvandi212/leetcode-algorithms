package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur1 := l1
	cur2 := l2

	cur3 := result
	complement := 0
	for cur1 != nil || cur2 != nil {
		var sum int
		if cur1 != nil && cur2 != nil {
			sum = cur1.Val + cur2.Val + complement
			cur1 = cur1.Next
			cur2 = cur2.Next
		} else if cur1 != nil {
			sum = cur1.Val + complement
			cur1 = cur1.Next
		} else {
			sum = cur2.Val + complement
			cur2 = cur2.Next
		}

		val := sum % 10
		complement = sum / 10
		cur3.Next = &ListNode{Val: val, Next: nil}
		cur3 = cur3.Next
	}

	if complement > 0 {
		cur3.Next = &ListNode{Val: complement}
	}

	return result.Next
}
