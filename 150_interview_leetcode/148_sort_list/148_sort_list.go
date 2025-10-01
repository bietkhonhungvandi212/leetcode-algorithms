package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var mergeSort func(head *ListNode) *ListNode
	mergeSort = func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}

		if head.Next == nil {
			return head
		}

		slow := head
		fast := head.Next

		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}

		head1 := slow.Next
		slow.Next = nil

		tmpHead1 := mergeSort(head)
		tmpHead2 := mergeSort(head1)

		dummy := &ListNode{}
		previous := dummy
		for tmpHead1 != nil && tmpHead2 != nil {
			if tmpHead1.Val < tmpHead2.Val {
				previous.Next = tmpHead1
				tmpHead1 = tmpHead1.Next
			} else {
				previous.Next = tmpHead2
				tmpHead2 = tmpHead2.Next
			}

			previous = previous.Next
		}

		for tmpHead1 != nil {
			previous.Next = tmpHead1
			tmpHead1 = tmpHead1.Next
			previous = previous.Next
		}

		for tmpHead2 != nil {
			previous.Next = tmpHead2
			tmpHead2 = tmpHead2.Next
			previous = previous.Next
		}

		return dummy.Next
	}

	return mergeSort(head)

}
