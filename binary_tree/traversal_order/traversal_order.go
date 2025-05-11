package main

import "fmt"

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func (node *TreeNode) FindLeftNode() *TreeNode {

	if node.Left != nil {
		return node.Left.FindLeftNode()

	}

	return node
}

func (node *TreeNode) FindRightNode() *TreeNode {
	if node.Right != nil {
		return node.Right.FindRightNode()
	}

	return node
}

func (node *TreeNode) FindTheLeftLowestNode() *TreeNode {
	if node.Parent != nil && node.Parent.Left == node {
		return node.Parent.FindTheLeftLowestNode()
	}

	return node.Parent
}

func (node *TreeNode) FindTheRightLowestNode() *TreeNode {
	if node.Parent != nil && node.Parent.Left == node {
		return node.Parent.FindTheRightLowestNode()
	}

	return node.Parent
}

func (node *TreeNode) FindSuccessor() *TreeNode {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return node.Right.FindLeftNode()

	}

	return node.FindTheRightLowestNode()
}

func (node *TreeNode) FindPredecessor() *TreeNode {
	if node == nil {
		return nil
	}

	if node.Left != nil {
		return node.Left.FindRightNode()

	}

	return node.FindTheRightLowestNode()
}

func (node *TreeNode) DeleteNode() *TreeNode {
	if node == nil {
		return nil
	}

	if node.Left != nil || node.Right != nil { // If node is not a leaf
		var tmpNode *TreeNode
		if node.Left != nil {
			tmpNode = node.FindPredecessor()
		} else {
			tmpNode = node.FindSuccessor()
		}

		tmpVal := node.Val
		node.Val = tmpNode.Val
		tmpNode.Val = tmpVal

		return tmpNode.DeleteNode()
	}

	if node.Parent != nil {
		if node.Parent != nil && node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	}

	return node
}

func main() {
	// Build the tree:
	//      4
	//     / \
	//    2   6
	//   / \ / \
	//  1  3 5  7
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}

	// Set parent pointers
	root.Left.Parent = root
	root.Right.Parent = root
	root.Left.Left.Parent = root.Left
	root.Left.Right.Parent = root.Left
	root.Right.Left.Parent = root.Right
	root.Right.Right.Parent = root.Right

	// Test successor
	root.DeleteNode()
	fmt.Printf("Value of this root is %d\n", root.Val) // Should print 3
	successor := root.FindSuccessor()
	fmt.Printf("Successor of %d is %d\n", root.Val, successor.Val) // Should print 3

	// Test predecessor
	predecessor := root.FindPredecessor()
	fmt.Printf("Predecessor of %d is %d\n", root.Val, predecessor.Val) // Should print 4
}
