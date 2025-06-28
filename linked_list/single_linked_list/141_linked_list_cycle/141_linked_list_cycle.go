package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	current := head
	reach := make(map[*ListNode]bool)

	for current != nil {
		if _, ok := reach[current]; ok {
			return true
		}

		reach[current] = true
		current = current.Next
	}

	return false
}

// Approach 2: Floyd's Cycle Detection (Tortoise & Hare) ⭐ MOST OPTIMAL
// Time: O(n), Space: O(1)
func hasCycleFloyd(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head      // Tortoise: moves 1 step
	fast := head.Next // Hare: moves 2 steps

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true // Cycle detected!
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

// Approach 3: Node Marking (Destructive but O(1) space)
// Time: O(n), Space: O(1) - but modifies the list!
func hasCycleDestructive(head *ListNode) bool {
	current := head

	for current != nil {
		if current.Val == -100001 { // Use impossible value as marker
			return true
		}
		current.Val = -100001 // Mark as visited
		current = current.Next
	}

	return false
}

// Approach 5: Optimized Hash Map (using addresses directly)
// Time: O(n), Space: O(n) - but faster than bool map
func hasCycleOptimizedMap(head *ListNode) bool {
	visited := make(map[*ListNode]struct{}) // Empty struct uses no memory

	for head != nil {
		if _, exists := visited[head]; exists {
			return true
		}
		visited[head] = struct{}{}
		head = head.Next
	}

	return false
}
