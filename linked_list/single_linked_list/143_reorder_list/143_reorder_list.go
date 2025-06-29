package reorderlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	current := head
	stack := make([]*ListNode, 0)
	cnt := 0

	for current != nil {
		stack = append(stack, current)
		cnt++
		current = current.Next
	}

	prev := head
	if cnt%2 != 0 {
		cnt += 1
	}
	cnt = cnt / 2
	fmt.Println(cnt)
	for cnt > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tmp := prev.Next
		prev.Next = top
		top.Next = tmp
		prev = tmp
		cnt--
		if cnt == 0 {
			top.Next = nil

		}
	}

}
