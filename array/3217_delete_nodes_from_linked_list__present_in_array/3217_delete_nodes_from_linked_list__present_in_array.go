package deletenodesfromlinkedlistpresentinarray3217_delete_nodes_from_linked_list__present

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	set := make(map[int]bool, 0)

	for i := range nums {
		set[nums[i]] = true
	}

	dump := &ListNode{Next: head}
	prev := dump
	curr := head

	for curr != nil {
		if set[curr.Val] {
			prev.Next = curr.Next
		} else {
			prev = curr
		}

		curr = curr.Next
	}

	return dump.Next
}
