package main

import "fmt"

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

func (this *Node) height() int {
	if this == nil {
		return -1
	}

	return this.Height
}

// update Height
func (this *Node) updateHeight() {
	this.Height = max(this.Left.height(), this.Right.height()) + 1
}

func max(num1, num2 int) int {
	if num1 < num2 {
		return num2
	}

	return num1
}

// ======================= SINGLE ROTATION =====================================
// rotate in left direction
func leftRotation(this *Node) *Node {
	tmp := this.Right
	this.Right = tmp.Left
	tmp.Left = this
	this.updateHeight()
	tmp.updateHeight()

	return tmp
}

// rotate in right direction
func rightRotation(this *Node) *Node {
	tmp := this.Left
	this.Left = tmp.Right
	tmp.Right = this
	this.updateHeight()
	tmp.updateHeight()

	return tmp
}

// ======================= DOUBLE ROTATION =====================================
// Left Right Rotation
func leftRightRotation(node *Node) *Node {
	node.Left = leftRotation(node.Left)
	return rightRotation(node)
}

// Right Left Rotation
func rightLeftRotation(node *Node) *Node {
	node.Left = rightRotation(node.Left)
	return leftRotation(node)
}

// ======================= INSERTION =====================================
func balanceFactor(node *Node) int {
	if node == nil {
		return 0
	}

	return node.Left.height() - node.Left.height()
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value, Height: 0}
	}

	//Insert node into binary search tree
	if value > node.Value {
		node.Right = insert(node.Right, value)
	} else if value < node.Value {
		node.Left = insert(node.Left, value)
	} else {
		return node
	}

	node.updateHeight()
	bf := balanceFactor(node)

	//right rotation
	if bf > 1 && value < node.Left.Value {
		return rightRotation(node)
	}

	//left rotation
	if bf < -1 && value > node.Right.Value {
		return leftRotation(node)
	}

	//right left rotation
	if bf < -1 && value < node.Right.Value {
		return rightLeftRotation(node)
	}

	//left right rotation
	if bf > 1 && value > node.Left.Value {
		return leftRightRotation(node)
	}

	return node
}

func leftMostNode(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

// ======================= DELETION =====================================
func delete(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	if val > node.Value {
		node.Right = delete(node.Right, val)
	} else if val < node.Value {
		node.Left = delete(node.Left, val)
	} else {
		if node.Left == nil && node.Right == nil {
			// Case 1: No children
			return nil
		} else if node.Left != nil && node.Right == nil {
			// Case 2: Only left child
			return node.Left
		} else if node.Left == nil && node.Right != nil {
			// Case 3: Only right child
			return node.Right
		} else {
			// Case 4: Both children exist
			tmp := leftMostNode(node.Right)
			node.Value = tmp.Value
			node.Right = delete(node.Right, tmp.Value)
		}
	}

	// Update height after deletion
	node.updateHeight()
	bf := balanceFactor(node)

	// Left Heavy (bf > 1)
	if bf > 1 {
		if node.Left != nil && balanceFactor(node.Left) >= 0 {
			// Left-Left case
			node = rightRotation(node)
		} else if node.Left != nil && balanceFactor(node.Left) < 0 {
			// Left-Right case
			node = leftRightRotation(node)
		}
	}

	// Right Heavy (bf < -1)
	if bf < -1 {
		if node.Right != nil && balanceFactor(node.Right) <= 0 {
			// Right-Right case
			node = leftRotation(node)
		} else if node.Right != nil && balanceFactor(node.Right) > 0 {
			// Right-Left case
			node = rightLeftRotation(node)
		}
	}

	return node
}

func main() {
	// Create an RR imbalance: 1 -> 2 -> 3
	root := &Node{Value: 1, Height: 0}
	root.Right = &Node{Value: 2, Height: 0}
	root.Right.Right = &Node{Value: 3, Height: 0}
	root.Right.updateHeight()
	root.updateHeight()
	fmt.Println("Before Left Rotation:")
	printTree(root, 0)
	root = leftRotation(root)
	fmt.Println("After Left Rotation:")
	printTree(root, 0)
}

func printTree(node *Node, level int) {
	if node == nil {
		return
	}
	printTree(node.Right, level+1)
	for range level {
		fmt.Print("  ")
	}
	fmt.Printf("%d (Height: %d)\n", node.Value, node.Height)
	printTree(node.Left, level+1)
}
