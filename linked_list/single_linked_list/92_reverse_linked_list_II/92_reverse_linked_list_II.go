package reverselinkedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil || left == right {
		return head
	}

	current := head
	count := 1
	var headPos *ListNode
	var tailPos *ListNode

	for count < left {
		headPos = current
		current = current.Next
		count++
	}

	for count < right {
		current = current.Next
		tailPos = current.Next
		count++
	}

	if current.Next != nil {
		current.Next = nil
	}

	var tmp *ListNode
	if headPos == nil && tailPos == nil {
		result, _ := reverseList(head)
		return result
	} else if headPos == nil {
		head, tmp = reverseList(head)
		tmp.Next = tailPos
	} else {
		(*headPos).Next, tmp = reverseList(headPos.Next)
		tmp.Next = tailPos
	}

	return head
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
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

	return tmp1, head
}
