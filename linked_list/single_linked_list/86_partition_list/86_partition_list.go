package partitionlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev != nil {
		if prev.Next != nil && prev.Next.Val >= x {
			break
		}

		prev = prev.Next
	}

	if prev == nil {
		return dummy.Next
	}

	cur := prev.Next
	var slow *ListNode
	for cur != nil {
		if cur.Val < x {
			slow.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
			prev = prev.Next
		}

		slow = cur
		cur = cur.Next
	}

	return dummy.Next
}

// partition uses the Two-List approach for optimal clarity and correctness.
func partition_Optimized(head *ListNode, x int) *ListNode {
	// 1. Setup: Create two dummy heads for the two new lists.
	// lessHead will be the start of the list with nodes < x.
	// greaterHead will be the start of the list with nodes >= x.
	lessHead := &ListNode{}
	greaterHead := &ListNode{}

	// Create tail pointers to efficiently append to the new lists.
	lessTail := lessHead
	greaterTail := greaterHead
	current := head

	// 2. Iteration: Traverse the original list.
	for current != nil {
		if current.Val < x {
			// Append current node to the 'less' list.
			lessTail.Next = current
			lessTail = lessTail.Next
		} else {
			// Append current node to the 'greater' list.
			greaterTail.Next = current
			greaterTail = greaterTail.Next
		}
		// Move to the next node in the original list.
		current = current.Next
	}

	// 3. Cleanup & Combining:
	// Terminate the 'greater' list to prevent cycles.
	// This is crucial if the last node of the original list was < x.
	greaterTail.Next = nil

	// Connect the tail of the 'less' list to the head of the 'greater' list.
	lessTail.Next = greaterHead.Next

	// 4. Return: The head of the combined list is the node after our 'less' dummy head.
	return lessHead.Next
}
