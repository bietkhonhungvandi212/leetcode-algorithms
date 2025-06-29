package linkedlistcycleii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			start := head
			for start != slow {
				start = start.Next
				slow = slow.Next
			}
			return start
		}
	}

	return nil
}
