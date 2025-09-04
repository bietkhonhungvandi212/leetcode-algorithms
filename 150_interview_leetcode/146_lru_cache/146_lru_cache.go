package lrucache

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Capacity int
	Count    int
	NodeMap  map[int]*Node // get and put O(1) -> use hashmap for efficient lookup
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		return LRUCache{} // Fixed: Return zero-value struct instead of nil
	}

	mapper := make(map[int]*Node, capacity)
	Head := &Node{Next: nil, Prev: nil}
	Tail := &Node{Next: nil, Prev: nil}
	Head.Next = Tail
	Tail.Prev = Head

	return LRUCache{Capacity: capacity, Count: 0, NodeMap: mapper, Head: Head, Tail: Tail}
}

func (this *LRUCache) Get(key int) int {
	// check hashmap
	if node, ok := this.NodeMap[key]; ok {
		// -> if exist -> delete node -> put exist node at head
		this.delete(node)
		this.putHead(node)

		return node.Val
	}

	// -> if not exist -> return -1
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// check hashmap
	if node, ok := this.NodeMap[key]; ok {
		// if exist key -> change the value of node -> delete node then put node at head
		node.Val = value
		this.delete(node)
		this.putHead(node)
	} else {
		// count < capacity
		node := &Node{Val: value, Key: key}
		if this.Count < this.Capacity {
			// count ++
			this.Count++
		} else {
			// k == capacity
			// evict the tail's previous node -> delete key hashmap
			lruNode := this.Tail.Prev
			delete(this.NodeMap, lruNode.Key)
			this.delete(lruNode)
		}

		// put new node at head   -> put key in hashmap
		this.putHead(node)
		this.NodeMap[key] = node
	}
}

// Helper
func (this *LRUCache) putHead(node *Node) {
	tmp := this.Head.Next
	this.Head.Next = node
	tmp.Prev = node
	node.Next = tmp
	node.Prev = this.Head
}

func (this *LRUCache) delete(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	node = nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
