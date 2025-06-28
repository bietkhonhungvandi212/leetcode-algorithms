package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	current := head
	set := make(map[int]bool)

	set[current.Val] = true

	for current.Next != nil {
		if set[current.Next.Val] {
			// Skip the duplicate node
			current.Next = current.Next.Next
		} else {
			set[current.Next.Val] = true
			current = current.Next
		}
	}

	return head
}

func deleteDuplicatesOptimal(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	current := head

	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
