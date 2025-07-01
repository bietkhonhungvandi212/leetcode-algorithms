package designlinkedlist

type node struct {
	val  int
	next *node
}
type MyLinkedList struct {
	head *node
	tail *node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: &node{}, tail: nil, size: 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	curr := l.head.next
	for range index {
		curr = curr.next
	}
	return curr.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	newNode := &node{val: val, next: l.head.next}
	l.head.next = newNode
	if l.size == 0 {
		l.tail = newNode
	}
	l.size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	newNode := &node{val: val}
	if l.size == 0 {
		l.head.next = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > l.size {
		return
	}
	if index == 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.size {
		l.AddAtTail(val)
		return
	}

	// traverse to predecessor
	prev := l.head
	for range index {
		prev = prev.next
	}
	newNode := &node{val: val, next: prev.next}
	prev.next = newNode
	l.size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}

	prev := l.head
	for range index {
		prev = prev.next
	}

	toDelete := prev.next
	prev.next = toDelete.next

	if toDelete == l.tail {
		l.tail = prev
		if l.tail == l.head {
			l.tail = nil
		}
	}

	l.size--
}
