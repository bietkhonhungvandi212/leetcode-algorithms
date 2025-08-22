package lrucache

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}
type LRUCache struct {
	Head  *Node
	Tail  *Node
	Cache map[int]*Node
	Size  int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{Head: head, Tail: tail, Cache: make(map[int]*Node), Size: capacity}
}

func (this *LRUCache) Get(key int) int {
	if currNode, ok := this.Cache[key]; ok {
		this.remove(currNode)
		this.pushToFront(currNode)

		return currNode.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//Update the existed node

	if currNode, ok := this.Cache[key]; ok {
		this.remove(currNode)
		this.pushToFront(currNode)
		currNode.Val = value
		return
	}

	if len(this.Cache) == this.Size {
		lru := this.Tail.Prev
		this.remove(lru)
		delete(this.Cache, lru.Key)
	}

	newNode := &Node{Key: key, Val: value}
	this.pushToFront(newNode)
	this.Cache[key] = newNode

}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) pushToFront(node *Node) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}
